package idutil

import (
	"fmt"
	"sync"
	"time"
)

const (
	epoch            = 1654041600000 // 自定义起始时间戳 (2022年6月1日)
	workerIDBits     = 5             // 工作机器ID位数
	datacenterIDBits = 5             // 数据中心ID位数
	sequenceBits     = 12            // 序列号位数

	maxWorkerID     = -1 ^ (-1 << workerIDBits)
	maxDatacenterID = -1 ^ (-1 << datacenterIDBits)
	maxSequence     = -1 ^ (-1 << sequenceBits)

	workerIDShift      = sequenceBits
	datacenterIDShift  = sequenceBits + workerIDBits
	timestampLeftShift = sequenceBits + workerIDBits + datacenterIDBits
)

type Snowflake struct {
	mutex         sync.Mutex
	lastTimestamp int64
	sequence      int64
	workerID      int64
	datacenterID  int64
}

func NewSnowflake(workerID, datacenterID int64) (*Snowflake, error) {
	if workerID > maxWorkerID || workerID < 0 {
		return nil, fmt.Errorf("workerID must be between 0 and %d", maxWorkerID)
	}
	if datacenterID > maxDatacenterID || datacenterID < 0 {
		return nil, fmt.Errorf("datacenterID must be between 0 and %d", maxDatacenterID)
	}
	return &Snowflake{
		workerID:     workerID,
		datacenterID: datacenterID,
	}, nil
}

func (s *Snowflake) GenerateSnowFlake() int64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	timestamp := time.Now().UnixMilli()
	if timestamp < s.lastTimestamp {
		panic("Clock moved backwards. Refusing to generate id")
	}

	if timestamp == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			for timestamp <= s.lastTimestamp {
				timestamp = time.Now().UnixMilli()
			}
		}
	} else {
		s.sequence = 0
	}

	s.lastTimestamp = timestamp

	id := ((timestamp - epoch) << timestampLeftShift) |
		(s.datacenterID << datacenterIDShift) |
		(s.workerID << workerIDShift) |
		s.sequence

	return id
}

func GenerateSnowFlake(workerID, datacenterID int64) (int64, error) {
	snowflake, err := NewSnowflake(workerID, datacenterID)
	if err != nil {
		return 0, err
	}
	return snowflake.GenerateSnowFlake(), nil
}
