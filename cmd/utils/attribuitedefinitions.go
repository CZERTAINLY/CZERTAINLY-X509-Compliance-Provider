package utils

import "CZERTAINLY-X509-Compliance-Provider/cmd/attributes"

func GetAttributeValue(attributeName string, attributes []attributes.Attributes) (value interface{}) {
	if len(attributes) == 0 {
		return nil
	}
	for _, attribute := range attributes {
		if attribute.Name == attributeName {
			return attribute.Value
		}
	}
	return nil
}

func GetRequestAttributeValue(attributeName string, attributes []attributes.RequestAttributes) (value interface{}) {
	if len(attributes) == 0 {
		return nil
	}
	for _, attribute := range attributes {
		if attribute.Name == attributeName {
			return attribute.Value
		}
	}
	return nil
}
