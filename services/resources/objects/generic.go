package objects 

import (
  "fmt"
  "context"
  "strings"
  "github.com/arldka/flammkuchen/internal/types"
  "github.com/arldka/flammkuchen/services/k8sclient"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func GetGeneric(entry types.Entry) *types.GenericObject {
  genericGVR := schema.GroupVersionResource{
    Group:    entry.APIGroup,
    Version:  entry.APIVersion,
    Resource: strings.ToLower(entry.Kind) + "s",
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
    status = conditions[len(conditions)-1].(map[string]interface{})["type"].(string)
  }

  age := generic.Object["metadata"].(map[string]interface{})["creationTimestamp"].(string)

  return &types.GenericObject{
    Name: entry.Name,
    Namespace: entry.Namespace,
    APIGroup: entry.APIGroup,
    APIVersion: entry.APIVersion,
    Kind: entry.Kind,
    Status: status,
    Age: age,
  }
}
