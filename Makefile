tag=lastest
templates=./hopp-cli/templates

build-prod:
	docker build -f Dockerfile-production -t hopp-api-doc:$(tag) .

# generate json to html
generate:
	./hopp-cli/cli gen --output $(output) $(templates)

# serve html to develop with docsify
serve:
	docsify serve $(path)
