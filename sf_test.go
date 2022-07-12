/* ----------------------------------
*  @author suyame 2022-07-12 11:07:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package snowflake

import (
	"sync"
	"testing"
)

func TestNewSF(t *testing.T) {
	expect := WorkerSpaceIdOverFlowErr
	if _, err := NewSF(200, 20); err != expect {
		t.Error("NewSF err! sf should be nil when ws_id biggther than 31.")
	}
	expect = NodeIdOverFlowErr
	if _, err := NewSF(20, 200); err != expect {
		t.Error("NewSF err! sf should be nil when node_id biggther than 31.")
	}
	if _, err := NewSF(20, 20); err != nil {
		t.Error("NewSF err! sf should not be nil when node_id ans ws_id less than 31.")
	}
}

func TestSnowFlakeGenerate(t *testing.T) {
	expect_wsid := uint8(31)
	expect_nid := uint8(28)
	sf, err := NewSF(expect_wsid, expect_nid)
	if err != nil {
		t.Error("SFGenerate err! new sf failed!")
	}
	wg := sync.WaitGroup{}
	N := 1000
	ids := make([]IDtpye, N)
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			id := sf.Generate()
			ids[j] = id
		}(i)
	}
	wg.Wait()

	// 判断是否唯一
	visit := make(map[IDtpye]struct{})
	// 判断sequece是否唯一
	visit_seq := make(map[uint32]struct{})
	for _, id := range ids {
		if id == 0 {
			t.Error("SFGenerate err! id generate failed!")
		}
		if _, ok := visit[id]; ok {
			t.Error("SFGenerate err! id is not unique!")
		}
		visit[id] = struct{}{}
		/*---------判断3个part------*/
		seq := id.GetSequence()
		ws_id := id.GetWorkerSpaceID()
		nid := id.GetNodeID()
		if ws_id != expect_wsid {
			t.Errorf("SFGenerate err! ws not match! ws_id: %v, expect_ws_id: %v", ws_id, expect_wsid)
		}
		if nid != expect_nid {
			t.Errorf("SFGenerate err! nid not match! nid: %v, expect_nid: %v", nid, expect_nid)
		}
		if _, ok := visit_seq[seq]; ok {
			t.Error("SFGenerate err! id is not unique!")
		}
		visit_seq[seq] = struct{}{}
	}
}
