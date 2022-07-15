package isnow

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

func InitializeSnow(startTime string, machineID int64) (node *snowflake.Node, err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return nil, err
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	return node, err
}
