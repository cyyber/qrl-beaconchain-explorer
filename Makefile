GITCOMMIT=`git describe --always`
VERSION=`git describe --always --tags`
GITDATE=`TZ=UTC git show -s --date=iso-strict-local --format=%cd HEAD`
BUILDDATE=`date -u +"%Y-%m-%dT%H:%M:%S%:z"`
PACKAGE=github.com/theQRL/zond-beaconchain-explorer
LDFLAGS="-X ${PACKAGE}/version.Version=${VERSION} -X ${PACKAGE}/version.BuildDate=${BUILDDATE} -X ${PACKAGE}/version.GitCommit=${GITCOMMIT} -X ${PACKAGE}/version.GitDate=${GITDATE} -s -w"
CGO_CFLAGS=""
CGO_CFLAGS_ALLOW=""

all: explorer stats frontend-data-updater eth1indexer rewards-exporter node-jobs-processor signatures misc

lint:
	echo 
	golint ./...

test:
	go test ./...

explorer:
	rm -rf bin/
	mkdir -p bin/
	go run cmd/bundle/main.go
	go install github.com/swaggo/swag/cmd/swag@v1.8.3 && swag init --exclude bin,_gitignore,.vscode,.idea --parseDepth 1 -g ./handlers/api.go
	CGO_CFLAGS=${CGO_CFLAGS} CGO_CFLAGS_ALLOW=${CGO_CFLAGS_ALLOW} go build --ldflags=${LDFLAGS} -o bin/explorer cmd/explorer/main.go

stats:
	CGO_CFLAGS=${CGO_CFLAGS} CGO_CFLAGS_ALLOW=${CGO_CFLAGS_ALLOW} go build --ldflags=${LDFLAGS} -o bin/statistics cmd/statistics/main.go

frontend-data-updater:
	CGO_CFLAGS=${CGO_CFLAGS} CGO_CFLAGS_ALLOW=${CGO_CFLAGS_ALLOW} go build --ldflags=${LDFLAGS} -o bin/frontend-data-updater cmd/frontend-data-updater/main.go

rewards-exporter:
	CGO_CFLAGS=${CGO_CFLAGS} CGO_CFLAGS_ALLOW=${CGO_CFLAGS_ALLOW} go build --ldflags=${LDFLAGS} -o bin/rewards-exporter cmd/rewards-exporter/main.go

eth1indexer:
	CGO_CFLAGS=${CGO_CFLAGS} CGO_CFLAGS_ALLOW=${CGO_CFLAGS_ALLOW} go build --ldflags=${LDFLAGS} -o bin/eth1indexer cmd/eth1indexer/main.go

node-jobs-processor:
	CGO_CFLAGS=${CGO_CFLAGS} CGO_CFLAGS_ALLOW=${CGO_CFLAGS_ALLOW} go build --ldflags=${LDFLAGS} -o bin/node-jobs-processor cmd/node-jobs-processor/main.go

signatures:
	CGO_CFLAGS=${CGO_CFLAGS} CGO_CFLAGS_ALLOW=${CGO_CFLAGS_ALLOW} go build --ldflags=${LDFLAGS} -o bin/signatures cmd/signatures/main.go

misc:
	CGO_CFLAGS=${CGO_CFLAGS} CGO_CFLAGS_ALLOW=${CGO_CFLAGS_ALLOW} go build --ldflags=${LDFLAGS} -o bin/misc cmd/misc/main.go

addhooks:
	git config core.hooksPath hooks

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative types/ens.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative types/eth1.proto