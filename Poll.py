from sys import argv


class Poll:
    def __init__(self, p_size, s_size, p):
        self.var = None
        self.p_size, self.s_size, self.p = p_size, s_size, p

    @staticmethod
    def print_help(exit_code: int = 0):
        print(
            "Usage\n" +
            f"\t{argv[0]}"
              )
        exit(exit_code)

    def calc_variance(self):
        self.var = ((self.p / 100) * (1 - (self.p / 100)) / self.s_size) * (
                    (self.p_size - self.s_size) / (self.p_size - 1))

    def print_results(self):
        print(f"Population Size:\t\t{self.p_size}")
        print(f"Sample Size:\t\t{self.s_size}")
        print(f"Voting intentions:\t\t{self.s_size}")
        print(f"Variance:\t\t{self.var}")
        print(f"95% confidence interval: [{0}; {0}]")
        print(f"99% confidence interval: [{0}; {0}]")

    def run(self):
        self.calc_variance()
        self.print_results()
