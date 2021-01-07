PKGS := release stream fedoracoreos

build:
	for pkg in $(PKGS); do (cd $$pkg && go build -mod=vendor); done
.PHONY: build

test:
	for pkg in $(PKGS); do (cd $$pkg && go test); done
.PHONY: test
