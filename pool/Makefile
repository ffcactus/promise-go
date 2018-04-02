all:
	@echo "Building pool ..."
	go build

clean:
	@rm -f ./pool

start: clean all
	./pool &

.PHONY: all clean start