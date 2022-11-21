package usecase

import (
	"context"

	"github.com/bartmika/stockyard/internal/domain/observation"
	oardomain "github.com/bartmika/stockyard/internal/domain/observation_analyzer_request"
)

func (uc observationUsecase) Insert(ctx context.Context, o *observation.Observation) (*observation.Observation, error) {
	if uc.HasAnalyzer {
		uc.KMutex.Lockf("entity-id-%d", o.EntityID)
		defer uc.KMutex.Unlockf("entity-id-%d", o.EntityID)
	}

	if err := uc.ObservationRepo.Insert(ctx, o); err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return nil, err
	}
	o, err := uc.ObservationRepo.GetByPrimaryKey(ctx, o.EntityID, o.Timestamp)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return nil, err
	}

	// DEVELOPERS NOTE:
	// If we have analyzer turned on then we will insert into the analyzer
	// the new `observation` into it to be processed.
	if uc.HasAnalyzer {
		oar := &oardomain.ObservationAnalyzerRequest{
			EntityID:    o.EntityID,
			UUID:        uc.UUID.NewUUID(),
			Timestamp:   o.Timestamp,
			Type:        oardomain.ObservationAnalyzerRequestInsertType,
			Observation: o,
		}
		if err := uc.ObservationAnalyzerRequestRepo.Insert(ctx, oar); err != nil {
			uc.Logger.Error().Err(err).Caller().Msg("database error")
			return nil, err
		}
		uc.Logger.Info().
			Str("app_service", "observation.insert").
			Uint64("entity_id", o.EntityID).
			Time("timestamp", o.Timestamp).
			Msg("request made to analyzer")
	}

	return o, nil
}
