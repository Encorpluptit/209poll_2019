##
## EPITECH PROJECT, 2020
## 209Polls
## File description:
## Makefile for Python.
##


#############################################################################################
NAME			=	209poll

LANGUAGE		=	Python

PY_NAME			=	poll.py

BONUS_PATH		=	bonus

GO_PATH			=	$(BONUS_PATH)/go

HASKELL_PATH	=	$(BONUS_PATH)/haskell

C_PATH			=	$(BONUS_PATH)/c

#############################################################################################

HEADER			=			'\033[95m'

END_HEADER		=			'\033[0m'

#############################################################################################
.DEFAULT: $(NAME)


all: $(NAME)

review: $(NAME)
	@-$(MAKE) -s -C $(GO_PATH) re
	@-$(MAKE) -s -C $(HASKELL_PATH) re
	@-$(MAKE) -s -C $(C_PATH) re


$(NAME):
	@python3 -c "print($(HEADER) + '='*120 + $(END_HEADER))"
	@printf $(HEADER)"Compiling $(NAME) <--> $(LANGUAGE)\n"$(END_HEADER)
	@cp $(PY_NAME) $(NAME)
	@chmod u+x $(NAME)


tests_run:
	@python3 -c "print($(HEADER) + '*'*120 + $(END_HEADER))"
	@printf $(HEADER)"Launching Python Tests\n\n"$(END_HEADER)
	@-python3 -m pytest -v --cov=Poll tests/tests.py
	@python3 -c "print($(HEADER) + '*'*120 + $(END_HEADER))"
	@printf $(HEADER)"Launching Functionnals Tests\n\n"$(END_HEADER)
	@-./tests/jenrik tests/test_204ducks.toml


all_tests: tests_run
	@-$(MAKE) -s go_tests_run
	@-$(MAKE) -s haskell_tests_run
	@-$(MAKE) -s c_tests_run


go_tests_run:
	@-$(MAKE) -s -C $(GO_PATH) tests_run


haskell_tests_run:
	@-$(MAKE) -s -C $(HASKELL_PATH) tests_run


c_tests_run:
	@-$(MAKE) -s -C $(C_PATH) tests_run


clean:
	@$(RM) -rd .pytest_cache
	@$(RM) .coverage

fclean: clean
	@printf $(HEADER)"Cleaning $(NAME) <--> $(LANGUAGE)\n"$(END_HEADER)
	@$(RM) $(NAME)

re: fclean all

.PHONY: all $(NAME) clean fclean re tests_run
