#!/bin/bash

# Copyright 2021, 2022 Hewlett Packard Enterprise Development LP
# Other additional copyright holders may be indicated within.
#
# The entirety of this work is licensed under the Apache License,
# Version 2.0 (the "License"); you may not use this file except
# in compliance with the License.
#
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


# Setup for nnf-sos based kind cluster

echo "$(tput bold)Installing prerequisite CRDs $(tput sgr 0)"
kubectl apply -f config/mpi/mpi-operator.yaml
kubectl apply -f vendor/github.com/NearNodeFlash/lustre-fs-operator/config/crd/bases/cray.hpe.com_lustrefilesystems.yaml

echo "$(tput bold)Installing sample LustreFileSystem $(tput sgr 0)"
cat <<-EOF | kubectl apply -f -
apiVersion: cray.hpe.com/v1alpha1
kind: LustreFileSystem
metadata:
  name: lustrefilesystem-sample-maui
  namespace: nnf-dm-system
spec:
  name: maui
  mgsNids:
  - 172.0.0.1@tcp
  mountRoot: /lus/maui
EOF
