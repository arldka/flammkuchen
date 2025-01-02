package handlers

import (
	"fmt"
	"net/http"
	"github.com/arldka/flammkuchen/services/resources"
  "github.com/arldka/flammkuchen/components"
  "github.com/arldka/flammkuchen/services/k8sclient"
)

func HandleHelmRelease(w http.ResponseWriter, r *http.Request) {
  namespace := r.PathValue("namespace")
  helmReleaseName := r.PathValue("helmReleaseName")

  if namespace == "" || helmReleaseName == "" {
        http.Error(w, "Missing namespace or HelmRelease", http.StatusBadRequest)
        return
  }

  fmt.Println("Namespace:", namespace)
  fmt.Println("HelmRelease Name:", helmReleaseName)
  inventory, _:= resources.GetHelmReleaseInventory(namespace, helmReleaseName)
  if inventory == nil {
    http.Error(w, "HelmRelease not found", http.StatusNotFound)
    return
  }
  objects, _ := resources.GetObjects(inventory)
  if objects == nil {
    http.Error(w, "No objects found", http.StatusNotFound)
  }
  serverVersion, _ := k8sclient.DiscoveryClient.ServerVersion()
	components.HelmRelease(serverVersion.String(), objects).Render(r.Context(), w)
}
