package poll

import (
	"fmt"
	"math"
	"sync"
)

const (
	ExitError   = 84
	ExitSuccess = 0
)

type Poll struct {
	Variance     float64
	PSize, SSize int64
	P            float64
	Ci95, Ci99   float64
}

func (p *Poll) CalculateVariance(wg *sync.WaitGroup) {
	p.Variance = ((p.P / 100.0) * (1 - p.P/100.0) / float64(p.SSize)) * (float64(p.PSize-p.SSize) / (float64(p.PSize) - 1.0))
	p.Ci95 = 1.96 * math.Sqrt(p.Variance) * 100
	p.Ci99 = 2.58 * math.Sqrt(p.Variance) * 100
	wg.Done()
}

func (p *Poll) Print() {
	fmt.Printf("Population size:\t\t%d\n", p.PSize)
	fmt.Printf("Sample size:\t\t\t%d\n", p.SSize)
	fmt.Printf("Voting intentions:\t\t%.2f%%\n", p.P)
	fmt.Printf("Variance:\t\t\t%.6f\n", p.Variance)
	fmt.Printf("95%% confidence interval:\t[%.2f%%; %.2f%%]\n",
		math.Max(p.P-p.Ci95, 0), math.Min(p.P+p.Ci95, 100))
	fmt.Printf("99%% confidence interval:\t[%.2f%%; %.2f%%]\n",
		math.Max(p.P-p.Ci99, 0), math.Min(p.P+p.Ci99, 100))
}
