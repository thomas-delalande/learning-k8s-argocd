#!/bin/bash

minikube start --container-runtime=docker
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
kubectl apply -f kubernetes/application.yaml

./scripts/build-docker-image.sh

echo "Default Password: "
argocd admin initial-password -n argocd
