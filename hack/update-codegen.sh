#!/usr/bin/env bash

# Copyright 2018 The Knative Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

if [ -z "${GOPATH:-}" ]; then
  export GOPATH=$(go env GOPATH)
fi

source $(dirname $0)/../vendor/github.com/knative/test-infra/scripts/library.sh

if [ -e ./vendor/k8s.io/code-generator ]; then
    chmod +x ./vendor/k8s.io/code-generator/*.sh
fi

CODEGEN_PKG=${CODEGEN_PKG:-$(cd ${REPO_ROOT_DIR}; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

KNATIVE_CODEGEN_PKG=${KNATIVE_CODEGEN_PKG:-$(cd ${REPO_ROOT_DIR}; ls -d -1 ./vendor/github.com/knative/pkg 2>/dev/null || echo ../pkg)}

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.
${CODEGEN_PKG}/generate-groups.sh "deepcopy,client,informer,lister" \
  github.com/knative/client/pkg/client github.com/knative/client/pkg/apis \
  "client:v1alpha1" \
  --go-header-file ${REPO_ROOT_DIR}/hack/boilerplate/boilerplate.go.txt

if [ ! ${GOPATH}/src/github.com/knative/client/pkg/client -ef ${REPO_ROOT_DIR}/pkg/client ]
then
    cp -r ${GOPATH}/src/github.com/knative/client/pkg/client ${REPO_ROOT_DIR}/pkg/client
    cp ${GOPATH}/src/github.com/knative/client/pkg/apis/client/v1alpha1/zz_generated.deepcopy.go ${REPO_ROOT_DIR}/pkg/apis/client/v1alpha1/
fi
