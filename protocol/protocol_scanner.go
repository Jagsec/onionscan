package protocol

import (
	"github.com/Jagsec/onionscan/config"
	"github.com/Jagsec/onionscan/report"
)

type Scanner interface {
	ScanProtocol(hiddenService string, onionscanConfig *config.OnionScanConfig, report *report.OnionScanReport)
}
