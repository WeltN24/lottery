VERSION := 0.3.3

export GITHUB_REPO := lottery
export GITHUB_USER := WeltN24
export TOKEN = `cat .token`

all: build

build: build/darwin build/linux

build/darwin:
	GO15VENDOREXPERIMENT=1 GOOS=darwin go build -o build/darwin/lottery

build/linux:
	GO15VENDOREXPERIMENT=1 GOOS=linux go build -o build/linux/lottery

clean:
	rm -rf build

release: clean build
	git tag $(VERSION) -f && git push --tags -f
	github-release release --tag $(VERSION) -s $(TOKEN)
	github-release upload --tag $(VERSION) -s $(TOKEN) --name lottery-osx --file build/darwin/lottery
	github-release upload --tag $(VERSION) -s $(TOKEN) --name lottery-linux --file build/linux/lottery