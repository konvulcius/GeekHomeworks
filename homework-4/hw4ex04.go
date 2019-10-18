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

func (pos knightPosition) OtherPoint(x int, y int, point canGo) canGo {
	pos.x += x
	pos.y += y
	if pos.x > 0 && pos.x <= 8 && pos.y > 0 && pos.y <= 8 {
		point = append(point, pos)
	}
	return point
}

func (pos *knightPosition) FindWay() (canGo, error) {
	movePoint := canGo{
		{1, 2,},
		{1, -2,},
		{-1, 2,},
		{-1, -2,},
	}
	var findPoints canGo

	if pos.x < 1 || pos.y < 1 || pos.x > 8 || pos.y > 8 {
		return nil, wrong()
	}
	for _, val := range movePoint {
		findPoints = pos.OtherPoint(val.y, val.x, findPoints)
		findPoints = pos.OtherPoint(val.x, val.y, findPoints)
	}
	return findPoints, nil
}

func (pos knightPosition) String() string {
	return fmt.Sprintf("[y = %v, x = %v]", pos.y, pos.x)
}

func main() {
	some := knightPosition{7, 7}
	arr, err := some.FindWay()
	if err == nil {
		fmt.Println(arr)
	} else {
		fmt.Println(err)
	}
}
