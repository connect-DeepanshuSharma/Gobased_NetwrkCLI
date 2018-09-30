package main
import (
  "fmt"
  "log"
	"net"
  "os"
"github.com/urfave/cli"
)



func main() {
	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = "Let's you query IPs, CNAMEs, MX records and Name Servers!"

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "www.facebook.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ns",
			Usage: "Looks Up the NameServers for a Particular Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
    },
    {
      Name: "ip",
      Usage: "lets you look up the ip address of the urls",
      Flags: myFlags,
      Action : func(c *cli.Context)error{
        ip,err := net.LookupIP(c.String("host"))
        if err != nil{
          return err
        }
        for i:=0;i<len(ip);i++{
          fmt.Println(ip[i])
        }
        return nil
      },
    },
    {
      Name: "cname",
      Usage: "Lets u get around the cname of the url",
      Flags: myFlags,
      Action: func(c *cli.Context)error{
        cn,err:= net.LookupCNAME(c.String("host"))
        if err != nil{
          return err
        }
        for i:=0;i<len(cn);i++{
          fmt.Println(cn[i])
        }
        return nil
      },
    },
    {
      Name:"mx",
      Usage:"let you go around the mx part of the network",
      Flags:myFlags,
      Action:func(c *cli.Context)error{
        mx,err:=net.LookupMX(c.String("host"))
        if err!=nil{
          return err
        }
        for i:=0;i<len(mx);i++{
          fmt.Println(mx[i])
        }
        return nil
      },
    },
  }

      err := app.Run(os.Args)
      	if err != nil {
      		log.Fatal(err)
      	}

}
