package filereader

import (
	tmn "Day16/transmission"
	"fmt"
	"io/ioutil"
)

func ReadTransmission(f string) (*tmn.Transmission, error) {
	transmissionFile, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf("Could not read file: %v", f)
	}

	transmission := tmn.NewHexTransmission(string(transmissionFile))
	return transmission, nil
}
