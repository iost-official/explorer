docker-build:
	@docker build . -f ./build/backend/api/Dockerfile -t asia.gcr.io/ginco-registry/iost-explorer-backend-api:latest
	@docker build . -f ./build/backend/task/Dockerfile -t asia.gcr.io/ginco-registry/iost-explorer-backend-task:latest
	@docker build . -f ./build/frontend/Dockerfile -t asia.gcr.io/ginco-registry/iost-explorer-frontend:latest

docker-push:
	@docker push asia.gcr.io/ginco-registry/iost-explorer-backend-api:latest
	@docker push asia.gcr.io/ginco-registry/iost-explorer-backend-task:latest
	@docker push asia.gcr.io/ginco-registry/iost-explorer-frontend:latest
