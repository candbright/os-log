TARGET_OS=$(os)
OS_WINDOWS=windows
OS_LINUX=linux

ifeq ($(TARGET_OS), $(OS_WINDOWS))
	CLEAN=del os-log
else ifeq ($(TARGET_OS), $(OS_LINUX))
	CLEAN=rm os-log
endif

all: os-log

os-log:
	go build -o $@ cmd/os-log/main.go

clean:
	$(CLEAN)
