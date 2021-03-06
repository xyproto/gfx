package gfx

import "image"

func ExampleNewScaledImage() {
	src := NewTile(Palette1Bit, 8, []uint8{
		1, 1, 1, 1, 1, 1, 1, 1,
		1, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 1, 1, 0, 0, 1,
		1, 0, 1, 1, 1, 1, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1,
	})

	dst := NewScaledImage(src, 2.0)

	func(images ...image.Image) {
		for _, m := range images {
			for y := 0; y < m.Bounds().Dy(); y++ {
				for x := 0; x < m.Bounds().Dx(); x++ {
					if r, _, _, _ := m.At(x, y).RGBA(); r == 0 {
						Printf("▓▓")
					} else {
						Printf("░░")
					}
				}
				Printf("\n")
			}
			Printf("\n")
		}
	}(src, dst)

	// Output:
	//
	// ░░░░░░░░░░░░░░░░
	// ░░▓▓▓▓▓▓▓▓▓▓▓▓░░
	// ░░▓▓▓▓░░░░▓▓▓▓░░
	// ░░▓▓░░░░░░░░▓▓░░
	// ░░▓▓▓▓▓▓▓▓▓▓▓▓░░
	// ░░░░░░░░░░░░░░░░
	//
	// ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
	// ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓░░░░░░░░▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓░░░░░░░░▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓░░░░░░░░░░░░░░░░▓▓▓▓░░░░
	// ░░░░▓▓▓▓░░░░░░░░░░░░░░░░▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░░░
	// ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
	// ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
	//
}
