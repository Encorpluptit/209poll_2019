##
## EPITECH PROJECT, 2020
## 209Polls
## File description:
## Makefile for Python.
##


#############################################################################################
NAME			=	209poll

LANGUAGE		=	Python

PY_NAME			=	209poll.py

BONUS_PATH		=	BONUS

GO_PATH			=	$(BONUS_PATH)/go

HASKELL_PATH	=	$(BONUS_PATH)/haskell

C_PATH			=	$(BONUS_PATH)/c

#############################################################################################

HEADER			=			'\033[95m'

END_HEADER		=			'\033[0m'

#############################################################################################
.DEFAULT: $(NAME)


all: $(NAME)


$(NAME):
	@printf $(HEADER)"Compiling $(NAME) <--> $(LANGUAGE)\n"$(END_HEADER)
	@cp $(PY_NAME) $(NAME)
	@chmod u+x $(NAME)


tests_run:
	@printf $(HEADER)"Launching Python Unit Tests\n"$(END_HEADER)
	@-python3 -m pytest -v --cov=Poll tests/tests.py
	@-$(MAKE) -s -C $(GO_PATH) tests_run
	@-$(MAKE) -s -C $(HASKELL_PATH) tests_run
	@-$(MAKE) -s -C $(C_PATH) tests_run
	@printf $(HEADER)"Launching Functionnals Tests\n"$(END_HEADER)
	@-./tests/jenrik tests/test_204ducks.toml

clean:

fclean:
	@$(RM) $(NAME)

re: fclean all

.PHONY: all $(NAME) clean fclean re tests_run
