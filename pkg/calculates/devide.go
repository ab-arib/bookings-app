package calculates

import "errors"

// devideValue add two float and return the
func DevideValue(x, y float32) (float32, error) {
	// error handler
	if y <= 0 {
		err := errors.New("cannot devide by zero")
		return 0, err
	}

	result := x / y
	return result, nil
}
