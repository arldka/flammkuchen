package objects

import (
  "github.com/arldka/flammkuchen/internal/types"
  "github.com/arldka/flammkuchen/internal/utils"
  "time"
  "sort"
  "fmt"
  "context"
  "github.com/arldka/flammkuchen/services/k8sclient"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

)

func GetWorkload(entry types.Entry) *types.WorkloadObject {
  resource, err := utils.GetResourceFromGroupVersionKind(entry.APIGroup, entry.APIVersion, entry.Kind)
	if err != nil {
		fmt.Printf("Error Getting Resource:%v, API Group:%v, API Version:%v, Kind:%v\n", err, entry.APIGroup, entry.APIVersion, entry.Kind)
	}
	genericGVR := schema.GroupVersionResource{
		Group:    entry.APIGroup,
		Version:  entry.APIVersion,
		Resource: resource,
	}
	generic, err := k8sclient.DynamicClient.Resource(genericGVR).Namespace(entry.Namespace).Get(context.TODO(), entry.Name, metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Error Getting Generic Object:%v, API Group:%v, API Version:%v, Resource:%v, Name:%v, Namespace:%v\n", err, entry.APIGroup, entry.APIVersion, entry.Kind, entry.Name, entry.Namespace)
		return nil
	}

	objectStatus := generic.Object["status"]
	var status = ""
  var lastTransitionTime = ""
	if objectStatus != nil && objectStatus.(map[string]interface{})["conditions"] != nil {
		conditions := objectStatus.(map[string]interface{})["conditions"].([]interface{})
		if conditions[0].(map[string]interface{})["lastTransitionTime"] == nil {
			status = "Unknown"
      lastTransitionTime, _ = utils.RelativeTime(generic.Object["metadata"].(map[string]interface{})["creationTimestamp"].(string))

		} else {
			sort.SliceStable(conditions, func(i, j int) bool {
				timeI, _ := time.Parse(time.RFC3339, conditions[i].(map[string]interface{})["lastTransitionTime"].(string))
				timeJ, _ := time.Parse(time.RFC3339, conditions[j].(map[string]interface{})["lastTransitionTime"].(string))
				return timeI.After(timeJ)
			})
			status = conditions[0].(map[string]interface{})["type"].(string)
      lastTransitionTime, err = utils.RelativeTime(conditions[len(conditions) - 1].(map[string]interface{})["lastTransitionTime"].(string))
      if err != nil {
        lastTransitionTime, _ = utils.RelativeTime(generic.Object["metadata"].(map[string]interface{})["creationTimestamp"].(string))
      }
    }
	} else {
    lastTransitionTime, _ = utils.RelativeTime(generic.Object["metadata"].(map[string]interface{})["creationTimestamp"].(string))
  }

  images := getImages(entry)

	return &types.WorkloadObject{
		Name:       entry.Name,
		Namespace:  entry.Namespace,
		APIGroup:   entry.APIGroup,
		APIVersion: entry.APIVersion,
		Kind:       entry.Kind,
    Details:    images,
		Status:     status,
		LastTransitionTime:        lastTransitionTime,
	}
}


func getImages(entry types.Entry) ([]string) {
  var images []string

  extractImages := func(containers []v1.Container) {
		for _, container := range containers {
			images = append(images, container.Image)
		}
	}

  switch entry.Kind {
    case "Deployment":
      deployment, _ := k8sclient.Clientset.AppsV1().Deployments(entry.Namespace).Get(context.TODO(), entry.Name, metav1.GetOptions{})
      extractImages(deployment.Spec.Template.Spec.Containers)
    case "StatefulSet":
      statefulset, _ := k8sclient.Clientset.AppsV1().StatefulSets(entry.Namespace).Get(context.TODO(), entry.Name, metav1.GetOptions{})
      extractImages(statefulset.Spec.Template.Spec.Containers)
    case "DaemonSet":
      daemonset, _ := k8sclient.Clientset.AppsV1().DaemonSets(entry.Namespace).Get(context.TODO(), entry.Name, metav1.GetOptions{})
      extractImages(daemonset.Spec.Template.Spec.Containers)
    case "Job":
      job, _ := k8sclient.Clientset.BatchV1().Jobs(entry.Namespace).Get(context.TODO(), entry.Name, metav1.GetOptions{})
      extractImages(job.Spec.Template.Spec.Containers)
    case "Cronjob":
      cronjob, _ := k8sclient.Clientset.BatchV1().CronJobs(entry.Namespace).Get(context.TODO(), entry.Name, metav1.GetOptions{})
      extractImages(cronjob.Spec.JobTemplate.Spec.Template.Spec.Containers)
  }
  return images
}
