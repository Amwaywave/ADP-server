package pkg

import "github.com/eightpigs/i18n"

func init() {
	_, e := i18n.NewLocale("zh-CN", "assets/locales/zh-CN.yaml")
	if e != nil {
		panic(e)
	}
}
