package darts

import "math"

func Score(x, y float64) int {
	dis := math.Sqrt(x*x +y*y)	
	res := 0
    if dis > 10{
        
    }else if dis > 5{
        res = 1
    } else if dis > 1{
        res = 5
    }else {
        res = 10
    }
    return res
}
