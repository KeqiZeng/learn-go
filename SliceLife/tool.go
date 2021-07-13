package main

func getHW(h, w int) (int, int) {
	if h < 0 {
		h += height
	} else if h >= height {
		h %= height
	}
	if w < 0 {
		w += width
	} else if w >= width {
		w %= width
	}
	return h, w
}

func Step(a Universe) Universe {
	b := make(Universe, height)
	for h := 0; h < height; h++ {
		tmpLine := make([]bool, 0, width)
		for w := 0; w < width; w++ {
			tmpLine = append(tmpLine, a.Next(h, w))
		}
		b[h] = tmpLine
	}
	return b
}
