all:
	@echo "Building auth ..."
	GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build

clean:
	@rm -f ./auth

image: all
	mv ./auth ./package/image/fs/
	docker build --no-cache -t auth.promise ./package/image

.PHONY: all clean start image

