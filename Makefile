DATE       := $(shell date +%FT%T%z)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
GIT_COMMIT := $(shell git rev-parse --short HEAD)
GIT_DIRTY  := $(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)

# Additional build variables, will be compiled into output binary
LDFLAGS := $(LDFLAGS) -X main.commit=$(GIT_COMMIT)${GIT_DIRTY} -X main.branch=$(GIT_BRANCH) -X main.date=$(DATE)

WORKSPACE := $(PWD)

# Build and run go-spam in development mode
run: build
	./go-spam dev

# Run tests
test: build
	go test -cover -race ./...

# Build binary using development configuration, used for local dev
build:
	go build -ldflags "-w -s $(LDFLAGS)" -tags dev

# Install binary using production configuration, used for production
install:
	go install -ldflags "-w -s $(LDFLAGS)" -tags prod

# Generate protobpf
proto: dep format proto_service

# Format protobuf code base
format:
	find . -type f -iname *.proto | xargs -I{} clang-format -i {}


# Since protobuf don't support ordered build (AFAIK),
# we list all build steps order manually.
proto_service:
	protoc \
		-I=. \
		-I=$(GOPATH)/src \
		-I=$(GOPATH)/src/github.com/gogo/protobuf/protobuf \
		--gogofaster_out=plugins=grpc,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types:.\
			./external/services/grpc/definition/*.proto

dep: clang_format

clang_format:
	@clang-format --version > /dev/null 2>&1 || (echo "ERROR: clang-format is required (brew install clang-format)"; exit 1)

# Rewrite example
# --gogofaster_out=plugins=grpc,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:.\
