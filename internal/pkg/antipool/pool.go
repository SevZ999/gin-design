// internal/pkg/antipool/pool.go
package antipool

import (
	"github.com/panjf2000/ants/v2"
)

func NewAntsPool(size int) (*ants.Pool, error) {
	return ants.NewPool(size)
}
