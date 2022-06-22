build:
	sh ./scripts/build.sh

test: ## run test including unit test and integration test
	bash ./scripts/test.sh
.PHONY: test