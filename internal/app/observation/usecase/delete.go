package usecase

import (
	"context"
	"errors"
	"time"

	oardomain "github.com/bartmika/stockyard/internal/domain/observation_analyzer_request"
)

func (uc observationUsecase) DeleteByPrimaryKey(ctx context.Context, entityID uint64, timestamp time.Time) error {
	if uc.HasAnalyzer {
		uc.KMutex.Lockf("entity-id-%d", entityID)
		defer uc.KMutex.Unlockf("entity-id-%d", entityID)
	}

	o, err := uc.ObservationRepo.GetByPrimaryKey(ctx, entityID, timestamp)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return err
	}
	if o == nil {
		err = errors.New("does not exist")
		uc.Logger.Warn().Err(err).Caller().Msg("does not exist")
		return err
	}
	if err := uc.ObservationRepo.DeleteByPrimaryKey(ctx, entityID, timestamp); err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return err
	}

	// DEVELOPERS NOTE:
	// If we have analyzer turned on then we will insert into the analyzer
	// the new `observation` into it to be processed.
	if uc.HasAnalyzer {
		oar := &oardomain.ObservationAnalyzerRequest{
			EntityID:    o.EntityID,
			UUID:        uc.UUID.NewUUID(),
			Timestamp:   o.Timestamp,
			Type:        oardomain.ObservationAnalyzerRequestDeleteType,
			Observation: o,
		}
		if err := uc.ObservationAnalyzerRequestRepo.Insert(ctx, oar); err != nil {
			uc.Logger.Error().Err(err).Caller().Msg("database error")
			return err
		}
		uc.Logger.Info().
			Str("app_service", "observation.delete").
			Uint64("entity_id", o.EntityID).
			Time("timestamp", o.Timestamp).
			Msg("request made to analyzer")
	}

	return nil
}
