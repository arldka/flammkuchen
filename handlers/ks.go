package handlers

import (
	"fmt"
	"net/http"
  "github.com/gorilla/mux"
	"github.com/arldka/flammkuchen/services/resources"
  "github.com/arldka/flammkuchen/components"
  "github.com/arldka/flammkuchen/services/k8sclient"
)

func HandleKustomization(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  namespace := vars["namespace"]
  kustomizationName := vars["kustomizationName"]

  if namespace == "" || kustomizationName == "" {
        http.Error(w, "Missing namespace or kustomizationName", http.StatusBadRequest)
        return
  }

  fmt.Println("Namespace:", namespace)
  fmt.Println("Kustomization Name:", kustomizationName)
  inventory, _:= resources.GetKustomizationInventory(namespace, kustomizationName)
  if inventory == nil {
    http.Error(w, "Kustomization not found", http.StatusNotFound)
    return
  }
  objects, _ := resources.GetObjects(inventory)
  if objects == nil {
    http.Error(w, "No objects found", http.StatusNotFound)
  }
  serverVersion, _ := k8sclient.DiscoveryClient.ServerVersion()
	components.Kustomization(serverVersion.String(), objects).Render(r.Context(), w)
}
