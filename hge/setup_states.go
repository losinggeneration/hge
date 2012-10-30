package hge

import (
	"log"
	"os"
)

func (h *HGE) setupTitle() error {
	setTitle()
	return nil
}

// TODO the log file likely needs close called on it at some point
func (h *HGE) setupLogfile() error {
	file, err := os.Create(stateStrings[LOGFILE])
	if err != nil {
		return h.postError(err)
	}
	h.log = log.New(file, "<< ", log.LstdFlags)
	return nil
}
