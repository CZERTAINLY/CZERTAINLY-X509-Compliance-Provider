package utils

import "CZERTAINLY-X509-Compliance-Provider/cmd/attributes"

// GetAttributeValue returns the value of the attribute with the given name.
func GetAttributeValue(attributeName string, attributes []attributes.Attributes, returnSingleEntry bool) (value interface{}) {
	if len(attributes) == 0 {
		return nil
	}
	for _, attribute := range attributes {
		if attribute.Name == attributeName {
			if returnSingleEntry {
				return attribute.Content[0].Data
			}
			return attribute.Content
		}
	}
	return nil
}

// GetRequiredAttributeValue returns the value of the required attribute with the given name.
func GetRequestAttributeValue(attributeName string, attrs []attributes.RequestAttributes, returnSingleEntry bool) (value interface{}) {
	if len(attrs) == 0 {
		return nil
	}
	for _, attribute := range attrs {
		if attribute.Name == attributeName {
			if returnSingleEntry {
				return attribute.Content[0].Data
			}
			return attribute.Content
		}
	}
	return nil
}
