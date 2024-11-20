package xoe

import (
	"encoding/json"
	"strings"
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

	Arg struct {
		Reporter  Reporter
		Logger    Logger
		Tag       string
		Msg       string
		Data      Data
		WithStack bool
		Stack     Stack
		StackSkip int
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
	if a.WithStack {
		a.Stack = NewStack(4 + a.StackSkip)
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
