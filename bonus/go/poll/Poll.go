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
	variance     float64
	pSize, sSize int64
	p            float64
	ci95, ci99   float64
}

func (p *Poll) CalculateVariance(wg *sync.WaitGroup) {
	p.variance = ((p.p / 100.0) * (1 - p.p/100.0) / float64(p.sSize)) * (float64(p.pSize-p.sSize) / (float64(p.pSize) - 1.0))
	p.ci95 = 1.96 * math.Sqrt(p.variance) * 100
	p.ci99 = 2.58 * math.Sqrt(p.variance) * 100
	wg.Done()
}

func (p *Poll) Print() {
	fmt.Printf("Population size:\t\t%d\n", p.pSize)
	fmt.Printf("Sample size:\t\t\t%d\n", p.sSize)
	fmt.Printf("Voting intentions:\t\t%.2f%%\n", p.p)
	fmt.Printf("Variance:\t\t\t%.6f\n", p.variance)
	fmt.Printf("95%% confidence interval:\t[%.2f%%; %.2f%%]\n", math.Max(p.p-p.ci95, 0), math.Min(p.p+p.ci95, 100))
	fmt.Printf("99%% confidence interval:\t[%.2f%%; %.2f%%]\n", math.Max(p.p-p.ci99, 0), math.Min(p.p+p.ci99, 100))
}
