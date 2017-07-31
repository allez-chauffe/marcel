GOCMD = go
GOCLEAN = $(GOCMD) clean
GOFMT = $(GOCMD) fmt
GOTEST = $(GOCMD) test
GOSWAGGER = $(GOCMD) generate
GOBUILD = $(GOCMD) build -o dist/backend

all: test
	$(GOCLEAN)
	$(GOFMT)
	$(GOSWAGGER)
	$(GOBUILD)

test:
	$(GOTEST) -v -covermode=count ./...

build: test
	$(GOBUILD)

doc:
	$(GOSWAGGER)

clean:
	$(GOCLEAN)