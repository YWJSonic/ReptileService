package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/YWJSonic/ReptileService/TWSEcom/stockday"
)

func TestDo(t *testing.T) {
	testcase := map[string]int{
		"getDelayTime": 3, // 請求延遲時間 /秒
		"startTime":    0, // 開始時間
		"endTime":      0, // 結束時間
		"successTime":  0, //成功次數
	}
	isEnd := false

	cacheTime := time.Now().Unix() * 1000
	testcase["startTime"] = int(time.Now().Unix())
	for {
		_, err := stockday.Get("0050", "20210601", cacheTime)
		if err != nil {
			fmt.Println("Error:", err)
			testcase["endTime"] = int(time.Now().Unix())
			isEnd = true
		}
		fmt.Println("Done")

		if isEnd {
			break
		} else {
			testcase["successTime"] += 1
			time.Sleep(time.Duration(testcase["getDelayTime"]) * time.Second)
		}
	}

	fmt.Println("startTime:", testcase["startTime"])
	fmt.Println("endTime:", testcase["endTime"])
	fmt.Println("getDelayTime:", testcase["getDelayTime"])
	fmt.Println("successTime:", testcase["successTime"])
}

func TestBlockTime(t *testing.T) {
	testcase := map[string]int{
		"getDelayTime": 3, // 請求延遲時間 /秒
		"startTime":    0, // 開始時間
		"endTime":      0, // 結束時間
		"successTime":  0, //成功次數
	}

	cacheTime := time.Now().Unix() * 1000
	testcase["startTime"] = int(time.Now().Unix())
	for {
		_, err := stockday.Get("0050", "20210601", cacheTime)
		if err != nil {
			fmt.Println("Error:", err)
			time.Sleep(3 * time.Second)
			continue
		}
		testcase["endTime"] = int(time.Now().Unix())
		fmt.Println("Done")
		break
	}

	fmt.Println("startTime:", testcase["startTime"])
	fmt.Println("endTime:", testcase["endTime"])
	fmt.Println("getDelayTime:", testcase["getDelayTime"])
	fmt.Println("successTime:", testcase["successTime"])
}
