package errorhandler

import (
	"errors"
	"fmt"
	"math/rand"
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
		eh.RetryAttempts++
		backoff := eh.backoffDuration()
		fmt.Println("Retrying connection in ", backoff)
		time.Sleep(backoff)
		return errors.New("connection timeout: retrying")
	}

	return err
}

func (eh *ErrorHandler) backoffDuration() time.Duration {
	// Calculate the backoff duration exponentially based on the retry attempt number
	retryAttemptNumber := eh.RetryAttempts
	backoff := time.Duration(1<<uint(retryAttemptNumber)) * time.Second

	// Add some jitter to avoid synchronized retries in case of multiple instances
	random := rand.Intn(100) // Generates a random number between 0 and 99
	backoffWithJitter := time.Duration(random)*time.Millisecond + backoff

	return backoffWithJitter
}
