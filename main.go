package main

import (
	"context"
	"fmt"

	"github.com/kimagliardi/kube-explain/pkg/genai"
	"github.com/kimagliardi/kube-explain/pkg/k8s"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {

	clientset, err := k8s.NewCLientSet()
	if err != nil {
		panic(err.Error())
	}

	// Replace "node-name" with your node's name
	nodeName := "docker-desktop"

	// Get the Node object
	node, err := clientset.CoreV1().Nodes().Get(context.TODO(), nodeName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	//Print the Node details
	/*fmt.Printf("Name: %s\n", node.Name)
	fmt.Printf("UID: %s\n", node.UID)
	fmt.Printf("CreationTimestamp: %s\n", node.CreationTimestamp)
	fmt.Printf("Labels: %v\n", node.Labels)
	fmt.Printf("Annotations: %v\n", node.Annotations)
	fmt.Printf("Taints: %v\n", node.Spec.Taints)
	fmt.Printf("Conditions: %v\n", node.Status.Conditions)
		fmt.Printf("Node: %v\n", node)*/
	nodeInfo := fmt.Sprintf("Limits: %v\n", node)

	client := genai.NewClient()
	resp, err := client.Summarize(nodeInfo)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)

}
