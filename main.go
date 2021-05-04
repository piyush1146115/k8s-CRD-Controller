package main

import (
	"flag"
	"time"

	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	clientset "github.com/piyush1146115/k8s-CRD-Controller/pkg/generated/clientset/versioned"
	informers "github.com/piyush1146115/k8s-CRD-Controller/pkg/generated/informers/externalversions"
	"github.com/piyush1146115/k8s-CRD-Controller/pkg/signals"
)


var (
	masterURL  string
	kubeconfig string
)

func main(){
	klog.InitFlags(nil)
	flag.Parse()
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}