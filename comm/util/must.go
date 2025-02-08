package util

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

func MustBool(str string) bool {
	return str == "true"
}

func MustInt64(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}

func MustInt64s(str string) (d []int64) {
	nbr := strings.Split(str, ",")
	for i := range nbr {
		d = append(d, MustInt64(nbr[i]))
	}
	return
}

func MustUInt64(str string) uint64 {
	i, _ := strconv.ParseUint(str, 10, 64)
	return i
}

func MustUInt64s(str string) (d []uint64) {
	nbr := strings.Split(str, ",")
	for i := range nbr {
		d = append(d, MustUInt64(nbr[i]))
	}
	return
}

func MustInt32(str string) int32 {
	i, _ := strconv.ParseInt(str, 10, 32)
	return int32(i)
}

func MustInt(str string) int {
	i, _ := strconv.ParseInt(str, 10, 32)
	return int(i)
}

func MustUInt32(str string) uint32 {
	i, _ := strconv.ParseUint(str, 10, 32)
	return uint32(i)
}

func MustFloat32(str string) float32 {
	i, _ := strconv.ParseFloat(str, 32)
	return float32(i)
}

func MustFloat64(str string) float64 {
	i, _ := strconv.ParseFloat(str, 64)
	return i
}

func NumBit(i uint64) (bit uint64) {
	for i > 0 {
		if i > 0 {
			bit += 1
		}
		i = i / 10
	}
	return bit
}

func MustAge(birth string) int {
	date, err := time.Parse("2006-01-02", birth)
	if err != nil {
		return 0
	}
	return time.Now().Year() - date.Year()
}

func MustJson(obj interface{}) string {
	vb, _ := json.Marshal(obj)
	return string(vb)
}
