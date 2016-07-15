LDFLAGS="-X \"main.FullVersion=$(COMPILE_VERSION)\""

all: clean compile test

binaries:
	@echo "==> Building binaries $(COMPILE_VERSION)"
	@rm -fr ./pkg
	@for os in linux darwin; do \
		echo "----> Building $$os binary"; \
		mkdir -p ./pkg/$$os; \
		CGO_ENABLED=0 GOOS=$$os GOARCH=amd64 go build -ldflags $(LDFLAGS) -o ./pkg/$$os/assumer ./assumer; \
	done

clean:
	@echo "==> Cleaning up previous builds."
	@rm -rf bin/assumer

compile:
	@echo "==> Compiling source code."
	@CGO_ENABLED=0 go build -v -o ./bin/assumer ./assumer
	@find ./vendor -type d -name .git | xargs rm -rf

coverage:
	@go test -coverprofile cover.out ./assumer/...
	@go tool cover -html=cover.out 

deps:
	@echo "==> Update dependencies."
	@godep save ./assumer/...
	@find ./vendor -type d -name .git | xargs rm -rf

fmt:
	@echo "==> Formatting source code."
	@gofmt -w ./assumer

test: fmt vet
	@echo "==> Running tests."
	@go test -cover ./assumer/...
	@echo "==> Tests complete."

vet:
	@go vet ./assumer/...

help:
	@echo "compile\t\tbuild the code"
	@echo "clean\t\tremove previous builds"
	@echo "deps\t\tdownload dependencies"
	@echo "fmt\t\tformat the code"
	@echo "test\t\ttest the code"
	@echo "vet\t\tvet the code"
	@echo "coverage\tshow detailed code coverage in your browser"
	@echo "godoc\t\tshow this libraries godoc in your browser"
	@echo ""
	@echo "default will test, format, and compile the code"
