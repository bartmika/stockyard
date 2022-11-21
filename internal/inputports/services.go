package inputports

import (
	"github.com/bartmika/stockyard/internal/app"
	"github.com/bartmika/stockyard/internal/config"
	"github.com/bartmika/stockyard/internal/inputports/cron"
	rpcs "github.com/bartmika/stockyard/internal/inputports/rpc"
	"github.com/bartmika/stockyard/internal/pkg/kmutex"
	"github.com/bartmika/stockyard/internal/pkg/time"
	"github.com/bartmika/stockyard/internal/pkg/uuid"
)

//Services contains the ports services
type Services struct {
	CronServer *cron.Server
	RPCServer  *rpcs.RPC
}

//NewServices instantiates the services of input ports
func NewServices(appConf *config.Conf, uuidProvider uuid.Provider, timeProvider time.Provider, kmutexProvider kmutex.Provider, appServices app.Services) Services {
	return Services{
		CronServer: cron.NewServer(appConf, uuidProvider, timeProvider, kmutexProvider, appServices),
		RPCServer:  rpcs.NewServer(appConf, uuidProvider, timeProvider, kmutexProvider, appServices),
	}
}
