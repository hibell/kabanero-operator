# The Docker image in format repository:tag. Repository may contain a remote reference.
# Override in order to customize
IMAGE ?= kabanero-operator:latest
REGISTRY_IMAGE ?= kabanero-operator-registry:latest
WEBHOOK_IMAGE ?= kabanero-operator-admission-webhook:latest
COLLECTION_CTRLR_IMAGE ?= kabanero-operator-collection-controller:latest


# For integration testing
# INTERNAL_REGISTRY: the public facing registry url. Set TRUE to enable and find the default address, or manually set to address itself
# INTERNAL_REGISTRY_SVC: the internal service image pull address. If not set, set to the default
INTERNAL_REGISTRY ?=
INTERNAL_REGISTRY_SVC ?=
ifdef INTERNAL_REGISTRY
ifeq ($(INTERNAL_REGISTRY),TRUE)
INTERNAL_REGISTRY := $(shell kubectl get route default-route -n openshift-image-registry --template='{{ .spec.host }}')
ifndef INTERNAL_REGISTRY_SVC
INTERNAL_REGISTRY_SVC=image-registry.openshift-image-registry.svc:5000
endif
# Public registry references
IMAGE=${INTERNAL_REGISTRY}/kabanero/kabanero-operator:latest
REGISTRY_IMAGE=${INTERNAL_REGISTRY}/openshift-marketplace/kabanero-operator-registry:latest
WEBHOOK_IMAGE=${INTERNAL_REGISTRY}/kabanero/kabanero-operator-admission-webhook:latest
COLLECTION_CTRLR_IMAGE=${INTERNAL_REGISTRY}/kabanero/kabanero-operator-collection-controller:latest
# Internal service registry references
IMAGE_SVC=${INTERNAL_REGISTRY_SVC}/kabanero/kabanero-operator:latest
REGISTRY_IMAGE_SVC=${INTERNAL_REGISTRY_SVC}/openshift-marketplace/kabanero-operator-registry:latest
WEBHOOK_IMAGE_SVC=${INTERNAL_REGISTRY_SVC}/kabanero/kabanero-operator-admission-webhook:latest
COLLECTION_CTRLR_IMAGE_SVC=${INTERNAL_REGISTRY_SVC}/kabanero/kabanero-operator-collection-controller:latest
endif
endif


# Computed repository name (no tag) including repository host/path reference
REPOSITORY=$(firstword $(subst :, ,${IMAGE}))
REGISTRY_REPOSITORY=$(firstword $(subst :, ,${REGISTRY_IMAGE}))
WEBHOOK_REPOSITORY=$(firstword $(subst :, ,${WEBHOOK_IMAGE}))
COLLECTION_CTRLR_REPOSITORY=$(firstword $(subst :, ,${COLLECTION_CTRLR_IMAGE}))


# Current release (used for CSV management)
CURRENT_RELEASE=0.5.0


.PHONY: build deploy deploy-olm build-image push-image int-test-install int-test-collections int-test-uninstall int-test-lifecycle

build: generate
	go install ./cmd/manager
	go install ./cmd/manager/collection
	go install ./cmd/admission-webhook

build-image: generate
  # These commands were taken from operator-sdk 0.8.1.  The sdk did not let us
  # pass the ldflags option.  The advice from operator-sdk was to run the 
  # commands separately here.
  # operator-sdk build ${IMAGE}
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/_output/bin/kabanero-operator -gcflags "all=-trimpath=$(GOPATH)" -asmflags "all=-trimpath=$(GOPATH)" -ldflags "-X main.GitTag=$(TRAVIS_TAG) -X main.GitCommit=$(TRAVIS_COMMIT) -X main.GitRepoSlug=$(TRAVIS_REPO_SLUG) -X main.BuildDate=`date -u +%Y%m%d.%H%M%S`" github.com/kabanero-io/kabanero-operator/cmd/manager
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/_output/bin/kabanero-operator-collection-controller -gcflags "all=-trimpath=$(GOPATH)" -asmflags "all=-trimpath=$(GOPATH)" -ldflags "-X main.GitTag=$(TRAVIS_TAG) -X main.GitCommit=$(TRAVIS_COMMIT) -X main.GitRepoSlug=$(TRAVIS_REPO_SLUG) -X main.BuildDate=`date -u +%Y%m%d.%H%M%S`" github.com/kabanero-io/kabanero-operator/cmd/manager/collection
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/_output/bin/admission-webhook -gcflags "all=-trimpath=$(GOPATH)" -asmflags "all=-trimpath=$(GOPATH)" -ldflags "-X main.GitTag=$(TRAVIS_TAG) -X main.GitCommit=$(TRAVIS_COMMIT) -X main.GitRepoSlug=$(TRAVIS_REPO_SLUG) -X main.BuildDate=`date -u +%Y%m%d.%H%M%S`" github.com/kabanero-io/kabanero-operator/cmd/admission-webhook

	docker build -f build/Dockerfile -t ${IMAGE} .

  # Build the Kananero collection controller image.
	docker build -f build/Dockerfile-collection-controller -t ${COLLECTION_CTRLR_IMAGE} .

  # Build the admission webhook.
	docker build -f build/Dockerfile-webhook -t ${WEBHOOK_IMAGE} .

  # Build an OLM private registry for Kabanero
	mkdir -p build/registry
	cp LICENSE build/registry/LICENSE
	cp -R registry/manifests build/registry/
	cp registry/Dockerfile build/registry/Dockerfile
	cp deploy/crds/kabanero_kabanero_crd.yaml deploy/crds/kabanero_collection_crd.yaml build/registry/manifests/kabanero-operator/$(CURRENT_RELEASE)/

ifdef INTERNAL_REGISTRY
	sed -e "s!kabanero/kabanero-operator:.*!${IMAGE_SVC}!" registry/manifests/kabanero-operator/$(CURRENT_RELEASE)/kabanero-operator.v$(CURRENT_RELEASE).clusterserviceversion.yaml > build/registry/manifests/kabanero-operator/$(CURRENT_RELEASE)/kabanero-operator.v$(CURRENT_RELEASE).clusterserviceversion.yaml
else
	sed -e "s!kabanero/kabanero-operator:.*!${IMAGE}!" registry/manifests/kabanero-operator/$(CURRENT_RELEASE)/kabanero-operator.v$(CURRENT_RELEASE).clusterserviceversion.yaml > build/registry/manifests/kabanero-operator/$(CURRENT_RELEASE)/kabanero-operator.v$(CURRENT_RELEASE).clusterserviceversion.yaml
endif

	docker build -t ${REGISTRY_IMAGE} -f build/registry/Dockerfile build/registry/

  # If we're doing a Travis build, need to build a second image because the CSV
  # in the registry image has to point to the tagged operator image.
ifdef TRAVIS_TAG
	sed -e "s!kabanero/kabanero-operator:.*!${REPOSITORY}:${TRAVIS_TAG}!" registry/manifests/kabanero-operator/$(CURRENT_RELEASE)/kabanero-operator.v$(CURRENT_RELEASE).clusterserviceversion.yaml > build/registry/manifests/kabanero-operator/$(CURRENT_RELEASE)/kabanero-operator.v$(CURRENT_RELEASE).clusterserviceversion.yaml
	docker build -t ${REGISTRY_REPOSITORY}:${TRAVIS_TAG} -f build/registry/Dockerfile build/registry/
endif

ifdef TRAVIS_BRANCH
	sed -e "s!kabanero/kabanero-operator:.*!${REPOSITORY}:${TRAVIS_BRANCH}!" registry/manifests/kabanero-operator/$(CURRENT_RELEASE)/kabanero-operator.v$(CURRENT_RELEASE).clusterserviceversion.yaml > build/registry/manifests/kabanero-operator/$(CURRENT_RELEASE)/kabanero-operator.v$(CURRENT_RELEASE).clusterserviceversion.yaml
	docker build -t ${REGISTRY_REPOSITORY}:${TRAVIS_BRANCH} -f build/registry/Dockerfile build/registry/
endif

	rm -R build/registry

push-image:
ifneq "$(IMAGE)" "kabanero-operator:latest"
  # Default push.  Make sure the namespace is there in case using local registry
	kubectl create namespace kabanero || true
	docker push $(IMAGE)
	docker push $(COLLECTION_CTRLR_IMAGE)
	docker push $(REGISTRY_IMAGE)
	docker push $(WEBHOOK_IMAGE)

ifdef TRAVIS_TAG
  # This is a Travis tag build. Pushing using Docker tag TRAVIS_TAG
	docker tag $(IMAGE) $(REPOSITORY):$(TRAVIS_TAG)
	docker push $(REPOSITORY):$(TRAVIS_TAG)
	docker tag $(WEBHOOK_IMAGE) $(WEBHOOK_REPOSITORY):$(TRAVIS_TAG)
	docker push $(WEBHOOK_REPOSITORY):$(TRAVIS_TAG)
	docker push $(REGISTRY_REPOSITORY):$(TRAVIS_TAG)
	docker tag $(COLLECTION_CTRLR_IMAGE) $(COLLECTION_CTRLR_REPOSITORY):$(TRAVIS_TAG)
	docker push $(COLLECTION_CTRLR_REPOSITORY):$(TRAVIS_TAG)
endif

ifdef TRAVIS_BRANCH
  # This is a Travis branch build. Pushing using Docker tag TRAVIS_BRANCH
	docker tag $(IMAGE) $(REPOSITORY):$(TRAVIS_BRANCH)
	docker push $(REPOSITORY):$(TRAVIS_BRANCH)
	docker tag $(WEBHOOK_IMAGE) $(WEBHOOK_REPOSITORY):$(TRAVIS_BRANCH)
	docker push $(WEBHOOK_REPOSITORY):$(TRAVIS_BRANCH)
	docker push $(REGISTRY_REPOSITORY):$(TRAVIS_BRANCH)
	docker tag $(COLLECTION_CTRLR_IMAGE) $(COLLECTION_CTRLR_REPOSITORY):$(TRAVIS_BRANCH)
	docker push $(COLLECTION_CTRLR_REPOSITORY):$(TRAVIS_BRANCH)
endif
endif

test: 
	go test ./cmd/... ./pkg/... 

format:
	go fmt ./cmd/... ./pkg/...

generate:
	operator-sdk generate k8s
	operator-sdk generate openapi
	go generate ./pkg/assets

install:
	kubectl config set-context $$(kubectl config current-context) --namespace=kabanero
	kubectl apply -f deploy/crds/kabanero_kabanero_crd.yaml
	kubectl apply -f deploy/crds/kabanero_collection_crd.yaml

deploy: 
	kubectl create namespace kabanero || true

ifneq "$(IMAGE)" "kabanero-operator:latest"
  # By default there is no image pull policy for local image. However, for other images
  # substitute current image name, and update the pull policy to always pull images.
	sed -i.bak -e 's!imagePullPolicy: Never!imagePullPolicy: Always!' deploy/operator.yaml
	sed -i.bak -e 's!image: kabanero-operator:latest!image: ${IMAGE}!' deploy/operator.yaml
endif
	rm deploy/operator.yaml.bak || true
	kubectl config set-context $$(kubectl config current-context) --namespace=kabanero
	kubectl apply -f deploy/

deploy-olm:
	kubectl create namespace kabanero || true

	kubectl apply -f deploy/olm/

# Update deployment to correct image 
ifdef INTERNAL_REGISTRY
	sed -e "s!image: KABANERO_REGISTRY_IMAGE!image: ${REGISTRY_IMAGE_SVC}!" deploy/olm/01-catalog-source.yaml | kubectl apply -f -
else
	sed -e "s!image: KABANERO_REGISTRY_IMAGE!image: ${REGISTRY_IMAGE}!" deploy/olm/01-catalog-source.yaml | kubectl apply -f -
endif

check: format build #test

dependencies: 
ifeq (, $(shell which dep))
	go get -u github.com/golang/dep/cmd/dep
endif
	dep ensure

  # Remove some creative commons licensed tests/samples
	rm vendor/golang.org/x/net/http2/h2demo/tmpl.go
	rm -r vendor/golang.org/x/text/internal/testtext

# Requires https://github.com/pmezard/licenses
dependency-report: 
	go get -u github.com/pmezard/licenses
	licenses ./pkg/... | cut -c49- > 3RD_PARTY


# Integration Tests
# Requires kube login context an existing cluster
# Requires internal registry with default route
# Requires jq

# Install Test
int-test-install: creds build-image push-image int-install int-config

creds:
	tests/00-credentials.sh

int-install:
# Update deployment to correct image 
ifdef INTERNAL_REGISTRY
	sed -e "s!image: kabanero/kabanero-operator-registry:.*!image: ${REGISTRY_IMAGE_SVC}!" deploy/kabanero-subscriptions.yaml > /tmp/kabanero-subscriptions.yaml
else
	sed -e "s!image: kabanero/kabanero-operator-registry:.*!image: ${REGISTRY_IMAGE}!" deploy/kabanero-subscriptions.yaml > /tmp/kabanero-subscriptions.yaml
endif

	KABANERO_SUBSCRIPTIONS_YAML=/tmp/kabanero-subscriptions.yaml KABANERO_CUSTOMRESOURCES_YAML=deploy/kabanero-customresources.yaml deploy/install.sh

int-config:
# Update config to correct image
ifdef INTERNAL_REGISTRY
	sed -e "s!image: kabanero/kabanero-operator-admission-webhook:.*!image: ${WEBHOOK_IMAGE_SVC}!; s!image: kabanero/kabanero-operator-collection-controller:.*!image: ${COLLECTION_CTRLR_IMAGE_SVC}!" config/samples/full.yaml | kubectl apply -f -
else
	sed -e "s!image: kabanero/kabanero-operator-admission-webhook:.*!image: ${WEBHOOK_IMAGE}!; s!image: kabanero/kabanero-operator-collection-controller:.*!image: ${COLLECTION_CTRLR_IMAGE}!" config/samples/full.yaml | kubectl apply -f -
endif

# Uninstall Test
int-test-uninstall: creds int-uninstall

int-uninstall:
	KABANERO_SUBSCRIPTIONS_YAML=deploy/kabanero-subscriptions.yaml KABANERO_CUSTOMRESOURCES_YAML=deploy/kabanero-customresources.yaml deploy/uninstall.sh

# Collections: Can be run in parallel ( -j ). Test manual pipeline run of collections.
int-test-collections: int-test-java-microprofile int-test-java-spring-boot2 int-test-nodejs int-test-nodejs-express int-test-nodejs-loopback

int-test-java-microprofile:
	tests/10-java-microprofile.sh

int-test-java-spring-boot2:
	tests/11-java-spring-boot2.sh

int-test-nodejs:
	tests/12-nodejs.sh

int-test-nodejs-express:
	tests/13-nodejs-express.sh

int-test-nodejs-loopback:
	tests/14-nodejs-loopback.sh

# Lifecycle: Lifecycle of Kabanero, Collection and owned objects
int-test-lifecycle: int-test-delete-pipeline int-test-delete-collection int-test-update-index int-test-deactivate-collection
int-test-delete-pipeline:
	tests/20-delete-pipeline.sh
int-test-delete-collection:
	tests/21-delete-collection.sh
int-test-update-index:
	tests/22-update-index.sh
int-test-deactivate-collection:
	tests/23-deactivate-collection.sh