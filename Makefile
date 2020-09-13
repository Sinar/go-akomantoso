run: build
	@./go-akn

build:
	@go build -o go-akn .

test:
	@cd  internal/parliament && gotest -v -run ^TestDebateProcessSinglePage

testpardap:
	@cd  internal/parliament && gotest -v -run ^TestDebateAnalyzer_Process

testdpp:
	@cd  internal/state-assembly && gotest -v -run ^TestDebateProcessPages

testdpsp:
	@cd  internal/state-assembly && gotest -v -run ^TestDebateProcessSinglePage

testdap:
	@cd  internal/state-assembly && gotest -v -run ^TestDebateAnalyzer_Process

testpared:
	@cd  internal/parliament && gotest -v -run ^Test_extractDebaters

testall:
	@cd  internal/parliament && gotest -v -run ^Test

testpar:
	@cd  internal/parliament && gotest -v -run ^Test_extractSectionMarkers$

testsa:
	@cd  internal/state-assembly && gotest -v -run ^Test_extractSectionMarkers$
