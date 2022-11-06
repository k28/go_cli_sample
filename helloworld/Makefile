
all:
	go build

test:
	go test $(go list ./... | grep -v /cmd/) 

install:
	go install

clean:
	rm ./helloworld

