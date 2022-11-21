package usecase

import (
	"context"

	oardomain "github.com/bartmika/stockyard/internal/domain/observation_analyzer_request"
)

func (uc observationAnalyzerRequestUsecase) RunAnalyzer(ctx context.Context) error {
	reqs, err := uc.ListAll(ctx)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return err
	}
	for _, req := range reqs {
		uc.KMutex.Lockf("entity-id-%d", req.EntityID)
		defer uc.KMutex.Unlockf("entity-id-%d", req.EntityID)

		if err := uc.analyzerProcessRequest(ctx, req); err != nil {
			uc.Logger.Error().Err(err).Caller().Msg("database error")
			return err
		}

		if err := uc.DeleteByPrimaryKey(ctx, req.EntityID, req.UUID); err != nil {
			uc.Logger.Error().Err(err).Caller().Msg("database error")
			return err
		}
	}
	return nil
}

func (uc observationAnalyzerRequestUsecase) analyzerProcessRequest(ctx context.Context, req *oardomain.ObservationAnalyzerRequest) error {
	uc.Logger.Info().
		Uint64("entity", req.EntityID).
		Str("uuid", req.UUID).
		Str("func", "analyzerProcessRequest").
		Msg("analyzer")

	counts, err := uc.analyzerProcessCountRequest(ctx, req)
	if err != nil {
		return err
	}
	sums, err := uc.analyzerProcessSummationRequest(ctx, req)
	if err != nil {
		return err
	}
	if err := uc.analyzerProcessAvgRequest(ctx, req, counts, sums); err != nil {
		return err
	}
	return nil
}
