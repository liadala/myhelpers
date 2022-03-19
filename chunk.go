package myhelpers

func ChunkSlice[Slice any](input []Slice, size int) [][]Slice {
	var chunk []Slice
	chunks := make([][]Slice, 0, len(input)/size+1)
	for len(input) >= size {
		chunk, input = input[:size], input[size:]
		chunks = append(chunks, chunk)
	}
	if len(input) > 0 {
		chunks = append(chunks, input)
	}
	return chunks
}
