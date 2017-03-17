#!/bin/bash

#HASH1=$(curl -X POST --data "password=angryMonkey" http://localhost:8080 &)
#HASH2=$(curl -X POST --data "password=angryMonkeyy" http://localhost:8080 &)

#echo "$HASH1"
#echo "$HASH2"

$ time  {
    array=();
    for i in {1..10}; do
      array+=(-X POST --data "password=angryMonkey" http://localhost:8080 ) ; 
    done; 
    curl "${array[@]}";
 }
