package k8s

import (
	"flag"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func getKubeConfig() *string {
	if home := homedir.HomeDir(); home != "" {
		return flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		return flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
}

func buildConfigFromFlags() (*rest.Config, error) {
	kubeconfig := getKubeConfig()
	// Use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return &rest.Config{}, err
	}
	return config, nil

}

func NewCLientSet() (*kubernetes.Clientset, error) {
	config, err := buildConfigFromFlags()
	if err != nil {
		return &kubernetes.Clientset{}, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return &kubernetes.Clientset{}, err
	}
	return clientset, nil
}
