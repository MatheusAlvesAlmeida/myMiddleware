package errorhandler

import (
	"errors"
	"net"
	"time"
)

type ErrorHandler struct {
	RetryAttempts int
}

func (eh *ErrorHandler) HandleError(conn *net.Conn, err error) error {
	if err == nil {
		return nil
	}

	if conn != nil {
		_ = (*conn).Close()
		*conn = nil
	}

	netErr, ok := err.(net.Error)
	if ok && netErr.Timeout() {
		if eh.RetryAttempts < 5 {
			backoff := eh.backoffDuration()
			time.Sleep(backoff)
			eh.RetryAttempts++
			return errors.New("connection timeout: disconnected")
		}
	}

	return err
}

func (eh *ErrorHandler) backoffDuration() time.Duration {
	// Exponential backoff formula: 5 * 2^(retryAttempts-1) seconds
	return time.Duration(5*(1<<uint(eh.RetryAttempts-1))) * time.Second
}
