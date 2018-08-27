GO = go

TARGET = explorer

all:$(TARGET)

explorer: main.go
	$(GO) build -o $@ $^

clean:
	rm -rf $(TARGET)