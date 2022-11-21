package usecase

import (
	"context"

	"github.com/bartmika/stockyard/internal/config"
	odomain "github.com/bartmika/stockyard/internal/domain/observation"
	oardomain "github.com/bartmika/stockyard/internal/domain/observation_analyzer_request"
	oadomain "github.com/bartmika/stockyard/internal/domain/observation_average"
	ocdomain "github.com/bartmika/stockyard/internal/domain/observation_count"
	osdomain "github.com/bartmika/stockyard/internal/domain/observation_summation"
	"github.com/bartmika/stockyard/internal/pkg/kmutex"
	timep "github.com/bartmika/stockyard/internal/pkg/time"
	"github.com/bartmika/stockyard/internal/pkg/uuid"
	"github.com/rs/zerolog"
)

// Usecase Provides interface for the observation use cases.
type Usecase interface {
	Insert(ctx context.Context, o *oardomain.ObservationAnalyzerRequest) (*oardomain.ObservationAnalyzerRequest, error)
	ListAll(ctx context.Context) ([]*oardomain.ObservationAnalyzerRequest, error)
	DeleteByPrimaryKey(ctx context.Context, entityID uint64, uuid string) error
	RunAnalyzer(ctx context.Context) error
}

type observationAnalyzerRequestUsecase struct {
	HasAnalyzer                    bool
	Logger                         *zerolog.Logger
	Time                           timep.Provider
	UUID                           uuid.Provider
	KMutex                         kmutex.Provider
	ObservationAnalyzerRequestRepo oardomain.Repository
	ObservationRepo                odomain.Repository
	ObservationCountRepo           ocdomain.Repository
	ObservationSummationRepo       osdomain.Repository
	ObservationAverageRepo         oadomain.Repository
}

// NewObservationAnalyzerRequestUsecase Constructor function for the `ObservationAnalyzerRequest` implementation.
func NewObservationAnalyzerRequestUsecase(
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

) *observationAnalyzerRequestUsecase {
	return &observationAnalyzerRequestUsecase{
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
