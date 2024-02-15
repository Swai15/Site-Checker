package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)



func Check(destination string, port string) string {
	address := destination + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)
	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable, \n Error: %v", destination, err)
	} else {
		conn.Close()
		status = fmt.Sprintf("[UP] %v is reachable", destination)

		resp, err := http.Get("http://" + destination)
		if err != nil {
			status += fmt.Sprintf(", but unable to retrieve status code: %v", err)
		} else {
			defer resp.Body.Close()
			statusCode := resp.StatusCode

			if statusCode != http.StatusOK {
				status += fmt.Sprintf(", but returned status code %d", statusCode)
			}

		}

	}
	return status
}
