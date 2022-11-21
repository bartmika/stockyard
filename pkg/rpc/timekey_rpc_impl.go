package rpc

import (
	"github.com/bartmika/stockyard/pkg/dtos"
)

func (s *StockYardRPCClient) InsertTimeKey(dto *dtos.TimeKeyInsertRequestDTO) (*dtos.TimeKeyResponseDTO, error) {
	s.Logger.Info().Msg("calling remote insert observation function...")

	// Create the response payload that will be filled out by the server.
	var reply dtos.TimeKeyResponseDTO

	// Make the remote procedure call and handle the result.
	err := s.call("RPC.InsertTimeKey", dto, &reply)
	if err != nil {
		s.Logger.Error().Err(err).Caller().Str("RemoteAddress", s.serverAddress).Msg("failed making remote proceedure call")
		return nil, err
	}

	s.Logger.Info().
		Uint64("entity_id", reply.EntityID).
		Str("meta", reply.Meta).
		Time("timestamp", reply.Timestamp).
		Str("value", reply.Value).
		Msg("succesfully called remote proceedure")

	return &reply, nil
}

func (s *StockYardRPCClient) ListTimeKeys(dto *dtos.TimeKeyFilterRequestDTO) (*dtos.TimeKeyListResponseDTO, error) {
	s.Logger.Info().Msg("calling remote list observation function...")

	// Create the response payload that will be filled out by the server.
	var reply dtos.TimeKeyListResponseDTO

	// Make the remote procedure call and handle the result.
	err := s.call("RPC.ListTimeKeys", dto, &reply)
	if err != nil {
		s.Logger.Error().Err(err).Caller().Str("RemoteAddress", s.serverAddress).Msg("failed making remote proceedure call")
		return nil, err
	}

	s.Logger.Info().
		Msg("succesfully called remote proceedure")

	return &reply, nil
}

func (s *StockYardRPCClient) DeleteTimeKeysByFilter(dto *dtos.TimeKeyFilterRequestDTO) error {
	s.Logger.Info().Msg("calling remote list observation function...")

	// Create the response payload that will be filled out by the server.
	var reply dtos.TimeKeyResponseDTO

	// Make the remote procedure call and handle the result.
	err := s.call("RPC.DeleteTimeKeysByFilter", dto, &reply)
	if err != nil {
		s.Logger.Error().Err(err).Caller().Str("RemoteAddress", s.serverAddress).Msg("failed making remote proceedure call")
		return err
	}

	s.Logger.Info().
		Msg("succesfully called remote proceedure")

	return nil
}
