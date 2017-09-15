#!/bin/bash

yes "" | curl https://glide.sh/get | sh
yes "" | sudo add-apt-repository ppa:masterminds/glide && sudo apt-get update
yes "" | sudo apt-get install glide

echo "glide install"
cd api && glide install
