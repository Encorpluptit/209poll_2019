##
## EPITECH PROJECT, 2020
## 209Polls
## File description:
## Makefile for Go.
##


#############################################################################################
NAME			=	209poll

MAIN			=	main.go

LANGUAGE		=	Go

COV_FILE		=	coverage.txt

#############################################################################################

HEADER			=	'\033[95m'

END_HEADER		=	'\033[0m'

#############################################################################################

.DEFAULT: $(NAME)


all: $(NAME)


$(NAME):
	@python3 -c "print($(HEADER) + '*'*120 + $(END_HEADER))"
	@printf $(HEADER)"Compiling $(NAME) <--> $(LANGUAGE)\n"$(END_HEADER)
	@go build -o $(NAME) $(MAIN)


tests_run: unit_tests func_tests


unit_tests:
	@python3 -c "print($(HEADER) + '*'*120 + $(END_HEADER))"
	@printf $(HEADER)"Launching $(LANGUAGE) Unit Tests\n\n"$(END_HEADER)
	@./tests/gotest -v -coverprofile=$(COV_FILE) -covermode=atomic -coverpkg=./poll ./tests/unit
#@./tests/gotest -v -coverpkg=./poll ./tests/unit
#@./tests/gotest -v -coverpkg=./... ./tests/unit
	@printf "\n\n"


func_tests:
	@python3 -c "print($(HEADER) + '*'*120 + $(END_HEADER))"
	@printf $(HEADER)"Launching $(LANGUAGE) Functionals Tests\n\n"$(END_HEADER)
	@printf "\n\n"


clean:


fclean: clean
	@printf $(HEADER)"Cleaning $(NAME) <--> $(LANGUAGE)\n"$(END_HEADER)
	@$(RM) $(NAME)
	@$(RM) $(COV_FILE)


re: fclean all


.PHONY: all $(NAME) clean fclean re tests_run
