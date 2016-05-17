package tools

import (
	"regexp"
	"strconv"
)

// @Title Generate md5
// @Description Generate 32-bit string md5
// @Success string
// bytes=0-1
//bytes=([0-9]*)-([0-9]*)
func SplitRange(rangestr string) (int64, int64, error){
	reg := regexp.MustCompile(`[0-9]+`)
	result := reg.FindAllString(rangestr, -1)
	start, err := strconv.ParseInt(result[0], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	end, err := strconv.ParseInt(result[1], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return start, end, nil
}

