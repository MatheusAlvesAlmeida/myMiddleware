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

func (eh *ErrorHandler) HandleConnectionError(conn *net.Conn, err error) error {
	if err == nil {
		return nil
	}

	if *conn != nil {
		_ = (*conn).Close()
		*conn = nil
	}

	if eh.RetryAttempts < 3 {
		backoff := eh.backoffDuration()
		time.Sleep(backoff)
		eh.RetryAttempts++
		return errors.New("connection timeout: retrying")
	}

	return err
}

func (eh *ErrorHandler) backoffDuration() time.Duration {
	retryAttemptNumber := eh.RetryAttempts
	return time.Duration(retryAttemptNumber) * time.Second
}
