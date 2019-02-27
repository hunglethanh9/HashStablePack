package marshalhash_test

//go:generate hsp -o=defgen_test.go -tests=false

type Blobs []Blob

type Blob struct {
	Name   string  `msg:"name"`
	Float  float64 `msg:"float"`
	Bytes  []byte  `msg:"bytes"`
	Amount int64   `msg:"amount"`
}
