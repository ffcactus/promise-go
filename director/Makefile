all:
	@echo "Building director ..."
	GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build

clean:
	@rm -f ./director

image: all
	mv ./director ./package/image/fs/
	cp ../package/docker/script/wait-for.sh ./package/image/fs/
	docker build  -t director.promise ./package/image

.PHONY: all clean image

