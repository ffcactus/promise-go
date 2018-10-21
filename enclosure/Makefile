all:
	@echo "Building enclosure ..."
	GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build

clean:
	@rm -f ./enclosure

image: all
	mv ./enclosure ./package/image/fs/
	cp ../package/docker/script/wait-for.sh ./package/image/fs/
	docker build  -t enclosure.promise ./package/image

.PHONY: all clean image
