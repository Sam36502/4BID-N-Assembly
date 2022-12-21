
BIN = 4nsm
VERSION = 0.1.0
BIN_DIR = bin/lin
SRC_DIR = src

RELEASE = 4BID-N-Assembler_v$(VERSION)_Linux.zip
REL_DIR = releases

run: build
	@echo '---> Running...'
	@echo ''
	@$(BIN_DIR)/$(BIN)

build:
	@echo '---> Building binary...'
	go build -o $(BIN_DIR)/$(BIN) $(SRC_DIR)/*.go

release: build
	@echo '---> Creating release...'
	cp $(BIN_DIR)/$(BIN) .
	zip -r -o $(RELEASE) \
		$(BIN) README.md examples
	mv $(RELEASE) $(REL_DIR)/$(RELEASE)
	rm $(BIN)
	@echo '---> Done'

install: build
	@echo '---> Installing...'
	ln -s $(PWD)/$(BIN_DIR)/$(BIN) /usr/bin/4nsm
