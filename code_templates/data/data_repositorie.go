package code_templates

import (
	"cleancode/code"

	"fmt"
	"log"
	"os"
	"strings"
)

func CreateDataRepository(featureName string) {
	lowerFeatureName := strings.ToLower(featureName)

	repositoryContent := fmt.Sprintf(code.RepositoryTemplate, lowerFeatureName, featureName, featureName, featureName, featureName, featureName, featureName, featureName, lowerFeatureName, featureName, featureName, lowerFeatureName, featureName)

	file, err := os.Create(fmt.Sprintf("%s/data/repositories/%s_repository.dart", featureName, lowerFeatureName))
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprint(file, repositoryContent)

}
