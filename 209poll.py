from sys import argv
from Poll import Poll

EXIT_ERROR = 84


def main():
    if "-h" in argv or "--help" in argv:
        Poll.print_help()
    if len(argv) != 4:
        Poll.print_help(EXIT_ERROR)
    try:
        poll = Poll(int(argv[1]), int(argv[2]), float(argv[3]))
        poll.run()
    except ValueError:
        exit(EXIT_ERROR)


if __name__ == '__main__':
    main()
