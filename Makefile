# Commnads declare
GOCMD=go
GOTEST=$(GOCMD) test
GOBUILD=$(GOCMD) build

# Params define
MAIN_PATH=mainC
PACKAGE_PATH=package
PACKAGE_BIN_PATH=package/bin
BIN=auth-server
FILENAME=auth-server.tar.gz

# Deploy Params
DEV_HOST=zy-dev
DEV_TAR_PATH=/home/worker/project/auth-server

PROD_HOST=zy-pro2
PROD_TAR_PATH=/home/worker/project/auth-server

default: clean build pack-dev

test: 
	- $(GOTEST) ./... -v

build: 
	# building
	mkdir $(PACKAGE_PATH)
	mkdir $(PACKAGE_BIN_PATH)
	cd $(MAIN_PATH) && CGO_ENABLE=false GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BIN) 
	mv "$(MAIN_PATH)/$(BIN)" $(PACKAGE_BIN_PATH)
	cp -r "configs" $(PACKAGE_PATH)
	cp "sh/start.sh" $(PACKAGE_BIN_PATH)

pack-dev:
	# packing dev
	cp "$(PACKAGE_PATH)/configs/config.dev.json" "$(PACKAGE_PATH)/configs/config.json"
	cd $(PACKAGE_PATH) && tar -zcvf $(FILENAME) ./*

pack-prod:
	# packing prod
	cp "$(PACKAGE_PATH)/configs/config.prod.json" "$(PACKAGE_PATH)/configs/config.json"
	cd $(PACKAGE_PATH) && tar -zcvf $(FILENAME) ./*

deploy: clean build pack-dev
	# deploy dev from dev
	cp $(PACKAGE_PATH)/$(FILENAME) $(DEV_TAR_PATH)
	cd $(DEV_TAR_PATH) && tar zxvf $(FILENAME) && supervisorctl -c configs/dev.supervisord.conf restart auth-server

deploy-dev: clean build pack-dev
	# deploy-dev from CI
	scp $(PACKAGE_PATH)/$(FILENAME) $(DEV_HOST):$(DEV_TAR_PATH)
	ssh $(DEV_HOST) "cd $(DEV_TAR_PATH) && tar zxvf $(FILENAME) && supervisorctl -c configs/dev.supervisord.conf restart auth-server"

deploy-prod: clean build pack-prod
	# deploying prod from dev or CI
	scp $(PACKAGE_PATH)/$(FILENAME) $(PROD_HOST):$(PROD_TAR_PATH)
	ssh $(PROD_HOST) "cd $(PROD_TAR_PATH) && tar zxvf $(FILENAME) && supervisorctl -c configs/prod.supervisord.conf restart auth-server"

clean:
	# cleaning
	rm -fr $(PACKAGE_PATH)
	rm -fr $(FILENAME)
