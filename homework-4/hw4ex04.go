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

func (pos knightPosition) OtherPoint(x, y int) {
	pos.x += x
	pos.y += y
	if pos.x > 0 && pos.x < 8 && pos.y > 0 && pos.y < 8 {
		allPoint = append(allPoint, pos)
	}
}

func (pos *knightPosition) FindWay() (canGo, error) {
	movePoint := canGo{
		{1, 2},
		{1, -2},
		{-1, 2,},
		{-1, -2,},
	}

	if pos.x < 1 || pos.y < 1 || pos.x > 8 || pos.y > 8 {
		return nil, wrong()
	}
	for _, val := range movePoint {
		pos.OtherPoint(val.y, val.x)
		pos.OtherPoint(val.x, val.y)
	}
	return allPoint, nil
}

func (pos knightPosition) String() string {
	return fmt.Sprintf("[y = %v, x = %v]", pos.y, pos.x)
}

func main() {
	some := knightPosition{1, 1}
	arr, err := some.FindWay()
	if err == nil {
		fmt.Println(arr)
	} else {
		fmt.Println(err)
	}
}
