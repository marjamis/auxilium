.DEFAULT_GOAL := helper
GIT_COMMIT ?= $(shell git rev-parse --short=12 HEAD || echo "NoGit")
BUILD_TIME ?= $(shell date -u '+%Y-%m-%d_%H:%M:%S')
TEXT_RED = \033[0;31m
TEXT_BLUE = \033[0;34;1m
TEXT_GREEN = \033[0;32;1m
TEXT_NOCOLOR = \033[0m

DEMO_FOLDER = /tmp/vhs_demo

helper: # Adapted from: https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
	@echo "Available targets..." # @ will not output shell command part to stdout that Makefiles normally do but will execute and display the output.
	@grep -hE '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Builds the application
	go build -o $$GOPATH/bin/auxilium main.go

test: build ## Builds and then runs tests against the application
	golint ./...
	go test -v -coverprofile cp.out ./...
	go tool cover -html=cp.out

	go test -cpuprofile cpu.prof -memprofile mem.prof 'github.com/marjamis/auxilium'
	go tool pprof -http=":8081" cpu.prof &
	go tool pprof -http=":8082" mem.prof &

demo: build
	# Creates a bit of a "blank environment" for the gif demo capture. If more isolation is required that can be done in the future
	mkdir $(DEMO_FOLDER) || echo "Temporary directory ($(DEMO_FOLDER)) already exists"
	cp -r ./test $(DEMO_FOLDER)
	cd $(DEMO_FOLDER); vhs < ./test/demo.tape
	cp $(DEMO_FOLDER)/demo.gif ./img/demo.gif

run: ## Runs the prod version of the application
	$(MAKE) command NAME="marjamis"

dev: ## Runs a dev version of the application
	$(MAKE) command NAME="devTester"
	go run main.go --config ./test/demo.aux.yml

clean: ## Cleans up any old/unneeded items
	# - means ignore this line on failure, such as if the file doesn't exist.
	-rm exampleFile
	-rm -r $(DEMO_FOLDER)
