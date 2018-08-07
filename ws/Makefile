all:
	@echo "Building ws ..."
	GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build

clean:
	@rm -f ./ws

image: all
	mv ./ws ./package/image/fs/
	cp ../package/docker/script/wait-for.sh ./package/image/fs/
	docker build  -t ws.promise ./package/image

.PHONY: all clean start image