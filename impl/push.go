package impl

import (
	"fmt"
	"gits/utils"
	"sync"
)

func GitPush(isAll bool, redirect string, args ...string) {
	urls := utils.CurEnv().GetStringMap("url")
	if isAll {
		var wg sync.WaitGroup
		for k, v := range urls {
			wg.Add(1)
			var fargs []string
			fargs = append(fargs, "add")
			fargs = append(fargs, k)
			fargs = append(fargs, v.(string))
			utils.Git(utils.GetArgs("remote", fargs...)...)
			go func(alias string) {
				for {
					var fargs []string
					fargs = append(fargs, "-u")
					fargs = append(fargs, alias)
					fargs = append(fargs, args...)
					stderr := utils.Git(utils.GetArgs("push", fargs...)...)
					if stderr == "" {
						break
					}
					fmt.Println("reload...")
				}
				wg.Done()
			}(k)
		}
		wg.Wait()
	}
}
