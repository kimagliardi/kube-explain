# Kubernetes Resource Explainer

This is a kubectl plugin that uses OpenAI to generate a human-friendly explanation of Kubernetes resources. It can explain different types of resources, such as nodes and deployments.

## Installation

To install the plugin, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/kimagliardi/kube-explain.git
```

2. Change into the project directory:
```
cd kube-explain
```

3. Build the project

```
go build -o kubectl-explain
```

4. Move the `kubectl-explain` binary to a directory in your PATH:
```
mv kubectl-explain /usr/local/bin
```


## Usage 
To use the plugin, run the `kubectl explain` command followed by the --type flag and the name of the resource you want to explain:

```
kubectl explain --type=node nodename
kubectl explain --type=deployment deploymentname
```

If you want to use a different kubeconfig file, you can set the `--kubeconfig` flag:

```
kubectl explain --type=node --kubeconfig=/path/to/kubeconfig nodename
```
## Contributing
Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## License
This project is licensed under the MIT License. See the LICENSE file for details.