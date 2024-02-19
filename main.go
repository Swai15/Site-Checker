package main

import (
	"fmt"
	"log"
	"os"
	"time"

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
			&cli.BoolFlag{
				Name: "checkAll",
				Aliases: []string{"ca"},
				Usage: "Check status of all tracked websites",
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
        Usage:    "Add a domain to be stored in a list of tracked websites",
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
      &cli.DurationFlag{
        Name: "interval",
        Aliases: []string{"i"},
        Usage: "Set Interval for automatically checking tracked websites",
        Value: 5 * time.Minute,
      },
    },
    Action: func(c *cli.Context) error {

      port := c.String("port")

      if c.IsSet("add") {
        // Add to tracked websites
				addDomain := c.String("add")
				err :=  Add(addDomain)
				if err != nil {
					return err
				}
        return nil
      } else if c.IsSet("list") {
        //list tracked websites
				ListTrackedWebsites()
        return nil
      } else if c.IsSet("delete"){
        // delete website
				deleteDomain := c.String("delete")
				Delete(deleteDomain)
        return nil
      } else if c.IsSet("checkAll"){
				fmt.Println("Checking status of all tracked websites")
				for _, website := range trackedWebsites {
					status := check(website, port)
					fmt.Printf("%s: %s\n", website, status)
				}
				return nil
			} else if c.IsSet("interval") {
       interval := c.Duration("interval") 
       fmt.Printf("Starting periodic checking at %v intervals\n", interval)
       go checkPeriodically(port, interval)
       select{}
      //  return nil
      } else {
        // Perform check on single site
				checkDomain := c.String("domain")
        status := check(checkDomain, port)
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


