package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// uses the current context in kubeconfig
	config, _ := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	// creates the clientset
	clientset, _ := kubernetes.NewForConfig(config)
	// access the API to list pods
	pods, _ := clientset.CoreV1().Pods("").List(v1.ListOptions{})
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
}
