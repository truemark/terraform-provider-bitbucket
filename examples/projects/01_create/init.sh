#!/bin/bash
source ../../_setup/setup_env.sh

rm -rf .terraform*

export TF_TEMP_LOG_PATH="./tmp-logs"

                                                        # TerraForm Related environment settings
export TF_LOG=TRACE                                     #    - Setup our log-level to TRACE - gives maximum output
export TF_LOG_PATH="./tf-logfile.log"

echo $TF_TEMP_LOG_PATH
echo $TF_LOG
echo $TF_LOG_PATH



~/go/bin/terraform init -plugin-dir ~/.terraform.d/plugins