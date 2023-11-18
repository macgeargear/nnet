package model

import "gonum.org/v1/gonum/mat"

type neuralNet struct {
	config  neuralNetConfig
	WHidden *mat.Dense
	bHidden *mat.Dense
	wOut    *mat.Dense
	bOut    *mat.Dense
}

type neuralNetConfig struct {
	inputNeurons  int
	outputNeurons int
	hiddenNeurons int
	numEpochs     int
	learningRate  float64
}

func newNetwork(config neuralNetConfig) *neuralNet {
	return &neuralNet{config: config}
}
