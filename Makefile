VERSION=$(shell git describe --tags --abbrev=0 2>/dev/null || echo "")
AST=$(shell find ast -type f | sed 's/ /\\ /g')
EVAL=$(shell find evaluator -type f | sed 's/ /\\ /g')
LEXER=$(shell find lexer -type f | sed 's/ /\\ /g')
OBJECT=$(shell find object -type f | sed 's/ /\\ /g')
PARSER=$(shell find parser -type f | sed 's/ /\\ /g')
TOKEN=$(shell find token -type f | sed 's/ /\\ /g')
TYPING=$(shell find typing -type f | sed 's/ /\\ /g')

all: help

help:	
	@echo "mwnci - Monkey Lang"
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

config: ## Configure and update build files
	@./configure

mwnci: ${AST} ${EVAL} ${LEXER} ${OBJECT} ${PARSER} ${TOKEN} ${TYPING} mwnci.go ## Build mwnci for current architecture
	@echo "Building mwnci"
	CGO_ENABLED=0 go build -ldflags="-X main.version=${VERSION}"

build: ## Configure and compile mwnci
	@make config
	@make mwnci

clean:  ## Clean untracked files
	@echo "Removing editor backup files"
	@find . -name \#\* -type f -exec rm {} \;
	@find . -name \*~ -type f -exec rm {} \;

install: mwnci    ## Install mwnci
	@make -f Mwnci.mk install

distclean: ## Clean untracked files, binaries, and build cache
	@echo "Cleaning build cache"
	@go clean
	@make clean
	@rm -f mwnci  Mwnci.mk includes/Makefile emacs/mwnci.el\
	 evaluator/include.go vim/syntax/mwnci.vim

test: mwnci ## Run unit tests
	@go test -v ./...
