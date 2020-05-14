package poll

import (
	"Poll/poll"
	"math"
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
	if exp.Variance != math.Round(res.Variance*1000000)/1000000 {
		return false
	}
	if ci.ci95min != math.Round(math.Max(res.P-res.Ci95, 0)*100)/100 {
		return false
	}
	if ci.ci95max != math.Round(math.Min(res.P+res.Ci95, 100)*100)/100 {
		return false
	}
	if ci.ci99min != math.Round(math.Max(res.P-res.Ci99, 0)*100)/100 {
		return false
	}
	if ci.ci99max != math.Round(math.Min(res.P+res.Ci99, 100)*100)/100 {
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
		{[]string{"10000", "100", "45"},
			&poll.Poll{PSize: 10000, SSize: 100, P: 45, Variance: 0.002450},
			&ciTests{35.30, 54.70, 32.23, 57.77},
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
