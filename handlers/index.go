package handlers

import (
	"fmt"
	"github.com/arldka/flammkuchen/components"
	"github.com/arldka/flammkuchen/services/k8sclient"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	serverVersion, _ := k8sclient.DiscoveryClient.ServerVersion()
	err := components.Index(serverVersion.String()).Render(r.Context(), w)
	if err != nil {
		fmt.Printf("Error rendering Index:%v\n", err)
	}
}
