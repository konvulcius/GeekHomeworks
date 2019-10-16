package main

import (
	"fmt"
)

type wrongPos struct {
	Some string
}

func (e *wrongPos) Error() string {
	return fmt.Sprintf("%v", e.Some)
}

func wrong() error {
	return &wrongPos{
		"Knight is outside the board!",
	}
}

type knightPosition struct {
	y int
	x int
}

type canGo []knightPosition

var allPoint canGo

func (pos knightPosition) FirstPoint() {
	pos.x += 1
	pos.y += 2
	if pos.x < 8 && pos.y < 8 {
		allPoint = append(allPoint, pos)
	}
}
func (pos knightPosition) SecondPoint() {
	pos.x += 1
	pos.y -= 2
	if pos.x < 8 && pos.y > 0 {
		allPoint = append(allPoint, pos)
	}
}
func (pos knightPosition) FirdPoint() {
	pos.x -= 1
	pos.y += 2
	if pos.x > 0 && pos.y < 8 {
		allPoint = append(allPoint, pos)
	}
}
func (pos knightPosition) FourthPoint() {
	pos.x -= 1
	pos.y -= 2
	if pos.x > 0 && pos.y > 0 {
		allPoint = append(allPoint, pos)
	}
}
func (pos knightPosition) FifthPoint() {
	pos.x += 2
	pos.y += 1
	if pos.x < 8 && pos.y < 8 {
		allPoint = append(allPoint, pos)
	}
}
func (pos knightPosition) SixthPoint() {
	pos.x -= 2
	pos.y += 1
	if pos.x > 0 && pos.y < 8 {
		allPoint = append(allPoint, pos)
	}
}
func (pos knightPosition) SeventhPoint() {
	pos.x += 2
	pos.y -= 1
	if pos.x < 8 && pos.y > 0 {
		allPoint = append(allPoint, pos)
	}
}
func (pos knightPosition) EigthPoint() {
	pos.x -= 2
	pos.y -= 1
	if pos.x > 0 && pos.y > 0 {
		allPoint = append(allPoint, pos)
	}
}

func (pos *knightPosition) FindWay() (canGo, error) {
	if pos.x < 1 || pos.y < 1 || pos.x > 8 || pos.y > 8 {
		return nil, wrong()
	}
	pos.FirstPoint()
	pos.SecondPoint()
	pos.FirdPoint()
	pos.FourthPoint()
	pos.FifthPoint()
	pos.SixthPoint()
	pos.SeventhPoint()
	pos.EigthPoint()
	return allPoint, nil
}

func (pos knightPosition) String() string {
	return fmt.Sprintf("[y = %v, x = %v]", pos.y, pos.x)
}

func main() {
	some := knightPosition{5, 5}
	arr, err := some.FindWay()
	if err == nil {
		fmt.Println(arr)
	} else {
		fmt.Println(err)
	}
}
