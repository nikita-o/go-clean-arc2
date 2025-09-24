package cron

import use_case "clean-arc-2/internal/use-case"

type Crons struct {
}

type DataInitCrons struct {
	UseCases *use_case.UseCases
}

func NewCrons(data DataInitCrons) *Crons {
	return &Crons{}
}
