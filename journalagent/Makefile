CONFIG_FILE = config.txt
KEYS_FILE = keys.txt

include $(CONFIG_FILE)
include $(KEYS_FILE)

EXECUTABLE = build/journalagent.exe
DEPS = $(shell find . -name '*.go' -not -path './vendor/*')
SRC = .
MODULE = github.com/axdeyoung/ed-wing-mission-helper/journalagent

LDFLAGS = \
	-X \"$(MODULE)/config.Version=$(VERSION)\" \
	-X \"$(MODULE)/config.Server_addr=$(SERVER_ADDR)\" \
	-X \"$(MODULE)/config.Server_port=$(SERVER_PORT)\" \
	-X \"$(MODULE)/config.Default_journal_dir_from_home=$(DEFAULT_JOURNAL_DIR_FROM_HOME)\" \
	-X \"$(MODULE)/config.Server_key=$(SERVER_KEY)\"

all: build

refresh: clean build

build: $(DEPS)
	@echo "Building journal agent version ${VERSION}"
	@mkdir -p build
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(EXECUTABLE) $(SRC)
clean:
	go clean
	@rm -rf build