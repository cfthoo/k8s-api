package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Get the kubeconfig file path
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		kubeconfig = filepath.Join(os.Getenv("HOME"), ".kube", "function")
	}

	// Create a Kubernetes clientset
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Read the ConfigMap
	cm, err := clientset.CoreV1().ConfigMaps("my-namespace").Get(context.TODO(), "configmap-name", metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			fmt.Printf("ConfigMap my-function not found in namespace my-namespace\n")
			return
		}
		panic(err.Error())
	}

	// Update the ConfigMap
	cm.Data["my-key"] = "new-value"
	_, err = clientset.CoreV1().ConfigMaps("my-namespace").Update(context.TODO(), cm, metav1.UpdateOptions{})
	if err != nil {
		panic(err.Error())
	}
}
