HOSTNAME=ozone.com
NAMESPACE=com
NAME=ozone
VERSION=1.0.1
BINARY=terraform-provider-${NAME}
OS_ARCH=darwin_amd64

build-terraform:
	go build -o bin/${BINARY} cmd/ozone-plugin/main.go

terraform: build-terraform
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	cp bin/${BINARY} ${GOBIN}
	# mv bin/${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
generate-terraform:
	swagger generate client -f ./swagger/swagger.yaml --template-dir templates -C config.yml > swagrun.log
generate-tfdocs:
	tfplugindocs generate --provider-name=ozone --provider-dir=./cmd/ozone-plugin