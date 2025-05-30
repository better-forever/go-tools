package easyid

import (
	"fmt"
	"strconv"
	"time"
)

// EasyId  生成唯一ID，16-20位
type EasyId struct {
	Id   int64 `json:"id"`
	Sign int   `json:"sign"` // 标志位，用于区分不同的业务
}

// new a easyid object, sign is a section number, recommend 1-9999
func New(sign int) EasyId {
	return EasyId{
		Sign: sign,
	}
}

// the style of id is 16-20 bit, the first 6 bit is date([year][month][day]), the next 2-4 bit is sign, the last 8 bit is the last 8 bit of timestamp
func (easyid *EasyId) GenId() *EasyId {
	now := time.Now()
	dateStr := now.Local().Format("20060102")
	timestampStr := fmt.Sprintf("%d", now.UnixMilli())
	var signStr string
	if easyid.Sign < 10 {
		signStr = fmt.Sprintf("0%d", easyid.Sign)
	} else {
		signStr = fmt.Sprintf("%d", easyid.Sign)
	}

	idStr := fmt.Sprintf(
		"%s%s%s",
		dateStr[2:],
		signStr,
		timestampStr[len(timestampStr)-8:],
	)

	id, _ := strconv.ParseInt(idStr, 10, 64)
	return &EasyId{
		Id:   id,
		Sign: easyid.Sign,
	}
}

func (easyid *EasyId) Int64() int64 {
	return easyid.Id
}

func (easyid *EasyId) String() string {
	return fmt.Sprintf("%d", easyid.Id)
}
