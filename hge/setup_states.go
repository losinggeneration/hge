package hge

import (
	"log"
	"os"
)

// TODO the log file likely needs close called on it at some point
func setupLogfile() (*log.Logger, error) {
	file, err := os.Create(stateStrings[LOGFILE])
	if err != nil {
		return nil, err
	}
	return log.New(file, "<< ", log.LstdFlags), nil
}
