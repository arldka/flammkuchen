package resources

import (
	"slices"
	"strings"
	"sync"

	"github.com/arldka/flammkuchen/internal/types"
	"github.com/arldka/flammkuchen/services/resources/objects"
)

var networking = []string{
	"Service",
	"Ingress",
	"Gateway",
	"VirtualService",
	"HTTPRoute",
	"TCPRoute",
	"UDPRoute",
	"TLSRoute",
	"NetworkPolicy",
}

var workload = []string{
	"Deployment",
	"StatefulSet",
	"DaemonSet",
	"Job",
	"CronJob",
}

func ObjectType(apiGroup string, kind string) string {
	if apiGroup == "rbac.authorization.k8s.io" || kind == "ServiceAccount" {
		return "rbac"
	} else if apiGroup == "apiextensions.k8s.io" && kind == "CustomResourceDefinition" {
		return "crd"
	} else if slices.Contains(workload, kind) {
		return "workload"
	} else if slices.Contains(networking, kind) {
		return "networking"
	} else if strings.HasSuffix(apiGroup, "toolkit.fluxcd.io") {
		return "flux"
	} else {
		return "generic"
	}
}

func InsertGenericObject(objectList []types.GenericObject, newObject types.GenericObject) []types.GenericObject {
	index := 0
	for i, g := range objectList {
		if g.Kind == newObject.Kind {
			if g.Name > newObject.Name {
				index = i
				break
			}
		} else {
			if g.Kind > newObject.Kind {
				index = i
				break
			}
		}
	}
	objectList = append(objectList[:index], append([]types.GenericObject{newObject}, objectList[index:]...)...)
	return objectList
}

func InsertWorkloadObject(objectList []types.WorkloadObject, newObject types.WorkloadObject) []types.WorkloadObject {
	index := 0
	for i, g := range objectList {
		if g.Kind == newObject.Kind {
			if g.Name > newObject.Name {
				index = i
				break
			}
		} else {
			if g.Kind > newObject.Kind {
				index = i
				break
			}
		}
	}
	objectList = append(objectList[:index], append([]types.WorkloadObject{newObject}, objectList[index:]...)...)
	return objectList
}

func GetObjects(inventory *types.Inventory) (*types.Objects, error) {
	objectList := &types.Objects{}
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, entry := range inventory.Entries {
		wg.Add(1)
		go func(entry types.Entry) {
			defer wg.Done()
			switch ObjectType(entry.APIGroup, entry.Kind) {
			case "generic":
				object := objects.GetGeneric(entry)
				if object != nil {
					mu.Lock()
					objectList.Generics = InsertGenericObject(objectList.Generics, *object)
					mu.Unlock()
				}
			case "rbac":
				object := objects.GetGeneric(entry)
				if object != nil {
					mu.Lock()
					objectList.RBACs = InsertGenericObject(objectList.RBACs, *object)
					mu.Unlock()
				}
			case "crd":
				object := objects.GetGeneric(entry)
				if object != nil {
					mu.Lock()
					objectList.CRDs = InsertGenericObject(objectList.CRDs, *object)
					mu.Unlock()
				}
			case "workload":
				object := objects.GetWorkload(entry)
				if object != nil {
					mu.Lock()
					objectList.Workloads = InsertWorkloadObject(objectList.Workloads, *object)
					mu.Unlock()
				}
			case "flux":
				object := objects.GetGeneric(entry)
				if object != nil {
					mu.Lock()
					objectList.Fluxes = InsertGenericObject(objectList.Fluxes, *object)
					mu.Unlock()
				}
			case "networking":
				object := objects.GetGeneric(entry)
				if object != nil {
					mu.Lock()
					objectList.Networkings = InsertGenericObject(objectList.Networkings, *object)
					mu.Unlock()
				}
			}
		}(entry)
	}
	wg.Wait()
	return objectList, nil
}
