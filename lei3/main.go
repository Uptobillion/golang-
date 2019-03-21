package main

import (
	"lei3/inandout"
	"math/rand"
	"time"
)

const (
	cover, uknown, left = iota, iota, iota
	bottom, known, right
	extra, flag = iota, iota
	mine        = 11
)

var xi, yi = []int{-1, 0, +1, -1, 0, +1, -1, 0, +1}, []int{-1, -1, -1, 0, 0, 0, +1, +1, +1}

type sl struct {
	a       [][][]int
	h, w, b int
}

func (s *sl) format(cl []int) {
	s.h, s.w, s.b = cl[0], cl[1], cl[2] //
	s.a = make([][][]int, 3)
	for k := 0; k < 3; k++ {
		s.a[k] = make([][]int, s.h+2)
		for i := 0; i < s.h+2; i++ {
			s.a[k][i] = make([]int, s.w+2)
			for j := 0; j < s.w+2; j++ {
				s.a[k][i][j] = 0
			}
		}
	}
}
func (s *sl) putmine(cl []int) {
	s.a[extra][cl[0]][cl[1]] = 1
	x, y := 0, 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < s.b; {
		x = r.Intn(s.h - 1)
		y = r.Intn(s.w - 1)
		x = x + 1
		y = y + 1
		if s.a[extra][x][y] != 1 {
			s.a[bottom][x][y] = mine
			s.a[extra][x][y] = 1
			i++
		}
	}
	for y := 1; y < s.w+1; y++ {
		for x := 1; x < s.h+1; x++ {
			if s.a[bottom][x][y] != mine {
				for i := 0; i < 9; i++ {
					if s.a[bottom][x+xi[i]][y+yi[i]] == mine {
						s.a[bottom][x][y] = s.a[bottom][x][y] + 1
					}
				}
			}
		}
	}
	for j := 0; j < s.h+2; j++ {
		s.a[bottom][j][s.w+1] = 12
		s.a[bottom][j][0] = 12
	}
	for i := 0; i < s.w+2; i++ {
		s.a[bottom][0][i] = 12
		s.a[bottom][s.h+1][i] = 12
	}
	s.a[cover][cl[0]][cl[1]] = known
	s.scan()
}
func (s *sl) scan() { //range
	m := s.h + s.w
	for n := 1; n < m; n++ {
		for y := 1; y < s.w+1; y++ {
			for x := 1; x < s.h+1; x++ {
				if s.a[cover][x][y] == known {
					continue
				}
				for i := 0; i < 9; i++ {
					if s.a[cover][x+xi[i]][y+yi[i]] == known && s.a[bottom][x+xi[i]][y+yi[i]] == uknown {
						s.a[cover][x][y] = known
					}
				}
			}
		}
	}
}
func (s *sl) digui(cl []int) bool {
	if cl[2] != left {
		switch s.a[cover][cl[0]][cl[1]] {
		case uknown:
			s.a[cover][cl[0]][cl[1]] = flag
		case known:
			return s.digui2(cl)
		case flag:
			s.a[cover][cl[0]][cl[1]] = uknown
		}
	}
	if s.a[cover][cl[0]][cl[1]] == flag {
		return false
	}
	s.a[cover][cl[0]][cl[1]] = known
	if s.a[bottom][cl[0]][cl[1]] == mine {
		return true
	}
	for i := 0; i < 9; i++ {
		if s.a[bottom][cl[0]+xi[i]][cl[1]+yi[i]] == 0 {
			s.a[cover][cl[0]+xi[i]][cl[1]+yi[i]] = known
		}
	}
	s.scan()
	return false //继续
}
func (s *sl) digui2(cl []int) bool {
	flagnum := 0
	for i := 0; i < 9; i++ {
		if s.a[bottom][cl[0]+xi[i]][cl[1]+yi[i]] == mine {
			flagnum++
		}
	}
	if flagnum == s.a[bottom][cl[0]][cl[1]] {
		for i := 0; i < 9; i++ {
			if s.a[cover][cl[0]+xi[i]][cl[1]+yi[i]] == uknown {
				if s.a[bottom][cl[0]+xi[i]][cl[1]+yi[i]] != mine {
					s.a[cover][cl[0]+xi[i]][cl[1]+yi[i]] = known
				} else {
					s.a[cover][cl[0]+xi[i]][cl[1]+yi[i]] = known
					return true
				}
			}
		}
	}
	s.scan()
	return false
}
func (s *sl) success() bool { //
	m := s.w * s.h
	sucess := 0
	for y := 1; y < s.w+1; y++ {
		for x := 1; x < s.h+1; x++ {
			if s.a[cover][x][y] == known {
				sucess++
			}
		}
	}
	if sucess == m-s.b {
		return true
	} else {
		return false
	}
}
func main() {
	a := sl{}
	a.format(inandout.In())
	inandout.Out(a.h, a.w, a.a[bottom], a.a[cover])
	a.putmine(inandout.Click())
	for {
		inandout.Out(a.h, a.w, a.a[bottom], a.a[cover])
		defer inandout.Out(a.h, a.w, a.a[bottom], a.a[cover])
		if a.digui(inandout.Click()) {
			inandout.Fail() //4
			break
		}
		if a.success() {
			inandout.Success()
			break
		}
	}
}
