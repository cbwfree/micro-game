package debug

import (
	"fmt"
	"github.com/cbwfree/micro-game/utils/color"
	"github.com/micro/go-micro/v2/logger"
	"path/filepath"
	"runtime"
	"time"
)

// 性能分析
type Prof struct {
	label string
	log   *logger.Helper
	t     time.Time
}

func (p *Prof) Result() {
	var mstat runtime.MemStats
	runtime.ReadMemStats(&mstat)

	p.log.Debug(color.Question.Text("[Prof]%s: Uptime: %s, Memory: %d, Threads: %d, GC: %d",
		p.label, time.Since(p.t), mstat.Alloc, runtime.NumGoroutine(), mstat.PauseTotalNs))
}

func NewProf(log *logger.Helper, label ...string) *Prof {
	p := &Prof{
		log: log,
		t:   time.Now(),
	}
	if len(label) > 0 {
		p.label = label[0]
	} else {
		caller := GetCaller(3)
		p.label = fmt.Sprintf("%s:%d", filepath.Base(caller.File), caller.Line)
	}
	return p
}
