package packb

import (
	"github.com/mosceo/counter"
)

func Inc() {
	counter.Inc()
}

func Dec() {
	counter.Dec()
}

func Counter() int {
	return counter.Counter
}

func init() {
	counter.Inc()

	counter.AddValue(5)
	counter.AddValue(6)
	counter.AddValue(7)
}
