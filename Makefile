VERSION = $$(cat VERSION)
IMG_NAME = gpaggi/dot_monitor

.PHONY: build build-docker clean distribute clean-image fmt

build:
	@CGO_ENABLED=0 go build -v -ldflags "-X github.com/gpaggi/dot_monitor/version.Version=$(VERSION) -s -w" -o dot_monitor .

clean:
	@rm dot_monitor

build-docker:
	@docker build --build-arg VERSION=$(VERSION) -t $(IMG_NAME):$(VERSION) .
	@docker tag $(IMG_NAME):$(VERSION) $(IMG_NAME):latest

distribute: build-docker
	@docker push $(IMG_NAME):$(VERSION)
	@docker push $(IMG_NAME):latest

clean-image:
	@docker image rm $(IMG_NAME):$(VERSION)

fmt:
	@go mod tidy
	@find . -type f -name '*.go' | xargs gofmt -w -s