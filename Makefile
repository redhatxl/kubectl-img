.PHONY: cmd clean

GOCMD=go
BINARY_NAME=scc_linux_agent

cmd: $(wildcard ./pkg/akcheck/*.go ./pkg/config/*.go ./pkg/executor/*.go ./pkg/health/*.go ./pkg/http/*.go ./pkg/pidfile/*.go ./pkg/scripts/*.go ./pkg/sshclient/*.go ./pkg/tasks/*.go ./pkg/utils/*.go ./pkg/tasks/gettask/*.go ./pkg/tasks/imagecreate/*.go ./pkg/tasks/imageupload/*.go ./pkg/tasks/migration/*.go ./pkg/tasks/srccheck/*.go ./pkg/tasks/sysrepair/*.go ./pkg/tasks/taskinit/*.go ./pkg/tasks/tasks.go/*.go ./pkg/tasks/taskstatus/*.go ./pkg/tasks/types.go/*.go ./cmd/app/*.go ./cmd/common/*.go ./*.go)
	$(GOCMD) build -o $(BINARY_NAME) ./main.go

clean:
	rm $(BINARY_NAME)