package conf

import (
	"testing"
)

func Testyaml(t *testing.T) {
	_ = LoadLocales("locales/zh-cn.yaml")
	t.Logf("%v", Dictinary)
}
