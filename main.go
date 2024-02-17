package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)
var trackedWebsites []string
var dataFileName = "tracked_websites.json"

func main() {


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
				Add(domain)
        fmt.Printf("Adding website: %s:%s\n", domain, port)
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

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}

func Add(domain string) {
	trackedWebsites = append(trackedWebsites, domain )
	fmt.Println("Added website %s\n", domain)
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
			return
		}
	}
	fmt.Printf("Domain %s not found in tracked list of websites\n", domainToDelete)
}

