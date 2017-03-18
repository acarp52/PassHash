#!/bin/bash

# Test suite that is not nearly as fleshed out as it could be. My bash knowledge is lacking.
# Mainly tests if server can handle concurrent connections.

PW1=`curl -s -X POST --data "password=angryMonkey" http://localhost:8080`
PW2=`curl -s -X POST --data "password=testPassword" http://localhost:8080`
PW3=`curl -s -X POST --data "password=foobar" http://localhost:8080`
HASH1 = "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q=="
HASH2 = "iluLRhHe5Gs9rzUx+rsqc6k6K+N26qJA3BFd1YGL0kpTPu7ppGqqJ8gGRRbkieYLdVM1Bud04ZeSKEKMkQrydQ=="
HASH3 = "ClAmHr0aOQ/tK/Mm8mc8FFWCpjQtUjIElz0CGTN/gWFqgGmwElh89WNfaSXxtWw2AjDBmyc1AO4BPgMGAb8kJQ=="

if [ $PW1!=$HASH1 ]
then
	echo "Problem with HASH1"
elif [ $PW2!=$HASH2 ]
then
	echo "Problem with HASH2"
elif [ $PW3!=$HASH3 ]
then
	echo "Problem with HASH3"
else
	echo "All HASH tests passed."
fi

CONCURRENT=$(
curl -s -X POST --data "password=foobar1" http://localhost:8080 &
curl -s -X POST --data "password=foobar2" http://localhost:8080 &
curl -s -X POST --data "password=foobar3" http://localhost:8080 &
curl -s -X POST --data "password=foobar4" http://localhost:8080 &
curl -s -X POST --data "password=foobar5" http://localhost:8080 &
curl -s -X POST --data "password=foobar6" http://localhost:8080 &
curl -s -X POST --data "password=foobar7" http://localhost:8080 &
curl -s -X POST --data "password=foobar8" http://localhost:8080 &
curl -s -X POST --data "password=foobar9" http://localhost:8080 &
)

wait
