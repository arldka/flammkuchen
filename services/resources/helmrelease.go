package resources

import (
	"context"
	"strings"

	"github.com/arldka/flammkuchen/internal/types"
	"github.com/arldka/flammkuchen/services/k8sclient"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func FilteredListHelmReleases(query string) ([]types.HelmRelease, error) {
  var helmreleaseList []types.HelmRelease
  helmreleaseGVR := schema.GroupVersionResource{
    Group: "helm.toolkit.fluxcd.io",
    Version: "v2",
    Resource: "helmreleases",
  }
  helmreleases, _ := k8sclient.DynamicClient.Resource(helmreleaseGVR).List(context.TODO(), metav1.ListOptions{})
  for _, helmrelease := range helmreleases.Items {
		meta := helmrelease.Object["metadata"].(map[string]interface{})
    conditions := helmrelease.Object["status"].(map[string]interface{})["conditions"].([]interface{})
    if strings.Contains(strings.ToLower(meta["name"].(string)), query) || strings.Contains(strings.ToLower(meta["namespace"].(string)), query) {
      helmreleaseList = append(helmreleaseList, types.HelmRelease{
      Name: meta["name"].(string),
      Namespace: meta["namespace"].(string),
      Status: conditions[0].(map[string]interface{})["type"].(string),
      Age: meta["creationTimestamp"].(string),
      })
    }
  }
  return helmreleaseList, nil
}
