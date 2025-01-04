package handlers

import (
	"github.com/arldka/flammkuchen/components"
	"github.com/arldka/flammkuchen/services/k8sclient"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	serverVersion, _ := k8sclient.DiscoveryClient.ServerVersion()
	components.Index(serverVersion.String()).Render(r.Context(), w)
}
