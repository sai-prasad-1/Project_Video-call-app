#DEV
build-dev:
	docker -t videochat -f containers/images/Dockerfile . && docker build  -t  -f containers/images/Dockerfile.turn . 
run-dev:
	docker-compose -f containers/composes/docker-compose.dev.yml up -d
clean-dev:
	docker-compose -f containers/composes/docker-compose.dev.yml down

