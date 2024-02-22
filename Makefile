test:
ifndef CI
	go run github.com/onsi/ginkgo/v2/ginkgo -tags debug -v ./...
else
	go run github.com/onsi/ginkgo/v2/ginkgo -tags debug -r --procs=$(shell nproc) --compilers=$(shell nproc) --randomize-all --randomize-suites --fail-on-pending --keep-going --race --trace --cover --coverprofile=cover.profile --junit-report=test-results/results.xml -v ./...
endif

dev-test:
	watchexec -e go sh -c "clear && make test"

.PHONY: test dev-test
