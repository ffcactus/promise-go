all:
	@echo "Building server ..."
	go build

clean:
	@rm -f ./server

start: clean all
	./server &

.PHONY: all clean start
