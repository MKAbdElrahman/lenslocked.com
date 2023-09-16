COVERAGE_THRESHOLD := 70

GO_VERSION := 1.21
INSTALL_DIR := /usr/local


COVERAGE_DIR := ./reports
BIN_DIR := ./bin


.PHONY: install-go init-go test coverage report build clean

install-go:
	wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz"
	sudo tar -C $(INSTALL_DIR) -xzf go$(GO_VERSION).linux-amd64.tar.gz
	rm go$(GO_VERSION).linux-amd64.tar.gz

init-go:
	echo 'export PATH=$$PATH:$(INSTALL_DIR)/go/bin' >> ${HOME}/.bashrc
	echo 'export PATH=$$PATH:$${HOME}/go/bin' >> ${HOME}/.bashrc

test: $(COVERAGE_DIR)
	go test ./... -coverprofile=$(COVERAGE_DIR)/coverage.out

coverage: test
	go tool cover -func $(COVERAGE_DIR)/coverage.out | grep "total:" | \
	awk '{print ((int($$3) > $(COVERAGE_THRESHOLD)) != 1) }'

report: test
	go tool cover -html=$(COVERAGE_DIR)/coverage.out -o $(COVERAGE_DIR)/coverage.html
	rm $(COVERAGE_DIR)/coverage.out

build: $(BIN_DIR)
	go build -o $(BIN_DIR)/lenslocked ./main/main.go

clean:
	rm -rf $(BIN_DIR)
	rm -rf $(COVERAGE_DIR)

$(COVERAGE_DIR):
	mkdir -p $(COVERAGE_DIR)

$(BIN_DIR):
	mkdir -p $(BIN_DIR)
