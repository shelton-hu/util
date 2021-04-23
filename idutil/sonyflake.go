package idutil

import (
	"sync"
	"time"

	"github.com/sony/sonyflake"
)

var sfs = new(sync.Map)

var SugStartTIme, _ = time.Parse(time.RFC3339, "2020-01-01T00:00:00Z00:00")

type SonyFlake struct {
	sf *sonyflake.Sonyflake
}

func NewSnowFlake(machineId uint16) *SonyFlake {
	if machineId >= 1024 {
		machineId = 1023
	}
	sf, _ := sfs.LoadOrStore(machineId, &SonyFlake{
		sf: sonyflake.NewSonyflake(sonyflake.Settings{
			StartTime: SugStartTIme,
			MachineID: func() (uint16, error) {
				return machineId, nil
			},
		}),
	})
	return sf.(*SonyFlake)
}

func (s *SonyFlake) NextID() uint64 {
	id, _ := s.sf.NextID()
	return id
}
