install:
	go get -v github.com/urfave/cli
	go get github.com/docker/docker/client
	go get -v github.com/laher/goxc

test:
	go test -v ./...

release:
	goxc -d=$TRAVIS_BUILD_DIR/dist -bc="linux darwin" -tasks=xc,archive