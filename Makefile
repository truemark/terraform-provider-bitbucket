.PHONY: clean setup build install go-format go-lint go-doc tf-doc

###
# Provider Variables
PROVIDER_PLUGINS_DIR = ~/.terraform.d/plugins
PROVIDER_HOSTNAME = truemark.io
PROVIDER_NAMESPACE = terraform
PROVIDER_NAME = truemark-bitbucket
PROVIDER_PROD_NAME = terraform-provider-$(PROVIDER_NAME)
PROVIDER_VERSION = 1.0.0
PROVIDER_OS_ARCH = darwin_amd64
PROVIDER_BUILT_FILENAME = $(PROVIDER_PROD_NAME)_$(PROVIDER_VERSION)_$(PROVIDER_ARCHOS)

### 
# Setup Sequence - Dependency Libs
DEP_BITBUCKET_ASSERT = github.com/stretchr/testify/assert
DEP_BITBUCKET_OAUTH = golang.org/x/oauth2
DEP_BITBUCKET_CONTEXT = golang.org/x/net/context
DEP_TFPLUGIN_LOGGING = github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging@v2.5.0
DEP_TFPLUGIN_RESOURCE = github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource@v2.5.0
DEP_TFPLUGIN_SCHEMA = github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema@v2.5.0
DEP_TFPLUGIN_ADDRS = github.com/hashicorp/terraform-plugin-sdk/v2/internal/addrs@v2.5.0
DEP_TFPLUGIN_HCL2SHIM = github.com/hashicorp/terraform-plugin-sdk/v2/internal/configs/hcl2shim@v2.5.0
DEP_TFPLUGIN_PLUGINTEST = github.com/hashicorp/terraform-plugin-sdk/v2/internal/plugintest@v2.5.0
DEP_TFPLUGIN_META = github.com/hashicorp/terraform-plugin-sdk/v2/meta@v2.5.0
DEP_TFPLUGIN_PLUGIN = github.com/hashicorp/terraform-plugin-sdk/v2/plugin@v2.5.0
DEP_TFPLUGIN_TF = github.com/hashicorp/terraform-plugin-sdk/v2/terraform@v2.5.0

DEPS_LIST=$(DEP_BITBUCKET_ASSERT) $(DEP_BITBUCKET_OAUTH) $(DEP_BITBUCKET_CONTEXT) $(DEP_TFPLUGIN_LOGGING) $(DEP_TFPLUGIN_RESOURCE) $(DEP_TFPLUGIN_SCHEMA) $(DEP_TFPLUGIN_ADDRS) $(DEP_TFPLUGIN_HCL2SHIM) $(DEP_TFPLUGIN_PLUGINTEST) $(DEP_TFPLUGIN_META) $(DEP_TFPLUGIN_PLUGIN) $(DEP_TFPLUGIN_TF)

###
# Build Variables
BUILD_COMPILER = go
BUILD_COMMAND = build

BUILD_OPTS_OUTDIR = build
BUILD_OPTS_VERBOSE = -v -work -x 
BUILD_OPTS_INSTALL_DEPS = -i 
BUILD_OPTS_REBUILD = -a 
BUILD_OPTS_MODULE = -mod=mod

BUILD_MKDIR = mkdir
BUILD_OPTS = $(BUILD_OPTS_VERBOSE)
BUILD_OUT = ./$(BUILD_OPTS_OUTDIR)/$(PROVIDER_PROD_NAME)


###
# Clean Commands
CLEAN_CMD = rm
CLEAN_OPTS = -rf 

###
# Install Commands
INSTALL_MK_CMD = mkdir
INSTALL_MK_OPTS = -p
INSTALL_MV_CMD = mv
INSTALL_MV_OPTS =
INSTALL_LOCATION = $(PROVIDER_PLUGINS_DIR)/$(PROVIDER_HOSTNAME)/$(PROVIDER_NAMESPACE)/$(PROVIDER_NAME)/$(PROVIDER_VERSION)/$(PROVIDER_OS_ARCH)

# Run this one time to setup the environment to build and develop in.
setup:
	$(foreach element,$(DEPS_LIST), go get -v $(element);) 

clean:
	$(CLEAN_CMD) $(CLEAN_OPTS) $(BUILD_OPTS_OUTDIR)

build: clean
	$(BUILD_MKDIR) ./$(BUILD_OPTS_OUTDIR)
	$(BUILD_COMPILER) $(BUILD_COMMAND) -o $(BUILD_OUT)

install: build
	$(INSTALL_MK_CMD) $(INSTALL_MK_OPTS) $(INSTALL_LOCATION)
	$(INSTALL_MV_CMD) $(INSTALL_MV_OPTS) $(BUILD_OUT) $(INSTALL_LOCATION)

go-format:
	go fmt ./...

go-lint:

go-doc:

tf-doc:
