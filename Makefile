crossbuild:
	goxz -pv=v$(shell gobump show -r) -d=./dist/v$(shell gobump show -r) ./cmd/songmu
