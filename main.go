package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/urfave/cli/v2"
)
var trackedWebsites []string
var dataFileName = "tracked_websites.json"

func main() {
	err := readTrackedWebsitesFromFile()
	if err != nil {
		log.Println("Error loading tracked websites: ", err)
	}

  app := &cli.App{
    Name:  "HealthChecker",
    Usage: "A tool for checking and managing website health",
    Flags: []cli.Flag{
      &cli.StringFlag{
        Name:     "domain",
        Aliases: []string{"d"},
        Usage:    "Domain name to check",
        // Required: true,
      },
      &cli.StringFlag{
        Name:     "port",
        Aliases: []string{"p"},
        Usage:    "Port number to check (default: 80)",
        Value:    "80",
      },
      &cli.StringFlag{
        Name:     "add",
        Aliases: []string{"a"},
        Usage:    "Add a website to track its health",
				// Required: true,
      },
      &cli.BoolFlag{
        Name:     "list",
        Aliases: []string{"l"},
        Usage:    "List tracked websites",
      },
      &cli.StringFlag{
        Name:     "delete",
        Aliases: []string{"del"},
        Usage:    "Delete a tracked website",
      },
    },
    Action: func(c *cli.Context) error {

      domain := c.String("domain")
      port := c.String("port")

      if c.String("add") != "" {
        // Add site
				err := Add(domain)
				if err != nil {
					return err
				}
        fmt.Printf("01 Adding website: %s \n", domain)
        return nil
      } else if c.IsSet("list") {
        //list tracked websites
				ListTrackedWebsites()
        fmt.Println("Listing tracked websites...")
        return nil
      } else if c.String("delete") != "" {
        // delete website
				Delete(domain)
        fmt.Printf("Deleting website: %s\n", domain)
        return nil
      } else {
        // Perform health check
        status := Check(domain, port)
        fmt.Println(status)
        return nil
      }
    },
  }

  err = app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}

func Add(domain string) error {
	validDomain := regexp.MustCompile(`^([a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9](\.[a-zA-Z0-9-]{1,61})*)?$`)
	if !validDomain.MatchString(domain) {
		return fmt.Errorf("Invalid domain format %s", domain)
	}

	for _, existingWebsite := range trackedWebsites {
		if existingWebsite == domain {
			return fmt.Errorf("Website '%s' already exists", domain)
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
			fmt.Printf("Deleted website %s\n", domainToDelete)
			err := writeTrackedWebsitesToFile()
			 
			if err != nil {
				fmt.Println("Error deleting website: ", err)
			}

			return
		}
	}
	fmt.Printf("Domain %s not found in tracked list of websites\n", domainToDelete)
}

func readTrackedWebsitesFromFile() error {
	file, err := os.Open(dataFileName)
	if err != nil {
		if os.IsNotExist(err) {
			trackedWebsites = []string{}
			return nil
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&trackedWebsites)
	if err != nil {
		return err
	}
	return nil

}

func writeTrackedWebsitesToFile() error {
	file, err := os.Create(dataFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// data, err := json.Marshal(trackedWebsites)
	// if err != nil {
	// 	return err
	// }
	encoder := json.NewEncoder(file)
	err = encoder.Encode(trackedWebsites)
	if err != nil {
		return err
	}

	// _, err = file.Write(data)
	// if err != nil {
	// 	return err
	// }
	return nil
}