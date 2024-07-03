package handlers

import (
	"net/http"
   "github.com/arldka/flammkuchen/services/resources"
    "github.com/arldka/flammkuchen/components"
)

// write a handler that returns a filtered list of Kustomizations & HelmReleases based on a search query
func HandleSearch(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query().Get("q")
  kustomizations, _ := resources.FilteredListKustomizations(query)
  helmReleases, _ := resources.FilteredListHelmReleases(query)
  components.Objects(kustomizations, helmReleases).Render(r.Context(), w)
}
