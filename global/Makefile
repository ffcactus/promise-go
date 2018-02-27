all:
	@echo "Building global ..."
	go build

clean:
	@rm -f ./global

start: clean all
	./global &

.PHONY: all clean start
