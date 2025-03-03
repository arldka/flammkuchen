package resources

import (
	"context"
	// "fmt"
	"github.com/arldka/flammkuchen/internal/types"
	"github.com/arldka/flammkuchen/internal/utils"
	"github.com/arldka/flammkuchen/services/k8sclient"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sort"
	"strings"
	"time"
)

func FilteredListKustomizations(query string) ([]types.Kustomization, error) {
	var kustomizationList []types.Kustomization
	kustomizationGVR := schema.GroupVersionResource{
		Group:    "kustomize.toolkit.fluxcd.io",
		Version:  "v1",
		Resource: "kustomizations",
	}
	kustomizations, _ := k8sclient.DynamicClient.Resource(kustomizationGVR).List(context.TODO(), metav1.ListOptions{})
	for _, kustomization := range kustomizations.Items {
		meta := kustomization.Object["metadata"].(map[string]interface{})
		conditions := kustomization.Object["status"].(map[string]interface{})["conditions"].([]interface{})
		sort.SliceStable(conditions, func(i, j int) bool {
			timeI, _ := time.Parse(time.RFC3339, conditions[i].(map[string]interface{})["lastTransitionTime"].(string))
			timeJ, _ := time.Parse(time.RFC3339, conditions[j].(map[string]interface{})["lastTransitionTime"].(string))
			return timeI.After(timeJ)
		})
		status := conditions[0].(map[string]interface{})["type"].(string)
		lastTransitionTime, _ := utils.RelativeTime(conditions[len(conditions)-1].(map[string]interface{})["lastTransitionTime"].(string))
		if strings.Contains(strings.ToLower(meta["name"].(string)), query) || strings.Contains(strings.ToLower(meta["namespace"].(string)), query) {
			kustomizationList = append(kustomizationList, types.Kustomization{
				Name:               meta["name"].(string),
				Namespace:          meta["namespace"].(string),
				Status:             status,
				LastTransitionTime: lastTransitionTime,
			})
		}
	}
	return kustomizationList, nil
}

func GetKustomizationInventory(namespace string, name string) (*types.Inventory, error) {
	kustomizationGVR := schema.GroupVersionResource{
		Group:    "kustomize.toolkit.fluxcd.io",
		Version:  "v1",
		Resource: "kustomizations",
	}
	kustomization, err := k8sclient.DynamicClient.Resource(kustomizationGVR).Namespace(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	inventory := kustomization.Object["status"].(map[string]interface{})["inventory"]
	var parsedEntries []types.Entry
	if inventory != nil {
		entries := inventory.(map[string]interface{})["entries"].([]interface{})
		for _, entry := range entries {
			entryMap := entry.(map[string]interface{})
			idParts := strings.Split(entryMap["id"].(string), "_")
			if len(idParts) >= 4 {
				parsedEntries = append(parsedEntries, types.Entry{
					Namespace:  idParts[0],
					Name:       idParts[1],
					APIGroup:   idParts[2],
					APIVersion: entryMap["v"].(string),
					Kind:       idParts[3],
				})
			}
		}
	}
	return &types.Inventory{
		Entries: parsedEntries,
	}, nil
}

func GetKustomization(name string, namespace string) (types.Kustomization, error) {
	kustomizationGVR := schema.GroupVersionResource{
		Group:    "kustomize.toolkit.fluxcd.io",
		Version:  "v1",
		Resource: "kustomizations",
	}
	kustomization, _ := k8sclient.DynamicClient.Resource(kustomizationGVR).Namespace(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	meta := kustomization.Object["metadata"].(map[string]interface{})
	conditions := kustomization.Object["status"].(map[string]interface{})["conditions"].([]interface{})
	sort.SliceStable(conditions, func(i, j int) bool {
		timeI, _ := time.Parse(time.RFC3339, conditions[i].(map[string]interface{})["lastTransitionTime"].(string))
		timeJ, _ := time.Parse(time.RFC3339, conditions[j].(map[string]interface{})["lastTransitionTime"].(string))
		return timeI.After(timeJ)
	})
	status := conditions[0].(map[string]interface{})["type"].(string)
	lastTransitionTime, _ := utils.RelativeTime(conditions[len(conditions)-1].(map[string]interface{})["lastTransitionTime"].(string))
	result := types.Kustomization{
		Name:               meta["name"].(string),
		Namespace:          meta["namespace"].(string),
		Status:             status,
		LastTransitionTime: lastTransitionTime,
	}
	return result, nil

}
