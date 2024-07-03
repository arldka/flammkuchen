package k8sclient

import (
	"os"
  "path/filepath"
  "k8s.io/client-go/tools/clientcmd"
  "k8s.io/client-go/kubernetes"
  "k8s.io/client-go/rest"
  "k8s.io/client-go/dynamic"
  "k8s.io/client-go/discovery"
	"k8s.io/client-go/util/homedir"
)

var (
  Clientset *kubernetes.Clientset
  DynamicClient *dynamic.DynamicClient
  DiscoveryClient *discovery.DiscoveryClient
)

func init() {
}

func initClients() (*kubernetes.Clientset, *dynamic.DynamicClient, *discovery.DiscoveryClient) {
  var err error
  var config *rest.Config
  if os.Getenv("LOCAL") == "true" {
    config, err = clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
  } else {
    config, err = rest.InClusterConfig()
  }
  if err != nil {
    panic(err.Error())
  }
  clientset, err := kubernetes.NewForConfig(config)
  if err != nil {
    panic(err.Error())
  }
  dynamicClient, err := dynamic.NewForConfig(config)
  if err != nil {
    panic(err.Error())
  }
  discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
  if err != nil {
    panic(err.Error())
  }
  return clientset, dynamicClient, discoveryClient
}
