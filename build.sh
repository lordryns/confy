#!/bin/bash


if [[ -z $1  ]]; then
    echo "Hello from build.sh, add arguments to actually do anything!"
    exit
fi 

if [[ $1 == "build" ]]; then 
    go build -o ./bin/confy ./cmd/confy
    echo "Done!"
elif [[ $1 == "clean"  ]]; then 
    rm ./bin/confy
else
    echo "Invalid command!"
fi 
