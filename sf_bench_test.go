/* ----------------------------------
*  @author suyame 2022-07-12 15:06:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package snowflake

import (
	"sync"
	"testing"
)

func BenchmarkSFGenerate(b *testing.B) {
	sf, err := NewSF(31, 28)
	if err != nil {
		b.Error("BM sf_generate err! new sf failed!")
	}
	wg := sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sf.Generate()
		}()
	}
	wg.Wait()
}
