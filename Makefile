# go makefile

program != basename $$(pwd)
go_version = go1.24.5
version != cat VERSION
latest_release != gh release list --json tagName --jq '.[0].tagName' | tr -d v
gitclean = $(if $(shell git status --porcelain),$(error git status is dirty),$(info git status is clean))
rstms_modules = $(shell awk <go.mod '/^module/{next} /rstms/{print $$1}')
$(program): build
logfile = /var/log/$(program)

build: fmt
	fix go build

fmt: go.sum
	fix go fmt . ./...

go.mod:
	$(go_version) mod init

go.sum: go.mod
	go mod tidy

install: build
	go install

test: fmt
	go test -v -failfast . ./...

debug: fmt
	go test -v -failfast -count=1 -run $(test) . ./...

release:
	$(gitclean)
	@$(if $(update),gh release delete -y v$(version),)
	gh release create v$(version) --notes "v$(version)"

latest_module_release = $(shell gh --repo $(1) release list --json tagName --jq '.[0].tagName')

update:
	@echo checking dependencies for updated versions 
	@$(foreach module,$(rstms_modules),go get $(module)@$(call latest_module_release,$(module));)
	curl -L -o cmd/common.go https://raw.githubusercontent.com/rstms/go-common/master/proxy_common_go
	sed <cmd/common.go >notify/common.go 's/^package cmd/package notify/'

logclean: 
	[ -f $(logfile) ] && echo >$(logfile) || true

clean: logclean
	rm -f $(program) *.core 
	go clean

sterile: clean
	which $(program) && go clean -i || true
	go clean
	go clean -cache
	go clean -modcache
	rm -f go.mod go.sum
