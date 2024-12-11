TINY_URL_SERVICE_NAME=tiny-url-service

ifeq ($(TINY_URL_SERVICE_TAG_VERSION),)
	TINY_URL_SERVICE_TAG_VERSION=9.0
endif

ifeq ($(IMAGE_REPO_USERNAME),)
	IMAGE_REPO_USERNAME=saurabhbelwal01
endif

TINY_SERVICE_IMAGE_URL=$(IMAGE_REPO_USERNAME)/$(TINY_URL_SERVICE_NAME):$(TINY_URL_SERVICE_TAG_VERSION)

# build image for service uses file based buld
.PHONY: build-service
build-service:
	docker build -f ./iac/Service.Dockerfile . -t $(TINY_SERVICE_IMAGE_URL)

# build image for service uses file based buld
.PHONY: push-service
 push-service:
	docker push $(TINY_SERVICE_IMAGE_URL)


# load local image to minikube, so that minikube can leverage the local build image
.PHONY: push-minikube
 push-minikube:
	minikube image load $(TINY_SERVICE_IMAGE_URL)


.PHONY: local-mongoSetup
local-mongoSetup:
	docker run -d  --name mongo \
	-e MONGO_INITDB_ROOT_USERNAME=mongouser \
	-e MONGO_INITDB_ROOT_PASSWORD=mongopassword \
	-e MONGO_INITDB_DATABASE=tiny-url \
	-p 27017:27017 \
	mongo:5.0

# Run the docker image in local
.PHONY: run-service
run-service:
	docker run -d -p 8080:8080 --name tiny-url-service $(TINY_SERVICE_IMAGE_URL)

