PKGS := $(shell go list ./...)
TXT_FILES := $(shell find * -type f -not -path 'vendor/**')
TESTFLAG=-race -cover

test:
	GOCACHE=off go test $(TESTFLAG) $(PKGS)

test-verbose:
	go test -v $(TESTFLAG) $(PKGS)

check: vet lint misspell staticcheck

lint:
	@echo "golint"
	golint -set_exit_status $(PKGS)

vet:
	@echo "vet"
	go vet $(PKGS)

misspell:
	@echo "misspell"
	misspell -source=text -error $(TXT_FILES)

staticcheck:
	@echo "staticcheck"
	staticcheck $(PKGS)

prepare:
	mkdir -p run

build: prepare
	go build -o run/holdem github.com/funpoker/holdem/cmd/holdem/

run: build
	./run/holdem
