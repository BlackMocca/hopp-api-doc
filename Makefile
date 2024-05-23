tag=lastest

build-prod:
	docker build -f Dockerfile-production -t hopp-api-doc:$(tag) .