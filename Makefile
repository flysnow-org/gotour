getdeps:
	@mkdir -p ${GOPATH}/bin
	@which golangci-lint 1>/dev/null || (echo "Installing golangci-lint" && go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.32.2)

fmt:
	@echo "Running $@ check"
	@GO111MODULE=on gofmt -d ./
lint:
	@echo "Running $@ check"
	@GO111MODULE=on ${GOPATH}/bin/golangci-lint cache clean
	@GO111MODULE=on ${GOPATH}/bin/golangci-lint run --timeout=5m --config ./.golangci.yml

verifiers: getdeps lint