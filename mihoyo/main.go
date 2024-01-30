package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("====== HoYoLab Auto Sign ======")
	Lang, Zone = LoadLangZone()
	//fmt.Println(Lang, Zone)
	cookie := LoadCookie()
	//fmt.Println(cookie[:10])
	ids := LoadActIds()
	//fmt.Println(ids)
	SetInterval(Interval, func() error {
		info := GetUserInfo(cookie)
		if info == nil {
			return errors.New("cookie has expired")
		}
		HoyoLabSign(cookie, ids, info)
		return nil
	})
}
