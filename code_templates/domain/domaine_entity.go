package code_templates

import (
	"cleancode/code"
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateDomainEntity(featureName string) {
	lowerFeatureName := strings.ToLower(featureName)

	entityContent := fmt.Sprintf(code.EntityTemplate, lowerFeatureName, featureName, featureName, featureName, featureName, featureName, featureName, featureName, lowerFeatureName, featureName, featureName, lowerFeatureName, featureName)

	file, err := os.Create(fmt.Sprintf("%s_entity.dart", lowerFeatureName))
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprint(file, entityContent)
}
