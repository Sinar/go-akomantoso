run: build
	@./go-akn

build:
	@go build -o go-akn .

test:
	@cd  internal/parliament && gotest -v -run ^Test_extractSectionMarkers$

testsa:
	@cd  internal/state-assembly && gotest -v -run ^Test_extractSectionMarkers$
