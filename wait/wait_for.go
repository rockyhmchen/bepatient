package wait

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func WaitFor(ns string) {
	client, _ := getClient()

	testNs, err := client.CoreV1().Namespaces().Get(ns, metav1.GetOptions{})
	if testNs.GetName() == "" || err != nil {
		fmt.Printf("%s \360\237\222\251\n", err.Error())
		return
	}

	timeout := time.Now().Add(300 * time.Second)
	for {
		pods, _ := client.CoreV1().Pods(ns).List(metav1.ListOptions{})
		if isAllReady(pods) {
			fmt.Printf("All pods are ready! \360\237\230\201\n")
			break
		}

		fmt.Printf("Waiting for pods to be ready... \360\237\230\264\n")
		time.Sleep(10 * time.Second)

		if time.Now().After(timeout) {
			fmt.Printf("Timeout! All pods are not ready yet. \360\237\244\256\n")
			break
		}
	}
}

func getClient() (*kubernetes.Clientset, error) {
	kubeconfig := getKubeconfig()
	cfg, _ := clientcmd.BuildConfigFromFlags("", kubeconfig)

	return kubernetes.NewForConfig(cfg)
}

func getKubeconfig() string {
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		kubeconfig = filepath.Join(os.Getenv("HOME"), ".kube", "config")
	}

	return kubeconfig
}

func isAllReady(pods *apiv1.PodList) bool {
	if len(pods.Items) == 0 {
		return true
	}

	for _, pod := range pods.Items {
		if !isReady(pod) {
			return false
		}
	}

	return true
}

func isReady(pod apiv1.Pod) bool {
	for _, con := range pod.Status.Conditions {
		if con.Type == "Ready" && con.Status == "True" {
			return true
		}
	}

	return false
}
