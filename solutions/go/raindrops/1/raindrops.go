package raindrops

func Convert(number int) string {
    str := ""
    if number % 3 == 0{
        str += "Pling"
    }

    if number % 5 == 0{
        str += "Plang"
    }

    if number % 7 == 0{
        str += "Plong"
    }

    if str != ""{
        return str
    }
    
    var a []rune
    for number > 9{
        a = append([]rune{rune((number % 10) + '0')}, a...)
        number /= 10
    }
    a = append([]rune{rune((number % 10) + '0')}, a...)
    
    return string(a)
}
