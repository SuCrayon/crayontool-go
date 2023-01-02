package mongoapi

import "sync"

const (
	CmdPing = "ping"
)

type DiagnosticCommander interface {
	Ping() *PingReq
}

type diagnosticCommander struct {
	commander
}

var (
	dcGetOnce sync.Once
	dc        DiagnosticCommander
)

func GetDiagnosticCommander(ctl MongoCtl) DiagnosticCommander {
	if dc == nil {
		dcGetOnce.Do(func() {
			dc = &diagnosticCommander{
				commander: commander{
					ctl: ctl,
				},
			}
		})
	}
	return dc
}

func (d *diagnosticCommander) Ping() *PingReq {
	return &PingReq{
		iCommandReq: newDefaultCommandReq(d.ctl).setCommandStr(CmdPing),
	}
}
