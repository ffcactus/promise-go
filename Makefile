all: director global auth ws task server enclosure pool

director:
	$(MAKE) -C director
global:
	$(MAKE) -C global
auth:
	$(MAKE) -C auth
task:
	$(MAKE) -C task
ws:
	$(MAKE) -C ws
server:
	$(MAKE) -C server
enclosure:
	$(MAKE) -C enclosure
pool:
	$(MAKE) -C pool
 
clean:
	$(MAKE) -C director clean
	$(MAKE) -C global clean
	$(MAKE) -C auth clean
	$(MAKE) -C task clean
	$(MAKE) -C ws clean
	$(MAKE) -C server clean
	$(MAKE) -C enclosure clean
	$(MAKE) -C pool clean

image:
	$(MAKE) -C director image
	$(MAKE) -C platform image
	$(MAKE) -C global image
	$(MAKE) -C auth image
	$(MAKE) -C task image
	$(MAKE) -C ws image
	$(MAKE) -C server image
	$(MAKE) -C enclosure image
	$(MAKE) -C pool image

deploy:
	docker stack deploy -c docker-compose.yml promise

undeploy:
	docker stack rm promise

.PHONY: director global auth task ws server enclosure pool all clean image deploy undeploy
