u := $(if $(update),-u)

export GO111MODULE=on

.PHONY: deps
deps:
	go get $(u) -d
	go mod tidy

pv=$(shell gobump show -r ./cmd/songmu)

crossbuild:
	goxz -pv=v$(pv) -d=./dist/v$(pv) ./cmd/songmu

release:
	_tools/releng
	_tools/upload_artifacts
