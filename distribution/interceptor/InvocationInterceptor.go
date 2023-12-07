package interceptor

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/miop"
)

type InvocationInterceptor struct {
	requestTime time.Time
}

func NewInvocationInterceptor() InvocationInterceptor {
	return InvocationInterceptor{}
}

func (i *InvocationInterceptor) Intercept(invocation miop.Packet, isRequest bool) {
	if isRequest {
		i.requestTime = time.Now()
		fmt.Println("Intercepting request")
	} else {
		elapsedTime := time.Since(i.requestTime)

		// Calculate estimated of the memory payload sizes and other metrics
		headerSize := unsafe.Sizeof(invocation.Header)
		reqHeaderSize := unsafe.Sizeof(invocation.Body.ReqHeader)
		reqBodySize := unsafe.Sizeof(invocation.Body.ReqBody)
		repHeaderSize := unsafe.Sizeof(invocation.Body.RepHeader)
		repBodySize := unsafe.Sizeof(invocation.Body.RepBody)

		// Get response code/status
		responseStatus := invocation.Body.RepHeader.Status

		// Print or log the elapsed time and other metrics
		fmt.Printf("Elapsed Time: %v\n", elapsedTime)
		fmt.Printf("Header Size: %d bytes\n", headerSize)
		fmt.Printf("Request Header Size: %d bytes\n", reqHeaderSize)
		fmt.Printf("Request Body Size: %d bytes\n", reqBodySize)
		fmt.Printf("Reply Header Size: %d bytes\n", repHeaderSize)
		fmt.Printf("Reply Body Size: %d bytes\n", repBodySize)
		fmt.Printf("Response Status: %d\n", responseStatus)
	}
}
