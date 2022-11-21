package rpc

import (
	"context"
	"errors"

	timekey_d "github.com/bartmika/stockyard/internal/domain/timekey"
	"github.com/bartmika/stockyard/pkg/dtos"
)

func (rpc *RPC) InsertTimeKey(req *dtos.TimeKeyInsertRequestDTO, res *dtos.TimeKeyResponseDTO) error {
	e := &timekey_d.TimeKey{
		EntityID:  req.EntityID,
		Meta:      req.Meta,
		Timestamp: req.Timestamp,
		Value:     req.Value,
	}
	ctx := context.Background()
	o, err := rpc.Services.TimeKeyUsecase.Insert(ctx, e)
	if err != nil {
		rpc.logger.Error().Err(err).Caller().Uint64("entity_id", req.EntityID).Msg("database error")
		return err
	}
	if o == nil {
		rpc.logger.Error().Caller().Uint64("entity_id", req.EntityID).Msg("does not exist error")
		return errors.New("does not exist error")
	}

	rpc.logger.Info().
		Str("func", "InsertTimeKey").
		Str("service", "rpc").
		Uint64("entity_id", o.EntityID).
		Str("meta", o.Meta).
		Time("timestamp", o.Timestamp).
		Str("value", o.Value).
		Msg("succesfully created entity")

	*res = dtos.TimeKeyResponseDTO{
		EntityID:  o.EntityID,
		Meta:      o.Meta,
		Timestamp: o.Timestamp,
		Value:     o.Value,
	}
	return nil
}

func (rpc *RPC) ListTimeKeys(req *dtos.TimeKeyFilterRequestDTO, res *dtos.TimeKeyListResponseDTO) error {

	////
	//// Convert DTO format into our database format.
	////

	f := &timekey_d.TimeKeyFilter{
		EntityIDs:                   req.EntityIDs,
		TimestampGreaterThen:        req.TimestampGreaterThen,
		TimestampGreaterThenOrEqual: req.TimestampGreaterThenOrEqual,
		TimestampLessThen:           req.TimestampLessThen,
		TimestampLessThenOrEqual:    req.TimestampLessThenOrEqual,
	}

	////
	//// Fetch results.
	////

	ctx := context.Background()
	es, count, err := rpc.Services.TimeKeyUsecase.ListAndCountByFilter(ctx, f)
	if err != nil {
		rpc.logger.Error().Err(err).Caller().Msg("database error")
		return err
	}
	if es == nil {
		rpc.logger.Error().Caller().Msg("does not exist error")
		return errors.New("does not exist error")
	}

	////
	//// Serialize and return the results.
	////

	var results []*dtos.TimeKeyResponseDTO
	for _, e := range es {
		er := &dtos.TimeKeyResponseDTO{
			EntityID:  e.EntityID,
			Meta:      e.Meta,
			Timestamp: e.Timestamp,
			Value:     e.Value,
		}
		results = append(results, er)
	}
	*res = dtos.TimeKeyListResponseDTO{
		Results: results,
		Count:   count,
	}
	rpc.logger.Info().
		Str("func", "ListTimeKeys").
		Str("service", "rpc").
		Msg("succesfully listed timekeys")
	return nil
}

func (rpc *RPC) DeleteTimeKeysByFilter(req *dtos.TimeKeyFilterRequestDTO, res *dtos.TimeKeyResponseDTO) error {

	////
	//// Convert DTO format into our database format.
	////

	f := &timekey_d.TimeKeyFilter{
		EntityIDs:                   req.EntityIDs,
		TimestampGreaterThen:        req.TimestampGreaterThen,
		TimestampGreaterThenOrEqual: req.TimestampGreaterThenOrEqual,
		TimestampLessThen:           req.TimestampLessThen,
		TimestampLessThenOrEqual:    req.TimestampLessThenOrEqual,
	}

	////
	//// Fetch results.
	////

	ctx := context.Background()
	err := rpc.Services.TimeKeyUsecase.DeleteByFilter(ctx, f)
	if err != nil {
		rpc.logger.Error().Err(err).Caller().Msg("database error")
		return err
	}

	////
	//// Return the results.
	////

	*res = dtos.TimeKeyResponseDTO{}
	rpc.logger.Info().
		Str("func", "DeleteTimeKeysByFilter").
		Str("service", "rpc").
		Msg("succesfully deleted timekeys")
	return nil
}
