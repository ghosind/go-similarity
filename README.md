# go-similarity

![test](https://github.com/ghosind/go-similarity/workflows/test/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/ghosind/go-similarity)](https://goreportcard.com/report/github.com/ghosind/go-similarity)
[![codecov](https://codecov.io/gh/ghosind/go-similarity/branch/main/graph/badge.svg)](https://codecov.io/gh/ghosind/go-similarity)
![Version Badge](https://img.shields.io/github/v/release/ghosind/go-similarity)
![License Badge](https://img.shields.io/github/license/ghosind/go-similarity)
[![Go Reference](https://pkg.go.dev/badge/github.com/ghosind/go-similarity.svg)](https://pkg.go.dev/github.com/ghosind/go-similarity)

Similarity or distance metrics for string implemented on Golang, inspired by [Sam Chapman's SimMetrics library](https://sourceforge.net/projects/simmetrics/).

## Installation

Run the following command to install the package:

```bash
go get -u github.com/ghosind/go-similarity
```

## Getting Started

To use the similarity package, you need to import it in your Go file and create a new instance of the similarity you want to use.

```go
sim := new(similarity.BlockDistance)
similarity := sim.Compare("Hello World", "Hello Go")
fmt.Println(similarity)
// 0.5
```

## Supported Metrics Algorithms

- [Block Distance](https://en.wikipedia.org/wiki/Taxicab_geometry)
- Chapman Length Deviation
- Cosine Similarity
- [Dice Coefficient](https://en.wikipedia.org/wiki/S%C3%B8rensen%E2%80%93Dice_coefficient)
- [Euclidean Distance](https://en.wikipedia.org/wiki/Euclidean_distance)
- [Jaccard Coefficient](https://en.wikipedia.org/wiki/Jaccard_index)
- [Levenshtein Distance](https://en.wikipedia.org/wiki/Levenshtein_distance)
- Matching Coefficient
- [Overlap Coefficient](https://en.wikipedia.org/wiki/Overlap_coefficient)

## Builtin Tokenizers

- Whitespace Tokenizer: Tokenize the input string by whitespace.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
