package settings

import (
	"fmt"
	"strings"

	_ "embed"

	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/darklab8/fl-configs/configs/configs_settings"
	"github.com/darklab8/go-utils/utils/enverant"
	"github.com/darklab8/go-utils/utils/utils_settings"
)

//go:embed version.txt
var version string

type DarkstatEnvVars struct {
	utils_settings.UtilsEnvs
	configs_settings.ConfEnvVars
	TractorTabName    string
	SiteRoot          string
	SiteRootAcceptors string
	AppHeading        string
	AppVersion        string
	IsDetailed        bool

	RelayHost     string
	RelayRoot     string
	RelayLoopSecs int
}

var Env DarkstatEnvVars

func init() {
	env := enverant.NewEnverant()
	Env = DarkstatEnvVars{
		UtilsEnvs:         utils_settings.GetEnvs(env),
		ConfEnvVars:       configs_settings.GetEnvs(env),
		TractorTabName:    env.GetStr("DARKSTAT_TRACTOR_TAB_NAME", enverant.OrStr("Tractors")),
		SiteRoot:          env.GetStr("SITE_ROOT", enverant.OrStr("/")),
		SiteRootAcceptors: env.GetStr("SITE_ROOT_ACCEPTORS", enverant.OrStr("")),
		AppHeading:        env.GetStr("FLDARKSTAT_HEADING", enverant.OrStr("")),
		AppVersion:        getAppVersion(),
		IsDetailed:        env.GetBoolOr("DARKSTAT_DETAILED", false),
		RelayHost:         env.GetStr("RELAY_HOST", enverant.OrStr("http://localhost:8080")),
		RelayRoot:         env.GetStr("RELAY_ROOT", enverant.OrStr("/")),
		RelayLoopSecs:     env.GetIntOr("RELAY_LOOP_SECS", 30),
	}

	fmt.Sprintln("conf=", Env)
}

func (e DarkstatEnvVars) GetSiteRootAcceptors() []string {
	if e.SiteRootAcceptors == "" {
		return []string{}
	}

	return strings.Split(e.SiteRootAcceptors, ",")
}

func IsRelayActive(configs *configs_mapped.MappedConfigs) bool {
	if configs.Discovery != nil {
		fmt.Println("discovery always wishes to see pobs")
		return true
	}

	fmt.Println("relay is disabled")
	return false
}

func getAppVersion() string {
	// cleaning up version from... debugging logs used during dev env
	lines := strings.Split(version, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "v") {
			return line
		}
	}
	return version
}
