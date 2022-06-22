package utils

import (
	"CZERTAINLY-X509-Compliance-Provider/cmd/attributes"
	"testing"
)

func TestGetRequestAttributeValue(t *testing.T) {
	attrValue := attributes.BaseAttributeContent{Value: "testing"}
	attrs := []attributes.Attributes{
		attributes.Attributes{
			UUID:    "df52f720-f239-11ec-b939-0242ac120002",
			Name:    "Test",
			Content: attrValue,
		},
	}

	valueOfAttribute := GetAttributeValue("Test", attrs)
	if valueOfAttribute != attrValue {
		t.Error("Attribute Value fetch failed")
	}
}

func TestGetAttributeValue(t *testing.T) {
	attrValue := attributes.BaseAttributeContent{Value: "testing"}
	attrs := []attributes.Attributes{
		attributes.Attributes{
			UUID:        "df52f720-f239-11ec-b939-0242ac120002",
			Name:        "Test",
			Label:       "Test",
			Type:        "STRING",
			Required:    false,
			ReadOnly:    false,
			Editable:    false,
			Visible:     false,
			MultiSelect: false,
			Description: false,
			Content:     attrValue,
			List:        false,
		},
	}

	valueOfAttribute := GetAttributeValue("Test", attrs)
	if valueOfAttribute != attrValue {
		t.Error("Attribute Value fetch failed")
	}
}
