#!/bin/bash

#build bin file and build docker image
make all

sudo docker tag mongoclient grzegorzpnk/mongodbclient:latest

sudo docker push grzegorzpnk/mongodbclient:latest

