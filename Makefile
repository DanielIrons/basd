
DATE=$(shell TZ=GMT date --rfc-3339="seconds")
HASH=$(shell printf "r%s.%s" $(shell git rev-list --count HEAD) $(shell git describe --always --abbrev=11 --dirty --exclude '*'))

DEPS_CONFIG := $(shell find config -name \*.go)

basd: main.go $(DEPS_CONFIG)
	go build -o basd -trimpath -ldflags=" \
		-s -w \
		-X 'main.buildDate=$(DATE)' \
		-X 'main.buildHash=$(HASH)'" \
	main.go