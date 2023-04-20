package main

import (
	"fmt"
	"git.woa.com/tme/international/infra/server/pkg/util/qua"
	"github.com/spf13/cast"
)

const (
	qua1 = "V1_AND_WESING_5.61.7_733_GOOGLEPLAY_A"
	qua2 = "V1_IPH_WESING_5.61.1_527_APPSTORE_A"
)

func main() {
	q, err := qua.Parse(qua1)
	if err != nil {
		fmt.Printf("parse qua error: %s qua: %s\n", err, qua1)
	} else {
		fmt.Printf("platform: %d\n", q.Platform)
		platformStr, errr := cast.ToStringE(q.Platform)
		if errr != nil {
			fmt.Printf("cast.ToStringE err: %s\n", errr)
			return
		}
		fmt.Printf("platform: %s\n", platformStr)
	}
}
