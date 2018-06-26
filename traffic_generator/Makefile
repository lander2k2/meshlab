default: release

release: build-image push-image

build-image:
	test -n "$(IMAGE_REPO)"  # must set $$IMAGE_REPO env var
	test -n "$(IMAGE_TAG)"  # must set $$IMAGE_TAG env var
	docker build --tag ${IMAGE_REPO}:${IMAGE_TAG} .

push-image:
	test -n "$(IMAGE_REPO)"  # must set $$IMAGE_REPO env var
	test -n "$(IMAGE_TAG)"  # must set $$IMAGE_TAG env var
	docker push ${IMAGE_REPO}:${IMAGE_TAG}

