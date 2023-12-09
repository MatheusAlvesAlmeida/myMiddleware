package interceptor

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	"unsafe"

	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/miop"
)

type InvocationInterceptor struct {
	requestTime          time.Time
	marshallingStartTime time.Time
	marshallingEndTime   time.Time
	requestContext       string
}

const timestampFormat = "2006-01-02 15:04:05"

func NewInvocationInterceptor() InvocationInterceptor {
	return InvocationInterceptor{}
}

func (i *InvocationInterceptor) Intercept(invocation miop.Packet, isRequest bool, getMarshallingTime bool) {
	context := invocation.Body.ReqHeader.Context
	// If context is not PercentageOperation, do not log
	if context == "NamingOperation" {
		return
	}
	filename := "logs.txt"
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return
	}

	// Check if the file exists, create it if not
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		if _, err := os.Create(absPath); err != nil {
			return
		}
	}

	file, err := os.OpenFile(absPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	var line string

	if isRequest && !getMarshallingTime {
		i.requestTime = time.Now()
		i.requestContext = context
		i.marshallingStartTime = time.Now()
		return
	} else if getMarshallingTime {
		i.marshallingEndTime = time.Now()
	} else {
		elapsedTime := time.Since(i.requestTime)

		// Calculate estimated of the memory payload sizes and other metrics
		headerSize := unsafe.Sizeof(invocation.Header)
		reqHeaderSize := unsafe.Sizeof(invocation.Body.ReqHeader)
		reqBodySize := unsafe.Sizeof(invocation.Body.ReqBody)
		repHeaderSize := unsafe.Sizeof(invocation.Body.RepHeader)
		repBodySize := unsafe.Sizeof(invocation.Body.RepBody)

		responseStatus := invocation.Body.RepHeader.Status
		responseContext := invocation.Body.RepHeader.Context
		marshallingTime := i.marshallingEndTime.Sub(i.marshallingStartTime)

		line = fmt.Sprintf("%v,%s,%v,%d,%d,%d,%d,%d,%d,%s,%d\n", i.requestTime.Format(timestampFormat), i.requestContext, elapsedTime, headerSize, reqHeaderSize, reqBodySize, repHeaderSize, repBodySize, responseStatus, responseContext, marshallingTime)
	}

	_, err = file.WriteString(line)
	if err != nil {
		return
	}
}
