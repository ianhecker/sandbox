RUN=go run
TARGET=main.go

default: run

dep:
	dep ensure -v

update:
	dep ensure -v -update

run: dep
	$(RUN) $(TARGET)

clean:
	-rm -rf ./vendor-new

very-clean: clean
	-rm -rf ./vendor

