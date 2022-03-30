install:
	go install github.com/swaggo/swag/cmd/swag@latest

generate:
	swag init -ot go --parseDependency

docker-clean:
	docker system prune -a

docker-build:
	docker buildx build --platform linux/amd64 -t kedai-authsvc .
	docker tag kedai-authsvc alganbr/kedai-authsvc
	docker push alganbr/kedai-authsvc

docker-rebuild:
	make docker-clean
	make docker-build