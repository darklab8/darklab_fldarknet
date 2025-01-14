package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/darklab8/fl-darkcore/darkcore/builder"
	"github.com/darklab8/fl-darkcore/darkcore/web"
	"github.com/darklab8/fl-darkstat/darkrelay/relayrouter"
	"github.com/darklab8/fl-darkstat/darkstat/router"
	"github.com/darklab8/fl-darkstat/darkstat/settings"
	"github.com/darklab8/fl-darkstat/darkstat/settings/logus"
	"github.com/darklab8/go-typelog/typelog"
	"github.com/darklab8/go-utils/utils/ptr"
)

type Action string

func (a Action) ToStr() string { return string(a) }

const (
	Build   Action = "build"
	Web     Action = "web"
	Version Action = "version"
	Relay   Action = "relay"
)

func GetRelayFs(app_data *router.AppData) *builder.Filesystem {
	relay_router := relayrouter.NewRouter(app_data)
	relay_builder := relay_router.Link()
	relay_fs := relay_builder.BuildAll(true, nil)
	relay_router = nil
	relay_builder = nil
	return relay_fs
}

func main() {
	fmt.Println("freelancer folder=", settings.Env.FreelancerFolder, settings.Env)
	defer func() {
		if r := recover(); r != nil {
			logus.Log.Error("Program crashed. Sleeping 10 seconds before exit", typelog.Any("recover", r))
			if !settings.Env.IsDevEnv {
				fmt.Println("going to sleeping")
				time.Sleep(10 * time.Second)
			}
			panic(r)
		}
	}()

	var action string
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 1 {
		action = argsWithoutProg[0]
	}
	fmt.Println("act:", action)

	web_darkstat := func() {
		app_data := router.NewAppData()

		stat_router := router.NewRouter(app_data)
		stat_builder := stat_router.Link()

		stat_fs := stat_builder.BuildAll(true, nil)

		go web.NewWeb(stat_fs,
			web.WithMutexableData(app_data),
			web.WithSiteRoot(settings.Env.SiteRoot),
		).Serve(web.WebServeOpts{})

		app_data.Lock()
		relay_fs := GetRelayFs(stat_router.AppData)
		app_data.Unlock()
		runtime.GC()

		if settings.IsRelayActive(app_data.Mapped) {
			go func() {
				for {
					time.Sleep(time.Second * time.Duration(settings.Env.RelayLoopSecs))
					app_data.Lock()
					app_data.Mapped.Discovery.PlayerOwnedBases.Refresh()
					app_data.Configs.PoBs = app_data.Configs.GetPoBs()
					app_data.Configs.PoBGoods = app_data.Configs.GetPoBGoods(app_data.Configs.PoBs)
					relay_fs2 := GetRelayFs(stat_router.AppData)
					relay_fs.Files = relay_fs2.Files
					logus.Log.Info("refreshed content")
					runtime.GC()
					app_data.Unlock()
				}
			}()
		}

		web.NewWeb(
			relay_fs,
			web.WithMutexableData(app_data),
			web.WithSiteRoot(settings.Env.SiteRoot),
		).Serve(web.WebServeOpts{Port: ptr.Ptr(8080)})
	}

	switch Action(action) {

	case Build:
		app_data := router.NewAppData()
		router.NewRouter(app_data).Link().BuildAll(false, nil)
	case Web:
		web_darkstat()
	case Version:
		fmt.Println("version=", settings.Env.AppVersion)
	default:
		web_darkstat()
	}
}
