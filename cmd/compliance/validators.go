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

func HashingAlgorithmValidation(certificate *x509.Certificate, requestRule RequestRules, request Request, ruleDefinition rules.RuleDefinition) ResponseRules {
	condition := utils.GetRequestAttributeValue("condition", requestRule.Attributes)
	value := utils.GetRequestAttributeValue("algorithm", requestRule.Attributes)

	signatureAlgorithm := certificate.SignatureAlgorithmName()
	isAlgorithmAvailable := utils.Contains(utils.InterfaceAsStringArray(value), signatureAlgorithm)

	switch condition {
	case "Equals":
		return ReturnRuleFramer(ruleDefinition.UUID, ruleDefinition.Name, boolToRuleStatusMapper(isAlgorithmAvailable))
	case "NotEquals":
		return ReturnRuleFramer(ruleDefinition.UUID, ruleDefinition.Name, boolToRuleStatusMapper(!isAlgorithmAvailable))
	}

	return ResponseRules{UUID: ruleDefinition.UUID, Name: ruleDefinition.Name, Status: rules.NA}
}

func PublicKeyAlgorithmValidation(certificate *x509.Certificate, requestRule RequestRules, request Request, ruleDefinition rules.RuleDefinition) ResponseRules {
	condition := utils.GetRequestAttributeValue("condition", requestRule.Attributes)
	value := utils.GetRequestAttributeValue("algorithm", requestRule.Attributes)
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

func EcCurveValidation(certificate *x509.Certificate, requestRule RequestRules, request Request, ruleDefinition rules.RuleDefinition) ResponseRules {
	condition := utils.GetRequestAttributeValue("condition", requestRule.Attributes)
	value := utils.GetRequestAttributeValue("algorithm", requestRule.Attributes)
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

func KeySizeValidator(certificate *x509.Certificate, requestRule RequestRules, request Request, ruleDefinition rules.RuleDefinition) ResponseRules {
	condition := utils.GetRequestAttributeValue("condition", requestRule.Attributes)
	value := utils.InterfaceAsInteger(utils.GetRequestAttributeValue("length", requestRule.Attributes))
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
