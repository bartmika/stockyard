package usecase

import (
	"context"
	"time"

	oardomain "github.com/bartmika/stockyard/internal/domain/observation_analyzer_request"
	osumdomain "github.com/bartmika/stockyard/internal/domain/observation_count"
	"github.com/bartmika/timekit"
)

func (uc observationAnalyzerRequestUsecase) analyzerProcessCountRequest(ctx context.Context, req *oardomain.ObservationAnalyzerRequest) (map[int8]*osumdomain.ObservationCount, error) {
	uc.Logger.Info().
		Uint64("entity", req.EntityID).
		Str("uuid", req.UUID).
		Str("func", "analyzerProcessCountRequest").
		Msg("analyzer")

	results := map[int8]*osumdomain.ObservationCount{}

	today, err := uc.analyzerProcessCountForTodayRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	thisWeek, err := uc.analyzerProcessCountForThisWeekRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	thisMonth, err := uc.analyzerProcessCountForThisMonthRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	thisYear, err := uc.analyzerProcessCountForThisYearRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	results[osumdomain.ObservationCountDayFrequency] = today
	results[osumdomain.ObservationCountWeekFrequency] = thisWeek
	results[osumdomain.ObservationCountMonthFrequency] = thisMonth
	results[osumdomain.ObservationCountYearFrequency] = thisYear

	return results, nil
}

func (uc observationAnalyzerRequestUsecase) analyzerProcessCountForTodayRequest(ctx context.Context, req *oardomain.ObservationAnalyzerRequest) (*osumdomain.ObservationCount, error) {
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
		Str("func", "analyzerProcessCountForTodayRequest").
		Msg("analyzer")

	start := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, month, day+1, 0, 0, 0, 0, time.UTC)
	oc := &osumdomain.ObservationCount{
		EntityID:  req.EntityID,
		Start:     start,
		Finish:    end,
		Day:       day,
		Week:      week,
		Month:     int(month),
		Year:      year,
		Frequency: osumdomain.ObservationCountDayFrequency,
		Result:    0,
	}

	oc, err := uc.ObservationCountRepo.InsertOrGetByPrimaryKey(ctx, oc)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return nil, err
	}

	////
	//// Perform our computation.
	////

	if req.Type == oardomain.ObservationAnalyzerRequestInsertType {
		oc.Result++
	} else if req.Type == oardomain.ObservationAnalyzerRequestDeleteType && oc.Result > 0 {
		oc.Result--
	}

	////
	//// Save record.
	////

	err = uc.ObservationCountRepo.UpdateByPrimaryKey(ctx, oc)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return nil, err
	}
	return oc, nil
}

func (uc observationAnalyzerRequestUsecase) analyzerProcessCountForThisWeekRequest(ctx context.Context, req *oardomain.ObservationAnalyzerRequest) (*osumdomain.ObservationCount, error) {
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
		Str("func", "analyzerProcessCountForThisWeekRequest").
		Msg("analyzer")

	start := timekit.GetFirstDateFromWeekAndYear(week, year, time.UTC)
	end := timekit.GetFirstDateFromWeekAndYear(week+1, year, time.UTC)
	oc := &osumdomain.ObservationCount{
		EntityID:  req.EntityID,
		Start:     start,
		Finish:    end,
		Day:       0,
		Week:      week,
		Month:     int(month),
		Year:      year,
		Frequency: osumdomain.ObservationCountWeekFrequency,
		Result:    0,
	}

	oc, err := uc.ObservationCountRepo.InsertOrGetByPrimaryKey(ctx, oc)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return nil, err
	}

	////
	//// Perform our computation.
	////

	if req.Type == oardomain.ObservationAnalyzerRequestInsertType {
		oc.Result++
	} else if req.Type == oardomain.ObservationAnalyzerRequestDeleteType && oc.Result > 0 {
		oc.Result--
	}

	////
	//// Save record.
	////

	err = uc.ObservationCountRepo.UpdateByPrimaryKey(ctx, oc)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return nil, err
	}
	return oc, nil
}

func (uc observationAnalyzerRequestUsecase) analyzerProcessCountForThisMonthRequest(ctx context.Context, req *oardomain.ObservationAnalyzerRequest) (*osumdomain.ObservationCount, error) {
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
		Str("func", "analyzerProcessCountForThisMonthRequest").
		Msg("analyzer")

	start := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
	oc := &osumdomain.ObservationCount{
		EntityID:  req.EntityID,
		Start:     start,
		Finish:    end,
		Day:       0,
		Week:      0,
		Month:     int(month),
		Year:      year,
		Frequency: osumdomain.ObservationCountMonthFrequency,
		Result:    0,
	}

	oc, err := uc.ObservationCountRepo.InsertOrGetByPrimaryKey(ctx, oc)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return nil, err
	}

	////
	//// Perform our computation.
	////

	if req.Type == oardomain.ObservationAnalyzerRequestInsertType {
		oc.Result++
	} else if req.Type == oardomain.ObservationAnalyzerRequestDeleteType && oc.Result > 0 {
		oc.Result--
	}

	////
	//// Save record.
	////

	err = uc.ObservationCountRepo.UpdateByPrimaryKey(ctx, oc)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return nil, err
	}
	return oc, nil
}

func (uc observationAnalyzerRequestUsecase) analyzerProcessCountForThisYearRequest(ctx context.Context, req *oardomain.ObservationAnalyzerRequest) (*osumdomain.ObservationCount, error) {
	////
	//// Fetch or create our record.
	////

	o := req.Observation
	year := o.Timestamp.Year()

	uc.Logger.Info().
		Uint64("entity", req.EntityID).
		Str("uuid", req.UUID).
		Int("year", year).
		Str("func", "analyzerProcessCountForThisYearRequest").
		Msg("analyzer")

	start := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(year+1, 1, 1, 0, 0, 0, 0, time.UTC)
	oc := &osumdomain.ObservationCount{
		EntityID:  req.EntityID,
		Start:     start,
		Finish:    end,
		Day:       0,
		Week:      0,
		Month:     0,
		Year:      year,
		Frequency: osumdomain.ObservationCountYearFrequency,
		Result:    0,
	}

	oc, err := uc.ObservationCountRepo.InsertOrGetByPrimaryKey(ctx, oc)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return nil, err
	}

	////
	//// Perform our computation.
	////

	if req.Type == oardomain.ObservationAnalyzerRequestInsertType {
		oc.Result++
	} else if req.Type == oardomain.ObservationAnalyzerRequestDeleteType && oc.Result > 0 {
		oc.Result--
	}

	////
	//// Save record.
	////

	err = uc.ObservationCountRepo.UpdateByPrimaryKey(ctx, oc)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return nil, err
	}
	return oc, nil
}
