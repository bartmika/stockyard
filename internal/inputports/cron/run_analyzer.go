package cron

import "context"

func (cronServer *Server) RunAnalyzer() error {
	cronServer.AppServices.Logger.Info().Str("func", "RunAnalyzer").Str("service", "cron").Msg("executing")
	defer cronServer.AppServices.Logger.Info().Str("func", "RunAnalyzer").Str("service", "cron").Msg("executed")
	ctx := context.Background()
	return cronServer.AppServices.ObservationAnalyzerRequestUsecase.RunAnalyzer(ctx)
}
