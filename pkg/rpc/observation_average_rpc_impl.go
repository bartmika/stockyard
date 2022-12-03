package rpc

import (
	"github.com/bartmika/stockyard/pkg/dtos"
)

func (s *StockYardRPCClient) ListObservationAverages(dto *dtos.ObservationAverageFilterRequestDTO) (*dtos.ObservationAverageListResponseDTO, error) {
	s.Logger.Info().Msg("calling remote list observation average function...")

	// Create the response payload that will be filled out by the server.
	var reply dtos.ObservationAverageListResponseDTO

	// Make the remote procedure call and handle the result.
	err := s.call("RPC.ListObservationAverages", dto, &reply)
	if err != nil {
		s.Logger.Error().Err(err).Caller().Str("RemoteAddress", s.serverAddress).Msg("failed making remote proceedure call")
		return nil, err
	}

	s.Logger.Info().
		Msg("succesfully called remote proceedure")

	return &reply, nil
}
