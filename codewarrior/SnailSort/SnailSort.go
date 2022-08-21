package SnailSort

type step struct {
	x int
	y int
}

func Snail(snaipMap [][]int) []int {
	length := len(snaipMap)
	//if length == 0 {
	//	return []int{}
	//}
	if len(snaipMap[0]) == 0 {
		return []int{}
	}

	steps := []step{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0}}

	//var walkMap = make [length][length]int{}
	walkMap := make([][]int, length)
	for i := range walkMap {
		walkMap[i] = make([]int, length)
	}

	//var walkMap [3][3]int
	var res []int
	posX, posY := 0, 0
	turns := 0

	MoveOn := func(x int, y int) (x2, y2 int) {
		return x + steps[turns%4].x, y + steps[turns%4].y
	}

	CantMove := func(x int, y int) bool {
		if x < 0 || x >= length ||
			y < 0 || y >= length ||
			walkMap[x][y] == 1 {
			return true
		}
		return false
	}

	for {
		res = append(res, snaipMap[posX][posY])
		//fmt.Println(posX, posY)
		//fmt.Println(snaipMap[posX][posY])
		walkMap[posX][posY] = 1
		posX2, posY2 := MoveOn(posX, posY)
		//fmt.Println(posX2, posY2)
		if CantMove(posX2, posY2) {
			turns = turns + 1
			// still cant move? the end
			posX, posY = MoveOn(posX, posY)
			if CantMove(posX, posY) {
				//Get Out , The End
				break
			}
		} else {
			posX, posY = posX2, posY2
		}
	}
	return res
}
