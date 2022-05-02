package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type event struct {
	URL   string `json:"url"`
	Param string `json:"param"`
}

func handler(e event) (*string, error) {
	u, err := url.Parse(e.URL)
	if err != nil {
		return nil, nil
	}

	val := u.Query().Get(e.Param)
	if len(val) == 0 {
		return nil, nil
	}

	return &val, nil
}

func main() {
	if len(os.Getenv("AWS_EXECUTION_ENV")) == 0 {
		url := os.Args[1]
		parameter := os.Args[2]
		p, err := handler(event{
			URL:   url,
			Param: parameter,
		})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(p)
		}
		return
	}

	lambda.Start(handler)
}
