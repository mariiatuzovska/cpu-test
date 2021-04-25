PROJECTNAME=$(shell basename "$(PWD)")

# Go related variables
GOBASE=$(shell pwd)
GOPATH=$(shell go env GOPATH)

show-cpu:
	go build -o test
	./test -cpu
	go tool pprof -top  ./cpu-profile.pb.gz

show-mem:
	go build -o test
	./test -mem
	go tool pprof -top  ./mem-profile.pb.gz

protoc:
	go get github.com/google/pprof
	protoc --decode=perftools.profiles.Profile  $(GOPATH)/src/github.com/google/pprof/proto/profile.proto --proto_path $(GOPATH)/src/github.com/google/pprof/proto/

clear:
	rm -f ./test
	rm -f ./mem-profile.pb.gz
	rm -f ./cpu-profile.pb.gz
	go mod tidy