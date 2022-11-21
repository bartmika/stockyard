package interfaceadapters

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	entity_r "github.com/bartmika/stockyard/internal/app/entity/repository"
	observation_r "github.com/bartmika/stockyard/internal/app/observation/repository"
	observation_analyz_r "github.com/bartmika/stockyard/internal/app/observation_analyzer_request/repository"
	observation_avg_r "github.com/bartmika/stockyard/internal/app/observation_average/repository"
	observation_cnt_r "github.com/bartmika/stockyard/internal/app/observation_count/repository"
	observation_sum_r "github.com/bartmika/stockyard/internal/app/observation_summation/repository"
	timekey_r "github.com/bartmika/stockyard/internal/app/timekey/repository"
	"github.com/bartmika/stockyard/internal/config"
	entity_d "github.com/bartmika/stockyard/internal/domain/entity"
	observation_d "github.com/bartmika/stockyard/internal/domain/observation"
	observation_analyz_d "github.com/bartmika/stockyard/internal/domain/observation_analyzer_request"
	observation_avg_d "github.com/bartmika/stockyard/internal/domain/observation_average"
	observation_cnt_d "github.com/bartmika/stockyard/internal/domain/observation_count"
	observation_sum_d "github.com/bartmika/stockyard/internal/domain/observation_summation"
	timekey_d "github.com/bartmika/stockyard/internal/domain/timekey"
	"github.com/bartmika/stockyard/internal/interfaceadapters/migrations"
	"github.com/bartmika/stockyard/internal/interfaceadapters/storage/postgres"
)

// Services contains the exposed services of interface adapters
type Services struct {
	Logger                         *zerolog.Logger
	EntityRepo                     entity_d.Repository
	TimeKeyRepo                    timekey_d.Repository
	ObservationAnalyzerRequestRepo observation_analyz_d.Repository
	ObservationRepo                observation_d.Repository
	ObservationSummationRepo       observation_sum_d.Repository
	ObservationCountRepo           observation_cnt_d.Repository
	ObservationAverageRepo         observation_avg_d.Repository
}

// NewServices Instantiates the interface adapter services
func NewServices(appConf *config.Conf) (*Services, error) {
	// Step 2: Connect to database.
	db, err := postgres.ConnectDB(appConf)
	if err != nil {
		return nil, err
	}

	// Step 2: Perform our automatic database migrations (if enabled)
	if appConf.DB.HasAutoMigrations {
		if err := migrations.RunOnDB(db); err != nil {
			return nil, err
		}
	} else {
		log.Warn().Msg("No migrations occured - you must do this manually.")
	}

	// Step 3: Default level for this example is info, unless debug flag is present
	var logger zerolog.Logger
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if appConf.Setting.HasDebugging {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// The following line of code adds a pretty output to the console.
	logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().
		Timestamp(). // Add timestamp to every call.
		Logger()

	log.Logger = logger

	return &Services{
		Logger:                         &logger,
		EntityRepo:                     entity_r.NewEntityRepoImpl(db, &logger),
		ObservationAnalyzerRequestRepo: observation_analyz_r.NewObservationAnalyzerRequestRepoImpl(db, &logger),
		ObservationRepo:                observation_r.NewObservationRepoImpl(db, &logger),
		ObservationSummationRepo:       observation_sum_r.NewObservationSummationRepoImpl(db, &logger),
		ObservationCountRepo:           observation_cnt_r.NewObservationCountRepoImpl(db, &logger),
		ObservationAverageRepo:         observation_avg_r.NewObservationAverageRepoImpl(db, &logger),
		TimeKeyRepo:                    timekey_r.NewTimeKeyRepoImpl(db, &logger),
	}, nil
}
