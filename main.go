package generate


import (
"math/rand"
"strconv"
)


func  GetId(n string) (string, error){
	var err error
	n1:=n[0]*n[2]+n[4]
	w1:=int(n1)
	/*Here we get the number of char in email and multiply by contenation of the indexes.*/
	d1:=len(n)
	n2:=d1*w1
	/*convert the multiplied value to string and concatenate with email*/
	n3:=strconv.Itoa(n2)
	j:=n[0:3]+n3+n[2:3]
	return j, err	

}


func GetCode(n string) (string, error){
	const maxA="abcd1e3f5ghij7klm9no6p2qrs5tu0v9wxyz"
	/*Take input from string and run algorithm from original string*/
	fore:= "0xcz"
	end:="shoppleverse"
	v:=string(len(n))
	v1:= n[0:4]+end[0:7]+v+maxA
	toRune:=[]rune(v1)
	rand.Shuffle(len(n),func (i,j int){
		toRune[i],toRune[j]=toRune[j],toRune[i]
	})
	roughchars:= string(toRune)
	mainchars:= roughchars[9:24]
	code:= fore+mainchars+end
	return code, nil
}
