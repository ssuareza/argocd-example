# argocd-example

This repository is just a example of how to use Argo CD.



## Structure

There are two repositories:

- `/app`: contains the application files. Is a simple HTTP server that prints a message showing the environment name.

- `deploy`: holds the Kubernetes manifests required by Argo CD to deploy the application.
