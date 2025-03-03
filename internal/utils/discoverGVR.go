package utils

import (
  "github.com/arldka/flammkuchen/services/k8sclient"
  "k8s.io/apimachinery/pkg/runtime/schema"
  "fmt"
)

func DiscoverGVR(kind string) (schema.GroupVersionResource, error) { 
	apiGroups, err := k8sclient.DiscoveryClient.ServerGroups()
	if err != nil {
		return schema.GroupVersionResource{}, err
	}

	for _, apiGroup := range apiGroups.Groups {
		groupVersion, err := k8sclient.DiscoveryClient.ServerResourcesForGroupVersion(apiGroup.PreferredVersion.GroupVersion)
		if err != nil {
			continue
		}

		for _, resource := range groupVersion.APIResources {
			if resource.Kind == kind {
				return schema.GroupVersionResource{
					Group:    apiGroup.Name,
					Version:  apiGroup.PreferredVersion.Version,
					Resource: resource.Name,
				}, nil
			}
		}
	}

	return schema.GroupVersionResource{}, fmt.Errorf("GVR not found for kind: %s", kind)
}
