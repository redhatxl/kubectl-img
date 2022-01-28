package kube

import (
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

// ClientSet k8s clientset
func ClientSet(configFlags *genericclioptions.ConfigFlags) *kubernetes.Clientset {
	config, err := configFlags.ToRESTConfig()
	if err != nil {
		panic("kube config load error")
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {

		panic("gen kube config error")
	}
	return clientSet
}
