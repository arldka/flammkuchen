package handlers

import (
	"fmt"
	"github.com/arldka/flammkuchen/components"
	"github.com/arldka/flammkuchen/services/k8sclient"
	"github.com/arldka/flammkuchen/services/resources"
	"net/http"
)

func HandleKustomization(w http.ResponseWriter, r *http.Request) {
	namespace := r.PathValue("namespace")
	kustomizationName := r.PathValue("kustomizationName")

	if namespace == "" || kustomizationName == "" {
		http.Error(w, "Missing namespace or kustomizationName", http.StatusBadRequest)
		return
	}

	inventory, _ := resources.GetKustomizationInventory(namespace, kustomizationName)
	if inventory == nil {
		http.Error(w, "Kustomization not found", http.StatusNotFound)
		return
	}
	objects, _ := resources.GetObjects(inventory)
	if objects == nil {
		http.Error(w, "No objects found", http.StatusNotFound)
	}
	serverVersion, _ := k8sclient.DiscoveryClient.ServerVersion()
	err := components.Kustomization(serverVersion.String(), objects).Render(r.Context(), w)
	if err != nil {
		fmt.Printf("Error rendering Kustomization:%v\n", err)
	}
}
