# JSON DUMMY API
This simulates the data the object detection system outputs. This simple API is
used as a placeholder until the real object detection API is finished.

## Usage
- Download and install [Go](https://golang.org/dl/)
- `git clone https://github.com/Gods-Eye-BTH/json-dummy-api.git`
- `go get github.com/gorilla/mux`
- `go run main.go`

## Config
You can change the port (currently 8080) in the file `main.go`
Change the variable named port

## Source
This is a fork of [SpaceLenore/Gestra](https://github.com/SpaceLenore/Gestra) that
has been modified (with permission from original author) to send example data
similar to the data object detection will send.
