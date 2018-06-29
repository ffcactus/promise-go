all:
	@echo "Building server ..."
	go build

clean:
	@rm -f ./server

image: all
	mv ./server ./package/docker
	docker build -t server.promise ./package/docker

install: all
	mkdir -p /opt/promise/apps/server
	cp -rf ./package/rpm/* /opt/promise/apps/server/

clean-install:
	rm -rf /opt/promise/apps/server


start: clean all
	./server &

.PHONY: all clean start
