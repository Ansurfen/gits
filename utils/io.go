package utils

import (
	"os"
	"os/user"
	"strings"

	"github.com/spf13/viper"
)

func getConf() (conf *viper.Viper) {
	dir := GetRootDir()
	if _, err := os.Stat(dir + "gitsconfig.ini"); err != nil {
		fp, err := os.Create(dir + "gitsconfig.ini")
		Panic(err)
		defer fp.Close()
	}
	conf = newConf("gitsconfig", "ini", dir)
	conf.SetDefault("env.git", "")
	conf.WriteConfigAs(dir + "gitsconfig.ini")
	return conf
}

func getEnv(username, filename string) *viper.Viper {
	return newConf(filename, "yaml", GetRootDir()+username+"/")
}

func newConf(confName, confType, dir string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigName(confName)
	conf.SetConfigType(confType)
	conf.AddConfigPath(dir)
	Panic(conf.ReadInConfig())
	return conf
}

func NewEnv(reponame string) *viper.Viper {
	dir := GetRootDir()
	envDir := dir + reponame + ".yml"
	repo := strings.Split(reponame, "/")
	if _, err := os.Stat(envDir); err != nil {
		MkDirs(dir + repo[0])
	}
	if _, err := os.Stat(envDir); err != nil {
		fp, err := os.Create(envDir)
		Panic(err)
		defer fp.Close()
	}
	envConf := newConf(repo[1], "yaml", dir+repo[0])
	return envConf
}

func CurEnv() *viper.Viper {
	path, err := os.Getwd()
	Panic(err)
	dir := GetRootDir()
	index := newConf("index", "ini", dir)
	path = index.GetString("index." + MD5(path))
	if path == "" {
		return nil
	}
	repo := strings.Split(path[strings.Index(path, dir)+len(dir):], "/")
	envConf := getEnv(repo[0], repo[1])
	return envConf
}

func RegisterIndex(key, value string) {
	dir := GetRootDir()
	if _, err := os.Stat(dir + "index.ini"); err != nil {
		fp, err := os.Create(dir + "index.ini")
		Panic(err)
		defer fp.Close()
	}
	conf := newConf("index", "ini", dir)
	conf.Set("index."+MD5(key), value)
	conf.WriteConfigAs(dir + "index.ini")
}

func GetRootDir() string {
	home, err := user.Current()
	Panic(err)
	return home.HomeDir + "/.gits/"
}
