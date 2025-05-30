package experiments

import (
	"fmt"

	"github.com/openfluke/pilot"
)

type MNISTDatasetStage struct {
	Setup *pilot.ExperimentSetup
}

func NewMNISTDatasetStage(basePath string) *MNISTDatasetStage {
	return &MNISTDatasetStage{
		Setup: &pilot.ExperimentSetup{
			Name:    "MNIST",
			BaseDir: basePath,
			MainURL: "https://storage.googleapis.com/cvdf-datasets/mnist/",
			Files: map[string]string{
				"train-images-idx3-ubyte.gz": "train-images-idx3-ubyte",
				"train-labels-idx1-ubyte.gz": "train-labels-idx1-ubyte",
				"t10k-images-idx3-ubyte.gz":  "t10k-images-idx3-ubyte",
				"t10k-labels-idx1-ubyte.gz":  "t10k-labels-idx1-ubyte",
			},
		},
	}
}

func (m *MNISTDatasetStage) Name() string { return "MNIST Dataset Prep" }

func (m *MNISTDatasetStage) Run() error {
	fmt.Println("🔧 Running MNIST dataset stage...")
	return m.Setup.Init()
}
