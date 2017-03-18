# passhash

This repository contains code for a Developer Assesment as a part of an unnamed company's hiring process.

The company and assesment requirements will not be specified, out of respect for confidentiality.

## Installation Instructions
### Clone from source
* `git clone https://github.com/acarp52/passhash.git`

### Build project and run web server

To build the Go package:
* `go build`

To run the server:
* `./passhash [port number]` or `passhash.exe [port number]`
* `[port number]` must be an open port that can accept connections.

To stop the server:
* `CTRL-C` or `^C`
* This will begin a graceful shutdown of the server, enabling pending requests to complete and rejecting all new requests.


User must have Go properly installed and configured.


## Usage
### HTTP Endpoints

The server is enabled to accept the following parameters in a POST request body:
* `password=[your password]`
  
  Returns the SHA512 hash of [your password].
* `shutdown`
  
   Initiates a graceful shutdown of the web server.
   
   NOTE: Currently, a shutdown request will result in an empty response. This is not intended functionality, but will remain in place until a solution is developed. More details are present in code comments.

Any other request will be rejected as invalid. In the event on an invalid request, the user should see the appropriate error code and message.

### Example server query

Assuming port number 8080:

`curl -X POST --data "password=testPassword" http://localhost:8080`

should return:

`iluLRhHe5Gs9rzUx+rsqc6k6K+N26qJA3BFd1YGL0kpTPu7ppGqqJ8gGRRbkieYLdVM1Bud04ZeSKEKMkQrydQ==`

You may also open `http://localhost:8080` in a web browser and use the built-in form to make a request.

### Testing web server
The supplied testing suite is contained within test.sh, and can be run with the command:
* `bash test.sh`
