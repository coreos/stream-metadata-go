PKGS := arch release stream stream/rhcos release/rhcos fedoracoreos

build:
	for pkg in $(PKGS); do (cd $$pkg && go build -mod=vendor); done
.PHONY: build

test:
	for pkg in $(PKGS); do (cd $$pkg && go test); done
.PHONY: test
