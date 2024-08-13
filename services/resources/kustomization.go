package resources

import (
	"context"
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
