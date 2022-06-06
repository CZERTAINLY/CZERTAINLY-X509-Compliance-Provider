package compliance

import (
	"CZERTAINLY-X509-Compliance-Provider/cmd/rules"
	"CZERTAINLY-X509-Compliance-Provider/cmd/utils"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3"
	"github.com/zmap/zlint/v3/lint"
	"strings"
)

type service struct {
	logger log.Logger
}

func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

func (s service) ComplianceCheck(ctx context.Context, kind string, request Request) (Response, error) {

	var ruleNamesForRegistry []string
	var complianceResponse []ResponseRules
	certificate := normalizeCertificate(request.Certificate)
	block, err := base64.StdEncoding.DecodeString(certificate)
	parsed, err := x509.ParseCertificate(block)
	fmt.Println(err)

	for _, rule := range request.Rules {
		ruleDefinition := rules.GetRuleFromUuid(rule.UUID)
		if ruleDefinition.Custom {
			evaluateCustomRule(parsed, rule, request, ruleDefinition)
			continue
		}
		ruleNamesForRegistry = append(ruleNamesForRegistry, ruleDefinition.Name)
	}

	registry, err := lint.GlobalRegistry().Filter(lint.FilterOptions{
		IncludeNames: ruleNamesForRegistry,
	})

	zlintResultSet := zlint.LintCertificateEx(parsed, registry)
	for d, s := range zlintResultSet.Results {
		complianceResponse = append(complianceResponse, ResponseRules{UUID: rules.GetRuleUuidFromName(d), Name: d, Status: getStatus(s.Status)})
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

func evaluateCustomRule(certificate *x509.Certificate, requestRule RequestRules, request Request, ruleDefinition rules.RuleDefinition) {
	if ruleDefinition.Name == "cus_hashing_algorithm_greater_than" {
		hashingAlgorithmValidation(certificate, requestRule, request, ruleDefinition)
	}
}

func hashingAlgorithmValidation(certificate *x509.Certificate, requestRule RequestRules, request Request, ruleDefinition rules.RuleDefinition) {
	condition := utils.GetRequestAttributeValue("condition", requestRule.Attributes)
	value := utils.GetRequestAttributeValue("algorithm", requestRule.Attributes)

	fmt.Println(certificate.SignatureAlgorithmName())
	fmt.Println(condition)
	fmt.Println(value)
}
