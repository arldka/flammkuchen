package types

type Kustomization struct {
  Name string
  Namespace string
  Status string
  Age string 
}

type HelmRelease struct {
  Name string
  Namespace string
  Status string
  Age string
}

type Inventory struct {
  Entries []Entry
}

type Entry struct {
  Name string
  Namespace string
  APIGroup string
  APIVersion string
  Kind string
}

type Objects struct {
  Generics []GenericObject
  RBACs    []GenericObject
  CRDs     []GenericObject
  Workloads []GenericObject
  Fluxes   []GenericObject
  Networkings []GenericObject
}

type GenericObject struct {
  Name string
  Namespace string
  Labels map[string]string
  APIGroup string
  APIVersion string
  Kind string
  Status string
  Age string
}
