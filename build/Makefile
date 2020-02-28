BUILDER=./bin/build-all.sh
VERSION?=0.0.1
BINARY_NAME?=my_binary
SERVER_TARGET?=main.go
CLIENT_TARGET?=main.go

build-client:
	@echo $(shell $(BUILDER) $(BINARY_NAME) $(VERSION) $(CLIENT_TARGET))

build-server:
	@echo $(shell $(BUILDER) $(BINARY_NAME) $(VERSION) $(SERVER_TARGET))

clean:
	rm -f $(BIN)/$(BINARY_NAME)