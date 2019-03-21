package main

import (
	"lei3/inandout"
	"math/rand"
	"time"
)

const (
	cover = iota
	bottom
	extra
)

/*cover为s.a[0]bottom为s.a[1]extra为s.a[2]在s.a[0]里0为未知1为开启2为插旗*/
var xi, yi = []int{-1, 0, +1, -1, 0, +1, -1, 0, +1}, []int{-1, -1, -1, 0, 0, 0, +1, +1, +1}

type sl struct {
	a       [][][]int
	h, w, b int
	c       []int
}

func (s *sl) format(cl1 []int) {
	s.h, s.w, s.b = cl1[0], cl1[1], cl1[2] //
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
	s.c = cl
	s.a[2][s.c[0]][s.c[1]] = 1
	x, y := 0, 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < s.b; {
		x = r.Intn(s.h - 1)
		y = r.Intn(s.w - 1)
		x = x + 1
		y = y + 1
		if s.a[2][x][y] != 1 {
			s.a[1][x][y] = 11
			s.a[2][x][y] = 1
			i++
		}
	}
	for y := 1; y < s.w+1; y++ {
		for x := 1; x < s.h+1; x++ {
			if s.a[1][x][y] != 11 {
				for i := 0; i < 9; i++ {
					if s.a[1][x+xi[i]][y+yi[i]] == 11 {
						s.a[1][x][y] = s.a[1][x][y] + 1
					}
				}
			}
		}
	}
	for j := 0; j < s.h+2; j++ {
		s.a[1][j][s.w+1] = 12
		s.a[1][j][0] = 12
	}
	for i2 := 0; i2 < s.w+2; i2++ {
		s.a[1][0][i2] = 12
		s.a[1][s.h+1][i2] = 12
	}
}
func (s *sl) scan() {
	m := s.h + s.w
	for n := 1; n < m; n++ {
		for y := 1; y < s.w+1; y++ {
			for x := 1; x < s.h+1; x++ {
				if s.a[0][x][y] == 1 {
					continue
				}
				for i := 0; i < 9; i++ {
					if s.a[0][x+xi[i]][y+yi[i]] == 1 && s.a[1][x+xi[i]][y+yi[i]] == 0 {
						s.a[0][x][y] = 1
					}
				}
			}
		}
	}
}
func (s *sl) digui(cl []int) bool {
	s.c = cl
	if s.c[2] == 1 {
		switch s.a[0][s.c[0]][s.c[1]] {
		case 0:
			s.a[0][s.c[0]][s.c[1]] = 2
		case 1:
			return s.digui2(s.c)
		case 2:
			s.a[0][s.c[0]][s.c[1]] = 0
		}
	}
	if s.a[0][s.c[0]][s.c[1]] == 2 {
		return true //失败
	}
	s.a[0][s.c[0]][s.c[1]] = 1
	if s.a[1][s.c[0]][s.c[1]] == 11 {
		return true
	}
	for i := 0; i < 9; i++ {
		if s.a[1][s.c[0]+xi[i]][s.c[1]+yi[i]] == 0 {
			s.a[0][s.c[0]+xi[i]][s.c[1]+yi[i]] = 1
		}
	}
	s.scan()
	return false //继续
}
func (s *sl) digui2(cl []int) bool {
	s.c = cl
	flagnum := 0
	for i := 0; i < 9; i++ {
		if s.a[1][s.c[0]+xi[i]][s.c[1]+yi[i]] == 11 {
			flagnum++
		}
	}
	if flagnum == s.a[1][s.c[0]][s.c[1]] {
		for i := 0; i < 9; i++ {
			if s.a[0][s.c[0]+xi[i]][s.c[1]+yi[i]] == 0 {
				if s.a[1][s.c[0]+xi[i]][s.c[1]+yi[i]] == 11 {
					s.a[0][s.c[0]+xi[i]][s.c[1]+yi[i]] = 1
					return true
				}
			}
			s.a[0][s.c[0]+xi[i]][s.c[1]+yi[i]] = 1
		}
		s.scan()
	}
	return false
}
func (s *sl) success() bool {
	m := s.w * s.h
	sucess := 0
	for y := 1; y < s.w+1; y++ {
		for x := 1; x < s.h+1; x++ {
			if s.a[0][x][y] == 1 {
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
	go a.format(inandout.In())
	inandout.Out(a.h, a.w, a.a[1], a.a[0])
	a.putmine(inandout.Click())
	for { //defer
		inandout.Out(a.h, a.w, a.a[1], a.a[0])
		defer inandout.Out(a.h, a.w, a.a[1], a.a[0])
		if a.digui(inandout.Click()) {
			inandout.Fail() //4
			return
		}
		if a.success() {
			inandout.Success()
			return
		}
	}
}
