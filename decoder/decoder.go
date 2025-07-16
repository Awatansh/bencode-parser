package decoder

//Plan Decode: Bencode byte to bencode byte slice
//decode: parse the Bencode byte :-> decodeint decodestring etc....
import (
	"bencode-parser/types"
	"bencode-parser/utils"
)


func Decode(data []byte)(type.BencodeType,error){
	if(len(data) == 0){
		return nil,utils.NewBencodeError("empty input")
	}
	result,_,err:=decode(data,0)
}

func decode(data []byte,pos int)(types.BencodeType,int , error){
	if(pos>=len(data)){
		return nil,pos,utils.NewBencodeError("unexpected end of input")
	}
	switch data[pos]{
	case 'i':
	case 'l':
	case 'd':
	default :

	}
}

//todo: helper functions to form bencode form
