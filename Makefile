.PHONY: all backend-go backend-node

all: backend-go backend-node

backend-go:
	$(MAKE) -C backend-go

backend-node:
	$(MAKE) -C backend-node
