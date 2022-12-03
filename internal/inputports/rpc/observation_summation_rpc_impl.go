package rpc

import (
	"context"
	"errors"

	observation_s_d "github.com/bartmika/stockyard/internal/domain/observation_summation"
	"github.com/bartmika/stockyard/pkg/dtos"
)

func (rpc *RPC) ListObservationSummations(req *dtos.ObservationSummationFilterRequestDTO, res *dtos.ObservationSummationListResponseDTO) error {

	////
	//// Convert DTO format into our database format.
	////

	f := &observation_s_d.ObservationSummationFilter{
		EntityIDs:               req.EntityIDs,
		Frequency:               req.Frequency,
		StartGreaterThen:        req.StartGreaterThen,
		StartGreaterThenOrEqual: req.StartGreaterThenOrEqual,
		FinishLessThen:          req.FinishLessThen,
		FinishLessThenOrEqual:   req.FinishLessThenOrEqual,
	}
	rpc.logger.Info().
		Str("func", "ListObservationSummations").
		Str("service", "rpc").
		Msg("beginning to list observation counts")

	////
	//// Fetch results.
	////

	ctx := context.Background()
	es, count, err := rpc.Services.ObservationSummationUsecase.ListAndCountByFilter(ctx, f)
	if err != nil {
		rpc.logger.Error().Err(err).Caller().Msg("database error")
		return err
	}
	if es == nil {
		rpc.logger.Error().Caller().Msg("observation sum list does not exist error")
		return errors.New("does not exist error")
	}

	////
	//// Serialize and return the results.
	////

	var results []*dtos.ObservationSummationResponseDTO
	for _, e := range es {
		er := &dtos.ObservationSummationResponseDTO{
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
	*res = dtos.ObservationSummationListResponseDTO{
		Results: results,
		Count:   count,
	}
	rpc.logger.Info().
		Str("func", "ListObservationSummations").
		Str("service", "rpc").
		Msg("succesfully listed observation counts")
	return nil
}
