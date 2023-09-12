package helpers

import (
	"fmt"
	"os"
)

func CreateFile(feature string, file string, content string) error {
	f, err := os.Create(fmt.Sprintf("%s/%s", feature, file))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
