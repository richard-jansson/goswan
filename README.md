# Goswan 
Yet another software rendering engine. This time written in go as a way for me to investigate, parallelisation of said process. 

## Build instructions
Build instructions, I am on Ubuntu 16.04. Do not hesitate to contact me regarding how to compile and run on different platforms / distributions.  

sudo apt-get install golang-go

cd $HOME

mkdir gocode 

export GOPATH="$HOME/gocode"

mkdir "$GOPATH/src"

cd "$GOPATH/src"

git clone "https://github.com/richard-jansson/goswan"

cd goswan 

chmod +x compile.sh

./compile.sh
