package pkg

import "github.com/eightpigs/i18n"

func init() {
	_, e := i18n.New()
	if e != nil {
		panic(e)
	}
}
