all:
	@echo "Building server ..."
	go build

clean:
	@rm -f ./server

install: all
	mkdir -p /opt/promise/apps/server
	cp -rf ./package/* /opt/promise/apps/server/

clean-install:
	rm -rf /opt/promise/apps/server


start: clean all
	./server &

.PHONY: all clean start
