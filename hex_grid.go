package maths

import (
	"math"
)

type Camera interface {
	GetPosition() Vector2D[float64]
	GetZoom() float64
	GetSize() Vector2D[float64]
}

// HexOrientation defines whether hexagons are flat-topped or pointy-topped
type HexOrientation struct {
	F0, F1, F2, F3 float64
	B0, B1, B2, B3 float64
	StartAngle     float64
}

// HexLayout represents the orientation and size of hexagons
type HexLayout struct {
	Orientation HexOrientation
	Size        Vector2D[float64]
	Origin      Vector2D[float64]
	Zoom        float64
}

var (
	LayoutPointy = HexOrientation{
		F0: math.Sqrt(3.0), F1: math.Sqrt(3.0) / 2.0, F2: 0.0, F3: 3.0 / 2.0,
		B0: math.Sqrt(3.0) / 3.0, B1: -1.0 / 3.0, B2: 0.0, B3: 2.0 / 3.0,
		StartAngle: 0.5,
	}
	LayoutFlat = HexOrientation{
		F0: 3.0 / 2.0, F1: 0.0, F2: math.Sqrt(3.0) / 2.0, F3: math.Sqrt(3.0),
		B0: 2.0 / 3.0, B1: 0.0, B2: -1.0 / 3.0, B3: math.Sqrt(3.0) / 3.0,
		StartAngle: 0.0,
	}
)

// NewHexLayout creates a new layout with specified parameters
func NewHexLayout(
	orientation HexOrientation,
	size Vector2D[float64],
	origin Vector2D[float64],
	zoom float64,
) HexLayout {
	return HexLayout{
		Orientation: orientation,
		Size:        size,
		Origin:      origin,
		Zoom:        zoom,
	}
}

// HexToVector2D converts hex coordinates to screen coordinates
func (layout HexLayout) HexToVector2D(h Hex[float64]) Vector2D[float64] {
	o := layout.Orientation
	size := Vector2D[float64]{
		X: float64(layout.Size.X) * layout.Zoom,
		Y: float64(layout.Size.Y) * layout.Zoom,
	}

	x := (o.F0*float64(h.Q) + o.F1*float64(h.R)) * float64(size.X)
	y := (o.F2*float64(h.Q) + o.F3*float64(h.R)) * float64(size.Y)

	return Vector2D[float64]{
		X: x + layout.Origin.X,
		Y: y + layout.Origin.Y,
	}
}

// Vector2DToHex converts screen coordinates to hex coordinates
func (layout HexLayout) Vector2DToHex(p Vector2D[float64]) Hex[float64] {
	o := layout.Orientation
	size := Vector2D[float64]{
		X: float64(layout.Size.X) * layout.Zoom,
		Y: float64(layout.Size.Y) * layout.Zoom,
	}

	pt := Vector2D[float64]{
		X: (float64(p.X - layout.Origin.X)) / float64(size.X),
		Y: (float64(p.Y - layout.Origin.Y)) / float64(size.Y),
	}

	q := o.B0*pt.X + o.B1*pt.Y
	r := o.B2*pt.X + o.B3*pt.Y

	return Hex[float64]{Q: q, R: r}
}

// HexGrid represents the entire hexagonal grid system with camera integration
type HexGrid struct {
	Layout HexLayout
}

// NewHexGrid creates a new hex grid with the given base configuration
func NewHexGrid(
	orientation HexOrientation,
	hexSize Vector2D[float64],
) *HexGrid {
	return &HexGrid{
		Layout: NewHexLayout(
			orientation,
			hexSize,
			NewVector2D[float64](0, 0),
			1.0,
		),
	}
}

// WorldToScreen converts world coordinates to screen coordinates using camera data
func (grid *HexGrid) WorldToScreen(
	worldPos Vector2D[float64],
	camera Camera,
) Vector2D[float64] {
	cameraPos := camera.GetPosition()
	cameraZoom := camera.GetZoom()
	cameraSize := camera.GetSize()

	// Convert to screen space considering camera position and zoom
	screenX := float64(worldPos.X-cameraPos.X)*cameraZoom + float64(cameraSize.X)/2
	screenY := float64(worldPos.Y-cameraPos.Y)*cameraZoom + float64(cameraSize.Y)/2

	return NewVector2D(screenX, screenY)
}

// ScreenToWorld converts screen coordinates to world coordinates using camera data
func (grid *HexGrid) ScreenToWorld(
	screenPos Vector2D[float64],
	camera Camera,
) Vector2D[float64] {
	cameraPos := camera.GetPosition()
	cameraZoom := camera.GetZoom()
	cameraSize := camera.GetSize()

	// Convert back to world space
	worldX := float64(screenPos.X-cameraSize.X/2)/cameraZoom + float64(cameraPos.X)
	worldY := float64(screenPos.Y-cameraSize.Y/2)/cameraZoom + float64(cameraPos.Y)

	return NewVector2D(worldX, worldY)
}

// HexToScreen converts hex coordinates to screen coordinates considering camera
func (grid *HexGrid) HexToScreen(
	hex Hex[float64],
	camera Camera,
) Vector2D[float64] {
	// First convert hex to world coordinates
	worldPos := grid.Layout.HexToVector2D(hex)
	// Then convert world coordinates to screen coordinates
	return grid.WorldToScreen(worldPos, camera)
}

// ScreenToHex converts screen coordinates to hex coordinates considering camera
func (grid *HexGrid) ScreenToHex(
	screenPos Vector2D[float64],
	camera Camera,
) Hex[float64] {
	// First convert screen to world coordinates
	worldPos := grid.ScreenToWorld(screenPos, camera)

	// Update layout zoom with camera zoom
	grid.Layout.Zoom = camera.GetZoom()

	// Convert world coordinates to hex
	fractionalHex := grid.Layout.Vector2DToHex(worldPos)
	return fractionalHex.Round()
}

// GetVisibleHexes returns all hexes that are currently visible in the camera view
func (grid *HexGrid) GetVisibleHexes(camera Camera) []Hex[float64] {
	cameraSize := camera.GetSize()

	// Calculate the bounds of the visible area in world coordinates
	topLeft := grid.ScreenToWorld(NewVector2D[float64](0, 0), camera)
	bottomRight := grid.ScreenToWorld(NewVector2D[float64](cameraSize.X, cameraSize.Y), camera)

	// Convert to hex coordinates and add some padding
	startHex := grid.Layout.Vector2DToHex(topLeft).Round()
	endHex := grid.Layout.Vector2DToHex(bottomRight).Round()

	// Calculate the range of q and r coordinates
	minQ := int(math.Min(float64(startHex.Q), float64(endHex.Q))) - 1
	maxQ := int(math.Max(float64(startHex.Q), float64(endHex.Q))) + 1
	minR := int(math.Min(float64(startHex.R), float64(endHex.R))) - 1
	maxR := int(math.Max(float64(startHex.R), float64(endHex.R))) + 1

	visibleHexes := make([]Hex[float64], 0)

	// Iterate through the range and collect visible hexes
	for q := minQ; q <= maxQ; q++ {
		for r := minR; r <= maxR; r++ {
			hex := NewHex[float64](float64(q), float64(r))
			worldPos := grid.Layout.HexToVector2D(hex)

			// Check if the hex center is within the visible area
			if float64(worldPos.X) >= float64(topLeft.X) &&
				float64(worldPos.X) <= float64(bottomRight.X) &&
				float64(worldPos.Y) >= float64(topLeft.Y) &&
				float64(worldPos.Y) <= float64(bottomRight.Y) {
				visibleHexes = append(visibleHexes, hex)
			}
		}
	}

	return visibleHexes
}

// HexCorners returns the corners of a hexagon in world coordinates
func (layout HexLayout) HexCorners(h Hex[float64]) []Vector2D[float64] {
	corners := make([]Vector2D[float64], 6)
	center := layout.HexToVector2D(h)
	size := Vector2D[float64]{
		X: float64(layout.Size.X) * layout.Zoom,
		Y: float64(layout.Size.Y) * layout.Zoom,
	}

	for i := 0; i < 6; i++ {
		angle := 2.0 * math.Pi * (float64(i) + layout.Orientation.StartAngle) / 6.0
		corners[i] = Vector2D[float64]{
			X: float64(center.X) + float64(size.X)*math.Cos(angle),
			Y: float64(center.Y) + float64(size.Y)*math.Sin(angle),
		}
	}
	return corners
}

// HexCornerScreen returns the corners of a hexagon in screen coordinates
func (grid *HexGrid) HexCornerScreen(
	hex Hex[float64],
	camera Camera,
) []Vector2D[float64] {
	corners := grid.Layout.HexCorners(hex)
	screenCorners := make([]Vector2D[float64], len(corners))

	for i, corner := range corners {
		screenCorners[i] = grid.WorldToScreen(corner, camera)
	}

	return screenCorners
}

// HexImageToScreen calculates the screen position and scale factor for an image in a hex cell
func (grid *HexGrid) HexImageToScreen(
	hex Hex[float64],
	imageDefaultSize Vector2D[float64],
	camera Camera,
) (centerPosition Vector2D[float64], imageScaleFactor Vector2D[float64]) {
	// Get the hex center position in screen coordinates
	centerPosition = grid.HexToScreen(hex, camera)

	// Calculate the current hex size in screen coordinates
	hexSize := Vector2D[float64]{
		X: grid.Layout.Size.X * grid.Layout.Zoom * camera.GetZoom(),
		Y: grid.Layout.Size.Y * grid.Layout.Zoom * camera.GetZoom(),
	}

	// Calculate scale factors needed to fit the image within the hex
	// We use the smaller scale factor to maintain aspect ratio
	scaleX := hexSize.X / imageDefaultSize.X
	scaleY := hexSize.Y / imageDefaultSize.Y
	scaleFactor := math.Min(scaleX, scaleY)

	// Apply the scale factor to both dimensions to maintain aspect ratio
	imageScaleFactor = Vector2D[float64]{
		X: scaleFactor,
		Y: scaleFactor,
	}

	return centerPosition, imageScaleFactor
}
