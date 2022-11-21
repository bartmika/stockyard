package app

import (
	entity_usecase "github.com/bartmika/stockyard/internal/app/entity/usecase"
	observation_usecase "github.com/bartmika/stockyard/internal/app/observation/usecase"
	observation_anlyz_usecase "github.com/bartmika/stockyard/internal/app/observation_analyzer_request/usecase"
	observation_avg_usecase "github.com/bartmika/stockyard/internal/app/observation_average/usecase"
	observation_cnt_usecase "github.com/bartmika/stockyard/internal/app/observation_count/usecase"
	observation_sum_usecase "github.com/bartmika/stockyard/internal/app/observation_summation/usecase"
	timekey_usecase "github.com/bartmika/stockyard/internal/app/timekey/usecase"
	"github.com/bartmika/stockyard/internal/config"
	"github.com/bartmika/stockyard/internal/interfaceadapters"
	"github.com/bartmika/stockyard/internal/pkg/kmutex"
	"github.com/bartmika/stockyard/internal/pkg/time"
	"github.com/bartmika/stockyard/internal/pkg/uuid"
	"github.com/rs/zerolog"
)

//Services contains all exposed services of the application layer
type Services struct {
	Logger                            *zerolog.Logger
	EntityUsecase                     entity_usecase.Usecase
	ObservationUsecase                observation_usecase.Usecase
	ObservationAnalyzerRequestUsecase observation_anlyz_usecase.Usecase
	ObservationSummationUsecase       observation_sum_usecase.Usecase
	ObservationCountUsecase           observation_cnt_usecase.Usecase
	ObservationAverageUsecase         observation_avg_usecase.Usecase
	TimeKeyUsecase                    timekey_usecase.Usecase
}

// NewAppServices Bootstraps Application Layer dependencies
func NewAppServices(appConf *config.Conf, uuidProvider uuid.Provider, timeProvider time.Provider, kmutexProvider kmutex.Provider, adapters *interfaceadapters.Services) Services {
	return Services{
		Logger: adapters.Logger,
		EntityUsecase: entity_usecase.NewEntityUsecase(
			uuidProvider,
			timeProvider,
			adapters.EntityRepo,
		),
		ObservationUsecase: observation_usecase.NewObservationUsecase(
			appConf,
			adapters.Logger,
			uuidProvider,
			timeProvider,
			kmutexProvider,
			adapters.ObservationRepo,
			adapters.ObservationCountRepo,
			adapters.ObservationSummationRepo,
			adapters.ObservationAverageRepo,
			adapters.ObservationAnalyzerRequestRepo,
		),
		ObservationAnalyzerRequestUsecase: observation_anlyz_usecase.NewObservationAnalyzerRequestUsecase(
			appConf,
			adapters.Logger,
			uuidProvider,
			timeProvider,
			kmutexProvider,
			adapters.ObservationRepo,
			adapters.ObservationCountRepo,
			adapters.ObservationSummationRepo,
			adapters.ObservationAverageRepo,
			adapters.ObservationAnalyzerRequestRepo,
		),
		ObservationSummationUsecase: observation_sum_usecase.NewObservationSummationUsecase(
			uuidProvider,
			timeProvider,
			adapters.ObservationSummationRepo,
		),
		ObservationCountUsecase: observation_cnt_usecase.NewObservationCountUsecase(
			uuidProvider,
			timeProvider,
			adapters.ObservationCountRepo,
		),
		ObservationAverageUsecase: observation_avg_usecase.NewObservationAverageUsecase(
			uuidProvider,
			timeProvider,
			adapters.ObservationAverageRepo,
		),
		TimeKeyUsecase: timekey_usecase.NewTimeKeyUsecase(
			uuidProvider,
			timeProvider,
			adapters.TimeKeyRepo,
		),
	}
}
