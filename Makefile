REGISTRY ?= obaydullahmhs
IMGAE_NAME ?= api-server
TAG ?= latest

push: build
	docker push ${REGISTRY}/${IMGAE_NAME}:${TAG}

build:
	docker build -t ${REGISTRY}/${IMGAE_NAME}:${TAG} .
	touch build

install: 
	helm install ${IMGAE_NAME} ./helm-charts

uninstall:
	helm uninstall ${IMGAE_NAME}

clean:
	rm -f build