#!/usr/bin/env python3
##
## EPITECH PROJECT, 2020
## 209Polls
## File description:
## Source file for main.
##

from sys import argv
from PollClass import Poll


EXIT_ERROR = 84


def main(args):
    if "-h" in args or "--help" in args:
        Poll.print_help()
    if len(args) != 4:
        Poll.print_help(EXIT_ERROR)
    try:
        p_size, s_size, p = int(args[1]), int(args[2]), float(args[3])
        if p_size < 0 or s_size < 0 or not (0 < p < 100):
            raise ValueError
        poll = Poll(int(args[1]), int(args[2]), float(args[3]))
        poll.calc_variance()
        poll.print_results()
    except ValueError:
        exit(EXIT_ERROR)


if __name__ == '__main__':
    main(argv)
