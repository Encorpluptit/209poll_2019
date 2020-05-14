package poll

import (
	"Poll/poll"
	"math"
	"reflect"
	"sync"
	"testing"
)

type ciTests struct {
	ci95min float64
	ci95max float64
	ci99min float64
	ci99max float64
}

func instanceRunPollEq(res *poll.Poll, exp *poll.Poll, ci *ciTests) bool {
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
	if ci.ci95min != math.Max(res.P-res.Ci95, 0) {
		return false
	}
	if ci.ci95max != math.Min(res.P+res.Ci95, 0) {
		return false
	}
	if ci.ci99min != math.Max(res.P-res.Ci99, 0) {
		return false
	}
	if ci.ci99max != math.Min(res.P+res.Ci99, 0) {
		return false
	}

	return true
}

func TestRunPoll(t *testing.T) {
	tables := []struct {
		arg []string
		exp *poll.Poll
		ci  *ciTests
	}{
		{[]string{"10000", "500", "42.24"},
			&poll.Poll{PSize: 10000, SSize: 500, P: 42.24, Variance: 0.000464},
			&ciTests{38.02, 46.46, 36.68, 47.80},
		},
	}
	var wg sync.WaitGroup

	for _, table := range tables {
		wg.Add(1)
		go func() {
			res, err := poll.InitPoll(table.arg)
			if err != nil {
				t.Errorf(
					"For argument(s) [%#v]), Error is not nil, got [%v]", table.arg, err)
			}

			res.CalculateVariance(&wg)
			if !instanceRunPollEq(res, table.exp, table.ci) {
				t.Errorf(
					"For argument(s) [%#v]), res is [%v] (Expected:[%v]\n[%v]\n", table.arg, res, table.exp, table.ci)
			}
		}()
	}
	wg.Wait()
}
