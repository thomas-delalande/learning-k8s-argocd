#!/bin/bash

eval $(minikube docker-env)
docker build -t learning-k8s-argocd -f server.dockerfile .
