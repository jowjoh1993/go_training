package conversion

import "strconv"

func StringsToFloat(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))
	for index, str := range strings {
		flt, err := strconv.ParseFloat(str, 64)

		if err != nil {
			return nil, err
		}

		floats[index] = flt
	}
	return floats, nil
}
