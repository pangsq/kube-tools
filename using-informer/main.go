package main

import (
	"flag"
	"fmt"
	"os/user"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

var namespace string
var kubeConfigFile string
var inCluster bool

func main() {
	curUser, _ := user.Current()
	flag.StringVar(&namespace, "n", "default", "namespace")
	flag.StringVar(&kubeConfigFile, "k", curUser.HomeDir+"/.kube/config", "kube config file")
	flag.BoolVar(&inCluster, "in", false, "whether it is in the cluster")
	flag.Parse()

	fmt.Println("for using informer")
	fmt.Println("+++++++++++++++++++++++++++++")

	var config *rest.Config
	var err error
	if inCluster {
		config, err = rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
	} else {
		config, err = clientcmd.BuildConfigFromFlags("", kubeConfigFile)
		if err != nil {
			panic(err.Error())
		}
	}
	var clientset *kubernetes.Clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	run(clientset)
}

func run(c *kubernetes.Clientset) {
	podListWatcher := cache.NewListWatchFromClient(c.CoreV1().RESTClient(), "pods", namespace, fields.Everything())

	var indexer cache.Indexer
	var informer cache.Controller
	indexer, informer = cache.NewIndexerInformer(podListWatcher, &v1.Pod{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				fmt.Println("Add a pod: ", key)
			}
		},
		UpdateFunc: func(old interface{}, new interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(new)
			if err == nil {
				fmt.Print("Update a pod: ", key)
			}
			item, _, _ := indexer.GetByKey(key)
			newStatus := item.(*v1.Pod).Status.Phase
			fmt.Println(" => ", newStatus)
		},
		DeleteFunc: func(obj interface{}) {
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			if err == nil {
				fmt.Println("Delete a pod: ", key)
			}
		},
	}, cache.Indexers{})
	informer.Run(make(chan struct{}))
	select {} // run forever
}
