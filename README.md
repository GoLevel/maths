# maths [![codecov](https://codecov.io/gh/GoLevel/maths/branch/main/graph/badge.svg?token=KGLHJWUBOO)](https://codecov.io/gh/GoLevel/maths)

A simple game maths package. Add later more and more functions.

- [2D Vector](#2d-vector)
- [Hex](#hex)
- [Hex Vector](#hex-vector)
- [Hex Grid](#hex-grid)

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

## Hex

Hex includes a hex vector with helpers and a hex grid to work with hex tiles.

### Hex Vector

```go
hex := NewHex[float64](10, 10)

hex = hex.Add(maths.NewHex[float64](1, 2))
hex = hex.Subtract(maths.NewHex[float64](1, 2))
hex = hex.Multiply(2)
hex = hex.Divide(2)
clone := hex.Clone()
hex = hex.Round()

distance := hex.Distance(maths.NewHex[float64](1, 2))
// neighbor hexes
neighbors := hex.Neighbors()

hexes := hex.Circle(2)
hexes := hex.Line(maths.NewHex[float64](1, 2))
hexes := hex.SpiralRing(2)
hexes := hex.Spiral(2)
```

### Hex Grid

Supports hex grid with flat and pointy layout. All Functions are chainable.

```go
hexSize := maths.NewVector2D[float64](32, 32)
hexGrid := maths.NewHexGrid(maths.LayoutFlat, hexSize)

// world position
screenPosition := hexGrid.WorldToScreen(maths.NewVector2D[float64](10, 10), camera)
worldPosition := hexGrid.ScreenToWorld(screenPosition, camera)

// hex position like tiles
screenPosition := hexGrid.HexToScreen(maths.NewVector2D[float64](10, 10), camera)
hexPosition := hexGrid.ScreenToHex(screenPosition, camera)

// all visible hexes in camera view
hexes := hexGrid.GetVisibleHexes(camera)

// all corners of a hex as world position
corners := hexGrid.HexCorners(maths.NewVector2D[float64](10, 10))
// all corners of a hex as screen position
corners := hexGrid.HexCornersScreen(maths.NewVector2D[float64](10, 10), camera)

// render position of a hex with scale factor.
imageDefaultSize := maths.NewVector2D[float64](64, 64)
hexScreenPosition, scaleFactor := hexGrid.HexImageToScreen(maths.NewVector2D[float64](10, 10), imageDefaultSize, camera)
```

## Dependencies

No external dependencies. Only for testing purposes.

## How to Contribute

Make a pull request...

## License

Distributed under MIT License, please see license file within the code for more details.