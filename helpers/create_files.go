package helpers

import (
	data "cleancode/code_templates/data"

	domain "cleancode/code_templates/domain"

	"github.com/fatih/color"
)

func CreateAllFiles(feature string) {
	data.CreateDataRepository(feature)
	color.Yellow("+ File %s_repository.dart created successfully...", feature)
	data.CreateDataModel(feature)
	color.Yellow("+ File %s_model.dart created successfully...", feature)
	domain.CreateDomainEntity(feature)
	color.Yellow("+ File %s_entity.dart created successfully...", feature)
	domain.CreateDomainUsecase(feature)
	color.Yellow("+ File %s_usecase.dart created successfully...", feature)
}
