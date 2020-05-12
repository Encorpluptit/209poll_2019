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

#############################################################################################

HEADER			=	'\033[95m'

END_HEADER		=	'\033[0m'

#############################################################################################
.DEFAULT: $(NAME)


all: $(NAME)


$(NAME):
	@printf $(HEADER)"Compiling $(NAME) <--> $(LANGUAGE)\n"$(END_HEADER)
	@go build -o $(NAME) $(MAIN)


tests_run:
	@printf $(HEADER)"Launching $(LANGUAGE) Tests\n"$(END_HEADER)
	@./gotest -v -coverpkg=./... ./tests


clean:


fclean: clean
	@printf $(HEADER)"Cleaning $(NAME) <--> $(LANGUAGE)\n"$(END_HEADER)
	@$(RM) $(NAME)


re: fclean all


.PHONY: all $(NAME) clean fclean re tests_run
