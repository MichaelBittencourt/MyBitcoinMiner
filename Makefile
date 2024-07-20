################## General Variables ##################

VERSION=v0.0.1

################## Build Variables ####################

BIN=MyBitcoinMiner
GO_C=go

SRC=main.go client.go
LD_FLAGS:=-X 'main.Version=$(VERSION)'

################# Build Targets #######################

$(BIN): $(SRC)
	$(GO_C) build -ldflags="$(LD_FLAGS)"

################ PHONY Targets ##################

.PHONY: clean run test

clean:
	rm $(BIN)

run: $(BIN)
	./$(BIN)

test: $(BIN)
	@ echo "Testing option \"version\""
	./$(BIN) version
	@ echo
	@ echo "Testing option \"-v\""
	./$(BIN) -v
	@ echo
	@ echo "Testing option \"help\""
	./$(BIN) help
	@ echo
	@ echo "Testing option \"-h\""
	./$(BIN) -h
	@ echo
	@ echo "Testing a invalid option:"
	./$(BIN) invalid && false || true
	
