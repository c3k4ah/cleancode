package code_templates

import (
	"cleancode/code"
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateDataModel(featureName string) {
	lowerFeatureName := strings.ToLower(featureName)

	modelContent := fmt.Sprintf(code.ModelTemplate, lowerFeatureName, featureName, featureName, featureName, featureName, featureName, featureName)
	file, err := os.Create(fmt.Sprintf("%s/data/models/%s_model.dart", featureName, lowerFeatureName))

	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprint(file, modelContent)
}
