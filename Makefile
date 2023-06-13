HOSTNAME=ozone.com
NAMESPACE=com
NAME=ozone
BINARY=terraform-provider-${NAME}
VERSION=1.0.0
OS_ARCH=darwin_amd64

build-terraform:
	go build -o bin/${BINARY} cmd/ozone-plugin/main.go

terraform: build-terraform
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	cp bin/${BINARY} ${GOBIN}
	mv bin/${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}