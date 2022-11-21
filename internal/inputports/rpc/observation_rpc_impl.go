package rpc

import (
	"context"
	"errors"

	observation_d "github.com/bartmika/stockyard/internal/domain/observation"
	"github.com/bartmika/stockyard/pkg/dtos"
)

func (rpc *RPC) InsertObservation(req *dtos.ObservationInsertRequestDTO, res *dtos.ObservationResponseDTO) error {
	ctx := context.Background()

	// Defensive code: Check if we previously submitted a similar observation.
	doesExist, err := rpc.Services.ObservationUsecase.CheckIfExistsByPrimaryKey(ctx, req.EntityID, req.Timestamp)
	if err != nil || doesExist {
		if err != nil {
			rpc.logger.Error().Err(err).Caller().Uint64("entity_id", req.EntityID).Msg("database error")
			return err
		}
		rpc.logger.Warn().Caller().Uint64("entity_id", req.EntityID).Time("timestamp", req.Timestamp).Msg("observation already exists for the `entity_id` and `timestamp`")
		return errors.New("observation already exists for the `entity_id` and `timestamp`")
	}

	e := &observation_d.Observation{
		EntityID:  req.EntityID,
		Meta:      req.Meta,
		Timestamp: req.Timestamp,
		Value:     req.Value,
	}

	o, err := rpc.Services.ObservationUsecase.Insert(ctx, e)
	if err != nil {
		rpc.logger.Error().Err(err).Caller().Uint64("entity_id", req.EntityID).Msg("database error")
		return err
	}
	if o == nil {
		rpc.logger.Warn().Caller().Uint64("entity_id", req.EntityID).Msg("observation does not exist error")
		return errors.New("does not exist error")
	}

	rpc.logger.Info().
		Str("func", "InsertObservation").
		Str("service", "rpc").
		Uint64("entity_id", o.EntityID).
		Str("meta", o.Meta).
		Time("timestamp", o.Timestamp).
		Float64("value", o.Value).
		Msg("succesfully created entity")

	*res = dtos.ObservationResponseDTO{
		EntityID:  o.EntityID,
		Meta:      o.Meta,
		Timestamp: o.Timestamp,
		Value:     o.Value,
	}
	return nil
}

func (rpc *RPC) ListObservations(req *dtos.ObservationFilterRequestDTO, res *dtos.ObservationListResponseDTO) error {

	////
	//// Convert DTO format into our database format.
	////

	f := &observation_d.ObservationFilter{
		EntityIDs:                   req.EntityIDs,
		TimestampGreaterThen:        req.TimestampGreaterThen,
		TimestampGreaterThenOrEqual: req.TimestampGreaterThenOrEqual,
		TimestampLessThen:           req.TimestampLessThen,
		TimestampLessThenOrEqual:    req.TimestampLessThenOrEqual,
	}
	rpc.logger.Info().
		Str("func", "ListObservations").
		Str("service", "rpc").
		Msg("beginning to list observations")

	////
	//// Fetch results.
	////

	ctx := context.Background()
	es, count, err := rpc.Services.ObservationUsecase.ListAndCountByFilter(ctx, f)
	if err != nil {
		rpc.logger.Error().Err(err).Caller().Msg("database error")
		return err
	}
	if es == nil {
		rpc.logger.Error().Caller().Msg("observation list does not exist error")
		return errors.New("does not exist error")
	}

	////
	//// Serialize and return the results.
	////

	var results []*dtos.ObservationResponseDTO
	for _, e := range es {
		er := &dtos.ObservationResponseDTO{
			EntityID:  e.EntityID,
			Meta:      e.Meta,
			Timestamp: e.Timestamp,
			Value:     e.Value,
		}
		results = append(results, er)
	}
	*res = dtos.ObservationListResponseDTO{
		Results: results,
		Count:   count,
	}
	rpc.logger.Info().
		Str("func", "ListObservations").
		Str("service", "rpc").
		Msg("succesfully listed observations")
	return nil
}

func (rpc *RPC) DeleteObservationByPrimaryKey(req *dtos.ObservationPrimaryKeyRequestDTO, res *dtos.ObservationResponseDTO) error {
	ctx := context.Background()
	err := rpc.Services.ObservationUsecase.DeleteByPrimaryKey(ctx, req.EntityID, req.Timestamp)
	if err != nil {
		rpc.logger.Error().Err(err).Caller().Msg("database error")
		return err
	}

	rpc.logger.Info().
		Str("func", "DeleteObservationByPrimaryKey").
		Str("service", "rpc").
		Uint64("entity_id", req.EntityID).
		Time("timestamp", req.Timestamp).
		Msg("succesfully deleted observation")

	*res = dtos.ObservationResponseDTO{}
	return nil
}
