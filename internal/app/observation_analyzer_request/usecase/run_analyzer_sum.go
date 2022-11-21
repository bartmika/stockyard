package usecase

import (
	"context"
	"time"

	oardomain "github.com/bartmika/stockyard/internal/domain/observation_analyzer_request"
	osumdomain "github.com/bartmika/stockyard/internal/domain/observation_summation"
	"github.com/bartmika/timekit"
)

func (uc observationAnalyzerRequestUsecase) analyzerProcessSummationRequest(ctx context.Context, req *oardomain.ObservationAnalyzerRequest) (map[int8]*osumdomain.ObservationSummation, error) {
	uc.Logger.Info().
		Uint64("entity", req.EntityID).
		Str("uuid", req.UUID).
		Str("func", "analyzerProcessSummationRequest").
		Msg("analyzer")

	results := map[int8]*osumdomain.ObservationSummation{}

	today, err := uc.analyzerProcessSummationForTodayRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	thisWeek, err := uc.analyzerProcessSummationForThisWeekRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	thisMonth, err := uc.analyzerProcessSummationForThisMonthRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	thisYear, err := uc.analyzerProcessSummationForThisYearRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	results[osumdomain.ObservationSummationDayFrequency] = today
	results[osumdomain.ObservationSummationWeekFrequency] = thisWeek
	results[osumdomain.ObservationSummationMonthFrequency] = thisMonth
	results[osumdomain.ObservationSummationYearFrequency] = thisYear

	return results, nil
}

func (uc observationAnalyzerRequestUsecase) analyzerProcessSummationForTodayRequest(ctx context.Context, req *oardomain.ObservationAnalyzerRequest) (*osumdomain.ObservationSummation, error) {
	////
	//// Fetch or create our record.
	////

	o := req.Observation
	year := o.Timestamp.Year()
	month := o.Timestamp.Month()
	week := timekit.GetWeekNumberFromDate(o.Timestamp)
	day := o.Timestamp.Day()

	uc.Logger.Info().
		Uint64("entity", req.EntityID).
		Str("uuid", req.UUID).
		Int("year", year).
		Int("month", int(month)).
		Int("week", week).
		Int("day", day).
		Str("func", "analyzerProcessSummationForTodayRequest").
		Msg("analyzer")

	start := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, month, day+1, 0, 0, 0, 0, time.UTC)
	oc := &osumdomain.ObservationSummation{
		EntityID:  req.EntityID,
		Start:     start,
		Finish:    end,
		Day:       day,
		Week:      week,
		Month:     int(month),
		Year:      year,
		Frequency: osumdomain.ObservationSummationDayFrequency,
		Result:    0,
	}

	oc, err := uc.ObservationSummationRepo.InsertOrGetByPrimaryKey(ctx, oc)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return nil, err
	}

	////
	//// Perform our computation.
	////

	if req.Type == oardomain.ObservationAnalyzerRequestInsertType {
		oc.Result += o.Value
	} else if req.Type == oardomain.ObservationAnalyzerRequestDeleteType && oc.Result > 0 {
		oc.Result -= o.Value
	}

	////
	//// Save record.
	////

	err = uc.ObservationSummationRepo.UpdateByPrimaryKey(ctx, oc)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return nil, err
	}
	return oc, nil
}

func (uc observationAnalyzerRequestUsecase) analyzerProcessSummationForThisWeekRequest(ctx context.Context, req *oardomain.ObservationAnalyzerRequest) (*osumdomain.ObservationSummation, error) {
	////
	//// Fetch or create our record.
	////

	o := req.Observation
	year := o.Timestamp.Year()
	month := o.Timestamp.Month()
	week := timekit.GetWeekNumberFromDate(o.Timestamp)

	uc.Logger.Info().
		Uint64("entity", req.EntityID).
		Str("uuid", req.UUID).
		Int("year", year).
		Int("month", int(month)).
		Int("week", week).
		Str("func", "analyzerProcessSummationForThisWeekRequest").
		Msg("analyzer")

	start := timekit.GetFirstDateFromWeekAndYear(week, year, time.UTC)
	end := timekit.GetFirstDateFromWeekAndYear(week+1, year, time.UTC)
	oc := &osumdomain.ObservationSummation{
		EntityID:  req.EntityID,
		Start:     start,
		Finish:    end,
		Day:       0,
		Week:      week,
		Month:     int(month),
		Year:      year,
		Frequency: osumdomain.ObservationSummationWeekFrequency,
		Result:    0,
	}

	oc, err := uc.ObservationSummationRepo.InsertOrGetByPrimaryKey(ctx, oc)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return nil, err
	}

	////
	//// Perform our computation.
	////

	if req.Type == oardomain.ObservationAnalyzerRequestInsertType {
		oc.Result += o.Value
	} else if req.Type == oardomain.ObservationAnalyzerRequestDeleteType && oc.Result > 0 {
		oc.Result -= o.Value
	}

	////
	//// Save record.
	////

	err = uc.ObservationSummationRepo.UpdateByPrimaryKey(ctx, oc)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return nil, err
	}
	return oc, nil
}

func (uc observationAnalyzerRequestUsecase) analyzerProcessSummationForThisMonthRequest(ctx context.Context, req *oardomain.ObservationAnalyzerRequest) (*osumdomain.ObservationSummation, error) {
	////
	//// Fetch or create our record.
	////

	o := req.Observation
	year := o.Timestamp.Year()
	month := o.Timestamp.Month()

	uc.Logger.Info().
		Uint64("entity", req.EntityID).
		Str("uuid", req.UUID).
		Int("year", year).
		Int("month", int(month)).
		Str("func", "analyzerProcessSummationForThisMonthRequest").
		Msg("analyzer")

	start := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
	oc := &osumdomain.ObservationSummation{
		EntityID:  req.EntityID,
		Start:     start,
		Finish:    end,
		Day:       0,
		Week:      0,
		Month:     int(month),
		Year:      year,
		Frequency: osumdomain.ObservationSummationMonthFrequency,
		Result:    0,
	}

	oc, err := uc.ObservationSummationRepo.InsertOrGetByPrimaryKey(ctx, oc)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return nil, err
	}

	////
	//// Perform our computation.
	////

	if req.Type == oardomain.ObservationAnalyzerRequestInsertType {
		oc.Result += o.Value
	} else if req.Type == oardomain.ObservationAnalyzerRequestDeleteType && oc.Result > 0 {
		oc.Result -= o.Value
	}

	////
	//// Save record.
	////

	err = uc.ObservationSummationRepo.UpdateByPrimaryKey(ctx, oc)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return nil, err
	}
	return oc, nil
}

func (uc observationAnalyzerRequestUsecase) analyzerProcessSummationForThisYearRequest(ctx context.Context, req *oardomain.ObservationAnalyzerRequest) (*osumdomain.ObservationSummation, error) {
	////
	//// Fetch or create our record.
	////

	o := req.Observation
	year := o.Timestamp.Year()

	uc.Logger.Info().
		Uint64("entity", req.EntityID).
		Str("uuid", req.UUID).
		Int("year", year).
		Str("func", "analyzerProcessSummationForThisYearRequest").
		Msg("analyzer")

	start := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(year+1, 1, 1, 0, 0, 0, 0, time.UTC)
	oc := &osumdomain.ObservationSummation{
		EntityID:  req.EntityID,
		Start:     start,
		Finish:    end,
		Day:       0,
		Week:      0,
		Month:     0,
		Year:      year,
		Frequency: osumdomain.ObservationSummationYearFrequency,
		Result:    0,
	}

	oc, err := uc.ObservationSummationRepo.InsertOrGetByPrimaryKey(ctx, oc)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return nil, err
	}

	////
	//// Perform our computation.
	////

	if req.Type == oardomain.ObservationAnalyzerRequestInsertType {
		oc.Result += o.Value
	} else if req.Type == oardomain.ObservationAnalyzerRequestDeleteType && oc.Result > 0 {
		oc.Result -= o.Value
	}

	////
	//// Save record.
	////

	err = uc.ObservationSummationRepo.UpdateByPrimaryKey(ctx, oc)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return nil, err
	}
	return oc, nil
}
