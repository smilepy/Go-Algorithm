package string

func Concatenate(first,second string) string {
	firstRune:=[]rune(first)
	secondRune:=[]rune(second)
	firstLen:=len(firstRune)
	secondLen:=len(secondRune)
	i:=firstLen-1
	j:=0
	for i>=0&&j<=secondLen-1{
		if firstRune[i]==secondRune[j]{
			i--
			j++
		}else{
			break
		}
	}
	result:=""
	if i>=0{
		result=string(firstRune[:i+1])
	}
	if j<=secondLen-1{
		result+=string(secondRune[j:])
	}

	return result
}
