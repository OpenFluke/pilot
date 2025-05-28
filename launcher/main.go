package main

import (
	"fmt"
	"os"

	"pilot"
	"pilot/experiments"
)

func main() {
	fmt.Println("🚀 Launcher starting PILOT...")

	mnist := experiments.NewMNISTDatasetStage("./data/mnist")
	exp := pilot.NewExperiment("MNIST", mnist)

	if err := exp.RunAll(); err != nil {
		fmt.Println("❌ Experiment failed:", err)
		os.Exit(1)
	}
}
