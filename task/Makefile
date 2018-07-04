all:
	@echo "Building task ..."
	GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build

clean:
	@rm -f ./task

image: all
	mv ./task ./package/image/fs/
	docker build  -t task.promise ./package/image

.PHONY: all clean start image
