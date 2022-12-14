#!/bin/bash


make all


sudo docker tag mongoclient grzegorzpnk/mongoclient:latest

sudo docker push "grzegorzpnk"/mongoclient:"latest"

