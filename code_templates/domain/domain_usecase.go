package code_templates

import (
	"cleancode/code"
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateDomainUsecase(featureName string) {
	lowerFeatureName := strings.ToLower(featureName)

	usecaseContent := fmt.Sprintf(code.UsecaseTemplate, lowerFeatureName, featureName, featureName, featureName, featureName, featureName, featureName, featureName, lowerFeatureName, featureName, featureName, lowerFeatureName, featureName)

	file, err := os.Create(fmt.Sprintf("%s_usecase.dart", lowerFeatureName))
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprint(file, usecaseContent)
}
