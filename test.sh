#!/bin/bash

HASH1=`curl -s -X POST --data "password=angryMonkey" http://localhost:8080 &`
HASH2=`curl -s -X POST --data "password=angryMonkeyy" http://localhost:8080 &`
HASH3=`curl -s -X POST --data "password=angryMonkeyyy" http://localhost:8080 &`


CONCURRENT=$(
curl -s -X POST --data "password=angryMonkey1" http://localhost:8080 &
curl -s -X POST --data "password=angryMonkey2" http://localhost:8080 &
curl -s -X POST --data "password=angryMonkey3" http://localhost:8080 &
curl -s -X POST --data "password=angryMonkey4" http://localhost:8080 &
curl -s -X POST --data "password=angryMonkey5" http://localhost:8080 &
curl -s -X POST --data "password=angryMonkey6" http://localhost:8080 &
curl -s -X POST --data "password=angryMonkey7" http://localhost:8080 &
curl -s -X POST --data "password=angryMonkey8" http://localhost:8080 &
curl -s -X POST --data "password=angryMonkey9" http://localhost:8080 &
)

wait
