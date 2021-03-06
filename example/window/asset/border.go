package asset

import (
	"github.com/Akatsuki-py/gbdk-go/api/gb"
	"github.com/Akatsuki-py/gbdk-go/api/macro"
)

var BorderBank = macro.Define(0)

var Border = []gb.UINT8{
	0x00, 0x00, 0x1F, 0x1F, 0x20, 0x20, 0x4F, 0x4F,
	0x50, 0x50, 0x50, 0x50, 0x50, 0x50, 0x50, 0x50,
	0x50, 0x50, 0x50, 0x50, 0x50, 0x50, 0x50, 0x50,
	0x4F, 0x4F, 0x20, 0x20, 0x1F, 0x1F, 0x00, 0x00,
	0x00, 0x00, 0xF8, 0xF8, 0x04, 0x04, 0xF2, 0xF2,
	0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A,
	0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A,
	0xF2, 0xF2, 0x04, 0x04, 0xF8, 0xF8, 0x00, 0x00,
	0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00,
	0x50, 0x50, 0x50, 0x50, 0x50, 0x50, 0x50, 0x50,
	0x50, 0x50, 0x50, 0x50, 0x50, 0x50, 0x50, 0x50,
	0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A,
	0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
}
