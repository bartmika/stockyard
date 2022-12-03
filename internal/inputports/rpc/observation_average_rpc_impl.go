package rpc

import (
	"context"
	"errors"

	observation_a_d "github.com/bartmika/stockyard/internal/domain/observation_average"
	"github.com/bartmika/stockyard/pkg/dtos"
)

func (rpc *RPC) ListObservationAverages(req *dtos.ObservationAverageFilterRequestDTO, res *dtos.ObservationAverageListResponseDTO) error {

	////
	//// Convert DTO format into our database format.
	////

	f := &observation_a_d.ObservationAverageFilter{
		EntityIDs:               req.EntityIDs,
		Frequency:               req.Frequency,
		StartGreaterThen:        req.StartGreaterThen,
		StartGreaterThenOrEqual: req.StartGreaterThenOrEqual,
		FinishLessThen:          req.FinishLessThen,
		FinishLessThenOrEqual:   req.FinishLessThenOrEqual,
	}
	rpc.logger.Info().
		Str("func", "ListObservationAverages").
		Str("service", "rpc").
		Msg("beginning to list observation averages")

	////
	//// Fetch results.
	////

	ctx := context.Background()
	es, count, err := rpc.Services.ObservationAverageUsecase.ListAndCountByFilter(ctx, f)
	if err != nil {
		rpc.logger.Error().Err(err).Caller().Msg("database error")
		return err
	}
	if es == nil {
		rpc.logger.Error().Caller().Msg("observation averages list does not exist error")
		return errors.New("does not exist error")
	}

	////
	//// Serialize and return the results.
	////

	var results []*dtos.ObservationAverageResponseDTO
	for _, e := range es {
		er := &dtos.ObservationAverageResponseDTO{
			EntityID:  e.EntityID,
			Start:     e.Start,
			Finish:    e.Finish,
			Day:       e.Day,
			Week:      e.Week,
			Month:     e.Month,
			Year:      e.Year,
			Frequency: e.Frequency,
			Result:    e.Result,
		}
		results = append(results, er)
	}
	*res = dtos.ObservationAverageListResponseDTO{
		Results: results,
		Count:   count,
	}
	rpc.logger.Info().
		Str("func", "ListObservationAverages").
		Str("service", "rpc").
		Msg("succesfully listed observation averages")
	return nil
}
