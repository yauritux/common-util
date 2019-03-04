package numeric

/**
 * Author: Yauri Attamimi (yauritux@gmail.com)
 * Version: 2.0.0-RC1
 * Description: Common utility for Numeric.
 *
 */

import (
	"fmt"

	"strconv"

	"github.com/rs/xid"
)

//const sixtyfour = uint64(^uint(0)) == ^uint64(0)

// GetRandomNumberAsString get random number as string value
func randomNumber() string {
	guid := xid.New()
	randomNum1 := guid.Counter()
	randomNum2 := guid.Counter()
	return fmt.Sprintf("%d%d", randomNum1, randomNum2)
}

// GetNextRandomInt get next random integer number
func GetNextRandomInt() (interface{}, error) {
	if Is64BitPlatform() == true {
		randomInt, err := strconv.ParseUint(randomNumber(), 0, 64)
		if err != nil {
			return 0, err
		}
		return randomInt, nil
	}
	randomInt, err := strconv.ParseUint(randomNumber(), 0, 32)
	if err != nil {
		return 0, err
	}
	return randomInt, nil
}

// Is64BitPlatform returns true if current platform is 64-bit, otherwise false is returned
func Is64BitPlatform() bool {
	return uint64(^uint(0)) == ^uint64(0)
}
