package main

import (
	"fmt"
	"net"
	"net/http"
	"regexp"
	"time"
)

func check(destination string, port string) string {
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

func Add(domain string) error {
	validDomain := regexp.MustCompile(`^([a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9](\.[a-zA-Z0-9-]{1,61})*)?$`)
	if !validDomain.MatchString(domain) {
		return fmt.Errorf("invalid domain format %s", domain)
	}

	for _, existingWebsite := range trackedWebsites {
		if existingWebsite == domain {
			return fmt.Errorf("website '%s' already exists", domain)
		}
	}

	trackedWebsites = append(trackedWebsites, domain )
	err := writeTrackedWebsitesToFile()
	if err != nil {
		return err
	}

	fmt.Printf("Added website %s\n", domain)
	return nil
}

func ListTrackedWebsites () {
	fmt.Println("Tracked websites: ")
	if len(trackedWebsites) == 0 {
		fmt.Println("No websites are currently being tracked")
	}
	for _, website := range trackedWebsites {
		fmt.Println(website)
	}
} 

func Delete (domainToDelete string) {
	for i, website := range trackedWebsites {
		if website == domainToDelete {		
			trackedWebsites = append(trackedWebsites[:i], trackedWebsites[i+1:]... )
			fmt.Printf("Deleted website '%s'\n", domainToDelete)
			err := writeTrackedWebsitesToFile()
			 
			if err != nil {
				fmt.Println("Error deleting website: ", err)
			}

			return
		}
	}
	fmt.Printf("Domain %s not found in tracked list of websites\n", domainToDelete)
}

func checkPeriodically(port string, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
  initialCheck(port)

	for{
		select {
		case <-ticker.C:
			fmt.Println("Checking status of all tracked websites...")
			for _, website := range trackedWebsites {
				status := check(website, port)
				fmt.Printf("%s: %s\n", website, status)
			}
		}
	}
}

// first check before interval begins on periodic checks
func initialCheck(port string) {
  	fmt.Println("Checking status of all tracked websites...")
	for _, website := range trackedWebsites {
		status := check(website, port)
		fmt.Printf("%s: %s\n", website, status)
	}
}