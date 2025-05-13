# Recognizing Handwritten Digits

This project is a simple implementation of a Neural Network to recognize handwritten digits using the MNIST dataset[^siramdasu]. The neural network is built from scratch in Go.

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

The model of neural network we'll be my interpretation of the _Multilayer Perceptron_[^sanderson], which I'll be refering as MLP. The MLP is defined as a tuple of the so called _layers_. We shall define those first.

A layer, informally, takes an vector input and produces an vector output, i.e., its a function. The way it operates by having a variation of the Pitts-McCulloch neuron[^mcculloch;pitts], which follow an "all-or-none" activation principle, our neurons have _activation values_ on the real numbers.

### Definition 1. Neuron

A neuron is a thing that possesses a _activation value_ $`a \in \Bbb{R}`$, connections called _weights_ $`\mathbf{w} = (w_i)`$, $`\mathbf{w} \in \Bbb{R}^n`$, a _bias_ $`b \in \Bbb{R}`$ and an _activation function_ $`\sigma : \Bbb{R} \to \Bbb{R}`$. The activation of a neuron is determined by its input $`\mathbf{u} = (u_i)`$, $`\mathbf{u} \in \Bbb{R}^n`$:

```math
a = \sigma \left( b + \sum_{i = 1}^{n} {w_i u_i} \right)
```

Or simply:

```math
a = \sigma \left( \mathbf{w} \cdot \mathbf{u} + b \right)
```

### Definition 2. Layer

A layer is a collection of neurons, since the activation of a single neuron can be thought of as a dot product added to a bias and passed through a activation function, the natural way to extend this idea is via matrices. A layer is a triple $`L = (\mathbf{W}, \mathbf{b}, \sigma)`$, where, for a layer with $`m_L`$ _input activation values_ and $`n_L`$ neurons:

- $`\mathbf{W}^{(L)} = [w^{(L)}_{i,j}]`$, $`\mathbf{W}^{(L)} \in \mathcal{M}_{n_L \times m_L} (\Bbb{R})`$, is a matrix of weights, whereas $`w^{(L)}_{p,q}`$ is the connection between the $`q`$-th input activation value and the $`p`$-th neuron;

- $`\mathbf{b}^{(L)} = [b^{(L)}_i]`$, $`\mathbf{b}^{(L)} \in \mathcal{M}_{n_L \times 1} (\Bbb{R})`$, is column vector of biases, whereas $`b^{(L)}_p`$ is the bias for the $`p`$-th neuron;

- $`\sigma_L : \Bbb{R} \to \Bbb{R}`$ is an activation function, we allow $`\sigma_L([v_{i,j}]) = [\sigma_L(v_{i,j})]`$.

Given an input activation vector $`\mathbf{u} = [u_i]`$, the activation of each neuron of the layer, $`\mathbf{a}^{(L)} = [a^{(L)}_j]`$ is defined:

```math
a^{(L)}_i = \sigma_L \left( b^{(L)}_{i} + \sum_{j = 1}^{m_L} {w^{(L)}_{i,j} u_j} \right)
```

Or:

```math
\mathbf{a}^{(L)} = \sigma_L(\mathbf{W}^{(L)}\mathbf{u} + \mathbf{b}^{(L)})
```

### Definition 3. Neural Network

A neural network is a collection of finite ordered layers $`\mathcal{N} = (L_i)`$. In notation, to reference, for example, the weight matrix of $`L_i`$, instead of $`\mathbf{W}^{(L_i)}`$, we can do simply $`\mathbf{W}^{(i)}`$. Note that $`m_L = n_{L - 1}`$ for $`1 \le L \le |\mathcal{N}|`$. With that said, the _output activation vector_ of the network is $\mathbf{y} = \mathbf{a}^{(|\mathcal{N}|)}$, achieved through:

```math
\begin{align*}
    a^{(0)}_i(\mathbf{u}) &= u_i \\
    a^{(L)}_i(\mathbf{u}) &= \sigma_L \left( b^{(L)}_{i} + \sum_{j = 1}^{m_L} {w^{(L)}_{i,j} a^{(L - 1)}_j(\mathbf{u})} \right) \\
\end{align*}
```

Or:

```math
\begin{align*}
    \mathbf{a}^{(0)}(\mathbf{u}) &= \mathbf{u} \\
    \mathbf{a}^{(L)}(\mathbf{u}) &= \sigma_L(\mathbf{W}^{(L)}\mathbf{a}^{(L - 1)}(\mathbf{u}) + \mathbf{b}^{(L)}) \\
\end{align*}
```

We will intruduce a variable $`\mathbf{z}^{(L)} = [z^{(L)}_i]`$ that represents the middle step, before the application of $`\sigma_L`$, that is, $`\mathbf{z}^{(L)} = \mathbf{W}^{(L)}\mathbf{a}^{(L - 1)} + \mathbf{b}^{(L)}`$ and $`z^{(L)}_i = b^{(L)}_{i} + \sum_{j = 1}^{m_L} {w^{(L)}_{i,j} a^{(L - 1)}_j}`$. Now we may proceed to training the network.

In order to train a network, we must have an explicit way to tell how good is the network doing. We will measure it through a cost function.

### Definition 4. Cost Function

Given a labeled dataset $`X = \left\{(\mathbf{u}, \mathbf{y}) | \mathbf{u} \in \mathcal{M}_{m_1 \times 1} (\Bbb{R}), \mathbf{y} \in \mathcal{M}_{n_{|\mathcal{N}|} \times 1} (\Bbb{R}) \right\}`$, where $`\mathbf{u}`$ is the _input_ and $`\mathbf{y}`$ is the _label_, the cost function will put the input through the network and analyze it against the label. It is done the following way:

```math
C(X, \mathcal{N}) = \sum_{(\mathbf{u}, \mathbf{y}) \in X} { {\left\lVert \mathbf{a}^{(|\mathcal{N}|)}(\mathbf{u}) - \mathbf{y} \right\rVert}^2 }
```

With that, we will calculate the gradient of the cost function in relation to each parameter in the neural network $`\mathcal{N}`$.

### Example 1

```math
\def\pd#1#2{\frac{\partial#1}{\partial#2}}
```

Consider and neural network $`\mathcal{N}`$, its cost function gradient would look like:

```math
\begin{align*}
    \nabla C
        &= \pd{C}{a^L} \\
        &= \pd{C}{z^L}\pd{z}{a^L} \\
\end{align*}
```

<!-- - https://www.cis.jhu.edu/~sachin/digit/digit.html -->

[^siramdasu]: SIRAMDASU, Vardhan. KAGGLE-DIGIT-RECOGNIZER. **Github repository Kaggle-Digit-Recognizer**. Avaliable in: https://github.com/vardhan-siramdasu/Kaggle-Digit-Recognizer/blob/main/data/. Access in: 11 of May, 2025.

[^sanderson]: SANDERSON, Grant. NEURAL NETWORKS. **Youtube channel 3Blue1Brown**. Available in: https://www.youtube.com/watch?v=aircAruvnKk&list=PLZHQObOWTQDNU6R1_67000Dx_ZCJB-3pi. Access in: 11 of May, 2025.

[^mcculloch;pitts]: McCULLOCH, Warren S.; PITTS, Walter. A LOGICAL CALCULUS OF THE IDEAS IMMANENT IN NERVOUS ACTIVITY\*. **Bulletin of Mathematical Biology**. Vol. 52, No. 1/2, pp. 99-115, 1990. Available in: https://www.cs.cmu.edu/~epxing/Class/10715/reading/McCulloch.and.Pitts.pdf. Access in: 11 of May, 2025.
