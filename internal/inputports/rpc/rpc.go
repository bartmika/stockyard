package rpc

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"syscall"
	t "time"

	"github.com/rs/zerolog"

	"github.com/bartmika/stockyard/internal/app"
	"github.com/bartmika/stockyard/internal/config"
	"github.com/bartmika/stockyard/internal/pkg/kmutex"
	"github.com/bartmika/stockyard/internal/pkg/time"
	"github.com/bartmika/stockyard/internal/pkg/uuid"
)

type RPC struct {
	Time        time.Provider
	UUID        uuid.Provider
	KMutext     kmutex.Provider
	logger      *zerolog.Logger
	Services    app.Services
	tcpListener *net.TCPListener
}

func NewServer(appConfig *config.Conf, uuidProvider uuid.Provider, timeProvider time.Provider, kmutexProvider kmutex.Provider, appServices app.Services) *RPC {

	applicationAddress := fmt.Sprintf("%s:%s", appConfig.Server.IP, appConfig.Server.Port)
	appServices.Logger.Info().Msgf("rpc api initializing address for %s", applicationAddress)

	tcpAddr, err := net.ResolveTCPAddr("tcp", applicationAddress)
	if err != nil {
		log.Fatal(err)
	}

	rpcServer := &RPC{
		Time:     timeProvider,
		UUID:     uuidProvider,
		KMutext:  kmutexProvider,
		logger:   appServices.Logger,
		Services: appServices,
	}

	// Attach the
	rpc.Register(rpcServer)
	rpc.HandleHTTP()

	appServices.Logger.Info().Msg("rpc api was initialized")
	l, e := net.ListenTCP("tcp", tcpAddr)
	if e != nil {
		l.Close()
		appServices.Logger.Fatal().Err(e).Msg("rpc api failed to initialize:")
	}

	appServices.Logger.Info().Msg("rpc api is listening now")
	rpcServer.tcpListener = l
	rpcServer.Services = appServices
	return rpcServer
}

func (rpcServer *RPC) ListenAndServe() error {
	rpcServer.logger.Info().Msg("rpc api is starting now")

	// The following code will attach a background handler to run when the
	// application detects a shutdown signal.
	// Special thanks via https://guzalexander.com/2017/05/31/gracefully-exit-server-in-go.html
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs // Block execution until signal from terminal gets triggered here.

		// Finish any RPC communication taking place at the moment before
		// shutting down the RPC server.
		rpcServer.tcpListener.Close()
	}()

	// Attach the following anonymous function to run on all cases (ex: panic,
	// termination signal, etc) so we can gracefully shutdown the service.
	defer func() {
		rpcServer.stopRuntimeLoop()
	}()

	// Safety net for 'too many open files' issue on legacy code.
	// Set a sane timeout duration for the http.DefaultClient, to ensure idle connections are terminated.
	// Reference: https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	http.DefaultClient.Timeout = t.Minute * 10

	// DEVELOPER NOTES:
	// If you get "too many open files" then please read the following article
	// http://publib.boulder.ibm.com/httpserv/ihsdiag/too_many_open_files.html
	// so you can run in your console:
	// $ ulimit -H -n 4096
	// $ ulimit -n 4096

	rpcServer.logger.Info().Msg("rpc api is ready and running")

	// Run the main loop blocking code.
	http.Serve(rpcServer.tcpListener, nil)

	return nil
}

func (rpcServer *RPC) stopRuntimeLoop() error {
	log.Printf("Starting graceful shutdown now...")
	rpcServer.tcpListener.Close()
	log.Printf("Terminated TCPListener.")
	log.Printf("Graceful shutdown finished.")
	return nil
}
