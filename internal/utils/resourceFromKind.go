package utils

import "github.com/arldka/flammkuchen/services/k8sclient"

func GetResourceFromGroupVersionKind(group string, version string, kind string) (string, error) {
	var searchInput string
	if group == "" {
		searchInput = version
	} else {
		searchInput = group + "/" + version
	}
	apiResourceList, err := k8sclient.DiscoveryClient.ServerResourcesForGroupVersion(searchInput)
	if err != nil {
		return "", err
	}
	for _, apiResource := range apiResourceList.APIResources {
		if apiResource.Kind == kind {
			return apiResource.Name, nil
		}
	}
	return "", nil
}
