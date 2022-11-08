package utils

import (
	"CZERTAINLY-X509-Compliance-Provider/cmd/attributes"
	"testing"
)

func TestGetRequestAttributeValue(t *testing.T) {
	attrValue := []attributes.BaseAttributeContent{{Data: "testing"}}
	attrs := []attributes.Attributes{
		attributes.Attributes{
			UUID:    "df52f720-f239-11ec-b939-0242ac120002",
			Name:    "Test",
			Content: attrValue,
		},
	}

	valueOfAttribute := GetAttributeValue("Test", attrs, true)
	if valueOfAttribute != attrValue[0].Data {
		t.Error("Attribute Value fetch failed")
	}
}

func TestGetAttributeValue(t *testing.T) {
	attrValue := []attributes.BaseAttributeContent{{Data: "testing"}}
	attrs := []attributes.Attributes{
		attributes.Attributes{
			UUID: "df52f720-f239-11ec-b939-0242ac120002",
			Name: "Test",
			Type: "STRING",
			Properties: attributes.AttributeProperties{
				Label:       "Test",
				Required:    false,
				ReadOnly:    false,
				Editable:    false,
				Visible:     false,
				MultiSelect: false,
				List:        false,
			},
			Description: "",
			Content:     attrValue,
		},
	}

	valueOfAttribute := GetAttributeValue("Test", attrs, true)
	if valueOfAttribute != attrValue[0].Data {
		t.Error("Attribute Value fetch failed")
	}
}
