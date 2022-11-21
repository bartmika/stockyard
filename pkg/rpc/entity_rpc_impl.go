package rpc

import (
	"github.com/bartmika/stockyard/pkg/dtos"
)

func (s *StockYardRPCClient) InsertEntity(name string, dataType int8, meta string) (*dtos.EntityResponseDTO, error) {
	s.Logger.Info().Msg("calling remote insert entity function...")

	// Create the request payload which we will send to the server.
	dto := &dtos.EntityInsertRequestDTO{
		Name:     name,
		DataType: dataType,
		Meta:     meta,
	}

	// Create the response payload that will be filled out by the server.
	var reply dtos.EntityResponseDTO

	// Make the remote procedure call and handle the result.
	err := s.call("RPC.InsertEntity", dto, &reply)
	if err != nil {
		s.Logger.Error().Err(err).Caller().Str("RemoteAddress", s.serverAddress).Msg("failed making remote proceedure call")
		return nil, err
	}

	s.Logger.Info().
		Uint64("id", reply.ID).
		Str("uuid", reply.UUID).
		Str("name", reply.Name).
		Int8("data_type", reply.DataType).
		Str("meta", reply.Meta).
		Msg("succesfully called remote proceedure")

	return &reply, nil
}

func (s *StockYardRPCClient) ListEntities(dto *dtos.EntityFilterRequestDTO) (*dtos.EntityListResponseDTO, error) {
	s.Logger.Info().Msg("calling remote list entity function...")

	// Create the response payload that will be filled out by the server.
	var reply dtos.EntityListResponseDTO

	// Make the remote procedure call and handle the result.
	err := s.call("RPC.ListEntities", dto, &reply)
	if err != nil {
		s.Logger.Error().Err(err).Caller().Str("RemoteAddress", s.serverAddress).Msg("failed making remote proceedure call")
		return nil, err
	}

	s.Logger.Info().
		Msg("succesfully called remote proceedure")

	return &reply, nil
}

func (s *StockYardRPCClient) DeleteEntityByPrimaryKey(entityID uint64) error {
	s.Logger.Info().Msg("calling remote delete entity function...")

	// Create the response payload that will be filled out by the server.
	var reply dtos.EntityResponseDTO

	// Make the remote procedure call and handle the result.
	err := s.call("RPC.DeleteEntityByPrimaryKey", &entityID, &reply)
	if err != nil {
		s.Logger.Error().Err(err).Caller().Str("RemoteAddress", s.serverAddress).Msg("failed making remote proceedure call")
		return err
	}

	s.Logger.Info().
		Uint64("id", entityID).
		Msg("succesfully called remote proceedure")

	return nil
}
