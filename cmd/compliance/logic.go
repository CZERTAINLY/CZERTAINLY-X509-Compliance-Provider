package compliance

import (
	"CZERTAINLY-X509-Compliance-Provider/cmd/rules"
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3"
	"github.com/zmap/zlint/v3/lint"
)

type service struct {
	logger log.Logger
}

func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

func (s service) ComplianceCheck(ctx context.Context, kind string, request Request) (response Response, err error) {
	defer func() {
		if p := recover(); p != nil {
			response = Response{Status: NON_COMPLIANT, Rules: nil}
			err = nil
		}
	}()

	var ruleNamesForRegistry []string
	var complianceResponse []ResponseRules
	certificate := normalizeCertificate(request.Certificate)
	block, err := base64.StdEncoding.DecodeString(certificate)
	parsed, err := x509.ParseCertificate(block)
	fmt.Println(err)

	for _, rule := range request.Rules {
		ruleDefinition := rules.GetRuleFromUuid(rule.UUID)
		if strings.HasPrefix(ruleDefinition.Name, "cus_") {
			complianceResponse = append(complianceResponse, evaluateCustomRule(parsed, rule, request, ruleDefinition))
		} else {
			ruleNamesForRegistry = append(ruleNamesForRegistry, ruleDefinition.Name)
		}
	}
	if len(ruleNamesForRegistry) > 0 {
		registry, _ := lint.GlobalRegistry().Filter(lint.FilterOptions{
			IncludeNames: ruleNamesForRegistry,
		})

		zlintResultSet := zlint.LintCertificateEx(parsed, registry)
		for d, s := range zlintResultSet.Results {
			complianceResponse = append(complianceResponse, ResponseRules{UUID: rules.GetRuleUuidFromName(d), Name: d, Status: getStatus(s.Status)})
		}
	}
	return Response{computeOverallStatus(complianceResponse), complianceResponse}, nil
}

func normalizeCertificate(certificate string) string {
	replacer := strings.NewReplacer("-----BEGIN CERTIFICATE-----", "", "-----END CERTIFICATE-----", "", "\n", "", "\r", "")
	return replacer.Replace(certificate)
}

func getStatus(status lint.LintStatus) rules.RuleStatus {
	switch status {
	case lint.NA, lint.NE:
		return rules.NOT_APPLICABLE
	case lint.Pass:
		return rules.COMPLIANT
	default:
		return rules.NON_COMPLIANT
	}
}

func computeOverallStatus(responseRules []ResponseRules) (status Status) {
	isNotApplicable := true
	isNonCompliant := false

	for _, s := range responseRules {
		if s.Status == rules.NON_COMPLIANT {
			isNonCompliant = true
			isNotApplicable = false
			break
		} else if s.Status == rules.COMPLIANT {
			isNotApplicable = false
		} else {
			//Nothing happens
		}
	}

	if isNotApplicable {
		return NOT_CHECKED
	} else if isNonCompliant {
		return NON_COMPLIANT
	} else {
		return COMPLIANT
	}
}

func evaluateCustomRule(certificate *x509.Certificate, requestRule RequestRules, request Request, ruleDefinition rules.RuleDefinition) (response ResponseRules) {
	defer func() {
		if p := recover(); p != nil {
			response = frameErrorValidation(ruleDefinition.UUID, ruleDefinition.Name, rules.NON_COMPLIANT)
			fmt.Println(p)
		}
	}()
	var compliant ResponseRules
	switch ruleDefinition.Name {
	case "cus_hashing_algorithm":
		compliant = HashingAlgorithmValidation(certificate, requestRule, request, ruleDefinition)
	case "cus_public_key_algorithm":
		compliant = PublicKeyAlgorithmValidation(certificate, requestRule, request, ruleDefinition)
	case "cus_elliptic_curve":
		compliant = EcCurveValidation(certificate, requestRule, request, ruleDefinition)
	case "cus_key_length":
		compliant = KeySizeValidator(certificate, requestRule, request, ruleDefinition)
	}
	return compliant
}

func frameErrorValidation(uuid string, name string, status rules.RuleStatus) (response ResponseRules) {
	return ReturnRuleFramer(uuid, name, status)
}
