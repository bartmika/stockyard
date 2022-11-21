package rpc

import (
	"context"
	"errors"

	entity_d "github.com/bartmika/stockyard/internal/domain/entity"
	"github.com/bartmika/stockyard/pkg/dtos"
)

func (rpc *RPC) InsertEntity(req *dtos.EntityInsertRequestDTO, res *dtos.EntityResponseDTO) error {
	e := &entity_d.Entity{
		Name:     req.Name,
		DataType: req.DataType,
		Meta:     req.Meta,
	}
	ctx := context.Background()
	e, err := rpc.Services.EntityUsecase.Insert(ctx, e)
	if err != nil {
		rpc.logger.Error().Err(err).Caller().Str("name", req.Name).Int8("data_type", req.DataType).Msg("database error")
		return err
	}
	if e == nil {
		rpc.logger.Warn().Caller().Str("name", req.Name).Int8("data_type", req.DataType).Msg("entity does not exist error")
		return errors.New("does not exist error")
	}

	rpc.logger.Info().
		Str("func", "InsertEntity").
		Str("service", "rpc").
		Uint64("id", e.ID).
		Str("uuid", e.UUID).
		Str("name", e.Name).
		Int8("data_type", e.DataType).
		Msg("succesfully created entity")

	*res = dtos.EntityResponseDTO{
		ID:       e.ID,
		UUID:     e.UUID,
		Name:     e.Name,
		DataType: e.DataType,
	}
	return nil
}

func (rpc *RPC) ListEntities(req *dtos.EntityFilterRequestDTO, res *dtos.EntityListResponseDTO) error {

	////
	//// Convert DTO format into our database format.
	////

	f := &entity_d.EntityFilter{
		SortOrder: req.SortOrder,
		SortField: req.SortField,
		Offset:    req.Offset,
		Limit:     req.Limit,
		IDs:       req.IDs,
		DataType:  req.DataType,
	}

	rpc.logger.Info().
		Str("func", "ListEntities").
		Str("service", "rpc").
		Msg("beginning to list entities")

	////
	//// Fetch results.
	////

	ctx := context.Background()
	es, count, err := rpc.Services.EntityUsecase.ListAndCountByFilter(ctx, f)
	if err != nil {
		rpc.logger.Error().Err(err).Caller().Int8("data_type", req.DataType).Msg("database error")
		return err
	}
	if es == nil {
		rpc.logger.Warn().Caller().Int8("data_type", req.DataType).Msg("does not exist error")
		return errors.New("does not exist error")
	}

	////
	//// Serialize and return the results.
	////

	var results []*dtos.EntityResponseDTO
	for _, e := range es {
		er := &dtos.EntityResponseDTO{
			ID:       e.ID,
			UUID:     e.UUID,
			Name:     e.Name,
			DataType: e.DataType,
		}
		results = append(results, er)
	}

	rpc.logger.Info().
		Str("func", "ListEntities").
		Str("service", "rpc").
		Msg("succesfully listed entities")

	*res = dtos.EntityListResponseDTO{
		Results: results,
		Count:   count,
	}
	return nil
}

func (rpc *RPC) DeleteEntityByPrimaryKey(entityID *uint64, res *dtos.EntityResponseDTO) error {
	ctx := context.Background()
	err := rpc.Services.EntityUsecase.Delete(ctx, *entityID)
	if err != nil {
		rpc.logger.Error().Err(err).Caller().Msg("database error")
		return err
	}
	*res = dtos.EntityResponseDTO{}
	rpc.logger.Info().
		Str("func", "DeleteEntityByPrimaryKey").
		Str("service", "rpc").
		Uint64("id", *entityID).
		Msg("succesfully deleted entity")
	return nil
}
