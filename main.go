package main

import (
	"flag"
	"fmt"
	"path/filepath"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	// List Deployments
	fmt.Printf("Listing deployments in namespace %q:\n", apiv1.NamespaceDefault)
	a := make(map[string]string)
	list, err := deploymentsClient.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	// Shouldn't be a map, needs to be a slice
	// We are interested in only the value, not the key
	// In this case we are interested in the value, so we
	// shouldn't use the blank identifier
	for _, d := range list.Items {
		fmt.Printf(" * %s with annotations: (%v)\n", d.Name, d.Annotations)
		a := d.Annotations
		_, ok := a["repo"]
		fmt.Println("repo:", ok)
	}

	_, ok := a["repo"]
	fmt.Println("repo:", ok)

	// if _, ok := a["repo"]; ok {
	// 	fmt.Println("Yeah buddy")
	// 	for k := range a {
	// 		fmt.Printf("key[%v] value[%v]\n", k, a[k])
	// 	}
	// }
}
