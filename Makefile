RUN=go run
PKG?=
TARGET=main.go

default: run

run:
	$(RUN) ./$(PKG)/$(TARGET)