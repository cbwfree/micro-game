package mongo

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
)

const (
	defaultScanSize = 20
	defaultScanCur  = 1
)

type Scan struct {
	Page   int64 `json:"page"`  // 当前页数
	Count  int64 `json:"count"` // 总数量
	Size   int64 `json:"size"`  // 每页大小
	Offset int64 `json:"-"`     // 跳过
}

func (sc *Scan) calc() {
	if sc.Size <= 0 {
		sc.Size = defaultScanSize
	}

	// 计算最大页数
	total := int64(math.Ceil(float64(sc.Count) / float64(sc.Size)))

	if sc.Page <= 0 {
		sc.Page = defaultScanCur
	}
	if sc.Page > total {
		sc.Page = total
	}

	sc.Offset = (sc.Page - 1) * sc.Size
}

func (sc *Scan) FindOptions() *options.FindOptions {
	return options.Find().SetSkip(sc.Offset).SetLimit(sc.Size)
}

func NewScan(count, page, size int64) *Scan {
	ms := &Scan{Count: count, Page: page, Size: size}
	ms.calc()
	return ms
}
