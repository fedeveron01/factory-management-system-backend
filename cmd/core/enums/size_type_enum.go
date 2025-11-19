package enums

import "strings"

const (
	Kid    Enum = "kid"
	Woman  Enum = "woman"
	Man    Enum = "man"
	Unisex Enum = "unisex"
)

var mapSizeTypeEnum = map[string]Enum{
	"kid":    Kid,
	"woman":  Woman,
	"man":    Man,
	"unisex": Unisex,
}

func StringToSizeTypeEnum(enum string) Enum {
	enum = strings.ToLower(enum)
	return mapSizeTypeEnum[enum]
}
