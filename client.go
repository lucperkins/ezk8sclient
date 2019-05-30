package ezk8sclient

import (
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

type Client struct {
	k8s *kubernetes.Clientset
}

func NewClient() (*Client, error) {
	flags := genericclioptions.NewConfigFlags()
	restConfig, err := flags.ToRESTConfig()
	if err != nil {
		return nil, err
	}

	k8s, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	return &Client{
		k8s: k8s,
	}, nil
}