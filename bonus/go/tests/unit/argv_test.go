package poll

import (
	"Poll/poll"
	"testing"
)

func TestArgv(t *testing.T) {
	tables := []struct {
		arg []string
		err *poll.Error
	}{
		{[]string{}, &poll.ArgvError},
		{[]string{"10000"}, &poll.ArgvError},
		{[]string{"10000", "500"}, &poll.ArgvError},
		{[]string{"10000", "500", "42.24", "500"}, &poll.ArgvError},
		{[]string{"-h", "500", "42.24"}, &poll.HelpError},
	}

	for _, table := range tables {
		err := poll.CheckArgv(table.arg)
		if err != table.err {
			t.Errorf(
				"For argument(s) [%#v]), err is [%v] (Expected [%v]\n", table.arg, err, table.err)
		}
	}
}

func TestInitPollError(t *testing.T) {
	tables := []struct {
		arg []string
		err *poll.Error
	}{
		{[]string{"10000a", "500", "42.24"}, &poll.IntError},
		{[]string{"10000.0", "500", "42.24"}, &poll.IntError},
		{[]string{"-1", "500", "42.24"}, &poll.IntError},
		{[]string{"10000", "500a", "42.24"}, &poll.IntError},
		{[]string{"10000", "500.0", "42.24"}, &poll.IntError},
		{[]string{"10000", "-1", "42.24"}, &poll.IntError},
		{[]string{"10000", "500", "42.24."}, &poll.FloatError},
		{[]string{"10000", "500", "42.24a"}, &poll.FloatError},
		{[]string{"10000", "500", "101"}, &poll.FloatError},
		{[]string{"10000", "500", "-1"}, &poll.FloatError},
	}

	for _, table := range tables {
		_, err := poll.InitPoll(table.arg)
		if err != table.err {
			t.Errorf(
				"For argument(s) [%#v]), err is [%v] (Expected [%v]\n", table.arg, err, table.err)
			continue
		}
	}
}
