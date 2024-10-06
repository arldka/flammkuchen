package resources

import (
	"context"
	// "fmt"
	"strings"
	"github.com/arldka/flammkuchen/internal/types"
	"github.com/arldka/flammkuchen/services/k8sclient"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func FilteredListKustomizations(query string) ([]types.Kustomization, error) {
  var kustomizationList []types.Kustomization
  kustomizationGVR := schema.GroupVersionResource{
    Group: "kustomize.toolkit.fluxcd.io",
    Version: "v1",
    Resource: "kustomizations",
  }
  kustomizations, _ := k8sclient.DynamicClient.Resource(kustomizationGVR).List(context.TODO(), metav1.ListOptions{})
  for _, kustomization := range kustomizations.Items {
		meta := kustomization.Object["metadata"].(map[string]interface{})
    conditions := kustomization.Object["status"].(map[string]interface{})["conditions"].([]interface{})
    if strings.Contains(strings.ToLower(meta["name"].(string)), query) || strings.Contains(strings.ToLower(meta["namespace"].(string)), query) {
      kustomizationList = append(kustomizationList, types.Kustomization{
        Name: meta["name"].(string),
        Namespace: meta["namespace"].(string),
        Status: conditions[0].(map[string]interface{})["type"].(string),
        Age: meta["creationTimestamp"].(string),
      })
    }
  }
  return kustomizationList, nil
}

// create a function to retrieve the list of objects in the inventory of the kustomization object

func GetKustomizationInventory(namespace string, name string) (*types.Inventory, error) {
  kustomizationGVR := schema.GroupVersionResource{
    Group: "kustomize.toolkit.fluxcd.io",
    Version: "v1",
    Resource: "kustomizations",
  }
  kustomization, err := k8sclient.DynamicClient.Resource(kustomizationGVR).Namespace(namespace).Get(context.TODO(), name, metav1.GetOptions{})
  if err != nil {
    return nil, err
  }
  inventory := kustomization.Object["status"].(map[string]interface{})["inventory"].(interface{})
  var parsedEntries []types.Entry
    if inventory != nil {
        entries := inventory.(map[string]interface{})["entries"].([]interface{})
        for _, entry := range entries {
            entryMap := entry.(map[string]interface{})
            idParts := strings.Split(entryMap["id"].(string), "_")
            if len(idParts) >= 4 {
                parsedEntries = append(parsedEntries, types.Entry{
                    Namespace:  idParts[0],
                    Name:       idParts[1],
                    APIGroup: idParts[2],
                    APIVersion: entryMap["v"].(string),
                    Kind:       idParts[3],
                })
              // fmt.Printf("Namespace: %s, Name: %s, APIGroup: %s,APIVersion: %s, Kind: %s\n", idParts[0], idParts[1], idParts[2],entryMap["v"].(string), idParts[3])
            }
        }
    }
  return &types.Inventory{
    Entries: parsedEntries,
  }, nil
}
