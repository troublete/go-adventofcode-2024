package main

type Map [][]*Tile

func (m *Map) Get(x, y int) *Tile {
	if y >= 0 && y <= len(*m)-1 {
		if x >= 0 && x <= len((*m)[y])-1 {
			return (*m)[y][x]
		}

	}
	return nil
}

type Tile struct {
	C        byte
	X, Y     int
	Map      *Map
	RegionID int
}

type Region []*Tile

func (m *Map) FindRegions() map[int]Region {
	regionID := 1
	for _, y := range *m {
		for _, x := range y {
			if x.RegionID == 0 {
				x.RegionID = regionID
				regionID++

				x.DiscoverRegion()
			}
		}
	}

	regions := map[int]Region{}
	for _, y := range *m {
		for _, x := range y {
			if v, ok := regions[x.RegionID]; ok {
				regions[x.RegionID] = append(v, x)
			} else {
				regions[x.RegionID] = Region{x}
			}
		}
	}
	return regions
}

func (r Region) FencePrice() int {
	lookup := [][]int{
		{1, 0},  // to right
		{-1, 0}, // to left
		{0, 1},  // to bottom
		{0, -1}, // to top
	}

	fences := 0
	for _, f := range r {
		for _, c := range lookup {
			x, y := f.X+c[0], f.Y+c[1]
			ti := f.Map.Get(x, y)
			if ti == nil || ti.RegionID != f.RegionID {
				fences++
			}
		}
	}

	return fences * len(r)
}

func (r Region) FencePriceBulk() int {
	lookup := [][]int{
		{1, 0},  // to right
		{-1, 0}, // to left
		{0, 1},  // to bottom
		{0, -1}, // to top
	}

	xsides, ysides := map[int]int{}, map[int]int{}
	for _, f := range r {
		for _, c := range lookup {
			x, y := f.X+c[0], f.Y+c[1]
			ti := f.Map.Get(x, y)
			if ti == nil || ti.RegionID != f.RegionID {
				if x != f.X {
					if v, ok := xsides[x]; ok {
						xsides[x] = v + 1
					} else {
						xsides[x] = 1
					}
				}
				if y != f.Y {
					if v, ok := ysides[y]; ok {
						ysides[y] = v + 1
					} else {
						ysides[y] = 1
					}
				}
			}
		}
	}

	return (len(xsides) + len(ysides)) * len(r)
}

func (t *Tile) DiscoverRegion() {
	lookup := [][]int{
		{1, 0},  // to right
		{-1, 0}, // to left
		{0, 1},  // to bottom
		{0, -1}, // to top
	}

	x, y := t.X, t.Y
	for _, l := range lookup {
		ti := t.Map.Get(x+l[0], y+l[1])
		if ti != nil {
			if ti.C == t.C && ti.RegionID == 0 {
				ti.RegionID = t.RegionID
				ti.DiscoverRegion()
			}
		}
	}
}
