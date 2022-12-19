#!/bin/bash

helm --kubeconfig ~/.kube/config uninstall mongoclient

sleep 1

cd deployments/helm && helm package mongoclient/ && cd ../..

helm --kubeconfig ~/.kube/config install mongoclient deployments/helm/mongoclient-0.1.0.tgz
