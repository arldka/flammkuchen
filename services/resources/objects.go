package resources

import (
  "fmt"
  "sync"
  "github.com/arldka/flammkuchen/internal/types"
  "github.com/arldka/flammkuchen/services/resources/objects"
)

func ObjectType(apiGroup string, kind string) (string){
  return "generic"
}

func GetObjects(inventory *types.Inventory) (*types.Objects, error) {
  objectList := &types.Objects{}
  var wg sync.WaitGroup
  var mu sync.Mutex
  for _, entry := range inventory.Entries {
    wg.Add(1)
    go func (entry types.Entry) {
      defer wg.Done()
      switch ObjectType(entry.APIGroup, entry.Kind) {
        case "generic":
          object := objects.GetGeneric(entry)
          if object != nil {
            mu.Lock()
            objectList.Generics = append(objectList.Generics, *object)
            fmt.Println("Generic Object Added")
            mu.Unlock()
          }
        }
    }(entry)
  }
  wg.Wait()
  return objectList, nil
}
