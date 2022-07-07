package cyclicnumber

import (
	"log"
	"sync"
	"time"

	"github.com/UedaTakeyuki/erapse"
)

/** Cyclic Number */
type CyclicNumber struct {
	Start     int64
	Present   int64
	Threshold int64
	Mu        sync.Mutex
}

func (n *CyclicNumber) Next() (present int64) {
	defer erapse.ShowErapsedTIme(time.Now())

	n.Mu.Lock()
	log.Println("n.Present", n.Present)
	log.Println("n.Threshold", n.Threshold)
	log.Println("n.Start", n.Start)
	log.Println("n.Present >= n.Threshold", n.Present >= n.Threshold)
	if n.Present >= n.Threshold {
		n.Present = n.Start
	} else {
		n.Present = n.Present + 1
	}
	present = n.Present
	n.Mu.Unlock()
	return
}

func (n *CyclicNumber) GetPresent() (present int64) {
	defer erapse.ShowErapsedTIme(time.Now())

	n.Mu.Lock()
	present = n.Present
	n.Mu.Unlock()
	return
}

func (n *CyclicNumber) GetPresentAndGoNext() (present int64) {
	defer erapse.ShowErapsedTIme(time.Now())

	n.Mu.Lock()
	present = n.Present
	n.Mu.Unlock()
	n.Next()
	log.Println("n.Present", n.Present)
	return
}
