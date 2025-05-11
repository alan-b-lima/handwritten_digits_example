# Recognizing Handwritten Digits

This project is a simple implementation of a Neural Network to recognize handwritten digits using the MNIST dataset. The neural network is built from scratch in Go.

# Project Structure

```
.
├── data
│   ├── mnist_labeled_test.csv
│   ├── mnist_training.csv
│   └── mnist_unlabeled_test.csv
├── src
│   ├── dataset
│   │   ├── digit.go
│   │   └── load_dataset.go
│   ├── model
│   │   ├── cost_function.go
│   │   └── neural_network.go
│   ├── nnmath
│   │   ├── gradient.go
│   │   └── matrix.go
│   └── main.go
├── go.mod
└── README.md
```
<!-- └──│  ├── -->

# The Project

The goal of this project is to learn about one of the most basic forms of neural network, as well as to train a neural network to recognize handwritten digits. To achive this last goal, shall we get into some math.

## The Neural Network

The model of neural network we'll be my interpretation of the _Multilayer Perceptron_, which I'll be refering as MLP. The MLP is defined as a tuple of the so called _layers_. We shall define those first.

A layer, informally, takes an vector input and produces an vector output, i.e., its a function. The way it operates by having a variation of the Pitts-McCulloch neuron[^1], which follow an "all-or-none" activation principle, our neurons have _activation values_ on the real numbers.

### Definition 1. Neuron

A neuron is a thing that possesses a _activation value_ $`x \in \Bbb{R}`$, connections called _weights_ $`\mathbf{w} = (w_i)`$, $`\mathbf{w} \in \Bbb{R}^n`$, a _bias_ $`b \in \Bbb{R}`$ and an _activation function_ $`\sigma : \Bbb{R} \to \Bbb{R}`$. The activation of a neuron is determined by its input $`\mathbf{u} = (u_i)`$, $`\mathbf{u} \in \Bbb{R}^n`$:

```math
x = \sigma \left( b + \sum_{i = 1}^{n} {w_i u_i} \right)
```

Or simply:

```math
x = \sigma \left( \mathbf{w} \cdot \mathbf{u} + b \right)
```

### Definition 2. Layer

A layer is a triple $L = (\mathbf{W}, \mathbf{b}, \sigma)$, where, for a layer with $`m`$ input activation values and $`n`$ neurons:

- $`\mathbf{W} = [w_{i,j}]`$, $`\mathbf{W} \in \mathcal{M}_{n \times m} (\Bbb{R})`$ is a weight matrix, the weight $`w_{p,q}`$ is the weight of the conection between the $`p`$-th position of the input vector and the $`q`$-th neuron;

- $`\mathbf{b} = [b_{i}]`$, $`\mathbf{b} \in \mathcal{M}_{n \times 1} (\Bbb{R})`$ is a bias column vector, whereas $`b_p`$ is the bias of the $`p`$-th neuron;

- $`\sigma : \Bbb{R} \to \Bbb{R}`$ is an activation function.

Given that, the activation column vector $`\mathbf{x} \in \mathcal{M}_{n \times 1} (\Bbb{R})`$ produced by this layer on the input column vector $`\mathbf{u} \in \mathcal{M}_{m \times 1} (\Bbb{R})`$ is:

```math
\mathbf{x} = L.A(\mathbf{u}) = \sigma \left( \mathbf{W}\mathbf{u} + \mathbf{b} \right)
```

### Definition 3. Neural Network

A neural network $`\mathcal{N}`$ is a ordered collection of possibly varied size layers. $`\mathcal{N} = (L_i)`$ for $`1 \le i \in \Bbb{N}`$, where $`L_i`$ is a layer such that is weight matrix height is equal to $`L_{i + 1}`$ weight matrix width, if it exists. The _input vector_ $`\mathbf{u} \in \mathcal{M}_{m \times 1} (\Bbb{R})`$ has $`m`$ as the width of $`L_1`$ and the _output vector_ $`\mathbf{x} \in \mathcal{M}_{n \times 1} (\Bbb{R})`$ has $`n`$ as the height of the last $`L_i`$. The aplication of the neural network is given:

```math
\begin{align*}
    \mathbf{x} &= ().f(\mathbf{u}) = \mathbf{u} \\
    \mathbf{x} &= (L_1, L_2, \dots, L_n).f(\mathbf{u}) = (L_2, \dots, L_n).(L_1.A(\mathbf{u}))
\end{align*}
```

# References

- https://www.cis.jhu.edu/~sachin/digit/digit.html
- https://github.com/vardhan-siramdasu/Kaggle-Digit-Recognizer/blob/main/data/
- https://www.youtube.com/watch?v=aircAruvnKk&list=PLZHQObOWTQDNU6R1_67000Dx_ZCJB-3pi

[^1]: McCULLOCH, Warren S.; PITTS, Walter. A LOGICAL CALCULUS OF THE IDEAS IMMANENT IN NERVOUS ACTIVITY\*. **Bulletin of Mathematical Biology**. Vol. 52, No. 1/2, pp. 99-115, 1990. Available in: https://www.cs.cmu.edu/~epxing/Class/10715/reading/McCulloch.and.Pitts.pdf. Access in: 11 of May, 2025.
