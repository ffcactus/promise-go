all:
	@echo "Building pool ..."
	GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build

clean:
	@rm -f ./pool

image: all
	mv ./pool ./package/image/fs/
	docker build --no-cache -t pool.promise ./package/image

.PHONY: all clean start image
