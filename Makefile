run: build
	@./go-akn

build:
	@go build -o go-akn .

test:
	@cd  internal/state-assembly && gotest -v -run ^TestDebateProcessPages

testdpsp:
	@cd  internal/state-assembly && gotest -v -run ^TestDebateProcessSinglePage

testdap:
	@cd  internal/state-assembly && gotest -v -run ^TestDebateAnalyzer_Process

tested:
	@cd  internal/parliament && gotest -v -run ^Test_extractDebaters

testda:
	@cd  internal/parliament && gotest -v -run ^TestDebateAnalyzer_Process$

testall:
	@cd  internal/parliament && gotest -v -run ^Test

testpar:
	@cd  internal/parliament && gotest -v -run ^Test_extractSectionMarkers$

testsa:
	@cd  internal/state-assembly && gotest -v -run ^Test_extractSectionMarkers$
