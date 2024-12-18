package xoe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-game-server-kit/xoe/h"
	"github.com/mohuani/notify/dingtalk/client"
	"github.com/mohuani/notify/dingtalk/message"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type (
	Reporter func(err error, arg *Arg)
)

var (
	_reporter = log.New(os.Stderr, "[xoe:report] ", log.LstdFlags)
	reporter  = func(err error, arg *Arg) {
		_reporter.Println(arg.String(err))
	}
)

func SetReporter(r Reporter) {
	reporter = r
}

func Report(tag string, err error, args []Args) {
	arg := NewArg(tag, err, args)
	arg.Reporter(err, arg)
}

func NewDingTalkReporter(token, title string) Reporter {
	if strings.Index(token, "access_token=") > 0 {
		token = strings.Split(token, "access_token=")[1]
		if strings.Index(token, "&") > 0 {
			token = strings.Split(token, "&")[0]
		}
	}
	cli := client.NewDingTalkClient(token, title)
	return func(err error, arg *Arg) {
		title0 := "# " + title + "\n\n"
		text0 := ""
		if arg.Tag != "" {
			title0 += " " + arg.Tag
			text0 += "* tag: " + arg.Tag + "\n"
		}
		if arg.Msg != "" {
			title0 += " " + arg.Msg
			text0 += "* msg: " + arg.Msg + "\n"
		}
		if err != nil {
			title0 += " " + err.Error()
			text0 += "* err: " + err.Error() + "\n"
		}
		if len(arg.Data) > 0 {
			db, _ := json.MarshalIndent(arg.Data, "", "  ")
			ds := string(db)
			text0 += "\ndata:\n\n```\n" + ds + "\n```\n\n"
		}
		if len(arg.Stack) > 0 {
			text0 += "stack:\n\n```\n" + arg.Stack.String() + "\n```\n\n"
		}
		text0 = title0 + "\n\n" + text0
		log.Println("Send DingTalk:", title0, "\n", text0)
		Loe(cli.SendMarkDownMessage(message.Markdown{
			Title: title0,
			Text:  text0,
		}, *arg.At))
	}
}

func NewDash3Reporter(url, app, service, env, group, channel string) Reporter {
	type dashRes struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	return func(err error, arg *Arg) {
		data := Data{
			"app":     app,
			"service": service,
			"env":     env,
			"group":   group,
			"channel": channel,
			"message": arg.Tag + " " + arg.Msg + " " + err.Error(),
			"content": Data{
				"data":  arg.Data,
				"stack": ShortStack(arg.Stack.Strings()),
			},
			"strategy": arg.Strategy,
			"action":   arg.Action,
			"at_all":   h.Bool2Int(arg.At.IsAtAll),
			"at_user":  strings.Join(arg.At.AtMobiles, ","),
		}
		b, err := json.Marshal(data)
		if Loe(err, WithMsg("json marshal error")) {
			return
		}
		resp, err := http.Post(url, "application/json", bytes.NewReader(b))
		if Loe(err, WithMsg("http post error")) {
			return
		}
		if resp.StatusCode != http.StatusOK {
			Loe(fmt.Errorf("http post status code error"), WithData(Data{"status_code": resp.StatusCode}))
			return
		}
		body := &bytes.Buffer{}
		defer LoeFn(resp.Body.Close)
		if _, err := io.Copy(body, resp.Body); Loe(err, WithMsg("io copy error")) {
			return
		}
		rs := &dashRes{}
		bb := body.Bytes()
		if Loe(json.Unmarshal(bb, rs), WithMsg("json unmarshal error")) {
			return
		}
		if rs.Code != 0 {
			Loe(fmt.Errorf(rs.Msg), WithMsg("dash3 error"), WithData(Data{"rs": string(bb)}))
		}
	}
}
