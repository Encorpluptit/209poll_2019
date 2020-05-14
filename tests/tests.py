#!/usr/bin/env python3
##
## EPITECH PROJECT, 2020
## 209Polls
## File description:
## Source file for Python tests.
##

from PollClass import Poll
from poll import main as main_fct
from poll import EXIT_ERROR
import pytest


@pytest.mark.parametrize(
    'param',
    [
        ["10000"],
        ["10000", "500"],
        ["10000", "500", "42.24", "500"],
        ["10000a", "500", "42.24"],
        ["10000.0", "500", "42.24"],
        ["-1", "500", "42.24"],
        ["10000", "500a", "42.24"],
        ["10000", "500.0", "42.24"],
        ["10000", "-1", "42.24"],
        ["10000", "500", "42.24."],
        ["10000", "500", "42.24a"],
        ["10000", "500", "101"],
        ["10000", "500", "-1"],
    ]
)
def test_bad_argv(param):
    with pytest.raises(SystemExit) as pytest_wrapped_e:
        main_fct(["204ducks"] + param)
    assert pytest_wrapped_e.type == SystemExit
    assert pytest_wrapped_e.value.code == EXIT_ERROR


@pytest.mark.parametrize(
    'p_size, s_size, p, variance, ci95min, ci95max, ci99min, ci99max',
    [
        [10000, 500, 42.24, 0.000464, 38.02, 46.46, 36.68, 47.80],
        [10000, 100, 45, 0.002450, 35.30, 54.70, 32.23, 57.77],
    ]
)
def test_value(capsys, p_size, s_size, p, variance, ci95min, ci95max, ci99min, ci99max):
    poll = Poll(p_size, s_size, p)
    poll.calc_variance()
    assert round(poll.var, 6) == variance
    assert round(max(poll.p - poll.ci95, 0), 2) == ci95min
    assert round(min(poll.p + poll.ci95, 100), 2) == ci95max
    assert round(max(poll.p - poll.ci99, 0), 2) == ci99min
    assert round(min(poll.p + poll.ci99, 100), 2) == ci99max
    poll.print_results()
    captured = capsys.readouterr()
