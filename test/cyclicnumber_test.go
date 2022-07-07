package cyclicnumber

import (
	"log"
	"testing"

	cp "github.com/UedaTakeyuki/compare"
	cn "local.packages/cyclicnumber"
)

func Test_01(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

	cyclicnumber := cn.CyclicNumber{Start: 0, Present: 0, Threshold: 3}

	cp.Compare(t, cyclicnumber.GetPresent(), int64(0))
	cp.Compare(t, cyclicnumber.GetPresentAndGoNext(), int64(0))
	cp.Compare(t, cyclicnumber.GetPresent(), int64(1))
	cp.Compare(t, cyclicnumber.Next(), int64(2))
	cp.Compare(t, cyclicnumber.Next(), int64(3))
	cp.Compare(t, cyclicnumber.Next(), int64(0))
}
