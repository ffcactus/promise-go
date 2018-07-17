all:
	@echo "Building server ..."
	cd producer && GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build
	cd consumer && GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build

clean:
	@rm -f ./mq

image: all
	mv ./producer/producer ./package/image/fs/
	mv ./consumer/consumer ./package/image/fs/
	docker build -t mq.promise ./package/image

.PHONY: all clean start image

