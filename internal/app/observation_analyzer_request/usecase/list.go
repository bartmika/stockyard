package usecase

import (
	"context"

	oardomain "github.com/bartmika/stockyard/internal/domain/observation_analyzer_request"
)

func (uc observationAnalyzerRequestUsecase) ListAll(ctx context.Context) ([]*oardomain.ObservationAnalyzerRequest, error) {
	arrCh := make(chan []*oardomain.ObservationAnalyzerRequest)

	go func() {
		arr, err := uc.ObservationAnalyzerRequestRepo.ListAll(ctx)
		if err != nil {
			arrCh <- nil
			return
		}
		arrCh <- arr[:]
	}()

	arr := <-arrCh

	return arr, nil
}
