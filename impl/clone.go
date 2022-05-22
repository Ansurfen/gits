package impl

import (
	"fmt"
	"gits/utils"
	"net/http"
	"os"
	"strings"
	"sync"
)

func GitClone(args []string) {
	if !isUrl(args[0]) {
		if strings.Index(args[0], "/") == -1 {
			fmt.Println("error")
			return
		}
		var wg sync.WaitGroup
		repo := []string{"gitee", "github", "gitlab"}
		tmpl := "https://{{repo}}.com/{{args[0]}}.git"
		wg.Add(len(repo))
		for i := 0; i < len(repo); i++ {
			repo[i] = strings.Replace(tmpl, "{{repo}}", repo[i], -1)
			repo[i] = strings.Replace(repo[i], "{{args[0]}}", args[0], -1)
			go func(i int) {
				res, err := http.Get(repo[i])
				if err != nil || res.StatusCode == 403 || res.StatusCode == 503 {
					repo = append(repo[:i], repo[i+1:]...)
					i--
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
		cnt := 0
		for {
			index := cnt % len(repo)
			stderr := utils.Git(utils.GetArgs("clone", repo[index])...)
			if stderr == "" || strings.Index(stderr, "already exists and is not an empty directory.") != -1 {
				first, second := strings.Index(repo[index], "https://"), strings.Index(repo[index], ".com/")
				reponame := repo[index][second+5 : len(repo[index])-4]
				envConf := utils.NewEnv(reponame)
				envConf.SetDefault("url."+repo[index][first+8:second], repo[index])
				dir := utils.GetRootDir()
				envConf.WriteConfigAs(dir + reponame + ".yml")
				path, err := os.Getwd()
				utils.Panic(err)
				path = path + "\\" + strings.Split(args[0], "/")[1]
				utils.RegisterIndex(path, dir+reponame)
				break
			}
			fmt.Printf("reload %d ...\n", cnt+1)
			cnt++
		}
	} else {
		cnt := 0
		for {
			stderr := utils.Git(utils.GetArgs("clone", args...)...)
			if stderr == "" || strings.Index(stderr, "already exists and is not an empty directory.") != -1 {
				break
			}
			fmt.Printf("reload %d ...\n", cnt+1)
			cnt++
		}
	}
}

func isUrl(url string) bool {
	if strings.Index(url, "https://") != -1 {
		return true
	}
	return false
}
