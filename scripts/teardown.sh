#!/bin/bash

kubectl delete -f kubernetes/application.yaml
kubectl delete -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
kubectl delete namespace argocd
minikube stop
