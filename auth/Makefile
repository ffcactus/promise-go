all:
	@echo "Building auth ..."
	GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build

clean:
	@rm -f ./auth

image: all
	mv ./auth ./package/image/fs/
	cp ../package/docker/script/wait-for.sh ./package/image/fs/
	docker build  -t auth.promise ./package/image

.PHONY: all clean start image

