package handlers

import (
	"fmt"
	"github.com/arldka/flammkuchen/components"
	"github.com/arldka/flammkuchen/services/k8sclient"
	"github.com/arldka/flammkuchen/services/resources"
	"net/http"
)

func HandleHelmRelease(w http.ResponseWriter, r *http.Request) {
	namespace := r.PathValue("namespace")
	helmReleaseName := r.PathValue("helmReleaseName")

	if namespace == "" || helmReleaseName == "" {
		http.Error(w, "Missing namespace or HelmRelease", http.StatusBadRequest)
		return
	}

	inventory, _ := resources.GetHelmReleaseInventory(namespace, helmReleaseName)
	if inventory == nil {
		http.Error(w, "HelmRelease not found", http.StatusNotFound)
		return
	}
	objects, _ := resources.GetObjects(inventory)
	if objects == nil {
		http.Error(w, "No objects found", http.StatusNotFound)
	}
	serverVersion, _ := k8sclient.DiscoveryClient.ServerVersion()

  helmrelease, _ := resources.GetHelmRelease(helmReleaseName, namespace)

	err := components.HelmRelease(helmrelease, serverVersion.String(), objects).Render(r.Context(), w)
	if err != nil {
		fmt.Printf("Error rendering HelmRelease:%v\n", err)
	}
}
