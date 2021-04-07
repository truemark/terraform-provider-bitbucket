#!/bin/bash

###
## from: https://www.terraform.io/docs/cli/config/environment-variables.html
## Sets up our operating environment. 
# TF_LOG
# TF_LOG_PATH
# TF_INPUT
# TF_VAR_name
# TF_DATA_DIR
# TF_WORKSPACE
# TF_IN_AUTOMATION
# TF_REGISTRY_DISCOVERY_RETRY
# TF_REGISTRY_CLIENT_TIMEOUT
# TF_CLI_CONFIG_FILE
# TF_IGNORE

export TF_TEMP_LOG_PATH="./tmp-logs"

                                                        # TerraForm Related environment settings
export TF_LOG=TRACE                                     #    - Setup our log-level to TRACE - gives maximum output
export TF_LOG_PATH="./tf-logfile.log"


                                                        # TrueMark.io Env-Vars for debugging/test/etc.
# export TMK_TF_PLUGINS_DIR="~/.terraform.d/plugins"      #    - where to find local plugins, overriding a search on registry.terraform.com. TODO: make this conditional perhaps?
# export TMK_TF_CLI_CMD=""                                #    - allows for adjusting the terraform binary in case running on a custom terraform cli build


#                                                         # Set Confluent Cloud related Settings here.
# export TRUEMARK_CONFLUENTCLOUD_USERNAME=                #    - confluent.cloud login username
# export TRUEMARK_CONFLUENTCLOUD_PASSWORD=                #    - confluent.cloud login password

     