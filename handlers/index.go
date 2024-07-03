package handlers

import (
	"net/http"
  "github.com/arldka/flammkuchen/services/k8sclient"
	"github.com/arldka/flammkuchen/components"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
  serverVersion, _ := k8sclient.DiscoveryClient.ServerVersion()
	components.Index(serverVersion.String()).Render(r.Context(), w)
}
