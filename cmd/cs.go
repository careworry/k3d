package cmd

import "strconv"

func openPE() (string, error) {
	_, err := strconv.Atoi("123")

	if err != nil {
		return "", err
	}

	_, err2 := strconv.Atoi("12a")

	if err2 != nil {
		return "", err
	}

	return "666", nil
}
