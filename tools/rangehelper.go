package tools

import (
	"regexp"
	"strconv"
//	"golangapi/models"
)

// @Title Split Range
// @Description split range to start and end
// @Success int64 int 64 error
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

