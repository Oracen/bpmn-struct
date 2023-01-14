package infrastructure

import (
	"bpmn-struct/constants"
	"bpmn-struct/validation"
	"fmt"
)

type Import struct {
	ImportType string   `xml:"importType" json:"importType"`
	Location   []string `xml:"location" json:"location"`
	Namespace  string   `xml:"namespace" json:"namespace"`
}

func CreateImport(importType, namespace string) Import {
	return Import{
		ImportType: importType,
		Location:   []string{},
		Namespace:  namespace,
	}
}

func (i Import) Validate() (errors []error) {
	name := fmt.Sprintf("Import:%s", i.ImportType)
	checks := []error{
		validation.ValNonzero(name, "ImportType", i.ImportType),
		validation.ArrZeroOne(name, "Location", i.Location),
		validation.ValNonzero(name, "Namespace", i.Namespace),
	}
	return validation.FilterErrors(checks)
}

func (i Import) setImportType(name string) Import {
	return Import{
		ImportType: name,
		Location:   i.Location,
		Namespace:  i.Namespace,
	}
}

func (i Import) SetImportTypeXML_1_0() Import {
	return i.setImportType(constants.ImportTypeXML_1_0)
}

func (i Import) SetImportTypeWSDL_2_0() Import {
	return i.setImportType(constants.ImportTypeWSDL_2_0)
}

func (i Import) SetImportTypeBPMN_2_0() Import {
	return i.setImportType(constants.ImportTypeBPMN_2_0)
}

func (i Import) SetImportTypeCustom(schemaLocation string) Import {
	return i.setImportType(schemaLocation)
}
