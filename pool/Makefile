all:
	@echo "Building pool ..."
	GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build

clean:
	@rm -f ./pool

image: all
	mv ./pool ./package/image/fs/
	cp ../package/docker/script/wait-for.sh ./package/image/fs/
	docker build  -t pool.promise ./package/image

.PHONY: all clean start image
