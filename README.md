# maths [![codecov](https://codecov.io/gh/GoLevel/maths/branch/main/graph/badge.svg?token=KGLHJWUBOO)](https://codecov.io/gh/GoLevel/maths)

A simple game maths package. Add later more and more functions.

## 2D Vector

Support int64 and float64 vectors with X and Y. All Functions are chainable.

```go
vector := maths.NewVector2D[int64]()

vector = vector.Add(maths.NewVector2D[int64](1, 2))
vector = vector.Subtract(maths.NewVector2D[int64](1, 2))
vector = vector.Multiply(2)
vector = vector.Divide(2)
vectorClone = vector.Clone()

distance := vector.Distance(maths.NewVector2D[int64](1, 2))

vectorFloat := vector.ToFloat()
```

## Dependencies

No external dependencies.

## How to Contribute

Make a pull request...

## License

Distributed under MIT License, please see license file within the code for more details.