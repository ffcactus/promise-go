all:
	@echo "Building auth ..."
	go build

clean:
	@rm -f ./auth

start: clean all
	./auth &

.PHONY: all clean start
