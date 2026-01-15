package day02

import (
	"math"

	"github.com/Its-Maniaco/advent-of-code-2025/internal/utils"
)

func getEvenRange(nR []utils.NumRange) []utils.NumRange {
	evenInterval := []utils.NumRange{}

	for _, r := range nR {
		lenL := utils.GetNumberOfDigits(r.Lower)
		lenU := utils.GetNumberOfDigits(r.Upper)
		// skip entire range: range has odd digits but also same number of digits
		if (lenL == lenU) && (lenL%2 != 0) {
			continue
		} else if (lenL == lenU) && (lenL%2 == 0) {
			// even range and of same length then add as is
			evenInterval = append(evenInterval, r)
		} else {
			for i := lenL - 1; i < lenU; i++ {
				if i%2 == 0 {
					continue
				}
				p := math.Pow10(i)
				lowerInterval := int(math.Max(p, float64(r.Lower)))
				upperInterval := int(math.Min(10*p-1, float64(r.Upper)))
				evenInterval = append(evenInterval, utils.NumRange{Lower: lowerInterval, Upper: upperInterval})
			}
		}
	}

	// fmt.Println("Final Intervals: ", evenInterval)
	return evenInterval
}
