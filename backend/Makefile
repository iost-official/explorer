GO = go

TARGET = explorer

all:$(TARGET)

explorer: main.go
	$(GO) build -mod vendor -o $@ $^

clean:
	rm -rf $(TARGET)

docker-build:
	@docker build . -f ./build/server/Dockerfile -t asia.gcr.io/ginco-registry/iost-explorer:latest
	@docker build . -f ./build/task/Dockerfile -t asia.gcr.io/ginco-registry/iost-explorer-task:latest

docker-push:
	@docker push asia.gcr.io/ginco-registry/iost-explorer:latest
	@docker push asia.gcr.io/ginco-registry/iost-explorer-task:latest
