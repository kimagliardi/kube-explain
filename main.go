package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	openai "github.com/sashabaranov/go-openai"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// Use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
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

	/* Print the Node details
	fmt.Printf("Name: %s\n", node.Name)
	fmt.Printf("UID: %s\n", node.UID)
	fmt.Printf("CreationTimestamp: %s\n", node.CreationTimestamp)
	fmt.Printf("Labels: %v\n", node.Labels)
	fmt.Printf("Annotations: %v\n", node.Annotations)
	fmt.Printf("Taints: %v\n", node.Spec.Taints)
	fmt.Printf("Conditions: %v\n", node.Status.Conditions)
	*/

	openai_key := os.Getenv("OPEN_AI_API_KEY")
	if openai_key == "" {
		fmt.Println("OPEN_AI_API_KEY is not set")
		return
	}

	client := openai.NewClient(openai_key)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a skilled Kubernetes administrator. Your task is to provide concise summaries of Kubernetes resources based on their descriptions. Focus on extracting and highlighting critical information such as status, configuration details, resource allocations, events, and any warnings or anomalies. Your summary should distill the essential details into a brief overview, offering quick insights into the resource's current state, usage, and health.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("I need a summary of this Kubernetes resource. Please highlight the key aspects like status, resource usage, and any critical events or warnings that need attention. Here's the output of the `kubectl describe` command: %v", node.Status.Conditions),
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
