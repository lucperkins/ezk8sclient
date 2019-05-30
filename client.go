package ezk8sclient

import (
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

func NewClient() (*kubernetes.Clientset, error) {
	flags := genericclioptions.NewConfigFlags()
	restConfig, err := flags.ToRESTConfig()
	if err != nil {
		return nil, err
	}

	k8s, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}

	return k8s, nil
}
