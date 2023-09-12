package helpers

import (
	"fmt"
	"os"
)

func CreateDir(feature string) error {

	folders := []string{
		fmt.Sprintf("%s/data/repositories", feature),
		fmt.Sprintf("%s/data/models", feature),
		fmt.Sprintf("%s/data/datasources", feature),
		fmt.Sprintf("%s/domain/usecases", feature),
		fmt.Sprintf("%s/domain/entities", feature),
		fmt.Sprintf("%s/domain/repositories", feature),
		fmt.Sprintf("%s/presentation/pages", feature),
		fmt.Sprintf("%s/presentation/blocs", feature),
	}

	for _, folder := range folders {
		err := os.MkdirAll(folder, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
