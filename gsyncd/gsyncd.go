package main

import (
	"database/sql"
	"fmt"
	simplejson "github.com/bitly/go-simplejson"
	"github.com/garfcat/filesync/api"
	"github.com/garfcat/filesync/index"
	"github.com/fsnotify/fsnotify"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"os"
)

func main() {

	input := args()
	if len(input) >= 1 {
		start(input[0])
	}
}

func start(configFile string) {
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println(configFile, " not found")
		return
	}
	json, _ := simplejson.NewJson(b)
	ip := json.Get("ip").MustString("127.0.0.1")
	port := json.Get("port").MustInt(6776)

	monitors := json.Get("monitors").MustMap()

	for _, v := range monitors {
		watcher, _ := fsnotify.NewWatcher()
		monitored, _ := v.(string)
		monitored = index.PathSafe(monitored)
		db, err := sql.Open("sqlite3", index.SlashSuffix(monitored)+".sync/index.db")
		if err != nil {
			panic(fmt.Errorf("init index error[%s]\n", err.Error()))
		}
		defer db.Close()
		db.Exec("VACUUM;")
		err = index.InitIndex(monitored, db)
		if err != nil {
			panic(fmt.Errorf("init index error[%s]\n", err.Error()))
		}
		err = index.WatchRecursively(watcher, monitored, monitored)
		if err != nil {
			panic(fmt.Errorf("WatchRecursively error[%s]\n", err.Error()))
		}
		go index.ProcessEvent(watcher, monitored)
	}

	api.RunWeb(ip, port, monitors)
	//watcher.Close()
}

func args() []string {
	ret := []string{}
	if len(os.Args) <= 1 {
		ret = append(ret, "gsyncd.json")
	} else {
		for i := 1; i < len(os.Args); i++ {
			ret = append(ret, os.Args[i])
		}
	}
	return ret
}
