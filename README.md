# Recognizing Handwritten Digits

This project is a simple implementation of a Neural Network to recognize handwritten digits using the MNIST dataset. The neural network is built from scratch in Go.

# Project Structure

```
.
├── data
│   └── mnist_dataset.csv
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

# References

- https://www.cis.jhu.edu/~sachin/digit/digit.html
- https://github.com/vardhan-siramdasu/Kaggle-Digit-Recognizer/blob/main/data/
- https://www.youtube.com/watch?v=aircAruvnKk&list=PLZHQObOWTQDNU6R1_67000Dx_ZCJB-3pi