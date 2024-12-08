#!/bin/bash

NAMESPACE=application
# create namespace first
# kubectl apply -f ./iac/k8s/application-namespace.yaml 
kubectl create namespace $NAMESPACE

# create required config and secret for the application
kubectl apply -f ./iac/k8s/application-configmap.yaml -n $NAMESPACE
kubectl apply -f ./iac/k8s/application-secret.yaml -n $NAMESPACE

# setup for mongodb
kubectl apply -f ./iac/k8s/mongo-5.0.yaml -n $NAMESPACE


# setup for service
kubectl apply -f ./iac/k8s/service.yaml -n $NAMESPACE