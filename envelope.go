package geo

type Box struct {
	MinX float64
	MinY float64
	MaxX float64
	MaxY float64
}

//在边界上也算在内部
func IsPointInBox(b Box, p Point) bool {
	if b.MinX <= p.X && p.X <= b.MaxX && b.MinY <= p.Y && p.Y <= b.MaxY {
		return true
	}
	return false
}

func BoxToGeo(b Box) Geometry {
	p1 := Point{b.MinX, b.MinY}
	p2 := Point{b.MinX, b.MaxY}
	p3 := Point{b.MaxX, b.MaxY}
	p4 := Point{b.MaxX, b.MinY}

	if p1.Equal(p3) {
		//元素是个点
		return p1
	} else if p1.Equal(p2) {
		//元素是条线 y坐标不同
		return LineString{p1, p3}
	} else if p2.Equal(p3) {
		//元素是条线 x坐标不同
		return LineString{p1, p2}
	}
	return *NewPolygon(LinearRing{p1, p2, p3, p4})
}

func calBox(points ...Point) Box {
	var minX, minY, maxX, maxY float64 = INF, INF, -INF, -INF
	for _, v := range points {
		if minX > v.X {
			minX = v.X
		}
		if minY > v.Y {
			minY = v.Y
		}
		if maxX < v.X {
			maxX = v.X
		}
		if maxY < v.Y {
			maxY = v.Y
		}
	}
	return Box{minX, minY, maxX, maxY}
}
