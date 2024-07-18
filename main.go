package main

import (
	"fmt"
	"strconv"
)

func main() {
	qrCode := "00020101021126660015ID.OR.GPNQR.WWW01189360042500000000060214PDP123456789150303UME51450015ID.OR.GPNQR.WWW0215IDSM110121022120303UME520458125303360540822000.005502015802ID5914BAKMIE GM AEON6007TANGSEL61051034062110707SGLWO016304FCC3"

	i := 0
	for i < len(qrCode) {
		if i+2 > len(qrCode) {
			break
		}
		tag := qrCode[i : i+2]
		i += 2

		if i+2 > len(qrCode) {
			break
		}
		length, err := strconv.Atoi(qrCode[i : i+2])
		if err != nil {
			break
		}
		i += 2

		if i+length > len(qrCode) {
			break
		}
		value := qrCode[i : i+length]
		i += length

		if len(value) >= 3 {
			merchantCriteria := value[len(value)-3:]
			fmt.Printf("Tag: %s, Length: %d, Value: %s, Merchant Criteria: %s\n", tag, length, value, merchantCriteria)
		}

		// Check for subtags within the value
		j := 0
		for j < len(value) {
			if j+2 > len(value) {
				break
			}
			subTag := value[j : j+2]
			j += 2

			if j+2 > len(value) {
				break
			}
			subLength, err := strconv.Atoi(value[j : j+2])
			if err != nil {
				break
			}
			j += 2

			if j+subLength > len(value) {
				break
			}
			subValue := value[j : j+subLength]
			j += subLength

			if len(subValue) >= 3 {
				merchantCriteria := subValue[len(subValue)-3:]
				fmt.Printf("  SubTag: %s, Length: %d, Value: %s, Merchant Criteria: %s\n", subTag, subLength, subValue, merchantCriteria)
			}
		}
	}
}
