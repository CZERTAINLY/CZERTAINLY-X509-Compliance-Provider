package compliance

import (
	"CZERTAINLY-X509-Compliance-Provider/cmd/rules"
	"CZERTAINLY-X509-Compliance-Provider/cmd/utils"

	"github.com/zmap/zcrypto/x509"
)

func ReturnRuleFramer(uuid string, name string, status rules.RuleStatus) (response ResponseRules) {
	return ResponseRules{UUID: uuid, Name: name, Status: status}
}

func boolToRuleStatusMapper(boolData bool) (status rules.RuleStatus) {
	if boolData {
		return rules.OK
	}
	return rules.NOK
}

// HashAlgorithmValidation validates the hash algorithm of the certificate
func HashingAlgorithmValidation(certificate *x509.Certificate, requestRule RequestRules, request Request, ruleDefinition rules.RuleDefinition) ResponseRules {
	condition := utils.GetRequestAttributeValue("condition", requestRule.Attributes, true)
	value := utils.GetRequestAttributeValue("algorithm", requestRule.Attributes, true)

	mappings := map[string]string{"MD5WITHRSA": "MD5-RSA", "SHA1WITHRSA": "SHA1-RSA", "SHA224WITHRSA": "SHA224-RSA", "SHA256WITHRSA": "SHA256-RSA", "SHA384WITHRSA": "SHA384-RSA", "SHA512WITHRSA": "SHA512-RSA", "SHA1WITHECDSA": "ECDSA-SHA1", "SHA224WITHECDSA": "ECDSA-SHA224", "SHA256WITHECDSA": "ECDSA-SHA256", "SHA384WITHECDSA": "ECDSA-SHA384", "SHA512WITHECDSA": "ECDSA-SHA512"}

	signatureAlgorithm := certificate.SignatureAlgorithm.String()
	isAlgorithmAvailable := utils.Contains(utils.InterfaceAsStringArray(mappings[value.(string)]), signatureAlgorithm)

	switch condition {
	case "Equals":
		return ReturnRuleFramer(ruleDefinition.UUID, ruleDefinition.Name, boolToRuleStatusMapper(isAlgorithmAvailable))
	case "NotEquals":
		return ReturnRuleFramer(ruleDefinition.UUID, ruleDefinition.Name, boolToRuleStatusMapper(!isAlgorithmAvailable))
	}

	return ResponseRules{UUID: ruleDefinition.UUID, Name: ruleDefinition.Name, Status: rules.NA}
}

// PublicKeyAlgorithmValidation validates the public key algorithm of the certificate
func PublicKeyAlgorithmValidation(certificate *x509.Certificate, requestRule RequestRules, request Request, ruleDefinition rules.RuleDefinition) ResponseRules {
	condition := utils.GetRequestAttributeValue("condition", requestRule.Attributes, true)
	value := utils.GetRequestAttributeValue("algorithm", requestRule.Attributes, true)
	pubKeyAlgorithm := certificate.PublicKeyAlgorithmName()
	isAlgorithmAvailable := utils.Contains(utils.InterfaceAsStringArray(value), pubKeyAlgorithm)
	switch condition {
	case "Equals":
		return ReturnRuleFramer(ruleDefinition.UUID, ruleDefinition.Name, boolToRuleStatusMapper(isAlgorithmAvailable))
	case "NotEquals":
		return ReturnRuleFramer(ruleDefinition.UUID, ruleDefinition.Name, boolToRuleStatusMapper(!isAlgorithmAvailable))
	}
	return ResponseRules{UUID: ruleDefinition.UUID, Name: ruleDefinition.Name, Status: rules.NA}
}

// EccCurveValidation validates the curve of the certificate
func EcCurveValidation(certificate *x509.Certificate, requestRule RequestRules, request Request, ruleDefinition rules.RuleDefinition) ResponseRules {
	condition := utils.GetRequestAttributeValue("condition", requestRule.Attributes, true)
	value := utils.GetRequestAttributeValue("algorithm", requestRule.Attributes, true)
	curve := certificate.PublicKey.(*x509.AugmentedECDSA).Pub.Curve.Params().Name
	isValid := utils.Contains(utils.InterfaceAsStringArray(value), curve)
	switch condition {
	case "Equals":
		return ReturnRuleFramer(ruleDefinition.UUID, ruleDefinition.Name, boolToRuleStatusMapper(isValid))
	case "NotEquals":
		return ReturnRuleFramer(ruleDefinition.UUID, ruleDefinition.Name, boolToRuleStatusMapper(!isValid))
	}
	return ResponseRules{UUID: ruleDefinition.UUID, Name: ruleDefinition.Name, Status: rules.NA}
}

// KeySizeValidation validates the key size of the certificate
func KeySizeValidator(certificate *x509.Certificate, requestRule RequestRules, request Request, ruleDefinition rules.RuleDefinition) ResponseRules {
	condition := utils.GetRequestAttributeValue("condition", requestRule.Attributes, true)
	value := utils.InterfaceAsInteger(utils.GetRequestAttributeValue("length", requestRule.Attributes, true))
	keySize := utils.GetPublicKeySize(certificate)

	switch condition {
	case "Equals":
		return ReturnRuleFramer(ruleDefinition.UUID, ruleDefinition.Name, boolToRuleStatusMapper(keySize == value))
	case "NotEquals":
		return ReturnRuleFramer(ruleDefinition.UUID, ruleDefinition.Name, boolToRuleStatusMapper(keySize != value))
	case "Greater":
		return ReturnRuleFramer(ruleDefinition.UUID, ruleDefinition.Name, boolToRuleStatusMapper(keySize >= value))
	case "Lesser":
		return ReturnRuleFramer(ruleDefinition.UUID, ruleDefinition.Name, boolToRuleStatusMapper(keySize <= value))
	}
	return ResponseRules{UUID: ruleDefinition.UUID, Name: ruleDefinition.Name, Status: rules.NA}
}
