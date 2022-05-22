package impl

import (
	"fmt"
	"gits/utils"
	"strings"
)

func Env(isShow bool, repo []string) {
	if len(repo) != 0 {
		if len(repo) == 1 {
			env := utils.CurEnv()
			if env == nil {
				fmt.Println("Not found file.")
				return
			}
			urls := env.GetStringMap("url")
			for k, v := range urls {
				AddRepo(repo[0], strings.Replace(v.(string), k, repo[0], -1))
				break
			}
		} else {
			AddRepo(repo[0], repo[1])
		}
		return
	}
	ShowCurEnv()
}

func ShowCurEnv() {
	env := utils.CurEnv()
	if env == nil {
		fmt.Println("Not found file.")
		return
	}
	urls := env.GetStringMap("url")
	for k, v := range urls {
		fmt.Println(k, v.(string))
	}
}

func AddRepo(alias, url string) {
	env := utils.CurEnv()
	if env == nil {
		fmt.Println("Not found file.")
		return
	}
	env.Set("url."+alias, url)
	env.WriteConfig()
}

func RemoveRepo() {

}
