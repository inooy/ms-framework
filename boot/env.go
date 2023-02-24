package boot

import (
	"github.com/inooy/ms-framework/pkg/log"
	"os"
)

type EnvConf struct {
	EnvId      string
	AppVersion string
	DeployType string
	InstanceId string
	ConfSrv    string
	AppName    string
	Port       string
}

var Env *EnvConf

func SetUpEnv() {
	result := &EnvConf{}
	if os.Getenv("ENV_ID") != "" {
		result.EnvId = os.Getenv("ENV_ID")
	} else {
		result.EnvId = "dev"
	}

	if os.Getenv("APP_NAME") != "" {
		result.AppName = os.Getenv("APP_NAME")
	} else {
		result.AppName = "write-sync"
	}
	result.Port = os.Getenv("PORT")
	result.ConfSrv = os.Getenv("CONF_SRV")
	result.AppVersion = os.Getenv("APP_VERSION")
	result.DeployType = os.Getenv("DEPLOY_TYPE")
	result.InstanceId = os.Getenv("HOSTNAME")
	log.Infof("env param: %+v", result)
	Env = result
}
