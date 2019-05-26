package main

import (
	"bufio"
	"flag"
	"log"
	"os"
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
	log.Printf("Searching deployments in namespace %q:\n", apiv1.NamespaceDefault)

	list, err := deploymentsClient.List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Could not get list deployments: %s", err)
	}

	// Get length of deployments returned
	i := len(list.Items)
	// Make slice of strings with length i
	repos := make([]string, i)

	// Iterate over deployments and get annotations
	for i, d := range list.Items {
		log.Printf(" * %s with annotations: (%v)\n", d.Name, d.Annotations)
		a := d.Annotations

		// If annotations have a key of "repo"...
		if _, ok := a["repo"]; ok {
			// Get value of key, "repo"
			repo := a["repo"]
			// Add value at index i to slice
			repos[i] = repo
		}
	}

	log.Printf("Full slice of values: %v", repos)

	// Write file
	writeFile(repos)
}

func writeFile(toWrite []string) {
	file, err := os.OpenFile("output", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for i, data := range toWrite {
		log.Printf("Writing file, index %v", i)
		_, _ = datawriter.WriteString(data + "\n")
	}

	defer datawriter.Flush()
	defer file.Close()
}
