package rpc

import (
	"github.com/bartmika/stockyard/pkg/dtos"
)

func (s *StockYardRPCClient) ListObservationSummations(dto *dtos.ObservationSummationFilterRequestDTO) (*dtos.ObservationSummationListResponseDTO, error) {
	s.Logger.Info().Msg("calling remote list observation function...")

	// Create the response payload that will be filled out by the server.
	var reply dtos.ObservationSummationListResponseDTO

	// Make the remote procedure call and handle the result.
	err := s.call("RPC.ListObservationSummations", dto, &reply)
	if err != nil {
		s.Logger.Error().Err(err).Caller().Str("RemoteAddress", s.serverAddress).Msg("failed making remote proceedure call")
		return nil, err
	}

	s.Logger.Info().
		Msg("succesfully called remote proceedure")

	return &reply, nil
}
