package auth

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
)

type KubernetesAccess struct {
	Config rest.Config
	ClientSet kubernetes.Clientset
}

func GetKubernetesAccess() (*KubernetesAccess, error) {
	kubeconfig, err := inClusterAuth()

	if kubeconfig == nil {
		kubeconfig, err = outOfClusterAuth()
	}
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	return &KubernetesAccess{
		Config: *kubeconfig,
		ClientSet: *clientset,
	}, nil
}

func inClusterAuth() (*rest.Config, error) {
	kubeconf, err := rest.InClusterConfig()

	if err != nil {
		return nil, nil
	}

	return kubeconf, nil
}

func outOfClusterAuth() (*rest.Config, error) {
	var kubeconfigFile *string

	homeKubeConf := filepath.Join(homedir.HomeDir(), ".kube", "config")
	if _, err := os.Stat(homeKubeConf); err == nil {
		kubeconfigFile = flag.String("c", homeKubeConf, "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfigFile = flag.String("c", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	if *kubeconfigFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	kubeconf, err := clientcmd.BuildConfigFromFlags("", *kubeconfigFile)
	if err != nil {
		panic(err.Error())
	}
	return kubeconf, nil
}