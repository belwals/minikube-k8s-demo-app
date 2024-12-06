TINY_URL_SERVICE_NAME=tiny-url-service

ifeq ($(TINY_URL_SERVICE_TAG_VERSION),)
	TINY_URL_SERVICE_TAG_VERSION=2.0
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

.PHONY: run-service
run-service:
	docker run -d -p 8080:8080 --name tiny-url-service $(TINY_SERVICE_IMAGE_URL)

