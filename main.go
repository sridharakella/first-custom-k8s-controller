package main

import (
	"context"
	"fmt"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Println("Falling back to in-cluster config")
		config, err = rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
	}

	dynClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	thefoothebar := schema.GroupVersionResource{Group: "myk8s.io", Version: "v1", Resource: "thefoosthebars"}

	informer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return dynClient.Resource(thefoothebar).Namespace("").List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return dynClient.Resource(thefoothebar).Namespace("").Watch(context.TODO(), options)
			},
		},
		&unstructured.Unstructured{},
		0, // Skip resync
		cache.Indexers{},
	)

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("Add event detected:", obj)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Println("Update event detected:", newObj)
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("Delete event detected:", obj)
		},
	})

	stop := make(chan struct{})
	defer close(stop)

	go informer.Run(stop)

	if !cache.WaitForCacheSync(stop, informer.HasSynced) {
		panic("Timeout waiting for cache sync")
	}

	fmt.Println("Custom Resource Controller started successfully")

	<-stop
}
