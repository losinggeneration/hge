package hge

import (
	"log"
	"os"
)

// TODO the log file likely needs close called on it at some point
func (h *HGE) setupLogfile() (*log.Logger, error) {
	file, err := os.Create(stateStrings[LOGFILE])
	if err != nil {
		return nil, h.lastError(err)
	}
	return log.New(file, "<< ", log.LstdFlags), nil
}
