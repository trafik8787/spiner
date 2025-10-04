package spiner

import (
	"fmt"
	"strings"
)


type ProgressBar struct {
	step int
	count int
	arrCount []int
}

func (p *ProgressBar) Start (count int) {
	p.count = count
	p.step = 0
	p.bar()
}


func (p *ProgressBar) Step () {

	if p.step < p.count {
		p.step++
		p.bar()
	}
}

func (p *ProgressBar) bar () {

	percent := float64(p.step) / float64(p.count) * 100
	bar := strings.Repeat("█", int(percent/10)) + strings.Repeat("-", 10-int(percent/10))
	fmt.Printf("\r[%s] %.0f%%", bar, percent)
	
}