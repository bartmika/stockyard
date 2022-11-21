package rpc

import (
	"github.com/bartmika/stockyard/pkg/dtos"
)

func (s *StockYardRPCClient) InsertObservation(dto *dtos.ObservationInsertRequestDTO) (*dtos.ObservationResponseDTO, error) {
	s.Logger.Info().Msg("calling remote insert observation function...")

	// Create the response payload that will be filled out by the server.
	var reply dtos.ObservationResponseDTO

	// Make the remote procedure call and handle the result.
	err := s.call("RPC.InsertObservation", dto, &reply)
	if err != nil {
		s.Logger.Error().Err(err).Caller().Str("RemoteAddress", s.serverAddress).Msg("failed making remote proceedure call")
		return nil, err
	}

	s.Logger.Info().
		Uint64("entity_id", reply.EntityID).
		Str("meta", reply.Meta).
		Time("timestamp", reply.Timestamp).
		Float64("value", reply.Value).
		Msg("succesfully called remote proceedure")

	return &reply, nil
}

func (s *StockYardRPCClient) ListObservations(dto *dtos.ObservationFilterRequestDTO) (*dtos.ObservationListResponseDTO, error) {
	s.Logger.Info().Msg("calling remote list observation function...")

	// Create the response payload that will be filled out by the server.
	var reply dtos.ObservationListResponseDTO

	// Make the remote procedure call and handle the result.
	err := s.call("RPC.ListObservations", dto, &reply)
	if err != nil {
		s.Logger.Error().Err(err).Caller().Str("RemoteAddress", s.serverAddress).Msg("failed making remote proceedure call")
		return nil, err
	}

	s.Logger.Info().
		Msg("succesfully called remote proceedure")

	return &reply, nil
}

func (s *StockYardRPCClient) DeleteObservationByPrimaryKey(dto *dtos.ObservationPrimaryKeyRequestDTO) error {
	s.Logger.Info().Msg("calling remote delete entity function...")

	// Create the response payload that will be filled out by the server.
	var reply dtos.EntityResponseDTO

	// Make the remote procedure call and handle the result.
	err := s.call("RPC.DeleteEntity", &dto, &reply)
	if err != nil {
		s.Logger.Error().Err(err).Caller().Str("RemoteAddress", s.serverAddress).Msg("failed making remote proceedure call")
		return err
	}

	s.Logger.Info().
		Uint64("entity_id", dto.EntityID).
		Time("timestamp", dto.Timestamp).
		Msg("succesfully called remote proceedure")

	return nil
}
