package compliance

import (
	"CZERTAINLY-X509-Compliance-Provider/cmd/rules"
	"encoding/base64"
	"strings"

	"go.uber.org/zap"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3"
	"github.com/zmap/zlint/v3/lint"
)

type service struct {
	logger *zap.SugaredLogger
}

func NewService(logger *zap.SugaredLogger) Service {
	return &service{
		logger: logger,
	}
}

func (s service) ComplianceCheck(kind string, request Request) (response Response, err error) {
	s.logger.Info("Entering ComplianceCheck with Kind:  ", kind, " request ", request)
	// Arrest any error in the process and return the status as NO
	defer func() {
		if p := recover(); p != nil {
			response = Response{Status: NOK, Rules: nil}
			s.logger.Error(err)
			err = nil
		}
	}()

	var ruleNamesForRegistry []string
	var complianceResponse []ResponseRules

	// Normalize the certificate string to remove the newline characters
	s.logger.Debug("Incoming certificate: ", request.Certificate)
	certificate := s.normalizeCertificate(request.Certificate)
	s.logger.Debug("Normalized Certificate: ", certificate)
	block, err := base64.StdEncoding.DecodeString(certificate)
	s.logger.Error(err)
	parsed, err := x509.ParseCertificate(block)
	s.logger.Warn(err)

	// Iterate and evaluate the rules, split them for custom rule and rules applicable with zlint
	for _, rule := range request.Rules {
		ruleDefinition := rules.GetRuleFromUuid(rule.UUID)
		s.logger.Debug("Rule definition: ", ruleDefinition)
		if strings.HasPrefix(ruleDefinition.Name, "cus_") {
			s.logger.Info("Custom rule identified. Custom logic will be applied")
			complianceResponse = append(complianceResponse, s.evaluateCustomRule(parsed, rule, request, ruleDefinition))
		} else {
			s.logger.Info("ZLint Compliance Rule for ", rule.UUID)
			ruleNamesForRegistry = append(ruleNamesForRegistry, ruleDefinition.Name)
		}
	}
	s.logger.Info("Total rules for ZLinter: ", len(ruleNamesForRegistry))

	//If there are rules to apply for zlint, apply them
	if len(ruleNamesForRegistry) > 0 {
		s.logger.Debug("Rule names for the registeration: ", ruleNamesForRegistry)
		registry, _ := lint.GlobalRegistry().Filter(lint.FilterOptions{
			IncludeNames: ruleNamesForRegistry,
		})

		zlintResultSet := zlint.LintCertificateEx(parsed, registry)
		for item, index := range zlintResultSet.Results {
			complianceResponse = append(complianceResponse, ResponseRules{UUID: rules.GetRuleUuidFromName(item), Name: item, Status: s.getStatus(index.Status)})
		}
	}
	s.logger.Debug("Compliance Response: ", complianceResponse)

	// Return the compliance response
	return Response{s.computeOverallStatus(complianceResponse), complianceResponse}, nil
}

// normalizeCertificate removes the newline characters from the certificate
func (s service) normalizeCertificate(certificate string) string {
	replacer := strings.NewReplacer("-----BEGIN CERTIFICATE-----", "", "-----END CERTIFICATE-----", "", "\n", "", "\r", "")
	return replacer.Replace(certificate)
}

// getStatus returns the status of the rule based on the result of zlint
func (s service) getStatus(status lint.LintStatus) rules.RuleStatus {
	s.logger.Debug("Entering getStatus with Status: ", status)
	switch status {
	case lint.NA, lint.NE:
		s.logger.Debug("Compliance Result: ", rules.NA)
		return rules.NA
	case lint.Pass:
		s.logger.Debug("Compliance Result: ", rules.OK)
		return rules.OK
	default:
		s.logger.Debug("Compliance Result: ", rules.NOK)
		return rules.NOK
	}
}

// evaluateCustomRule evaluates the custom rules and produces the response
func (s service) computeOverallStatus(responseRules []ResponseRules) (status Status) {
	s.logger.Info("Entering Overall Compliance Status Calculation")
	isNotApplicable := true
	isNonCompliant := false

	for _, s := range responseRules {
		if s.Status == rules.NOK {
			isNonCompliant = true
			isNotApplicable = false
			break
		} else if s.Status == rules.OK {
			isNotApplicable = false
		}
	}
	s.logger.Info("Compliant Status: ", isNonCompliant)

	if isNotApplicable {
		return NA
	} else if isNonCompliant {
		return NOK
	} else {
		return OK
	}
}

// evaluateCustomRule evaluates the custom rules and produces the response
func (s service) evaluateCustomRule(certificate *x509.Certificate, requestRule RequestRules, request Request, ruleDefinition rules.RuleDefinition) (response ResponseRules) {
	s.logger.Info("Evaluating custom rule: ", requestRule, " Rule Definition ", ruleDefinition)
	defer func() {
		if p := recover(); p != nil {
			response = s.frameErrorValidation(ruleDefinition.UUID, ruleDefinition.Name, rules.NOK)
			s.logger.Error(p)
		}
	}()
	var compliant ResponseRules
	switch ruleDefinition.Name {
	case "cus_signature_algorithm":
		s.logger.Info("Evaluating Signature Algorithm with parameters: ", request)
		compliant = HashingAlgorithmValidation(certificate, requestRule, request, ruleDefinition)
	case "cus_public_key_algorithm":
		s.logger.Info("Evaluating Public Key Algorithm with parameters: ", request)
		compliant = PublicKeyAlgorithmValidation(certificate, requestRule, request, ruleDefinition)
	case "cus_elliptic_curve":
		s.logger.Info("Evaluating Elliptic Curve with parameters: ", request)
		compliant = EcCurveValidation(certificate, requestRule, request, ruleDefinition)
	case "cus_key_length":
		s.logger.Info("Evaluating Key length with parameters: ", request)
		compliant = KeySizeValidator(certificate, requestRule, request, ruleDefinition)
	}
	return compliant
}

func (s service) frameErrorValidation(uuid string, name string, status rules.RuleStatus) (response ResponseRules) {
	return ReturnRuleFramer(uuid, name, status)
}
