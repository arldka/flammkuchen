package resources

import (
	"context"
	"fmt"
	"strings"
  "fmt"
  "k8s.io/apimachinery/pkg/util/yaml"
	"github.com/arldka/flammkuchen/internal/types"
	"github.com/arldka/flammkuchen/internal/utils"
	"github.com/arldka/flammkuchen/services/k8sclient"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func FilteredListHelmReleases(query string) ([]types.HelmRelease, error) {
	var helmreleaseList []types.HelmRelease
	helmreleaseGVR := schema.GroupVersionResource{
		Group:    "helm.toolkit.fluxcd.io",
		Version:  "v2",
		Resource: "helmreleases",
	}
	helmreleases, _ := k8sclient.DynamicClient.Resource(helmreleaseGVR).List(context.TODO(), metav1.ListOptions{})
	for _, helmrelease := range helmreleases.Items {
		meta := helmrelease.Object["metadata"].(map[string]interface{})
		relativeTime, _ := utils.RelativeTime(meta["creationTimestamp"].(string))
		conditions := helmrelease.Object["status"].(map[string]interface{})["conditions"].([]interface{})
		if strings.Contains(strings.ToLower(meta["name"].(string)), query) || strings.Contains(strings.ToLower(meta["namespace"].(string)), query) {
			helmreleaseList = append(helmreleaseList, types.HelmRelease{
				Name:      meta["name"].(string),
				Namespace: meta["namespace"].(string),
				Status:    conditions[0].(map[string]interface{})["type"].(string),
				Age:       relativeTime,
			})
		}
	}
	return helmreleaseList, nil
}

func GetHelmReleaseInventory(namespace string, name string) (*types.Inventory, error) {
	settings := cli.New()
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), namespace, "secret", func(format string, v ...interface{}) {
		fmt.Printf(format, v...)
	}); err != nil {
		return nil, err
	}

	client := action.NewGet(actionConfig)
	rel, err := client.Run(name)
	if err != nil {
		return nil, err
	}
	var parsedEntries []types.Entry
	manifests := strings.Split(strings.TrimSpace(rel.Manifest), "---")
	for _, manifest := range manifests {
		var object map[string]interface{}
		if err := yaml.Unmarshal([]byte(manifest), &object); err == nil {
			gv := strings.Split(object["apiVersion"].(string), "/")
			version := ""
			group := ""
			if len(gv) > 1 {
				version = gv[1]
				group = gv[0]
			} else {
				version = gv[0]
				group = ""
			}
			parsedEntries = append(parsedEntries, types.Entry{
				Namespace:  namespace,
				Name:       object["metadata"].(map[string]interface{})["name"].(string),
				APIVersion: version,
				APIGroup:   group,
				Kind:       object["kind"].(string),
			})
		}
	}
	return &types.Inventory{
		Entries: parsedEntries,
	}, nil
}

func GetHelmReleaseInventory(namespace string, name string) (*types.Inventory, error) {
  // helmreleaseGVR := schema.GroupVersionResource{
  //   Group: "helm.toolkit.fluxcd.io",
  //   Version: "v2",
  //   Resource: "helmreleases",
  // }
  // helmrelease, _ := k8sclient.DynamicClient.Resource(helmreleaseGVR).Namespace(namespace).Get(context.TODO(), name, metav1.GetOptions{})
  settings := cli.New()
  actionConfig := new(action.Configuration)
  if err := actionConfig.Init(settings.RESTClientGetter(), namespace, "secret", func(format string, v ...interface{}) {
      fmt.Sprintf(format, v...)
  }); err != nil {
      return nil, err
  }

  client := action.NewGet(actionConfig)
  rel, err := client.Run(name)
  if err != nil {
      return nil, err
  }
  var parsedEntries []types.Entry
    manifests := strings.Split(strings.TrimSpace(rel.Manifest), "---")
  for _, manifest := range manifests {
      var object map[string]interface{}
      if err := yaml.Unmarshal([]byte(manifest), &object); err == nil {
          gv := strings.Split(object["apiVersion"].(string), "/")
          version := ""
          group := ""
          if len(gv) > 1 {
            version = gv[1]
            group = gv[0]
          } else {
            version = gv[0]
            group = ""
          }
          parsedEntries = append(parsedEntries, types.Entry{
            Namespace: namespace,
            Name: object["metadata"].(map[string]interface{})["name"].(string),
            APIVersion: version,
            APIGroup: group,
            Kind: object["kind"].(string),
          })
      }
  }
  return &types.Inventory{
    Entries: parsedEntries,
  }, nil
}
