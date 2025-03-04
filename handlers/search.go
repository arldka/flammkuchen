package handlers

import (
	"fmt"
	"github.com/arldka/flammkuchen/components"
	"github.com/arldka/flammkuchen/services/resources"
	"net/http"
)

// write a handler that returns a filtered list of Kustomizations & HelmReleases based on a search query
func HandleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	kustomizations, _ := resources.FilteredListKustomizations(query)
	helmReleases, _ := resources.FilteredListHelmReleases(query)
	err := components.Objects(kustomizations, helmReleases).Render(r.Context(), w)
	if err != nil {
		fmt.Printf("Error rendering Objects:%v\n", err)
	}
}
