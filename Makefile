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

COV_FILE		=	.coverage

PY_CACHE		=	.pytest_cache

#############################################################################################

HEADER			=			'\033[95m'

END_HEADER		=			'\033[0m'

#############################################################################################
.DEFAULT: $(NAME)


all: $(NAME)


review: re
	@-$(MAKE) -s -C $(GO_PATH) re


$(NAME):
	@python3 -c "print($(HEADER) + '*'*120 + $(END_HEADER))"
	@printf $(HEADER)"Compiling $(NAME) <--> $(LANGUAGE)\n"$(END_HEADER)
	@cp $(PY_NAME) $(NAME)
	@chmod u+x $(NAME)


tests_run: unit_tests func_tests

unit_tests:
	@python3 -c "print($(HEADER) + '*'*120 + $(END_HEADER))"
	@printf $(HEADER)"Launching Python Tests\n\n"$(END_HEADER)
	@-python3 -m pytest -v --cov=PollClass --cov=poll tests/tests.py
	@printf "\n\n"


func_tests:
	@python3 -c "print($(HEADER) + '*'*120 + $(END_HEADER))"
	@printf $(HEADER)"Launching Functionnals Tests\n\n"$(END_HEADER)
	@-./tests/jenrik tests/test_209poll.toml
	@printf "\n\n"


all_unit_tests: review
	@-$(MAKE) -s  unit_tests
	@-$(MAKE) -s -C $(GO_PATH) unit_tests


all_func_tests: review
	@$(MAKE) -s  func_tests
	@$(MAKE) -s -C $(GO_PATH) func_tests


all_tests: tests_run
	@-$(MAKE) -s go_tests_run


go_tests_run:
	@-$(MAKE) -s -C $(GO_PATH) tests_run


haskell_tests_run:
	@-$(MAKE) -s -C $(HASKELL_PATH) tests_run


clean:
	@$(RM) -rd .pytest_cache
	@$(RM) .coverage


fclean: clean
	@-$(MAKE) -s -C $(GO_PATH) fclean
	@-$(MAKE) -s -C $(HASKELL_PATH) fclean
	@-$(MAKE) -s -C $(C_PATH) fclean
	@printf $(HEADER)"Cleaning $(NAME) <--> $(LANGUAGE)\n"$(END_HEADER)
	@$(RM) $(NAME)
	@$(RM) $(COV_FILE)

re: fclean all


.PHONY: all $(NAME) clean fclean re tests_run
