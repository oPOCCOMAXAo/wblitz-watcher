.PHONY: build test generate build-docker upload-docker test-docker prod

generate:
	go generate ./...

build-docker:
	docker build -t poccomaxa/wbwatcher:latest .

upload-docker:
	docker push poccomaxa/wbwatcher:latest

update-prod:
	docker tag poccomaxa/wbwatcher:latest poccomaxa/wbwatcher:prod
	docker push poccomaxa/wbwatcher:prod

test-docker:
	docker run -it --rm poccomaxa/wbwatcher:latest

prod: build-docker upload-docker update-prod
