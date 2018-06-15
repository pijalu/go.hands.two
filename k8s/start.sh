#!/bin/bash

set -e

echo "Starting minikube (if needed)"
minikube status > /dev/null 2>&1 || minikube start

echo "Enabling ingress"
minikube addons enable ingress

echo "Setting up namespace"
kubectl apply -f namespace.yml

echo "Setting up PostgreSQL"
kubectl apply -f postgres.yml

echo "Setting up RBAC service account"
kubectl apply -f rbac.yml

echo "Setting up micro services"
echo "* Data"
kubectl apply -f frinsultdata.yml
echo "* Gate"
kubectl apply -f frinsultgate.yml

echo "Setting up frontend"
kubectl apply -f frinsultfront.yml

echo "Setting up ingress"
kubectl apply -f ingress.yml


