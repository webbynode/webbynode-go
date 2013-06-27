#!/bin/bash
cd /tmp
wget https://go.googlecode.com/files/go1.1.1.linux-amd64.tar.gz
tar vzxf go1.1.1.linux-amd64.tar.gz
mv /tmp/go /usr/local
sed -i '4iexport PATH=$PATH:$GOROOT/bin' ~/.bashrc
sed -i '4iexport GOROOT=/usr/local/go' ~/.bashrc

