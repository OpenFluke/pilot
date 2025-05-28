package pilot

import "fmt"

type Stage interface {
	Name() string
	Run() error
}

type Experiment struct {
	Name   string
	Stages []Stage
}

func NewExperiment(name string, stages ...Stage) *Experiment {
	return &Experiment{
		Name:   name,
		Stages: stages,
	}
}

func (e *Experiment) RunAll() error {
	fmt.Printf("🚀 Running experiment: %s\n", e.Name)
	for _, stage := range e.Stages {
		fmt.Printf("⚙️  Stage: %s\n", stage.Name())
		if err := stage.Run(); err != nil {
			return fmt.Errorf("stage %s failed: %w", stage.Name(), err)
		}
	}
	fmt.Println("✅ All stages completed successfully.")
	return nil
}
