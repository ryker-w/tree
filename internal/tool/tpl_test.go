package tool

import (
	"errors"
	"testing"
	"time"
)

func TestTplResolve(t *testing.T) {
	var tpl = "aaa/bbb/{para}/d"
	var topic = "aaa/bbb/cccc/d"
	m, err := TopicResolver(tpl, topic)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(m)
}

func TestTplBuilder(t *testing.T) {
	var format = "aaa/bbb/%s/d/%s/m"
	var key = "Device"
	var key2 = "gateway"
	var tpl = TopicBuilder(format, key, key2)
	t.Log(tpl)
	if tpl != "aaa/bbb/{Device}/d/{gateway}/m" {
		t.Fatal(errors.New(tpl))
	}
}

func TestTime(t *testing.T) {
	var s = time.Now().Format("2006-002")
	t.Log(s)
	t.Logf("%d", time.Now().Weekday())
}
