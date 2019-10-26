VERSION := $(shell godzil show-version ./cmd/songmu)
u := $(if $(update),-u)

export GO111MODULE=on

.PHONY: deps
deps:
	go get $(u) -d
	go mod tidy

.PHONY: devel-deps
devel-deps:
	GO111MODULE=off go get $(u) \
      golang.org/x/ling/golint \
      github.com/Songmu/godzil/cmd/godzil \
      github.com/Songmu/goxz/cmd/goxz     \
      github.com/tcnksm/ghr

.PHONY: bump
bump: devel-deps
	godzil release

.PHONY: crossbuild
crossbuild:
	goxz -pv=v$(VERSION) -d=./dist/v$(VERSION) ./cmd/songmu

.PHONY: upload
upload:
	ghr v$(VERSION) dist/v$(VERSION)

.PHONY: release
release: crossbuild upload
