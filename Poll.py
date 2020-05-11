from sys import argv
from math import sqrt


class Poll:
    def __init__(self, p_size, s_size, p):
        self.var = None
        self.p_size, self.s_size, self.p = p_size, s_size, p
        self.ci95, self.ci99 = 0.0, 0.0

    @staticmethod
    def print_help(exit_code: int = 0):
        print(
            "Usage\n" +
            f"\t{argv[0]}" +
            ""
              )
        exit(exit_code)

    def calc_variance(self):
        self.var = ((self.p / 100) * (1 - (self.p / 100)) / self.s_size) * (
                    (self.p_size - self.s_size) / (self.p_size - 1))
        self.ci95 = 196 * sqrt(self.var)
        self.ci99 = 258 * sqrt(self.var)

    def print_results(self):
        print(f"Population Size:\t\t{self.p_size}")
        print(f"Sample Size:\t\t\t{self.s_size}")
        print(f"Voting intentions:\t\t{self.p}%")
        print(f'Variance:\t\t\t\t{self.var:.6f}')
        print(f"95% confidence interval: [{max(self.p - self.ci95, 0):.2f}%; {min(self.p + self.ci95, 100):.2f}%]")
        print(f"99% confidence interval: [{max(self.p - self.ci99, 0):.2f}%; {min(self.p + self.ci99, 100):.2f}%]")
