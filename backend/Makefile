GO = go

TARGET = explorer

all:$(TARGET)

explorer: main.go
	$(GO) build -mod vendor -o $@ $^

clean:
	rm -rf $(TARGET)
