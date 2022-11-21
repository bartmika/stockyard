package usecase

import (
	"context"
	"time"

	oardomain "github.com/bartmika/stockyard/internal/domain/observation_analyzer_request"
	oadomain "github.com/bartmika/stockyard/internal/domain/observation_average"
	ocdomain "github.com/bartmika/stockyard/internal/domain/observation_count"
	osdomain "github.com/bartmika/stockyard/internal/domain/observation_summation"
	"github.com/bartmika/timekit"
)

func (uc observationAnalyzerRequestUsecase) analyzerProcessAvgRequest(ctx context.Context, req *oardomain.ObservationAnalyzerRequest, counts map[int8]*ocdomain.ObservationCount, sums map[int8]*osdomain.ObservationSummation) error {
	uc.Logger.Info().
		Uint64("entity", req.EntityID).
		Str("uuid", req.UUID).
		Str("func", "analyzerProcessAvgRequest").
		Msg("analyzer")

	if err := uc.analyzerProcessAverageForTodayRequest(ctx, req, counts[ocdomain.ObservationCountDayFrequency], sums[osdomain.ObservationSummationDayFrequency]); err != nil {
		return err
	}
	if err := uc.analyzerProcessAverageForThisWeekRequest(ctx, req, counts[ocdomain.ObservationCountWeekFrequency], sums[osdomain.ObservationSummationWeekFrequency]); err != nil {
		return err
	}
	if err := uc.analyzerProcessAverageForThisMonthRequest(ctx, req, counts[ocdomain.ObservationCountMonthFrequency], sums[osdomain.ObservationSummationMonthFrequency]); err != nil {
		return err
	}
	if err := uc.analyzerProcessAverageForThisYearRequest(ctx, req, counts[ocdomain.ObservationCountYearFrequency], sums[osdomain.ObservationSummationYearFrequency]); err != nil {
		return err
	}

	return nil
}

func (uc observationAnalyzerRequestUsecase) analyzerProcessAverageForTodayRequest(ctx context.Context, req *oardomain.ObservationAnalyzerRequest, count *ocdomain.ObservationCount, sum *osdomain.ObservationSummation) error {
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
		Str("func", "analyzerProcessAverageForTodayRequest").
		Msg("analyzer")

	start := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, month, day+1, 0, 0, 0, 0, time.UTC)
	oc := &oadomain.ObservationAverage{
		EntityID:  req.EntityID,
		Start:     start,
		Finish:    end,
		Day:       day,
		Week:      week,
		Month:     int(month),
		Year:      year,
		Frequency: oadomain.ObservationAverageDayFrequency,
		Result:    0,
	}

	oc, err := uc.ObservationAverageRepo.InsertOrGetByPrimaryKey(ctx, oc)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return err
	}

	////
	//// Perform our computation.
	////

	if count.Result == 0 {
		oc.Result = 0 // Protect again divide by zero error.
	} else {
		oc.Result = sum.Result / count.Result
	}

	////
	//// Save record.
	////

	err = uc.ObservationAverageRepo.UpdateByPrimaryKey(ctx, oc)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return err
	}
	return nil
}

func (uc observationAnalyzerRequestUsecase) analyzerProcessAverageForThisWeekRequest(ctx context.Context, req *oardomain.ObservationAnalyzerRequest, count *ocdomain.ObservationCount, sum *osdomain.ObservationSummation) error {
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
		Str("func", "analyzerProcessAverageForThisWeekRequest").
		Msg("analyzer")

	start := timekit.GetFirstDateFromWeekAndYear(week, year, time.UTC)
	end := timekit.GetFirstDateFromWeekAndYear(week+1, year, time.UTC)
	oc := &oadomain.ObservationAverage{
		EntityID:  req.EntityID,
		Start:     start,
		Finish:    end,
		Day:       0,
		Week:      week,
		Month:     int(month),
		Year:      year,
		Frequency: oadomain.ObservationAverageWeekFrequency,
		Result:    0,
	}

	oc, err := uc.ObservationAverageRepo.InsertOrGetByPrimaryKey(ctx, oc)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return err
	}

	////
	//// Perform our computation.
	////

	if count.Result == 0 {
		oc.Result = 0
	} else {
		oc.Result = sum.Result / count.Result
	}

	////
	//// Save record.
	////

	err = uc.ObservationAverageRepo.UpdateByPrimaryKey(ctx, oc)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return err
	}
	return nil
}

//////

func (uc observationAnalyzerRequestUsecase) analyzerProcessAverageForThisMonthRequest(ctx context.Context, req *oardomain.ObservationAnalyzerRequest, count *ocdomain.ObservationCount, sum *osdomain.ObservationSummation) error {
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
		Str("func", "analyzerProcessAverageForThisMonthRequest").
		Msg("analyzer")

	start := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
	oc := &oadomain.ObservationAverage{
		EntityID:  req.EntityID,
		Start:     start,
		Finish:    end,
		Day:       0,
		Week:      0,
		Month:     int(month),
		Year:      year,
		Frequency: oadomain.ObservationAverageMonthFrequency,
		Result:    0,
	}

	oc, err := uc.ObservationAverageRepo.InsertOrGetByPrimaryKey(ctx, oc)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return err
	}

	////
	//// Perform our computation.
	////

	if count.Result == 0 {
		oc.Result = 0
	} else {
		oc.Result = sum.Result / count.Result
	}

	////
	//// Save record.
	////

	err = uc.ObservationAverageRepo.UpdateByPrimaryKey(ctx, oc)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return err
	}
	return nil
}

func (uc observationAnalyzerRequestUsecase) analyzerProcessAverageForThisYearRequest(ctx context.Context, req *oardomain.ObservationAnalyzerRequest, count *ocdomain.ObservationCount, sum *osdomain.ObservationSummation) error {
	////
	//// Fetch or create our record.
	////

	o := req.Observation
	year := o.Timestamp.Year()

	uc.Logger.Info().
		Uint64("entity", req.EntityID).
		Str("uuid", req.UUID).
		Int("year", year).
		Str("func", "analyzerProcessAverageForThisYearRequest").
		Msg("analyzer")

	start := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(year+1, 1, 1, 0, 0, 0, 0, time.UTC)
	oc := &oadomain.ObservationAverage{
		EntityID:  req.EntityID,
		Start:     start,
		Finish:    end,
		Day:       0,
		Week:      0,
		Month:     0,
		Year:      year,
		Frequency: oadomain.ObservationAverageYearFrequency,
		Result:    0,
	}

	oc, err := uc.ObservationAverageRepo.InsertOrGetByPrimaryKey(ctx, oc)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return err
	}

	////
	//// Perform our computation.
	////

	if count.Result == 0 {
		oc.Result = 0
	} else {
		oc.Result = sum.Result / count.Result
	}

	////
	//// Save record.
	////

	err = uc.ObservationAverageRepo.UpdateByPrimaryKey(ctx, oc)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return err
	}
	return nil
}
