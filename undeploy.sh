#!/bin/bash

helm --kubeconfig ~/.kube/core.config uninstall mongoclient

sleep 1
