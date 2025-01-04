package objects

import (
	"context"
	"fmt"
	"github.com/arldka/flammkuchen/internal/types"
	"github.com/arldka/flammkuchen/internal/utils"
	"github.com/arldka/flammkuchen/services/k8sclient"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sort"
	"time"
)

func GetGeneric(entry types.Entry) *types.GenericObject {
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
	if objectStatus != nil && objectStatus.(map[string]interface{})["conditions"] != nil {
		conditions := objectStatus.(map[string]interface{})["conditions"].([]interface{})
		if conditions[0].(map[string]interface{})["lastTransitionTime"] == nil {
			status = "Unknown"
		} else {
			sort.SliceStable(conditions, func(i, j int) bool {
				timeI, _ := time.Parse(time.RFC3339, conditions[i].(map[string]interface{})["lastTransitionTime"].(string))
				timeJ, _ := time.Parse(time.RFC3339, conditions[j].(map[string]interface{})["lastTransitionTime"].(string))
				return timeI.After(timeJ)
			})
			status = conditions[0].(map[string]interface{})["type"].(string)
		}
	}
	relativeAge, _ := utils.RelativeTime(generic.Object["metadata"].(map[string]interface{})["creationTimestamp"].(string))

	return &types.GenericObject{
		Name:       entry.Name,
		Namespace:  entry.Namespace,
		APIGroup:   entry.APIGroup,
		APIVersion: entry.APIVersion,
		Kind:       entry.Kind,
		Status:     status,
		Age:        relativeAge,
	}
}
