package function

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

// getConfigMapContent this func returns the cm content of a given function map name.
func getConfigMapContent(client kubernetes.Interface, nameSpace, cmName string) (map[string]string, error) {
	cm, err := client.CoreV1().ConfigMaps(nameSpace).Get(context.TODO(), cmName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return cm.Data, nil
}

// updateConfigMapContent this func updates the cm content of a given function map by key.
func updateConfigMapContent(client kubernetes.Interface, cm *v1.ConfigMap, key string, newValue string) {
	cm.Data[key] = newValue
	_, err := client.CoreV1().ConfigMaps(cm.Namespace).Update(context.TODO(), cm, metav1.UpdateOptions{})
	if err != nil {
		panic(err.Error())
	}

}

// CreateK8sClientInCluster this func creates a k8s client
func CreateK8sClientInCluster() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return client, err
}

// CreateK8sClientCmd this func creates a k8s client from kubeconfig file.
func CreateK8sClientCmd(kubeconfig string) (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}
	if client, err := kubernetes.NewForConfig(config); err != nil {
		return nil, err
	} else {
		return client, err
	}
}
