package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

var Version = "v0.0.0"

func main() {
	version := flag.Bool("v", false, "prints current version")
	flag.Parse()
	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}
	if len(os.Args) < 4 || len(os.Args) > 5 {
		fmt.Fprintf(os.Stderr, "Usage: %s zone-id domain ip [ttl]\n", os.Args[0])
		os.Exit(1)
	}
	var ttl int64 = 1
	var err error
	if len(os.Args) == 5 {
		ttl, err = strconv.ParseInt(os.Args[4], 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Usage: %s zone-id domain ip [ttl]\n", os.Args[0])
			os.Exit(1)
		}
	}
	s := route53.New(session.New())
	_, err = s.ChangeResourceRecordSets(&route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String(route53.ChangeActionUpsert),
					ResourceRecordSet: &route53.ResourceRecordSet{
						Name: aws.String(os.Args[2]),
						Type: aws.String(route53.RRTypeA),
						TTL:  aws.Int64(ttl),
						ResourceRecords: []*route53.ResourceRecord{
							{Value: aws.String(os.Args[3])},
						},
					},
				},
			},
		},
		HostedZoneId: aws.String(os.Args[1]),
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
