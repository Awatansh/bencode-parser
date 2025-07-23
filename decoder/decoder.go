package decoder

//Plan Decode: Bencode byte to bencode byte slice
//decode: parse the Bencode byte :-> decodeint decodestring etc....
import (
	"bencode-parser/types"
	"bencode-parser/utils"
	"strconv"
)

func Decode(data []byte) (types.BencodeType, error) {
	if len(data) == 0 {
		return nil, utils.NewBencodeError("empty input")
	}
	result, _, err := decode(data, 0)
	return result, err
}

func decode(data []byte, pos int) (types.BencodeType, int, error) {
	if pos >= len(data) {
		return nil, pos, utils.NewBencodeError("unexpected end of input")
	}
	switch data[pos] {
	case 'i':
		return decodeInt(data, pos)
	case 'l':
	case 'd':
	default:

	}
}

// todo: helper functions to form bencode form
func decodeInt(data []byte, pos int) (types.BencodeInt, int, error) {
	if pos >= len(data) {
		return 0, pos, utils.NewBencodeError("unexpected end during integer")
	}
	if data[pos] != 'i' {
		return 0, pos, utils.NewBencodeError("invalid integer start at position %d", pos)
	}
	pos++
	start := pos
	for pos <= len(data) && data[pos] != 'e' {
		if (data[pos] < '0' || data[pos] >= '9') && (data[pos] != '-') {
			return 0, pos, utils.NewBencodeError("invalid character in integer at position %d", pos)
		}
		pos++
	}
	if pos >= len(data) || data[pos] != 'e' {
		return 0, pos, utils.NewBencodeError("unterminated integer at pos: %d", pos)
	}
	numstr := string(data[start:pos])
	num, err := strconv.ParseInt(numstr, 10, 64)
	if err != nil {
		return 0, pos, utils.NewBencodeError("invalid integer value: %v", err)
	}
	return types.BencodeInt(num), pos + 1, nil
}

func decodeString(data []byte, pos int) (types.Bencodestring, int, error) {
	start := pos
	for pos <= len(data) && data[pos] != ':' {
		if data[pos] < '0' || data[pos] > '9' {
			return "", pos, utils.NewBencodeError("Invalid lenght in string at position : %d", pos)
		}
		pos++
	}
	if pos >= len(data) || data[pos] != ':' {
		return "", pos, utils.NewBencodeError("unterminated string lenght at pos %d", pos)
	}
	lenghtStr := string(data[start:pos])
	length, err := strconv.Atoi(lenghtStr)
	if err != nil {
		return "", pos, utils.NewBencodeError("invalid string length %v", err)
	}
	pos++
	if pos+length > len(data) {
		return "", pos, utils.NewBencodeError("string length exceeds data at position %d", pos)
	}
	return types.Bencodestring(data[pos : pos+length]), pos + length, nil
}
