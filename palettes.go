package gfx

// Palettes is a slice of Palette.
type Palettes []Palette

// PaletteName is a palette name.
type PaletteName string

// PaletteLookup is a map of PaletteName to Palette.
type PaletteLookup map[PaletteName]Palette

// PaletteByName is a map of all palettes by name.
var PaletteByName = PaletteLookup{
	// Basic palettes
	"1Bit":          Palette1Bit,
	"2BitGrayScale": Palette2BitGrayScale,
	"3Bit":          Palette3Bit,
	"CGA":           PaletteCGA,

	// Palettes by GrafxKid.
	"15PDX": Palette15PDX,
	"20PDX": Palette20PDX,

	// Palettes by Adigun Polack.
	"AAP16":       PaletteAAP16,
	"AAP64":       PaletteAAP64,
	"Splendor128": PaletteSplendor128,

	// Palettes by Arne.
	"Arne16":   PaletteArne16,
	"Famicube": PaletteFamicube,

	// Palettes by Endesga.
	"EDG16": PaletteEDG16,
	"EDG32": PaletteEDG32,
	"EDG36": PaletteEDG36,
	"EDG64": PaletteEDG64,
	"EDG8":  PaletteEDG8,
	"EN4":   PaletteEN4,
	"ARQ4":  PaletteARQ4,

	// Palettes by others.
	"Ink":     PaletteInk,
	"Ammo8":   PaletteAmmo8,
	"NYX8":    PaletteNYX8,
	"Night16": PaletteNight16,
	"PICO8":   PalettePICO8,

	// Palettes from open source projects
	"Tango": PaletteTango,
	"Go":    PaletteGo,

	// Palettes for geo tile maps
	"Basemaps": PaletteBasemaps,
}

// PalettesByNumberOfColors is a map of int to Palettes.
var PalettesByNumberOfColors = func() map[int]PaletteLookup {
	pnc := map[int]PaletteLookup{}

	for n, p := range PaletteByName {
		c := len(p)

		if pnc[c] == nil {
			pnc[c] = PaletteLookup{}
		}

		pnc[c][n] = p
	}

	return pnc
}()

// Palette1Bit is a basic 1-bit (black and white) palette.
var Palette1Bit = Palette{
	{0x00, 0x00, 0x00, 0xFF},
	{0xFF, 0xFF, 0xFF, 0xFF},
}

// Palette2BitGrayScale is a grayscale palette calculated using 2-bits per color.
//
// It was used by the original gameboy and a few other computer systems.
//
// https://lospec.com/palette-list/2-bit-grayscale
var Palette2BitGrayScale = Palette{
	{0x00, 0x00, 0x00, 0xFF},
	{0x67, 0x67, 0x67, 0xFF},
	{0xB6, 0xB6, 0xB6, 0xFF},
	{0xFF, 0xFF, 0xFF, 0xFF},
}

// Palette3Bit is the 3-Bit palette.
//
// A calculated palette using 1 bit for each RGB value.
// It was used by a number of early computers.
//
//
var Palette3Bit = Palette{
	{0x00, 0x00, 0x00, 0xFF},
	{0xFF, 0x00, 0x00, 0xFF},
	{0x00, 0xFF, 0x00, 0xFF},
	{0x00, 0x00, 0xFF, 0xFF},
	{0x00, 0xFF, 0xFF, 0xFF},
	{0xFF, 0x00, 0xFF, 0xFF},
	{0xFF, 0xFF, 0x00, 0xFF},
	{0xFF, 0xFF, 0xFF, 0xFF},
}

// PaletteARQ4 is the ARQ4 palette.
//
// Created by Endesga. #ARQ4
//
// https://lospec.com/palette-list/arq4
var PaletteARQ4 = Palette{
	{0xFF, 0xFF, 0xFF, 0xFF},
	{0x67, 0x72, 0xA9, 0xFF},
	{0x3A, 0x32, 0x77, 0xFF},
	{0x00, 0x00, 0x00, 0xFF},
}

// PaletteCGA is the Color Graphics Adapter palette.
//
// CGA was a graphics card released in 1981 for the IBM PC.
// The standard mode uses one of two 4-color palettes
// (each with a low-intensity and high-intensity mode),
// but a hack allows use of all 16. #cga
//
// https://lospec.com/palette-list/color-graphics-adapter
var PaletteCGA = Palette{
	{0x00, 0x00, 0x00, 0xFF},
	{0x55, 0x55, 0x55, 0xFF},
	{0xAA, 0xAA, 0xAA, 0xFF},
	{0xFF, 0xFF, 0xFF, 0xFF},
	{0x00, 0x00, 0xAA, 0xFF},
	{0x55, 0x55, 0xFF, 0xFF},
	{0x00, 0xAA, 0x00, 0xFF},
	{0x55, 0xFF, 0x55, 0xFF},
	{0x00, 0xAA, 0xAA, 0xFF},
	{0x55, 0xFF, 0xFF, 0xFF},
	{0xAA, 0x00, 0x00, 0xFF},
	{0xFF, 0x55, 0x55, 0xFF},
	{0xAA, 0x00, 0xAA, 0xFF},
	{0xFF, 0x55, 0xFF, 0xFF},
	{0xAA, 0x55, 0x00, 0xFF},
	{0xFF, 0xFF, 0x55, 0xFF},
}

// PaletteEDG8 is the Endesga 8 palette.
//
// Created by Endesga. #EDG8
//
// https://lospec.com/palette-list/endesga-8
var PaletteEDG8 = Palette{
	{0xFD, 0xFD, 0xF8, 0xFF},
	{0xD3, 0x27, 0x34, 0xFF},
	{0xDA, 0x7D, 0x22, 0xFF},
	{0xE6, 0xDA, 0x29, 0xFF},
	{0x28, 0xC6, 0x41, 0xFF},
	{0x2D, 0x93, 0xDD, 0xFF},
	{0x7B, 0x53, 0xAD, 0xFF},
	{0x1B, 0x1C, 0x33, 0xFF},
}

// PaletteEDG16 is the Endesga 16 palette.
//
// Created by Endesga. #EDG16
//
// https://lospec.com/palette-list/endesga-16
var PaletteEDG16 = Palette{
	{0xE4, 0xA6, 0x72, 0xFF},
	{0xB8, 0x6F, 0x50, 0xFF},
	{0x74, 0x3F, 0x39, 0xFF},
	{0x3F, 0x28, 0x32, 0xFF},
	{0x9E, 0x28, 0x35, 0xFF},
	{0xE5, 0x3B, 0x44, 0xFF},
	{0xFB, 0x92, 0x2B, 0xFF},
	{0xFF, 0xE7, 0x62, 0xFF},
	{0x63, 0xC6, 0x4D, 0xFF},
	{0x32, 0x73, 0x45, 0xFF},
	{0x19, 0x3D, 0x3F, 0xFF},
	{0x4F, 0x67, 0x81, 0xFF},
	{0xAF, 0xBF, 0xD2, 0xFF},
	{0xFF, 0xFF, 0xFF, 0xFF},
	{0x2C, 0xE8, 0xF4, 0xFF},
	{0x04, 0x84, 0xD1, 0xFF},
}

// PaletteEDG32 is the Endesga 32 palette.
//
// Created by Endesga for his game NYKRA. #EDG32
//
// https://lospec.com/palette-list/endesga-32
var PaletteEDG32 = Palette{
	{0xBE, 0x4A, 0x2F, 0xFF},
	{0xD7, 0x76, 0x43, 0xFF},
	{0xEA, 0xD4, 0xAA, 0xFF},
	{0xE4, 0xA6, 0x72, 0xFF},
	{0xB8, 0x6F, 0x50, 0xFF},
	{0x73, 0x3E, 0x39, 0xFF},
	{0x3E, 0x27, 0x31, 0xFF},
	{0xA2, 0x26, 0x33, 0xFF},
	{0xE4, 0x3B, 0x44, 0xFF},
	{0xF7, 0x76, 0x22, 0xFF},
	{0xFE, 0xAE, 0x34, 0xFF},
	{0xFE, 0xE7, 0x61, 0xFF},
	{0x63, 0xC7, 0x4D, 0xFF},
	{0x3E, 0x89, 0x48, 0xFF},
	{0x26, 0x5C, 0x42, 0xFF},
	{0x19, 0x3C, 0x3E, 0xFF},
	{0x12, 0x4E, 0x89, 0xFF},
	{0x00, 0x99, 0xDB, 0xFF},
	{0x2C, 0xE8, 0xF5, 0xFF},
	{0xFF, 0xFF, 0xFF, 0xFF},
	{0xC0, 0xCB, 0xDC, 0xFF},
	{0x8B, 0x9B, 0xB4, 0xFF},
	{0x5A, 0x69, 0x88, 0xFF},
	{0x3A, 0x44, 0x66, 0xFF},
	{0x26, 0x2B, 0x44, 0xFF},
	{0x18, 0x14, 0x25, 0xFF},
	{0xFF, 0x00, 0x44, 0xFF},
	{0x68, 0x38, 0x6C, 0xFF},
	{0xB5, 0x50, 0x88, 0xFF},
	{0xF6, 0x75, 0x7A, 0xFF},
	{0xE8, 0xB7, 0x96, 0xFF},
	{0xC2, 0x85, 0x69, 0xFF},
}

// PaletteEDG36 is the Endesga 36 palette.
//
// Created by Endesga. #EDG36
//
// https://lospec.com/palette-list/endesga-36
var PaletteEDG36 = Palette{
	{0xDB, 0xE0, 0xE7, 0xFF},
	{0xA3, 0xAC, 0xBE, 0xFF},
	{0x67, 0x70, 0x8B, 0xFF},
	{0x4E, 0x53, 0x71, 0xFF},
	{0x39, 0x3A, 0x56, 0xFF},
	{0x26, 0x24, 0x3A, 0xFF},
	{0x14, 0x10, 0x20, 0xFF},
	{0x7B, 0xCF, 0x5C, 0xFF},
	{0x50, 0x9B, 0x4B, 0xFF},
	{0x2E, 0x6A, 0x42, 0xFF},
	{0x1A, 0x45, 0x3B, 0xFF},
	{0x0F, 0x27, 0x38, 0xFF},
	{0x0D, 0x2F, 0x6D, 0xFF},
	{0x0F, 0x4D, 0xA3, 0xFF},
	{0x0E, 0x82, 0xCE, 0xFF},
	{0x13, 0xB2, 0xF2, 0xFF},
	{0x41, 0xF3, 0xFC, 0xFF},
	{0xF0, 0xD2, 0xAF, 0xFF},
	{0xE5, 0xAE, 0x78, 0xFF},
	{0xC5, 0x81, 0x58, 0xFF},
	{0x94, 0x55, 0x42, 0xFF},
	{0x62, 0x35, 0x30, 0xFF},
	{0x46, 0x21, 0x1F, 0xFF},
	{0x97, 0x43, 0x2A, 0xFF},
	{0xE5, 0x70, 0x28, 0xFF},
	{0xF7, 0xAC, 0x37, 0xFF},
	{0xFB, 0xDF, 0x6B, 0xFF},
	{0xFE, 0x97, 0x9B, 0xFF},
	{0xED, 0x52, 0x59, 0xFF},
	{0xC4, 0x2C, 0x36, 0xFF},
	{0x78, 0x1F, 0x2C, 0xFF},
	{0x35, 0x14, 0x28, 0xFF},
	{0x4D, 0x23, 0x52, 0xFF},
	{0x7F, 0x3B, 0x86, 0xFF},
	{0xB4, 0x5E, 0xB3, 0xFF},
	{0xE3, 0x8D, 0xD6, 0xFF},
}

// PaletteEDG64 is the Endesga 64 palette.
//
// "Honed over years of palette creation, refined for materialistic pixelart
// and design. High contrast, high saturation, shaped around painting the
// organic and structured life of the heptaverse." Created by Endesga. #EDG64
//
// https://lospec.com/palette-list/endesga-64
var PaletteEDG64 = Palette{
	{0xFF, 0x00, 0x40, 0xFF},
	{0x13, 0x13, 0x13, 0xFF},
	{0x1B, 0x1B, 0x1B, 0xFF},
	{0x27, 0x27, 0x27, 0xFF},
	{0x3D, 0x3D, 0x3D, 0xFF},
	{0x5D, 0x5D, 0x5D, 0xFF},
	{0x85, 0x85, 0x85, 0xFF},
	{0xB4, 0xB4, 0xB4, 0xFF},
	{0xFF, 0xFF, 0xFF, 0xFF},
	{0xC7, 0xCF, 0xDD, 0xFF},
	{0x92, 0xA1, 0xB9, 0xFF},
	{0x65, 0x73, 0x92, 0xFF},
	{0x42, 0x4C, 0x6E, 0xFF},
	{0x2A, 0x2F, 0x4E, 0xFF},
	{0x1A, 0x19, 0x32, 0xFF},
	{0x0E, 0x07, 0x1B, 0xFF},
	{0x1C, 0x12, 0x1C, 0xFF},
	{0x39, 0x1F, 0x21, 0xFF},
	{0x5D, 0x2C, 0x28, 0xFF},
	{0x8A, 0x48, 0x36, 0xFF},
	{0xBF, 0x6F, 0x4A, 0xFF},
	{0xE6, 0x9C, 0x69, 0xFF},
	{0xF6, 0xCA, 0x9F, 0xFF},
	{0xF9, 0xE6, 0xCF, 0xFF},
	{0xED, 0xAB, 0x50, 0xFF},
	{0xE0, 0x74, 0x38, 0xFF},
	{0xC6, 0x45, 0x24, 0xFF},
	{0x8E, 0x25, 0x1D, 0xFF},
	{0xFF, 0x50, 0x00, 0xFF},
	{0xED, 0x76, 0x14, 0xFF},
	{0xFF, 0xA2, 0x14, 0xFF},
	{0xFF, 0xC8, 0x25, 0xFF},
	{0xFF, 0xEB, 0x57, 0xFF},
	{0xD3, 0xFC, 0x7E, 0xFF},
	{0x99, 0xE6, 0x5F, 0xFF},
	{0x5A, 0xC5, 0x4F, 0xFF},
	{0x33, 0x98, 0x4B, 0xFF},
	{0x1E, 0x6F, 0x50, 0xFF},
	{0x13, 0x4C, 0x4C, 0xFF},
	{0x0C, 0x2E, 0x44, 0xFF},
	{0x00, 0x39, 0x6D, 0xFF},
	{0x00, 0x69, 0xAA, 0xFF},
	{0x00, 0x98, 0xDC, 0xFF},
	{0x00, 0xCD, 0xF9, 0xFF},
	{0x0C, 0xF1, 0xFF, 0xFF},
	{0x94, 0xFD, 0xFF, 0xFF},
	{0xFD, 0xD2, 0xED, 0xFF},
	{0xF3, 0x89, 0xF5, 0xFF},
	{0xDB, 0x3F, 0xFD, 0xFF},
	{0x7A, 0x09, 0xFA, 0xFF},
	{0x30, 0x03, 0xD9, 0xFF},
	{0x0C, 0x02, 0x93, 0xFF},
	{0x03, 0x19, 0x3F, 0xFF},
	{0x3B, 0x14, 0x43, 0xFF},
	{0x62, 0x24, 0x61, 0xFF},
	{0x93, 0x38, 0x8F, 0xFF},
	{0xCA, 0x52, 0xC9, 0xFF},
	{0xC8, 0x50, 0x86, 0xFF},
	{0xF6, 0x81, 0x87, 0xFF},
	{0xF5, 0x55, 0x5D, 0xFF},
	{0xEA, 0x32, 0x3C, 0xFF},
	{0xC4, 0x24, 0x30, 0xFF},
	{0x89, 0x1E, 0x2B, 0xFF},
	{0x57, 0x1C, 0x27, 0xFF},
}

// PaletteEN4 is the EN4 palette.
//
// Created by Endesga. #EN4
//
// https://lospec.com/palette-list/en4
var PaletteEN4 = Palette{
	{0xFB, 0xF7, 0xF3, 0xFF},
	{0xE5, 0xB0, 0x83, 0xFF},
	{0x42, 0x6E, 0x5D, 0xFF},
	{0x20, 0x28, 0x3D, 0xFF},
}

// PaletteInk is the Ink palette.
//
// Created by AprilSundae.
//
// https://lospec.com/palette-list/ink
var PaletteInk = Palette{
	{0x1F, 0x1F, 0x29, 0xFF},
	{0x41, 0x3A, 0x42, 0xFF},
	{0x59, 0x60, 0x70, 0xFF},
	{0x96, 0xA2, 0xB3, 0xFF},
	{0xEA, 0xF0, 0xD8, 0xFF},
}

// PalettePICO8 is the palette used by PICO-8.
//
// The PICO-8 is a virtual video game console created by Lexaloffle Games.
//
// https://lospec.com/palette-list/pico-8
var PalettePICO8 = Palette{
	{0x00, 0x00, 0x00, 0xFF},
	{0x5F, 0x57, 0x4F, 0xFF},
	{0xC2, 0xC3, 0xC7, 0xFF},
	{0xFF, 0xF1, 0xE8, 0xFF},
	{0xFF, 0xEC, 0x27, 0xFF},
	{0xFF, 0xA3, 0x00, 0xFF},
	{0xFF, 0xCC, 0xAA, 0xFF},
	{0xAB, 0x52, 0x36, 0xFF},
	{0xFF, 0x77, 0xA8, 0xFF},
	{0xFF, 0x00, 0x4D, 0xFF},
	{0x83, 0x76, 0x9C, 0xFF},
	{0x7E, 0x25, 0x53, 0xFF},
	{0x29, 0xAD, 0xFF, 0xFF},
	{0x1D, 0x2B, 0x53, 0xFF},
	{0x00, 0x87, 0x51, 0xFF},
	{0x00, 0xE4, 0x36, 0xFF},
}

// PaletteAmmo8 is the Ammo-8 palette.
//
// Created by rsvp asap.
//
// https://lospec.com/palette-list/ammo-8
var PaletteAmmo8 = Palette{
	{0x04, 0x0C, 0x06, 0xFF},
	{0x11, 0x23, 0x18, 0xFF},
	{0x1E, 0x3A, 0x29, 0xFF},
	{0x30, 0x5D, 0x42, 0xFF},
	{0x4D, 0x80, 0x61, 0xFF},
	{0x89, 0xA2, 0x57, 0xFF},
	{0xBE, 0xDC, 0x7F, 0xFF},
	{0xEE, 0xFF, 0xCC, 0xFF},
}

// PaletteNYX8 is the NYX8 palette.
//
// Palette created by Javier Guerrero.
//
// https://lospec.com/palette-list/nyx8
var PaletteNYX8 = Palette{
	{0x08, 0x14, 0x1E, 0xFF},
	{0x0F, 0x2A, 0x3F, 0xFF},
	{0x20, 0x39, 0x4F, 0xFF},
	{0xF6, 0xD6, 0xBD, 0xFF},
	{0xC3, 0xA3, 0x8A, 0xFF},
	{0x99, 0x75, 0x77, 0xFF},
	{0x81, 0x62, 0x71, 0xFF},
	{0x4E, 0x49, 0x5F, 0xFF},
}

// Palette15PDX is the 15P DX palette.
//
// Palette created by GrafxKid.
//
// https://lospec.com/palette-list/15p-dx
var Palette15PDX = Palette{
	{0x6E, 0x32, 0x32, 0xFF},
	{0xBB, 0x57, 0x35, 0xFF},
	{0xDF, 0x92, 0x45, 0xFF},
	{0xEC, 0xD2, 0x74, 0xFF},
	{0x83, 0xA8, 0x16, 0xFF},
	{0x27, 0x72, 0x24, 0xFF},
	{0x17, 0x3B, 0x47, 0xFF},
	{0x04, 0x68, 0x94, 0xFF},
	{0x17, 0xA1, 0xA9, 0xFF},
	{0x81, 0xDB, 0xCD, 0xFF},
	{0xFD, 0xF9, 0xF1, 0xFF},
	{0xC7, 0xB2, 0x95, 0xFF},
	{0x87, 0x71, 0x5B, 0xFF},
	{0x46, 0x3F, 0x3C, 0xFF},
	{0x20, 0x17, 0x08, 0xFF},
}

// Palette20PDX is the 20P DX palette.
//
// Palette created by GrafxKid.
//
// https://lospec.com/palette-list/20p-dx
var Palette20PDX = Palette{
	{0x17, 0x0D, 0x20, 0xFF},
	{0x47, 0x47, 0x57, 0xFF},
	{0x78, 0x78, 0x76, 0xFF},
	{0xB1, 0xB9, 0xA6, 0xFF},
	{0xEB, 0xFF, 0xDA, 0xFF},
	{0x68, 0x29, 0x3E, 0xFF},
	{0xA9, 0x44, 0x00, 0xFF},
	{0xD9, 0x7E, 0x00, 0xFF},
	{0xEB, 0xD0, 0x00, 0xFF},
	{0x52, 0x3C, 0x14, 0xFF},
	{0x81, 0x60, 0x31, 0xFF},
	{0xBC, 0x8B, 0x57, 0xFF},
	{0xEB, 0xCD, 0x93, 0xFF},
	{0x0E, 0x4C, 0x58, 0xFF},
	{0x04, 0x6E, 0x92, 0xFF},
	{0x01, 0xA3, 0xC3, 0xFF},
	{0x55, 0xDE, 0xB7, 0xFF},
	{0x17, 0x79, 0x47, 0xFF},
	{0x5A, 0xB2, 0x17, 0xFF},
	{0xB1, 0xE3, 0x29, 0xFF},
}

// PaletteArne16 is the Arne 16 palette.
//
// Created by Arne.
//
// https://lospec.com/palette-list/arne-16
var PaletteArne16 = Palette{
	{0x00, 0x00, 0x00, 0xFF},
	{0x49, 0x3C, 0x2B, 0xFF},
	{0xBE, 0x26, 0x33, 0xFF},
	{0xE0, 0x6F, 0x8B, 0xFF},
	{0x9D, 0x9D, 0x9D, 0xFF},
	{0xA4, 0x64, 0x22, 0xFF},
	{0xEB, 0x89, 0x31, 0xFF},
	{0xF7, 0xE2, 0x6B, 0xFF},
	{0xFF, 0xFF, 0xFF, 0xFF},
	{0x1B, 0x26, 0x32, 0xFF},
	{0x2F, 0x48, 0x4E, 0xFF},
	{0x44, 0x89, 0x1A, 0xFF},
	{0xA3, 0xCE, 0x27, 0xFF},
	{0x00, 0x57, 0x84, 0xFF},
	{0x31, 0xA2, 0xF2, 0xFF},
	{0xB2, 0xDC, 0xEF, 0xFF},
}

// PaletteNight16 is the Night 16 palette.
//
// 3rd place winner of the PixelJoint 16 color palette competition (2015).
// Created by Night.
//
// https://lospec.com/palette-list/night-16
var PaletteNight16 = Palette{
	{0x0F, 0x0F, 0x1E, 0xFF},
	{0xFF, 0xF8, 0xBC, 0xFF},
	{0x0C, 0x21, 0x33, 0xFF},
	{0x48, 0x58, 0x6D, 0xFF},
	{0x79, 0xA0, 0xB0, 0xFF},
	{0xB0, 0xCE, 0x9D, 0xFF},
	{0x65, 0x7F, 0x49, 0xFF},
	{0x3F, 0x45, 0x36, 0xFF},
	{0xB9, 0x9D, 0x6A, 0xFF},
	{0xFF, 0xDD, 0x91, 0xFF},
	{0xDD, 0x94, 0x5B, 0xFF},
	{0x9A, 0x51, 0x42, 0xFF},
	{0x64, 0x4B, 0x48, 0xFF},
	{0x33, 0x30, 0x33, 0xFF},
	{0x76, 0x70, 0x88, 0xFF},
	{0xC5, 0xA3, 0xB3, 0xFF},
}

// PaletteAAP16 is the AAP-16 palette.
//
// Created by Adigun Polack, meant for beginners.
//
// https://lospec.com/palette-list/aap-16
var PaletteAAP16 = Palette{
	{0x07, 0x07, 0x08, 0xFF},
	{0x33, 0x22, 0x22, 0xFF},
	{0x77, 0x44, 0x33, 0xFF},
	{0xCC, 0x88, 0x55, 0xFF},
	{0x99, 0x33, 0x11, 0xFF},
	{0xDD, 0x77, 0x11, 0xFF},
	{0xFF, 0xDD, 0x55, 0xFF},
	{0xFF, 0xFF, 0x33, 0xFF},
	{0x55, 0xAA, 0x44, 0xFF},
	{0x11, 0x55, 0x22, 0xFF},
	{0x44, 0xEE, 0xBB, 0xFF},
	{0x33, 0x88, 0xDD, 0xFF},
	{0x55, 0x44, 0xAA, 0xFF},
	{0x55, 0x55, 0x77, 0xFF},
	{0xAA, 0xBB, 0xBB, 0xFF},
	{0xFF, 0xFF, 0xFF, 0xFF},
}

// PaletteAAP64 is the AAP-64 palette.
//
// Created by Adigun Polack.
//
// https://lospec.com/palette-list/aap-16
var PaletteAAP64 = Palette{
	{0x06, 0x06, 0x08, 0xFF},
	{0x14, 0x10, 0x13, 0xFF},
	{0x3B, 0x17, 0x25, 0xFF},
	{0x73, 0x17, 0x2D, 0xFF},
	{0xB4, 0x20, 0x2A, 0xFF},
	{0xDF, 0x3E, 0x23, 0xFF},
	{0xFA, 0x6A, 0x0A, 0xFF},
	{0xF9, 0xA3, 0x1B, 0xFF},
	{0xFF, 0xD5, 0x41, 0xFF},
	{0xFF, 0xFC, 0x40, 0xFF},
	{0xD6, 0xF2, 0x64, 0xFF},
	{0x9C, 0xDB, 0x43, 0xFF},
	{0x59, 0xC1, 0x35, 0xFF},
	{0x14, 0xA0, 0x2E, 0xFF},
	{0x1A, 0x7A, 0x3E, 0xFF},
	{0x24, 0x52, 0x3B, 0xFF},
	{0x12, 0x20, 0x20, 0xFF},
	{0x14, 0x34, 0x64, 0xFF},
	{0x28, 0x5C, 0xC4, 0xFF},
	{0x24, 0x9F, 0xDE, 0xFF},
	{0x20, 0xD6, 0xC7, 0xFF},
	{0xA6, 0xFC, 0xDB, 0xFF},
	{0xFF, 0xFF, 0xFF, 0xFF},
	{0xFE, 0xF3, 0xC0, 0xFF},
	{0xFA, 0xD6, 0xB8, 0xFF},
	{0xF5, 0xA0, 0x97, 0xFF},
	{0xE8, 0x6A, 0x73, 0xFF},
	{0xBC, 0x4A, 0x9B, 0xFF},
	{0x79, 0x3A, 0x80, 0xFF},
	{0x40, 0x33, 0x53, 0xFF},
	{0x24, 0x22, 0x34, 0xFF},
	{0x22, 0x1C, 0x1A, 0xFF},
	{0x32, 0x2B, 0x28, 0xFF},
	{0x71, 0x41, 0x3B, 0xFF},
	{0xBB, 0x75, 0x47, 0xFF},
	{0xDB, 0xA4, 0x63, 0xFF},
	{0xF4, 0xD2, 0x9C, 0xFF},
	{0xDA, 0xE0, 0xEA, 0xFF},
	{0xB3, 0xB9, 0xD1, 0xFF},
	{0x8B, 0x93, 0xAF, 0xFF},
	{0x6D, 0x75, 0x8D, 0xFF},
	{0x4A, 0x54, 0x62, 0xFF},
	{0x33, 0x39, 0x41, 0xFF},
	{0x42, 0x24, 0x33, 0xFF},
	{0x5B, 0x31, 0x38, 0xFF},
	{0x8E, 0x52, 0x52, 0xFF},
	{0xBA, 0x75, 0x6A, 0xFF},
	{0xE9, 0xB5, 0xA3, 0xFF},
	{0xE3, 0xE6, 0xFF, 0xFF},
	{0xB9, 0xBF, 0xFB, 0xFF},
	{0x84, 0x9B, 0xE4, 0xFF},
	{0x58, 0x8D, 0xBE, 0xFF},
	{0x47, 0x7D, 0x85, 0xFF},
	{0x23, 0x67, 0x4E, 0xFF},
	{0x32, 0x84, 0x64, 0xFF},
	{0x5D, 0xAF, 0x8D, 0xFF},
	{0x92, 0xDC, 0xBA, 0xFF},
	{0xCD, 0xF7, 0xE2, 0xFF},
	{0xE4, 0xD2, 0xAA, 0xFF},
	{0xC7, 0xB0, 0x8B, 0xFF},
	{0xA0, 0x86, 0x62, 0xFF},
	{0x79, 0x67, 0x55, 0xFF},
	{0x5A, 0x4E, 0x44, 0xFF},
	{0x42, 0x39, 0x34, 0xFF},
}

// PaletteSplendor128 is the Splendor 128 palette.
//
// Created by Adigun Polack as a successor to his AAP-64 palette. #Splendor128
//
// https://lospec.com/palette-list/aap-splendor128
var PaletteSplendor128 = Palette{
	{0x05, 0x04, 0x03, 0xFF},
	{0x0E, 0x0C, 0x0C, 0xFF},
	{0x2D, 0x1B, 0x1E, 0xFF},
	{0x61, 0x27, 0x21, 0xFF},
	{0xB9, 0x45, 0x1D, 0xFF},
	{0xF1, 0x64, 0x1F, 0xFF},
	{0xFC, 0xA5, 0x70, 0xFF},
	{0xFF, 0xE0, 0xB7, 0xFF},
	{0xFF, 0xFF, 0xFF, 0xFF},
	{0xFF, 0xF0, 0x89, 0xFF},
	{0xF8, 0xC5, 0x3A, 0xFF},
	{0xE8, 0x8A, 0x36, 0xFF},
	{0xB0, 0x5B, 0x2C, 0xFF},
	{0x67, 0x39, 0x31, 0xFF},
	{0x27, 0x1F, 0x1B, 0xFF},
	{0x4C, 0x3D, 0x2E, 0xFF},
	{0x85, 0x5F, 0x39, 0xFF},
	{0xD3, 0x97, 0x41, 0xFF},
	{0xF8, 0xF6, 0x44, 0xFF},
	{0xD5, 0xDC, 0x1D, 0xFF},
	{0xAD, 0xB8, 0x34, 0xFF},
	{0x7F, 0x8E, 0x44, 0xFF},
	{0x58, 0x63, 0x35, 0xFF},
	{0x33, 0x3C, 0x24, 0xFF},
	{0x18, 0x1C, 0x19, 0xFF},
	{0x29, 0x3F, 0x21, 0xFF},
	{0x47, 0x72, 0x38, 0xFF},
	{0x61, 0xA5, 0x3F, 0xFF},
	{0x8F, 0xD0, 0x32, 0xFF},
	{0xC4, 0xF1, 0x29, 0xFF},
	{0xD0, 0xFF, 0xEA, 0xFF},
	{0x97, 0xED, 0xCA, 0xFF},
	{0x59, 0xCF, 0x93, 0xFF},
	{0x42, 0xA4, 0x59, 0xFF},
	{0x3D, 0x6F, 0x43, 0xFF},
	{0x27, 0x41, 0x2D, 0xFF},
	{0x14, 0x12, 0x1D, 0xFF},
	{0x1B, 0x24, 0x47, 0xFF},
	{0x2B, 0x4E, 0x95, 0xFF},
	{0x27, 0x89, 0xCD, 0xFF},
	{0x42, 0xBF, 0xE8, 0xFF},
	{0x73, 0xEF, 0xE8, 0xFF},
	{0xF1, 0xF2, 0xFF, 0xFF},
	{0xC9, 0xD4, 0xFD, 0xFF},
	{0x8A, 0xA1, 0xF6, 0xFF},
	{0x45, 0x72, 0xE3, 0xFF},
	{0x49, 0x41, 0x82, 0xFF},
	{0x78, 0x64, 0xC6, 0xFF},
	{0x9C, 0x8B, 0xDB, 0xFF},
	{0xCE, 0xAA, 0xED, 0xFF},
	{0xFA, 0xD6, 0xFF, 0xFF},
	{0xEE, 0xB5, 0x9C, 0xFF},
	{0xD4, 0x80, 0xBB, 0xFF},
	{0x90, 0x52, 0xBC, 0xFF},
	{0x17, 0x15, 0x16, 0xFF},
	{0x37, 0x33, 0x34, 0xFF},
	{0x69, 0x5B, 0x59, 0xFF},
	{0xB2, 0x8B, 0x78, 0xFF},
	{0xE2, 0xB2, 0x7E, 0xFF},
	{0xF6, 0xD8, 0x96, 0xFF},
	{0xFC, 0xF7, 0xBE, 0xFF},
	{0xEC, 0xEB, 0xE7, 0xFF},
	{0xCB, 0xC6, 0xC1, 0xFF},
	{0xA6, 0x9E, 0x9A, 0xFF},
	{0x80, 0x7B, 0x7A, 0xFF},
	{0x59, 0x57, 0x57, 0xFF},
	{0x32, 0x32, 0x32, 0xFF},
	{0x4F, 0x34, 0x2F, 0xFF},
	{0x8C, 0x5B, 0x3E, 0xFF},
	{0xC6, 0x85, 0x56, 0xFF},
	{0xD6, 0xA8, 0x51, 0xFF},
	{0xB4, 0x75, 0x38, 0xFF},
	{0x72, 0x4B, 0x2C, 0xFF},
	{0x45, 0x2A, 0x1B, 0xFF},
	{0x61, 0x68, 0x3A, 0xFF},
	{0x93, 0x94, 0x46, 0xFF},
	{0xC6, 0xB8, 0x58, 0xFF},
	{0xEF, 0xDD, 0x91, 0xFF},
	{0xB5, 0xE7, 0xCB, 0xFF},
	{0x86, 0xC6, 0x9A, 0xFF},
	{0x5D, 0x9B, 0x79, 0xFF},
	{0x48, 0x68, 0x59, 0xFF},
	{0x2C, 0x3B, 0x39, 0xFF},
	{0x17, 0x18, 0x19, 0xFF},
	{0x2C, 0x34, 0x38, 0xFF},
	{0x46, 0x54, 0x56, 0xFF},
	{0x64, 0x87, 0x8C, 0xFF},
	{0x8A, 0xC4, 0xC3, 0xFF},
	{0xAF, 0xE9, 0xDF, 0xFF},
	{0xDC, 0xEA, 0xEE, 0xFF},
	{0xB8, 0xCC, 0xD8, 0xFF},
	{0x88, 0xA3, 0xBC, 0xFF},
	{0x5E, 0x71, 0x8E, 0xFF},
	{0x48, 0x52, 0x62, 0xFF},
	{0x28, 0x2C, 0x3C, 0xFF},
	{0x46, 0x47, 0x62, 0xFF},
	{0x69, 0x66, 0x82, 0xFF},
	{0x9A, 0x97, 0xB9, 0xFF},
	{0xC5, 0xC7, 0xDD, 0xFF},
	{0xE6, 0xE7, 0xF0, 0xFF},
	{0xEE, 0xE6, 0xEA, 0xFF},
	{0xE3, 0xCD, 0xDF, 0xFF},
	{0xBF, 0xA5, 0xC9, 0xFF},
	{0x87, 0x73, 0x8F, 0xFF},
	{0x56, 0x4F, 0x5B, 0xFF},
	{0x32, 0x2F, 0x35, 0xFF},
	{0x36, 0x28, 0x2B, 0xFF},
	{0x65, 0x49, 0x56, 0xFF},
	{0x96, 0x68, 0x88, 0xFF},
	{0xC0, 0x90, 0xA9, 0xFF},
	{0xD4, 0xB8, 0xB8, 0xFF},
	{0xEA, 0xE0, 0xDD, 0xFF},
	{0xF1, 0xEB, 0xDB, 0xFF},
	{0xDD, 0xCE, 0xBF, 0xFF},
	{0xBD, 0xA4, 0x99, 0xFF},
	{0x88, 0x6E, 0x6A, 0xFF},
	{0x59, 0x4D, 0x4D, 0xFF},
	{0x33, 0x27, 0x2A, 0xFF},
	{0xB2, 0x94, 0x76, 0xFF},
	{0xE1, 0xBF, 0x89, 0xFF},
	{0xF8, 0xE3, 0x98, 0xFF},
	{0xFF, 0xE9, 0xE3, 0xFF},
	{0xFD, 0xC9, 0xC9, 0xFF},
	{0xF6, 0xA2, 0xA8, 0xFF},
	{0xE2, 0x72, 0x85, 0xFF},
	{0xB2, 0x52, 0x66, 0xFF},
	{0x64, 0x36, 0x4B, 0xFF},
	{0x2A, 0x1E, 0x23, 0xFF},
}

// PaletteFamicube is the Famicube palette.
//
// Created by Arne as part of his Famicube Project.
//
// https://lospec.com/palette-list/famicube
var PaletteFamicube = Palette{
	{0x00, 0x00, 0x00, 0xFF},
	{0x00, 0x17, 0x7D, 0xFF},
	{0x02, 0x4A, 0xCA, 0xFF},
	{0x00, 0x84, 0xFF, 0xFF},
	{0x5B, 0xA8, 0xFF, 0xFF},
	{0x98, 0xDC, 0xFF, 0xFF},
	{0x9B, 0xA0, 0xEF, 0xFF},
	{0x62, 0x64, 0xDC, 0xFF},
	{0x3D, 0x34, 0xA5, 0xFF},
	{0x21, 0x16, 0x40, 0xFF},
	{0x5A, 0x19, 0x91, 0xFF},
	{0x6A, 0x31, 0xCA, 0xFF},
	{0xA6, 0x75, 0xFE, 0xFF},
	{0xE2, 0xC9, 0xFF, 0xFF},
	{0xFE, 0xC9, 0xED, 0xFF},
	{0xD5, 0x9C, 0xFC, 0xFF},
	{0xCC, 0x69, 0xE4, 0xFF},
	{0xA3, 0x28, 0xB3, 0xFF},
	{0x87, 0x16, 0x46, 0xFF},
	{0xCF, 0x3C, 0x71, 0xFF},
	{0xFF, 0x82, 0xCE, 0xFF},
	{0xFF, 0xE9, 0xC5, 0xFF},
	{0xF5, 0xB7, 0x84, 0xFF},
	{0xE1, 0x82, 0x89, 0xFF},
	{0xDA, 0x65, 0x5E, 0xFF},
	{0x82, 0x3C, 0x3D, 0xFF},
	{0x4F, 0x15, 0x07, 0xFF},
	{0xE0, 0x3C, 0x28, 0xFF},
	{0xE2, 0xD7, 0xB5, 0xFF},
	{0xC5, 0x97, 0x82, 0xFF},
	{0xAE, 0x6C, 0x37, 0xFF},
	{0x5C, 0x3C, 0x0D, 0xFF},
	{0x23, 0x17, 0x12, 0xFF},
	{0xAD, 0x4E, 0x1A, 0xFF},
	{0xF6, 0x8F, 0x37, 0xFF},
	{0xFF, 0xE7, 0x37, 0xFF},
	{0xFF, 0xBB, 0x31, 0xFF},
	{0xCC, 0x8F, 0x15, 0xFF},
	{0x93, 0x97, 0x17, 0xFF},
	{0xB6, 0xC1, 0x21, 0xFF},
	{0xEE, 0xFF, 0xA9, 0xFF},
	{0xBE, 0xEB, 0x71, 0xFF},
	{0x8C, 0xD6, 0x12, 0xFF},
	{0x6A, 0xB4, 0x17, 0xFF},
	{0x37, 0x6D, 0x03, 0xFF},
	{0x17, 0x28, 0x08, 0xFF},
	{0x00, 0x4E, 0x00, 0xFF},
	{0x13, 0x9D, 0x08, 0xFF},
	{0x58, 0xD3, 0x32, 0xFF},
	{0x20, 0xB5, 0x62, 0xFF},
	{0x00, 0x60, 0x4B, 0xFF},
	{0x00, 0x52, 0x80, 0xFF},
	{0x0A, 0x98, 0xAC, 0xFF},
	{0x25, 0xE2, 0xCD, 0xFF},
	{0xBD, 0xFF, 0xCA, 0xFF},
	{0x71, 0xA6, 0xA1, 0xFF},
	{0x41, 0x5D, 0x66, 0xFF},
	{0x0D, 0x20, 0x30, 0xFF},
	{0x15, 0x15, 0x15, 0xFF},
	{0x34, 0x34, 0x34, 0xFF},
	{0x7B, 0x7B, 0x7B, 0xFF},
	{0xA8, 0xA8, 0xA8, 0xFF},
	{0xD7, 0xD7, 0xD7, 0xFF},
	{0xFF, 0xFF, 0xFF, 0xFF},
}

// PaletteTango is the Tango palette.
//
// http://en.wikipedia.org/wiki/Tango_Desktop_Project#Palette
var PaletteTango = Palette{
	{0xFC, 0xE9, 0x4F, 0xFF},
	{0xED, 0xD4, 0x00, 0xFF},
	{0xC4, 0xA0, 0x00, 0xFF},
	{0xFC, 0xAF, 0x3E, 0xFF},
	{0xF5, 0x79, 0x00, 0xFF},
	{0xCE, 0x5C, 0x00, 0xFF},
	{0xE9, 0xB9, 0x6E, 0xFF},
	{0xC1, 0x7D, 0x11, 0xFF},
	{0x8F, 0x59, 0x02, 0xFF},
	{0x8A, 0xE2, 0x34, 0xFF},
	{0x73, 0xD2, 0x16, 0xFF},
	{0x4E, 0x9A, 0x06, 0xFF},
	{0x72, 0x9F, 0xCF, 0xFF},
	{0x34, 0x65, 0xA4, 0xFF},
	{0x20, 0x4A, 0x87, 0xFF},
	{0xAD, 0x7F, 0xA8, 0xFF},
	{0x75, 0x50, 0x7B, 0xFF},
	{0x5C, 0x35, 0x66, 0xFF},
	{0xEF, 0x29, 0x29, 0xFF},
	{0xCC, 0x00, 0x00, 0xFF},
	{0xA4, 0x00, 0x00, 0xFF},
	{0xEE, 0xEE, 0xEC, 0xFF},
	{0xD3, 0xD7, 0xCF, 0xFF},
	{0xBA, 0xBD, 0xB6, 0xFF},
	{0x88, 0x8A, 0x85, 0xFF},
	{0x55, 0x57, 0x53, 0xFF},
	{0x2E, 0x34, 0x36, 0xFF},
}

// PaletteGo is the Go palette.
//
// https://golang.org/s/brandbook
var PaletteGo = Palette{
	{0x00, 0xAD, 0xD8, 0xFF}, //  0 Gopher Blue
	{0x0B, 0xB5, 0xDB, 0xFF}, //  1
	{0x31, 0xBE, 0xE0, 0xFF}, //  2
	{0x4D, 0xC7, 0xE4, 0xFF}, //  3
	{0x68, 0xCC, 0xE7, 0xFF}, //  4
	{0x82, 0xD2, 0xE8, 0xFF}, //  5
	{0x9C, 0xDB, 0xED, 0xFF}, //  6
	{0xB5, 0xE3, 0xF0, 0xFF}, //  7
	{0xE9, 0xF3, 0xF9, 0xFF}, //  8
	{0x5D, 0xC9, 0xE2, 0xFF}, //  9 Light Blue
	{0x7D, 0xD1, 0xE6, 0xFF}, // 10
	{0x98, 0xD9, 0xEA, 0xFF}, // 11
	{0x98, 0xD5, 0xEC, 0xFF}, // 12
	{0xC5, 0xE9, 0xF2, 0xFF}, // 13
	{0xD5, 0xEE, 0xF5, 0xFF}, // 14
	{0xE3, 0xF4, 0xF8, 0xFF}, // 15
	{0xEE, 0xF8, 0xFB, 0xFF}, // 16
	{0xF7, 0xFC, 0xFD, 0xFF}, // 17
	{0x00, 0xA2, 0x9C, 0xFF}, // 18 Aqua
	{0x5B, 0xC4, 0xBA, 0xFF}, // 19
	{0x77, 0xCB, 0xC5, 0xFF}, // 20
	{0x94, 0xD5, 0xD1, 0xFF}, // 21
	{0xAD, 0xDE, 0xDB, 0xFF}, // 22
	{0xC4, 0xE7, 0xE4, 0xFF}, // 23
	{0xD7, 0xEE, 0xED, 0xFF}, // 24
	{0xE8, 0xF5, 0xF4, 0xFF}, // 25
	{0xD8, 0xEE, 0xEB, 0xFF}, // 26
	{0xCE, 0x32, 0x62, 0xFF}, // 27 Fuchsia
	{0xD7, 0x5C, 0x7E, 0xFF}, // 28
	{0xDE, 0x7B, 0x96, 0xFF}, // 29
	{0xE4, 0x97, 0xAD, 0xFF}, // 30
	{0xEB, 0xB1, 0xC1, 0xFF}, // 31
	{0xF2, 0xC9, 0xD4, 0xFF}, // 32
	{0xF6, 0xDC, 0xE3, 0xFF}, // 33
	{0xF9, 0xEA, 0xEE, 0xFF}, // 34
	{0xF1, 0xD2, 0xD3, 0xFF}, // 35
	{0x00, 0x00, 0x00, 0xFF}, // 36 Black
	{0x1B, 0x1A, 0x1A, 0xFF}, // 37
	{0x2E, 0x2D, 0x2C, 0xFF}, // 38
	{0x40, 0x3D, 0x3D, 0xFF}, // 39
	{0x53, 0x50, 0x50, 0xFF}, // 40
	{0x68, 0x64, 0x64, 0xFF}, // 41
	{0x7F, 0x7C, 0x7B, 0xFF}, // 42
	{0x9A, 0x97, 0x96, 0xFF}, // 43
	{0xB5, 0xB2, 0xB3, 0xFF}, // 44
	{0xFD, 0xDD, 0x00, 0xFF}, // 45 Yellow
	{0xFE, 0xE3, 0x3D, 0xFF}, // 46
	{0xFF, 0xE9, 0x67, 0xFF}, // 47
	{0xFF, 0xED, 0x88, 0xFF}, // 48
	{0xFE, 0xF1, 0xA4, 0xFF}, // 49
	{0xFE, 0xF5, 0xBE, 0xFF}, // 50
	{0xFE, 0xF9, 0xD5, 0xFF}, // 51
	{0xFF, 0xFB, 0xE6, 0xFF}, // 52
	{0xFF, 0xFE, 0xF3, 0xFF}, // 53
	{0x00, 0x75, 0x8D, 0xFF}, // 54 Dark Blue
	{0x55, 0x57, 0x59, 0xFF}, // 55 Gray
	{0x40, 0x2B, 0x56, 0xFF}, // 56 Purple
	{0xDB, 0xD9, 0xD6, 0xFF}, // 57 Light Gray
}

// PaletteBasemaps is the Basemaps palette.
var PaletteBasemaps = Palette{
	{0xEE, 0xDD, 0xDE, 0xFF},
	{0xE1, 0xC7, 0xC9, 0xFF},
	{0xEA, 0xD5, 0xD7, 0xFF},
	{0x8B, 0x94, 0xA0, 0xFF},
	{0x9B, 0xA3, 0xAD, 0xFF},
	{0x87, 0x91, 0x9D, 0xFF},
	{0xA4, 0xAB, 0xB3, 0xFF},
	{0x95, 0x9F, 0xAA, 0xFF},
	{0xF2, 0xF5, 0xF8, 0xFF},
	{0xD7, 0xD8, 0xD9, 0xFF},
	{0xBE, 0xC3, 0xC8, 0xFF},
	{0x41, 0x5C, 0x77, 0xFF},
	{0xB9, 0xBF, 0xC5, 0xFF},
	{0x56, 0x6E, 0x85, 0xFF},
	{0x68, 0x7D, 0x91, 0xFF},
	{0x96, 0xA4, 0xB0, 0xFF},
	{0xC2, 0xC9, 0xCE, 0xFF},
	{0xD3, 0xE6, 0xEA, 0xFF},
	{0xC3, 0xDC, 0xE1, 0xFF},
	{0xD1, 0xE5, 0xE9, 0xFF},
	{0xAE, 0xCE, 0xD4, 0xFF},
	{0xBB, 0xD6, 0xDB, 0xFF},
	{0xC0, 0xDB, 0xE0, 0xFF},
	{0xCB, 0xE1, 0xE5, 0xFF},
	{0xB3, 0xD1, 0xD6, 0xFF},
	{0xC6, 0xDE, 0xE2, 0xFF},
	{0xC8, 0xE0, 0xE4, 0xFF},
	{0xD5, 0xE8, 0xEB, 0xFF},
	{0xCB, 0xE0, 0xE3, 0xFF},
	{0xCE, 0xE3, 0xE6, 0xFF},
	{0xBC, 0xD4, 0xD5, 0xFF},
	{0xE2, 0xEB, 0xEB, 0xFF},
	{0xDD, 0xE0, 0xE0, 0xFF},
	{0xE5, 0xEE, 0xED, 0xFF},
	{0xEE, 0xF2, 0xF1, 0xFF},
	{0xE3, 0xE5, 0xE4, 0xFF},
	{0xEC, 0xF1, 0xEE, 0xFF},
	{0xD5, 0xE3, 0xD7, 0xFF},
	{0xC9, 0xD6, 0xC9, 0xFF},
	{0xDC, 0xE9, 0xDB, 0xFF},
	{0xD9, 0xE7, 0xCE, 0xFF},
	{0xDA, 0xE8, 0xCF, 0xFF},
	{0xE2, 0xED, 0xD8, 0xFF},
	{0xE6, 0xEF, 0xDD, 0xFF},
	{0xE5, 0xEE, 0xDB, 0xFF},
	{0xE3, 0xEC, 0xD9, 0xFF},
	{0xE8, 0xF0, 0xDD, 0xFF},
	{0xE9, 0xF0, 0xDF, 0xFF},
	{0xED, 0xF2, 0xE4, 0xFF},
	{0xF3, 0xF6, 0xED, 0xFF},
	{0xF1, 0xF2, 0xEE, 0xFF},
	{0xF5, 0xF5, 0xF2, 0xFF},
	{0xEB, 0xEB, 0xE8, 0xFF},
	{0xFF, 0xFF, 0xFE, 0xFF},
	{0xF6, 0xF6, 0xF5, 0xFF},
	{0xEE, 0xEE, 0xEC, 0xFF},
	{0xFE, 0xFC, 0xDA, 0xFF},
	{0xEE, 0xE9, 0xCE, 0xFF},
	{0xF6, 0xF4, 0xEB, 0xFF},
	{0xEF, 0xEE, 0xEA, 0xFF},
	{0xE6, 0xE5, 0xE1, 0xFF},
	{0xF1, 0xEE, 0xE3, 0xFF},
	{0xF0, 0xEE, 0xE7, 0xFF},
	{0xE3, 0xE1, 0xDA, 0xFF},
	{0xE4, 0xE3, 0xE0, 0xFF},
	{0xF2, 0xF0, 0xEA, 0xFF},
	{0xFE, 0xE8, 0xA9, 0xFF},
	{0xE8, 0xDB, 0xB8, 0xFF},
	{0xF9, 0xF6, 0xEE, 0xFF},
	{0xF4, 0xEB, 0xD4, 0xFF},
	{0xFD, 0xFB, 0xF6, 0xFF},
	{0xF9, 0xF6, 0xEF, 0xFF},
	{0xFC, 0xE1, 0xA4, 0xFF},
	{0xF8, 0xF4, 0xEB, 0xFF},
	{0xFF, 0xE8, 0xB7, 0xFF},
	{0xFA, 0xDD, 0xA1, 0xFF},
	{0xFA, 0xDC, 0x9E, 0xFF},
	{0xFA, 0xE7, 0xC0, 0xFF},
	{0xE3, 0xE2, 0xE0, 0xFF},
	{0xFC, 0xFA, 0xF6, 0xFF},
	{0xF4, 0xF2, 0xEE, 0xFF},
	{0xEF, 0xEA, 0xE0, 0xFF},
	{0xE5, 0xDB, 0xC7, 0xFF},
	{0xCD, 0xCC, 0xCA, 0xFF},
	{0xFE, 0xF9, 0xEF, 0xFF},
	{0xF9, 0xE0, 0xAF, 0xFF},
	{0xFB, 0xED, 0xD2, 0xFF},
	{0xFB, 0xEE, 0xD5, 0xFF},
	{0xF8, 0xDD, 0xAA, 0xFF},
	{0xF9, 0xE6, 0xC3, 0xFF},
	{0xFE, 0xF8, 0xED, 0xFF},
	{0xF6, 0xF0, 0xE5, 0xFF},
	{0xFA, 0xE4, 0xBC, 0xFF},
	{0xFD, 0xEF, 0xD6, 0xFF},
	{0xED, 0xE7, 0xDD, 0xFF},
	{0xF2, 0xE9, 0xDA, 0xFF},
	{0xFB, 0xF8, 0xF3, 0xFF},
	{0xF7, 0xF1, 0xE7, 0xFF},
	{0xFD, 0xEC, 0xD0, 0xFF},
	{0xFD, 0xEB, 0xCE, 0xFF},
	{0xE7, 0xE0, 0xD5, 0xFF},
	{0xF5, 0xF3, 0xF0, 0xFF},
	{0xFC, 0xF2, 0xE3, 0xFF},
	{0xF4, 0xEC, 0xE0, 0xFF},
	{0xEA, 0xE2, 0xD6, 0xFF},
	{0xEE, 0xE1, 0xCE, 0xFF},
	{0xF1, 0xE6, 0xD6, 0xFF},
	{0xF6, 0xEF, 0xE5, 0xFF},
	{0xEB, 0xDD, 0xC9, 0xFF},
	{0xEB, 0xE4, 0xDA, 0xFF},
	{0xE7, 0xDB, 0xCA, 0xFF},
	{0xE7, 0xDA, 0xC8, 0xFF},
	{0xE4, 0xDC, 0xD1, 0xFF},
	{0xF5, 0xED, 0xE3, 0xFF},
	{0xDF, 0xDE, 0xDD, 0xFF},
	{0xE1, 0xE0, 0xDF, 0xFF},
	{0xEA, 0xE4, 0xDF, 0xFF},
	{0xF8, 0xF2, 0xEE, 0xFF},
	{0xEC, 0xE2, 0xDD, 0xFF},
	{0xEE, 0xE5, 0xE3, 0xFF},
	{0xF2, 0xE6, 0xE5, 0xFF},
	{0xE9, 0xE9, 0xE9, 0xFF},
	{0xE9, 0xD6, 0xD6, 0xFF},
}
