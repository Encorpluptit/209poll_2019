package tests

import (
	"Poll/poll"
	"testing"
)

func TestArgv(t *testing.T) {
	tables := []struct {
		arg []string
		exp float64
		err *poll.Error
	}{
		{[]string{"-0.5"}, 0., &argv.ExitError},
		{[]string{"-84"}, 0., &argv.ExitError},
		{[]string{"50"}, 0., &argv.ExitError},
		{[]string{"3"}, 0., &argv.ExitError},
		{[]string{"a"}, 0., &argv.ExitError},
		{[]string{"2a"}, 0., &argv.ExitError},
		{[]string{"3."}, 0., &argv.ExitError},
		{[]string{""}, 0., &argv.ExitError},
		{[]string{"3", "4"}, 0., &argv.ExitError},
		{[]string{"2."}, 2., nil},
		{[]string{"0.2"}, 0.2, nil},
		{[]string{"1.7"}, 1.7, nil},
		{[]string{"1.8"}, 1.8, nil},
		{[]string{"2.4"}, 2.4, nil},
		{[]string{"2.5"}, 2.5, nil},
		{[]string{"0"}, 0, nil},
	}

	for _, table := range tables {
		res, err := argv.Check(table.arg)
		if err != table.err {
			t.Errorf(
				"For argument(s) [%v]), err is [%v] (Expected [%v]\n", table.arg, err, table.err)
		}
		if res != table.exp {
			t.Errorf(
				"For argument(s) [%v]), res is [%v] (Expected [%v]\n", table.arg, res, table.exp)
		}
	}
}
