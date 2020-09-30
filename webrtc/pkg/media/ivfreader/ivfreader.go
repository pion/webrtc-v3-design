// Package ivfreader implements IVF media container reader
package ivfreader

import (
	"io"
)

// IVFFileHeader 32-byte header for IVF files
type IVFFileHeader struct{}

// IVFFrameHeader 12-byte header for IVF frames
type IVFFrameHeader struct{}

// IVFReader is used to read IVF files and return frame payloads
type IVFReader struct{}

// NewWith returns a new IVF reader and IVF file header
func NewWith(in io.Reader) (*IVFReader, *IVFFileHeader, error) { return nil, nil, nil }

// ParseNextFrame reads from stream and returns IVF frame payload, header,
func (i *IVFReader) ParseNextFrame() ([]byte, *IVFFrameHeader, error) { return nil, nil, nil }
