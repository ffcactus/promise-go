all:
	@echo "Building student ..."
	go build

clean:
	@rm -f ./student

start: clean all
	./student &

.PHONY: all clean start
