package types

type Kustomization struct {
	Name      string
	Namespace string
	Status    string
	LastTransitionTime       string
}

type HelmRelease struct {
	Name      string
	Namespace string
	Status    string
	LastTransitionTime       string
}

type Inventory struct {
	Entries []Entry
}

type Entry struct {
	Name       string
	Namespace  string
	APIGroup   string
	APIVersion string
	Kind       string
}

type Objects struct {
	Generics    []GenericObject
	RBACs       []GenericObject
	CRDs        []GenericObject
	Workloads   []WorkloadObject
	Fluxes      []GenericObject
	Networkings []GenericObject
}

type GenericObject struct {
	Name       string
	Namespace  string
	APIGroup   string
	APIVersion string
	Kind       string
	Status     string
	LastTransitionTime        string
}

type WorkloadObject struct {
  Name string
  Namespace string
  APIGroup string 
  APIVersion string 
  Kind string 
  Details []string 
  Status string 
  LastTransitionTime string
}
