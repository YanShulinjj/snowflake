/* ----------------------------------
*  @author suyame 2022-07-11 19:16:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package snowflake

import (
	"sync/atomic"
	"time"
)

// 定义一些选择器
// WORKERSPACESELECTOR 从id中选择出工作区编号
// NODESELECTOR 从id中选择出节点编号
// SEQSELECTOR 从id中选择出顺序编号
const (
	WORKERSPACESELECTOR IDtpye = 0x3e0000
	NODESELECTOR        IDtpye = 0x1f000
	SEQSELECTOR         IDtpye = 0xfff
)

// IDtpye 是分布式id的类型，是uint64的别名
type IDtpye uint64

// SnowFlake 是雪花算法维护的数据结构
type SnowFlake struct {
	workerSpaceID uint8
	nodeID        uint8
	sequence      uint32
}

// NewSF 新建一个sf对象
func NewSF(workerSpaceID, nodeID uint8) (*SnowFlake, error) {
	if workerSpaceID > 31 {
		return nil, WorkerSpaceIdOverFlowErr
	}
	if nodeID > 31 {
		return nil, NodeIdOverFlowErr
	}
	return &SnowFlake{
		workerSpaceID: workerSpaceID,
		nodeID:        nodeID,
		sequence:      0,
	}, nil
}

// Generate 生成一个分布式id
func (sf *SnowFlake) Generate() IDtpye {
	id := IDtpye(0)
	// fmt.Println(strconv.FormatUint(uint64(id), 2))
	// 获取当前的时间戳（微秒）
	t := uint64(time.Now().UnixMicro())
	// part2 时间戳
	id |= IDtpye((t << 23) >> 1)
	// part3 机房id
	id |= IDtpye(sf.workerSpaceID<<3) << 14
	// part4 机器id
	id |= IDtpye(sf.nodeID<<3) << 9
	// part5 sequence
	atomic.AddUint32(&sf.sequence, 1)
	id |= IDtpye((sf.sequence << 20) >> 20)
	return id
}

// GetWorkerSpaceID 得到工作区编号
func (it IDtpye) GetWorkerSpaceID() uint8 {
	// 获取it 的低17bit ~ 22bit
	return uint8(it & WORKERSPACESELECTOR >> 17)
}

// GetNodeID 得到节点编号
func (it IDtpye) GetNodeID() uint8 {
	// 获取it 的低12bit ~ 17bit
	return uint8(it & NODESELECTOR >> 12)
}

// GetSequence 得到顺序编号
func (it IDtpye) GetSequence() uint32 {
	return uint32(it & SEQSELECTOR)
}
