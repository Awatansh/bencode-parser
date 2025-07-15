package types

type BencodeType interface{}

type Bencodestring string

type BencodeInt int64

type BencodeList []BencodeType

type BencodeDict map[string]BencodeType
