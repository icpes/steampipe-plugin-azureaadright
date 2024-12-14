STEAMPIPE_INSTALL_DIR ?= ~/.steampipe
install:
	go build -o $(STEAMPIPE_INSTALL_DIR)/plugins/hub.steampipe.io/plugins/icpes/azureaadright@latest/steampipe-plugin-azureaadright.plugin *.go
