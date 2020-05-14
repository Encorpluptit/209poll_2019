package poll

import (
	"Poll/poll"
	"reflect"
	"testing"
)

func instancePollEq(res *poll.Poll, exp *poll.Poll, t *testing.T) bool {
	v := reflect.ValueOf(*res)
	w := reflect.ValueOf(*exp)

	for i := 0; i < v.NumField(); i++ {
		if !v.Field(i).CanInterface() || !w.Field(i).CanInterface() {
			continue
		}
		if v.Field(i).Interface() != w.Field(i).Interface() {
			return false
		}
	}
	return true
}

func TestRunPoll(t *testing.T) {
	tables := []struct {
		arg []string
		exp *poll.Poll
	}{
		{[]string{"10000", "500", "42.24"}, &poll.Poll{PSize: 10000, SSize: 500, P: 42.24}},
		{[]string{"1000", "500", "42.24"}, &poll.Poll{PSize: 1000, SSize: 500, P: 42.24}},
		{[]string{"10000", "50", "42.24"}, &poll.Poll{PSize: 10000, SSize: 50, P: 42.24}},
		{[]string{"10000", "500", "99"}, &poll.Poll{PSize: 10000, SSize: 500, P: 99}},
		{[]string{"5", "1", "1"}, &poll.Poll{PSize: 5, SSize: 1, P: 1}},
	}

	for _, table := range tables {
		res, err := poll.InitPoll(table.arg)
		if err != nil {
			t.Errorf(
				"For argument(s) [%#v]), Error is not nill, got [%v]", table.arg, err)
		}
		if !instancePollEq(res, table.exp, t) {
			t.Errorf(
				"For argument(s) [%#v]), res is [%v] (Expected [%v]\n", table.arg, res, table.exp)
		}
	}
}
