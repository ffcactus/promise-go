all: global auth ws task server

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

clean:
	$(MAKE) -C global clean
	$(MAKE) -C auth clean
	$(MAKE) -C task clean
	$(MAKE) -C ws clean
	$(MAKE) -C server clean

start:
	# Make sure the start order is right.
	$(MAKE) -C auth start
	$(MAKE) -C ws start
	$(MAKE) -C task start
	$(MAKE) -C server start

.PHONY: global auth task ws server clean start
