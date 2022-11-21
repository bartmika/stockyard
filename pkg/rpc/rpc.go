package rpc

import (
	"net/rpc"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type StockYardRPCClient struct {
	serverAddress string
	Logger        *zerolog.Logger
	Client        *rpc.Client
	RetryLimit    uint8
	retryCount    uint8
	DelayDuration time.Duration
	addr          string
}

func NewClient(addr string, retryLimit uint8, delayDuration time.Duration) (*StockYardRPCClient, error) {
	// The following line of code adds a pretty output to the console.
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().
		Timestamp(). // Add timestamp to every call.
		Logger()

	logger.Info().Str("Address", addr).Msg("dialing remote server...")

	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		logger.Error().Err(err).Caller().Str("Address", addr).Msg("failed dialing remote server")
		return nil, err
	}

	logger.Info().Str("Address", addr).Msg("successfully connected to remote server")

	return &StockYardRPCClient{
		Logger:        &logger,
		serverAddress: addr,
		Client:        client,
		RetryLimit:    retryLimit,
		retryCount:    0,
		DelayDuration: delayDuration,
		addr:          addr,
	}, err
}

// Function used to make RPC calls with retry functionality in case the RPC
// server has been shutdown and the connection was lost.
func (s *StockYardRPCClient) call(serviceMethod string, args interface{}, reply interface{}) error {
	err := s.Client.Call(serviceMethod, args, reply)

	// Detect the `connection is shut down` error.
	if err == rpc.ErrShutdown {
		if s.retryCount < s.RetryLimit {
			s.retryCount += 1
			s.Logger.Error().Err(err).Caller().Str("Address", s.serverAddress).Uint8("RetryCount", s.retryCount).Msg("detected `connection is shutdown`, trying...")

			// We need to apply an artifical delay in case we need to give time
			// for the server is starting up.
			time.Sleep(s.DelayDuration)

			// Attempt to re-connected and if successful then retry calling the
			// RPC endpoint, else return with error.
			client, err := rpc.DialHTTP("tcp", s.addr)
			if err != nil {
				s.Logger.Error().Err(err).Caller().Str("Address", s.serverAddress).Uint8("RetryCount", s.retryCount).Msg("detected `connection is shutdown`, retrying...")

				// Note: Use recursion to retry the call.
				return s.call(serviceMethod, args, reply)
			}

			s.Logger.Info().Str("Address", s.serverAddress).Uint8("RetryCount", s.retryCount).Msg("reconnected")
			s.Client = client

			// Note: Use recursion to retry the call.
			return s.call(serviceMethod, args, reply)
		}
		s.Logger.Error().Err(err).Caller().Str("Address", s.serverAddress).Uint8("RetryCount", s.retryCount).Msg("detected `connection is shutdown`, too many retries, shutting down now")
		return err
	}

	// If success then nil will be returned, else the specific error will be returned.
	return err
}
