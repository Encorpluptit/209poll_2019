![Build](https://github.com/Encorpluptit/209poll_2019/workflows/Build/badge.svg)
[![codecov](https://codecov.io/gh/Encorpluptit/209poll_2019/branch/master/graph/badge.svg?token=ID0NHLIUAG)](https://codecov.io/gh/Encorpluptit/209poll_2019)

# 209poll_2019

### [Subject](Project/B-MAT-400_209poll.pdf) :
The goal of this project is to compute the 95% and 99% confidence intervals, given (as arguments):
 - Population size.
 - Size of the sample of population.
 - Percentage of voting intentions for a specific candidate.

(made with [Killian Clette](https://github.com/Skerilyo))

#### Mathematical Resolution:
People follows a Bernoulli process, and therefore that a binomial distribution (converging toward a normal distribution) is a good model for the results.

#### Usage:
    make
    ./poll 10000 500 42.24

#### Help:
    ./poll -h
    ./poll --help


## Skills Learned/Upgraded

- ### General
    - [Python](/)
    - [Go](#go)
    - [Haskell](#haskell)
    - [Codecov](#codecov-coverage) ([Official Site](https://codecov.io/)) in CI.
    - Markdown

- ### Python
    - [Python Tests with PyTests FrameWork.](https://www.python-course.eu/python3_pytest.php)

- ### Go
    - <ins>Use of Go routine and WaitGroup</ins> :
        - [Go Routine](https://tutorialedge.net/golang/concurrency-with-golang-goroutines/)
        - [Wait Group](https://tutorialedge.net/golang/go-waitgroup-tutorial/)
    - <ins>Upgrade skills on go coverage</ins> :
        - [Introduction and CI Integration](https://blog.seriesci.com/how-to-measure-code-coverage-in-go/)
        - [Cover Tool](https://blog.golang.org/cover)
        - [Testing Race Detector](https://blog.golang.org/race-detector) (Non-synchronised data)

- ### Haskell (Unfinished)
    - Argument management.

#### <ins>CodeCov Coverage</ins> :
 - [Python](https://github.com/codecov/example-python)
 - [Go](https://github.com/codecov/example-go)

