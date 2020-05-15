##
## EPITECH PROJECT, 2020
## FUN_imageCompressor_2019
## File description:
## Makefile
##


package = 209poll


stack_yaml = STACK_YAML="stack.yaml"
stack = $(stack_yaml) stack
executable := $(shell stack path --local-install-root)


all: build sign

re: all

#profiling:
#	stack build --profile
#	stack run --profile 2 0.8 your_file -- +RTS -p
#	cat imageCompressor.prof

clean:
	@echo "Cleaning Stack Project"

fclean: clean
	$(RM) $(package)

build:
	@$(stack) build $(package)
	@cp $(executable)/bin/poll ./209poll

build-dirty:
	$(stack) build --ghc-options=-fforce-recomp $(package)

build-profile:
	$(stack) --work-dir .stack-work-profiling --profile build

run:
	$(stack) build --fast && $(stack) exec -- $(package)

install:
	$(stack) install

ghci:
	$(stack) ghci $(package):lib --ghci-options='-j6 +RTS -A128m'

test:
	$(stack) test $(package)

test-ghci:
	$(stack) ghci $(package):test:$(package)-tests --ghci-options='-j6 +RTS -A128m'

bench:
	$(stack) bench $(package)

ghcid:
	$(stack) exec -- ghcid -c "stack ghci $(package):lib --test --ghci-options='-fobject-code -fno-warn-unused-do-bind -j4 +RTS -A128m' --main-is $(package):$(package)"

dev-deps:
	stack install ghcid

sign:
	@echo ""
	@echo "*******************"
	@echo "* Damien Bernard  *"
	@echo "* Killian Clette  *"
	@echo "*                 *"
	@echo "*  Epitech 2020   *"
	@echo "*******************"
	@echo ""

.PHONY : build build-dirty run install ghci test test-ghci ghcid dev-deps re cle