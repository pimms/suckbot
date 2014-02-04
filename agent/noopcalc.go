package agent

import (
	"fmt"
	"github.com/pimms/suckbot/arg"
)

const (
	MAX_SUCK = 0.33
	MIN_SUCK = 0.25
)

type t_noopstate struct {
	max int
	min int
	inc int
	cur int
}

func (n *t_noopstate) init() {
	n.max = 4
	n.min = 2
	n.inc = 1
	n.cur = 4
}

func (n *t_noopstate) onRoundComplete(history *t_history, totalTiles int) int {
	var p float64 = history.getActionPercent(SUCK, totalTiles)

	// See "alg_outline" for details here
	if p < MIN_SUCK {
		if n.inc != n.min/2 {
			if n.inc == 0 {
				n.max = n.min
				n.min /= 2
				n.inc = n.min / 2
				n.cur = n.max
			} else {
				n.cur += n.inc
				n.inc /= 2
			}
		} else {
			n.min = n.max
			n.max *= 2
			n.inc = n.min / 2
			n.cur = n.max
		}
	} else if p > MAX_SUCK {
		if n.inc == 0 {
			n.min = n.max
			n.max *= 2
			n.inc = n.min / 2
			n.cur = n.max
		} else {
			n.cur -= n.inc
			n.inc /= 2
		}
	}

	if arg.Verbose() {
		fmt.Printf("[noopcalc]: P:%f min:%d  max:%d  inc:%d  cur:%d\n",
			p, n.min, n.max, n.inc, n.cur)
	}

	return n.cur
}
