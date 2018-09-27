all:
	@echo "Building global ..."
	GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build

clean:
	@rm -f ./global

image: all
	mv ./global ./package/image/fs/
	cp ../package/docker/script/wait-for.sh ./package/image/fs/
	docker build  -t global.promise ./package/image

.PHONY: all clean image

