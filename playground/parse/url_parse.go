package main

import (
	"fmt"
	"net/url"
)

const (
	schemeKeyRoomType = "roomtype"
	schemeKeyRoomID   = "roomid"
)

func main() {

	scheme := "wesing://kege.com?action=ktvroom&roomid=zora1&roomid=zora2&roomtype=ktv&need_redirect=true"

	// 1. 解析url
	raw, params, err := parseURL(scheme)
	if err != nil {
		fmt.Printf("parseURL err: %s\n", err)
	}
	fmt.Printf("roomIDs: %+v\n", params["roomid"])
	fmt.Printf("roomType: %+v\n", params["roomtype"])

	roomID := "zora3"
	// 设置roomid
	params.Set(schemeKeyRoomID, roomID)
	params.Set(schemeKeyRoomType, params[schemeKeyRoomType][0])
	raw.RawQuery = params.Encode()
	fmt.Printf("raw: %s\n", raw.String())
}

// 解析url
func parseURL(scheme string) (*url.URL, url.Values, error) {
	// url解析
	raw, err := url.Parse(scheme)
	if err != nil {
		fmt.Printf("url Parse err: %s\n", err)
		return nil, nil, err
	}

	// 从RawQuery解析参数
	params, err := url.ParseQuery(raw.RawQuery)
	if err != nil {
		fmt.Printf("url ParseQuery err: %s\n", err)
		return nil, nil, err
	}

	return raw, params, nil
}
