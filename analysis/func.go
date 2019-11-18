package analysis

import (
	"github.com/YWJSonic/ReptileService/analysis/analysisday"
	"github.com/YWJSonic/ReptileService/routineswitch"
)

// GetAnalysisManager Get
func GetAnalysisManager() *Manager {
	if analysisManager == nil {
		manager := &Manager{}
		manager.PriceDetails = make(map[string]*analysisday.Info)
		manager.routingswitchs = make(map[string]*routineswitch.Info)
		analysisManager = manager
	}

	return analysisManager
}
