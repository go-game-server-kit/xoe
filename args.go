package xoe

import (
	"encoding/json"
	"github.com/mohuani/notify/dingtalk/message"
	"strings"
)

const (
	// StrategyLinearBackoff
	// 0: Linear Backoff: First send, next send interval is 60 seconds, the send interval after that is 300 seconds, and so on, up to a maximum of 600 seconds.
	// 0：线性退避：首次发送，下次发送间隔60秒，下下次发送间隔300秒，依次类推最大600
	StrategyLinearBackoff = 0
	// StrategyRateLimiting
	// 1: Rate Limiting: Default notification is sent after exceeding 10 times in 60 seconds.
	// 1：速率限制：默认60秒超过10次才通知
	StrategyRateLimiting = 1
	// StrategyImmediate
	// 2: Immediate Notification: Notify every time.
	// 2：即时通知：每次都通知
	StrategyImmediate = 2
	// StrategyOneTime
	// 3: One-Time Notification: Do not notify again if unresolved.
	// 3：一次性通知：未解决则不通知
	StrategyOneTime = 3
)

const (
	// ActionResolve 通知需要确认解决
	ActionResolve = "resolve"
	// ActionNotice 通知无需任何操作
	ActionNotice = "notice"
)

var (
	DefaultWithStack = true
)

type (
	Data map[string]any

	Args interface {
		Apply(arg *Arg)
	}
	ReporterArgs struct {
		reporter Reporter
	}
	LoggerArgs struct {
		logger Logger
	}
	DataArgs struct {
		data Data
	}
	StackArgs struct {
	}
	StackSkipArgs struct {
		skip int
	}
	MsgArgs struct {
		msg string
	}
	AtArgs struct {
		at *message.At
	}
	StrategyArgs struct {
		strategy int
	}
	ActionArgs struct {
		action string
	}

	Arg struct {
		Reporter  Reporter
		Logger    Logger
		Tag       string
		Msg       string
		Data      Data
		WithStack bool
		Stack     Stack
		StackSkip int
		At        *message.At
		Strategy  int
		Action    string
	}
)

func (d Data) String() string {
	b, _ := json.Marshal(d)
	return string(b)
}

func (r *ReporterArgs) Apply(arg *Arg) {
	arg.Reporter = r.reporter
}

func WithReporter(reporter Reporter) Args {
	return &ReporterArgs{
		reporter: reporter,
	}
}

func (l *LoggerArgs) Apply(arg *Arg) {
	arg.Logger = l.logger
}

func WithLogger(logger Logger) Args {
	return &LoggerArgs{
		logger: logger,
	}
}

func (d *DataArgs) Apply(arg *Arg) {
	if arg.Data == nil {
		arg.Data = make(Data)
	}
	for k, v := range d.data {
		arg.Data[k] = v
	}
}

func WithData(data Data) Args {
	return &DataArgs{
		data: data,
	}
}

func WithKV(k string, v any) Args {
	return &DataArgs{
		data: Data{k: v},
	}
}

func (s *StackArgs) Apply(arg *Arg) {
	arg.WithStack = true
}

func WithStack() Args {
	return &StackArgs{}
}

func (s *StackSkipArgs) Apply(arg *Arg) {
	arg.StackSkip = s.skip
}

func WithStackSkip(skip int) Args {
	return &StackSkipArgs{
		skip: skip,
	}
}

func (m *MsgArgs) Apply(arg *Arg) {
	arg.Msg = m.msg
}

func WithMsg(msg string) Args {
	return &MsgArgs{
		msg: msg,
	}
}

func (a *AtArgs) Apply(arg *Arg) {
	arg.At = a.at
}

func WithAt(at *message.At) Args {
	return &AtArgs{
		at: at,
	}
}

func (s *StrategyArgs) Apply(arg *Arg) {
	arg.Strategy = s.strategy
}

func WithStrategy(strategy int) Args {
	return &StrategyArgs{
		strategy: strategy,
	}
}

func (a *ActionArgs) Apply(arg *Arg) {
	arg.Action = a.action
}

func WithAction(action string) Args {
	return &ActionArgs{
		action: action,
	}
}

func NewArg(tag string, args []Args) *Arg {
	a := &Arg{Tag: tag}
	for _, arg := range args {
		arg.Apply(a)
	}
	if a.Reporter == nil {
		a.Reporter = reporter
	}
	if a.Logger == nil {
		a.Logger = logger
	}
	if a.WithStack || DefaultWithStack {
		a.Stack = NewStack(4 + a.StackSkip)
	}
	if a.At == nil {
		a.At = &message.At{}
	}
	if a.Action == "" {
		a.Action = ActionResolve
	}
	return a
}

func (a *Arg) String(err error) string {
	var str []string
	if a.Tag != "" {
		str = append(str, "tag="+a.Tag)
	}
	if a.Msg != "" {
		str = append(str, "msg="+a.Msg)
	}
	if err != nil {
		str = append(str, "err="+err.Error())
	}
	if len(a.Data) > 0 {
		str = append(str, "data="+a.Data.String())
	}
	if len(a.Stack) > 0 {
		str = append(str, "\n\tstack:\n\t"+strings.ReplaceAll(a.Stack.String(), "\n", "\n\t"))
	}
	return strings.Join(str, " ")
}
