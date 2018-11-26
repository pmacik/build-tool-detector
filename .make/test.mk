# By default no go test calls will use the -v switch when running tests.
# But if you want you can enable that by setting GO_TEST_VERBOSITY_FLAG=-v
GO_TEST_VERBOSITY_FLAG ?=

.PHONY: test-contracts-consumer
## Runs the consumer part of the contract tests to re-generate the local pact file
test-contracts-consumer:
	$(call log-info,"Running test: $@")
	$(eval TEST_PACKAGES:=$(shell go list ./... | grep 'contracts/consumer'))
	PACT_DIR=$(PWD)/test/contracts/pacts \
	PACT_CONSUMER=Fabric8BuildToolDetectorGenericConsumer \
	PACT_PROVIDER=Fabric8BuildToolDetector \
	PACT_VERSION=1.0.0 \
	go test -count=1 $(GO_TEST_VERBOSITY_FLAG) $(TEST_PACKAGES)

.PHONY: test-contracts-no-coverage
## Runs the contract tests WITHOUT producing coverage files for each package.
## Make sure you ran "make run" before you run this target.
## The Chrome or Chromium browser with headless feature 
## as well as the [chromedriver](http://chromedriver.chromium.org/) is required 
## to be installed for the user login part of the tests.
## The following env variables needs to be set in environment:
## - RHD account credentials:
##   OSIO_USERNAME
##   OSIO_PASSWORD
## - Service account credentials (according to https://github.com/fabric8-services/fabric8-auth/blob/master/configuration/conf-files/service-account-secrets.conf#L30)
##   AUTH_SERVICE_ACCOUNT_CLIENT_ID
##   AUTH_SERVICE_ACCOUNT_CLIENT_SERCRET
test-contracts-no-coverage:
	$(call log-info,"Running test: $@")
	$(eval TEST_PACKAGES:=$(shell go list ./... | grep 'contracts/provider'))
	PACT_DIR=$(PWD)/test/contracts/pacts \
	PACT_CONSUMER=Fabric8BuildToolDetectorGenericConsumer \
	PACT_PROVIDER=Fabric8BuildToolDetector \
	PACT_VERSION=1.0.0 \
	PACT_PROVIDER_BASE_URL=http://localhost:8099 \
	PACT_AUTH_URL=http://auth.openshift.io \
	go test -count=1 $(GO_TEST_VERBOSITY_FLAG) $(TEST_PACKAGES)