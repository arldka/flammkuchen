package objects 

import (
  "fmt"
  "context"
  "github.com/arldka/flammkuchen/internal/types"
  "github.com/arldka/flammkuchen/services/k8sclient"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func GetGeneric(entry types.Entry) *types.GenericObject {
  genericGVR := schema.GroupVersionResource{
    Group: entry.APIGroup,
    Version: entry.APIVersion,
    Resource: entry.Kind, 
  }
  fmt.Println("Generic GVR:", genericGVR)
  fmt.Println("Generic Name:", entry.Name)
  fmt.Println("Generic Namespace:", entry.Namespace)
  generic, err := k8sclient.DynamicClient.Resource(genericGVR).Namespace(entry.Namespace).Get(context.TODO(), entry.Name, metav1.GetOptions{})
  if err != nil {
    fmt.Println("Error Getting Generic Object:", err)
    return nil
  }
  fmt.Println("Generic Object:", generic)
  status := generic.Object["status"].(map[string]interface{})["conditions"].([]interface{})[0].(map[string]interface{})["type"].(string)
  fmt.Println("Generic Object Status:", status)
  return &types.GenericObject{
    Name: entry.Name,
    Namespace: entry.Namespace,
    APIGroup: entry.APIGroup,
    APIVersion: entry.APIVersion,
    Kind: entry.Kind,
    Status: status,
  }
}
