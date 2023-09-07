package helpers

import (
	"log"
	"os"
)

func createDir(feature string) {
	
	folders := []string{
		feature+"/domain/usecases",
		feature+"/domain/entities",
		feature+"/domain/repositories",
		feature+"/data/repositories",
		feature+"/data/models",
		feature+"/data/datasources",
		feature+"/presentation/pages",
		feature+"/presentation/blocs",
	}

	for _, folder := range folders {
		err := os.MkdirAll(folder, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}