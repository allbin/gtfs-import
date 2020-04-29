package time

import (
	"fmt"
	"strconv"
	"strings"
)

func FormatTimeString(time string, sep string) string {
	parts := strings.Split(time, sep)
	int_parts := []int{}
	for i, _ := range parts {
		val, err := strconv.Atoi(parts[i])
		if err != nil {
			panic(err)
		}
		int_parts = append(int_parts, val)
	}
	return fmt.Sprintf("%02d"+sep+"%02d"+sep+"%02d", int_parts[0], int_parts[1], int_parts[2])
}
