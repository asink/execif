WREN_DIR=src/wren
BUILD_DIR=build
DEPS_FILE=vendor/.deps
APPLICATION=execif
INSTALL_PREFIX=/usr/local

all: execif

execif: deps
	mkdir -p build
	go build -o ${BUILD_DIR}/${APPLICATION} -v

# Update and pull submodules
deps: vendor/
	@if [ -a ${DEPS_FILE} ] ; \
	then \
		echo "Skipping dependencies - run 'make clean' first if needed" ; \
	else \
		echo "Updating dependencies..." ; \
		git submodule init ; \
		git submodule update ; \
		git submodule foreach git pull origin master ; \
		touch ${DEPS_FILE} ; \
	fi;

clean:
	rm -f ${BUILD_DIR}/${APPLICATION}
	rm -f ${DEPS_FILE}

.PHONY: install
install: ${BUILD_DIR}/${APPLICATION}
	@install -m 0755 $^ ${INSTALL_PREFIX}/bin
	@echo "Wrengo has been successfully installed!"
	@echo "Execute 'execif --version' to verify."