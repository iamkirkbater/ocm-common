#
# Copyright (c) 2023 Red Hat, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

# Ensure go modules are enabled:
export GO111MODULE=on
export GOPROXY=https://proxy.golang.org

# Constants
GOPATH := $(shell go env GOPATH)

# Version
revision:=$(shell git rev-parse --short HEAD)
build_time:=$(shell date +%D@%T)
version_stamp:=$(revision)-$(build_time)
import_path:=github.com/openshift-online/ocm-support-cli
ldflags:=-X $(import_path)/pkg/info.VersionStamp=$(version_stamp)

# Disable CGO so that we always generate static binaries:
export CGO_ENABLED=0

# Unset GOFLAG for CI and ensure we've got nothing accidently set
unexport GOFLAGS

# Builds the CLI tool located in /cmd/
.PHONY: build-cli
build-cli: clean
	go build -o rosa-helper -ldflags="$(ldflags)" ./cmd/rosa-helper || exit 1

# Installs the CLI tool located in /cmd/ to your bin folder for easy execution
.PHONY: install-cli
install-cli: clean
	go build -o $(GOPATH)/bin/rosa-helper -ldflags="$(ldflags)" ./cmd/rosa-helper || exit 1

.PHONY: test
test:
	go test ./...

.PHONY: coverage
coverage:
	go test -coverprofile=cover.out  ./...

.PHONY: fmt
fmt:
	gofmt -s -l -w cmd pkg

.PHONY: lint
lint:
	golangci-lint run --timeout 5m0s

.PHONY: clean
clean:
	rm -rf \
		ocm-common \
		$(NULL)
