package usecase

import (
	"context"
	"time"

	"github.com/rs/zerolog"

	"github.com/bartmika/stockyard/internal/config"
	"github.com/bartmika/stockyard/internal/domain/observation"
	odomain "github.com/bartmika/stockyard/internal/domain/observation"
	oardomain "github.com/bartmika/stockyard/internal/domain/observation_analyzer_request"
	oadomain "github.com/bartmika/stockyard/internal/domain/observation_average"
	ocdomain "github.com/bartmika/stockyard/internal/domain/observation_count"
	osdomain "github.com/bartmika/stockyard/internal/domain/observation_summation"
	"github.com/bartmika/stockyard/internal/pkg/kmutex"
	timep "github.com/bartmika/stockyard/internal/pkg/time"
	"github.com/bartmika/stockyard/internal/pkg/uuid"
)

// Usecase Provides interface for the observation use cases.
type Usecase interface {
	Insert(ctx context.Context, e *observation.Observation) (ee *observation.Observation, err error)
	CheckIfExistsByPrimaryKey(ctx context.Context, entityID uint64, timestamp time.Time) (bool, error)
	ListAndCountByFilter(ctx context.Context, ef *observation.ObservationFilter) ([]*observation.Observation, uint64, error)
	DeleteByPrimaryKey(ctx context.Context, entityID uint64, timestamp time.Time) error
}

type observationUsecase struct {
	HasAnalyzer                    bool
	Logger                         *zerolog.Logger
	Time                           timep.Provider
	UUID                           uuid.Provider
	KMutex                         kmutex.Provider
	ObservationRepo                odomain.Repository
	ObservationCountRepo           ocdomain.Repository
	ObservationSummationRepo       osdomain.Repository
	ObservationAverageRepo         oadomain.Repository
	ObservationAnalyzerRequestRepo oardomain.Repository
}

// NewObservationUsecase Constructor function for the `UserUsecase` implementation.
func NewObservationUsecase(
	appConf *config.Conf,
	logger *zerolog.Logger,
	uuidp uuid.Provider,
	tp timep.Provider,
	kmutexp kmutex.Provider,
	o odomain.Repository,
	oc ocdomain.Repository,
	os osdomain.Repository,
	oa oadomain.Repository,
	oar oardomain.Repository,

) *observationUsecase {
	return &observationUsecase{
		HasAnalyzer:                    appConf.Setting.HasAnalyzer,
		Logger:                         logger,
		Time:                           tp,
		UUID:                           uuidp,
		KMutex:                         kmutexp,
		ObservationRepo:                o,
		ObservationCountRepo:           oc,
		ObservationSummationRepo:       os,
		ObservationAverageRepo:         oa,
		ObservationAnalyzerRequestRepo: oar,
	}
}
