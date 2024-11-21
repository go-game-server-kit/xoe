package xoe

import (
	"fmt"
	"os"
	"testing"
)

func TestDingTalk(t *testing.T) {
	SetReporter(NewDingTalkReporter("https://oapi.dingtalk.com/robot/send?access_token=faa31de7e160caed7662ca6bbb74f86739746b033aa6201f2ca26f279f9e201d", "通知告警"))
	//Roe(fmt.Errorf("error1"))
	Roe(fmt.Errorf("error1"), WithData(Data{"key": "value"}), WithMsg("msg"))
}

func TestDash3(t *testing.T) {
	url := os.Getenv("DASH3_URL")
	fmt.Println("url:", url)
	r := NewDash3Reporter(url, "test", "test", "test", "test", "test")
	SetReporter(r)
	Roe(fmt.Errorf("error1"), WithData(Data{"key": "value"}), WithMsg("msg"))
}
