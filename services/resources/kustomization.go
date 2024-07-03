package resources

import (
	"context"
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
    kustomizationList = append(kustomizationList, types.Kustomization{
      Name: meta["name"].(string),
      Namespace: meta["namespace"].(string),
      Status: conditions[0].(map[string]interface{})["type"].(string),
      Age: meta["creationTimestamp"].(string),
    })
  }
  // apiGroupList, _ := k8sclient.DiscoveryClient.ServerGroups()
  // var versions []metav1.GroupVersionForDiscovery
  // for _, apiGroup := range apiGroupList.Groups {
  //   if apiGroup.Name == "kustomize.toolkit.fluxcd.io" {
  //     versions = apiGroup.Versions
  //   }
  // }

  // for _, version := range versions {
  //   gvr := schema.GroupVersionResource{
  //     Group: "kustomize.toolkit.fluxcd.io",
  //     Version: version.Version,
  //     Resource: "kustomizations",
  //   }
  //   // kustomization seesms to be nil
  //   fmt.Println("gvr: ", gvr)
  //   
  //   kustomizations, err := k8sclient.DynamicClient.Resource(gvr).List(context.TODO(), metav1.ListOptions{})
  //   if err != nil {
  //     fmt.Println("error: ", err)
  //     return nil, err
  //   }
  //   fmt.Println(len(kustomizations.Items))
  //   fmt.Println("I was here")
    // for _, kustomization := range kustomizations.Items {
    //   fmt.Println("I was here")
		  // meta := kustomization.Object["metadata"].(map[string]interface{})
    //   conditions := kustomization.Object["status"].(map[string]interface{})["conditions"].([]interface{})
    //   kustomizationList = append(kustomizationList, types.Kustomization{
    //     Name: meta["name"].(string),
    //     Namespace: meta["namespace"].(string),
    //     Status: conditions[0].(map[string]interface{})["type"].(string),
    //     Age: meta["creationTimestamp"].(string),
    //   })
    // }
  // }
  return kustomizationList, nil
}
