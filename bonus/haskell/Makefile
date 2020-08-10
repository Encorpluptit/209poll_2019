##
## EPITECH PROJECT, 2020
## FUN_imageCompressor_2019
## File description:
## Makefile
##


package = poll
NAME = 209poll


stack_yaml				=	STACK_YAML="stack.yaml"
stack					=	$(stack_yaml) stack
executable				:=	$(shell stack path --local-install-root)
profiling_executable	:=	$(shell stack --work-dir .stack-work-profiling path --local-install-root)

LANGUAGE		=			Haskell

HEADER			=			'\033[95m'

END_HEADER		=			'\033[0m'

#############################################################################################


all: build sign

re: all

unit_tests:
	$(stack) --work-dir .stack-work-profiling test --coverage

profiling: build-profile
	$(stack) --work-dir .stack-work-profiling path --local-install-root
	./test/profiling/profile.sh

clean:
	@printf $(HEADER)"Cleaning $(NAME) <--> $(LANGUAGE)\n"$(END_HEADER)

fclean: clean
	$(RM) $(package)
	$(RM) $(NAME)

build:
	@python3 -c "print($(HEADER) + '*'*120 + $(END_HEADER))"
	@printf $(HEADER)"Compiling $(NAME) <--> $(LANGUAGE)\n"$(END_HEADER)
	@$(stack) build $(package)
	@cp $(executable)/bin/poll ./$(NAME)

build-dirty:
	$(stack) build --ghc-options=-fforce-recomp $(package)

build-profile:
	$(stack) --work-dir .stack-work-profiling --profile build $(package)

run:
	$(stack) build --fast && $(stack) exec -- $(NAME)

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