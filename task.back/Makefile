all:
	@echo "Building task ..."
	go build

clean:
	@rm -f ./task

start: clean all
	./task &

.PHONY: all clean start
