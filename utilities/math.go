/*
Copyright © 2021-2024 Cameron Esfahani
*/

package utilities

import (
	"iter"
	"slices"
	"sort"

	"golang.org/x/exp/constraints"
)

func PrimeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

func AbsoluteDifference(a, b int) int {
	if a < b {
		return b - a
	}

	return a - b
}

type Size2D struct {
	Width  int
	Height int
}

func NewSize2D(Width int, Height int) Size2D {
	return Size2D{Width, Height}
}

type Point2D struct {
	X int
	Y int
}

func NewPoint2D(X int, Y int) Point2D {
	return Point2D{X, Y}
}

func (p Point2D) Up() Point2D {
	return NewPoint2D(p.X, p.Y-1)
}

func (p Point2D) UpRight() Point2D {
	return NewPoint2D(p.X+1, p.Y-1)
}

func (p Point2D) UpLeft() Point2D {
	return NewPoint2D(p.X-1, p.Y-1)
}

func (p Point2D) Down() Point2D {
	return NewPoint2D(p.X, p.Y+1)
}

func (p Point2D) DownRight() Point2D {
	return NewPoint2D(p.X+1, p.Y+1)
}

func (p Point2D) DownLeft() Point2D {
	return NewPoint2D(p.X-1, p.Y+1)
}

func (p Point2D) Left() Point2D {
	return NewPoint2D(p.X-1, p.Y)
}

func (p Point2D) Right() Point2D {
	return NewPoint2D(p.X+1, p.Y)
}

type SetPoint2D struct {
	Points map[Point2D]bool
}

func NewSetPoint2D() *SetPoint2D {
	set := &SetPoint2D{
		Points: make(map[Point2D]bool),
	}

	return set
}

func (s *SetPoint2D) Add(point Point2D) {
	s.Points[point] = true
}

func (s *SetPoint2D) Remove(point Point2D) {
	delete(s.Points, point)
}

func (s *SetPoint2D) Exists(point Point2D) bool {
	_, ok := s.Points[point]

	return ok
}

func (s *SetPoint2D) Size() int {
	return len(s.Points)
}

func (s *SetPoint2D) All() iter.Seq[Point2D] {
	return func(yield func(Point2D) bool) {
		for v := range s.Points {
			if !yield(v) {
				return
			}
		}
	}
}

func SortPoints(points []Point2D) {
	sort.Slice(points, func(i, j int) bool {

		if points[i].Y < points[j].Y {
			return true
		}

		if points[i].Y == points[j].Y && points[i].X < points[j].X {
			return true
		}

		return false
	})
}

func ManhattanDistance(p1, p2 Point2D) int {
	return AbsoluteDifference(p1.X, p2.X) + AbsoluteDifference(p1.Y, p2.Y)
}

func IsLeft(start Point2D, end Point2D, point Point2D) int {
	val := (end.X-start.X)*(point.Y-start.Y) - (point.X-start.X)*(end.Y-start.Y)
	return val
}

func PointInPolyCrossing(point Point2D, vertices []Point2D) bool {
	crossing := 0

	for i := 0; i < len(vertices)-1; i++ {
		if (vertices[i].Y <= point.Y && vertices[i+1].Y > point.Y) ||
			(vertices[i].Y > point.Y && vertices[i+1].Y <= point.Y) {
			vt := (point.Y - vertices[i].Y) / (vertices[i+1].Y - vertices[i].Y)
			if point.X < vertices[i].X+vt*(vertices[i+1].X-vertices[i].X) {
				crossing++
			}
		}

	}

	return (crossing & 1) != 0
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func MakeHistogram(list []int) map[int]int {
	hist := make(map[int]int, 0)

	for _, id := range list {
		if existing, ok := hist[id]; ok {
			hist[id] = existing + 1
		} else {
			hist[id] = 1
		}
	}

	return hist
}

func Concatenate(a int64, b int64) int64 {
	padding := func(n int64) int64 {
		p := int64(10)
		for p <= n {
			p *= 10
		}
		return p
	}

	return a*padding(b) + b
}

func CalculateSlopeIntercept(pointA Point2D, pointB Point2D) (m float64, b float64) {
	// y = mx + b
	m = float64(pointB.Y-pointA.Y) / float64(pointB.X-pointA.X)
	b = float64(pointA.Y) - m*float64(pointA.X)
	return
}

type PointPair struct {
	One Point2D
	Two Point2D
}

func GenerateUniquePointPairs(list []Point2D) []PointPair {

	uniqueSet := make(map[PointPair]bool)
	pairs := make([]PointPair, 0)

	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list); j++ {
			if j == i {
				continue
			}

			// Sort pairs so left is always smaller than right.
			pairList := []Point2D{list[i], list[j]}
			slices.SortFunc(pairList, func(a, b Point2D) int {
				if a.X < b.X {
					return -1
				}
				if a.X > b.X {
					return 1
				}
				// X are equal
				if a.Y < b.Y {
					return -1
				}
				if a.Y > b.Y {
					return 1
				}
				// X and Y are equal
				return 0
			})

			pair := PointPair{pairList[0], pairList[1]}

			if _, ok := uniqueSet[pair]; ok {
				continue
			}

			uniqueSet[pair] = true
		}
	}

	for p := range uniqueSet {
		pairs = append(pairs, p)
	}

	return pairs
}

func DigitCount(number int) int {
	if number == 0 {
		return 1
	}

	count := 0

	for number != 0 {
		number /= 10
		count += 1
	}

	return count
}

func Abs[T constraints.Integer](a T) T {
	if a < 0 {
		return -a
	}
	return a
}
