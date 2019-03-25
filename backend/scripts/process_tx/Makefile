GO = go

TARGET = processTx

all:$(TARGET)

processTx: voteTx.go
	$(GO) build -o $@ $^

clean:
	rm -rf $(TARGET)
