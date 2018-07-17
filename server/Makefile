all:
	@echo "Building server ..."
	GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build

clean:
	@rm -f ./server

image: all
	mv ./server ./package/image/fs/
	docker build -t server.promise ./package/image

.PHONY: all clean start image

