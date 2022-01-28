.PHONY: cmd clean

GOCMD=go
BINARY_NAME=kubectl-img

cmd: $(wildcard ./pkg/kube/*.go ./pkg/mtable/*.go ./cmd/*.go ./*.go)
	$(GOCMD) build -o $(BINARY_NAME) ./main.go

clean:
	rm $(BINARY_NAME)