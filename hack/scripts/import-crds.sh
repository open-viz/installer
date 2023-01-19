#!/bin/bash

# Copyright AppsCode Inc. and Contributors.
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

set -eou pipefail

crd_dir=${1:-}

api_repo_url=https://github.com/open-viz/grafana-tools.git
api_repo_tag=${OPEN_VIZ_GRAFANA_TOOLS:-master}

if [ "$#" -ne 1 ]; then
    if [ "${api_repo_tag}" == "master" ]; then
        echo "Error: missing path_to_input_crds_directory"
        echo "Usage: import-crds.sh <path_to_input_crds_directory>"
        exit 1
    fi

    tmp_dir=$(mktemp -d -t api-XXXXXXXXXX)
    # always cleanup temp dir
    # ref: https://opensource.com/article/20/6/bash-trap
    trap \
        "{ rm -rf "${tmp_dir}"; }" \
        SIGINT SIGTERM ERR EXIT

    mkdir -p ${tmp_dir}
    pushd $tmp_dir
    git clone $api_repo_url
    repo_dir=$(ls -b1)
    cd $repo_dir
    git checkout $api_repo_tag
    popd
    crd_dir=${tmp_dir}/${repo_dir}/crds
fi

crd-importer \
    --input=${crd_dir} \
    --out=. --output-yaml=crds/openviz-crds.yaml \
    --gk=GrafanaDashboard.openviz.dev \
    --gk=GrafanaDatasource.openviz.dev

crd-importer \
    --input=${crd_dir} \
    --out=./charts/grafana-operator/crds \
    --gk=GrafanaDashboard.openviz.dev \
    --gk=GrafanaDatasource.openviz.dev

crd-importer \
    --input=https://github.com/kmodules/custom-resources/raw/v0.25.1/crds/appcatalog.appscode.com_appbindings.yaml \
    --out=./charts/grafana-operator/crds

crd-importer \
    --input=${crd_dir} \
    --out=./charts/grafana-ui-server/crds \
    --gk=GrafanaDashboard.openviz.dev

crd-importer \
    --input=${crd_dir} \
    --out=./charts/kube-grafana-dashboards/crds \
    --gk=GrafanaDashboard.openviz.dev

crd-importer \
    --input=https://github.com/open-viz/trickster-config/raw/master/config/crd/bases/trickstercache.org_backends.yaml \
    --input=https://github.com/open-viz/trickster-config/raw/master/config/crd/bases/trickstercache.org_caches.yaml \
    --input=https://github.com/open-viz/trickster-config/raw/master/config/crd/bases/trickstercache.org_requestrewriters.yaml \
    --input=https://github.com/open-viz/trickster-config/raw/master/config/crd/bases/trickstercache.org_rules.yaml \
    --input=https://github.com/open-viz/trickster-config/raw/master/config/crd/bases/trickstercache.org_tracingconfigs.yaml \
    --input=https://github.com/open-viz/trickster-config/raw/master/config/crd/bases/trickstercache.org_tricksters.yaml \
    --out=./charts/trickster/crds
