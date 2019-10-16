package main

import (
	"fmt"
	"math"
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

func (pos *knightPosition) FindWay() (canGo, error) {
	arr := make(canGo, 0, 8)
	var result canGo
	var newPos knightPosition

	if pos.x < 1 || pos.y < 1 || pos.x > 8 || pos.y > 8 {
		return nil, wrong()
	}
	for distance := -2; distance <= 2; distance++ {
		if distance != 0 {
			if math.Sqrt(float64(distance*distance)) == 2 {
				newPos.x = pos.x + distance
				newPos.y = pos.y + 1
				arr = append(arr, newPos)
				newPos.y = pos.y - 1
				arr = append(arr, newPos)
			} else {
				newPos.x = pos.x + distance
				newPos.y = pos.y + 2
				arr = append(arr, newPos)
				newPos.y = pos.y - 2
				arr = append(arr, newPos)
			}
		}
	}
	for _, val := range arr {
		if val.x > 0 && val.y > 0 {
			result = append(result, val)
		}
	}
	return result, nil
}

func (pos knightPosition) String() string {
	return fmt.Sprintf("[y = %v, x = %v]", pos.y, pos.x)
}

func main() {
	some := knightPosition{0, 0}
	arr, err := some.FindWay()
	if err == nil {
		fmt.Println(arr)
	} else {
		fmt.Println(err)
	}
}
