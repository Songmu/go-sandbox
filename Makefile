crossbuild:
	goxz -pv=v$(shell gobump show -r) -d=./dist/v$(shell gobump show -r) ./cmd/songmu

release:
	_tools/releng
	_tools/upload_artifacts
