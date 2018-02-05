package main

// InputNeuron is an interface for a neuron that takes input from the outside world
type InputNeuron interface {
	GetWeightedValue() float64
}

// Neuron is a single neuron in a neural network that has a series of weights behind it and an activation value
type Neuron struct {
	activation float64
	inputs     []InputNeuron
}

// NewNeuron creates a neuron based on the array of input neurons behind it
func NewNeuron(inputs []InputNeuron) *Neuron {
	return &Neuron{0.0, inputs}
}
