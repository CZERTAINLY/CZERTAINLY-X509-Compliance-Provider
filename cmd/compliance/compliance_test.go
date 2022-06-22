package compliance

import (
	"CZERTAINLY-X509-Compliance-Provider/cmd/attributes"
	"CZERTAINLY-X509-Compliance-Provider/cmd/rules"
	"go.uber.org/zap"
	"testing"
)

var logger *zap.Logger
var tService Service
var rService rules.Service

func init() {
	logger, _ = zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	tService = NewService(sugar)
	rService = rules.NewService(sugar)
}

func TestService_ComplianceCheck(t *testing.T) {
	certificate := "MIIDfjCCAmagAwIBAgIUFanTqP1qmACQDMYlbsvJhyvjvIYwDQYJKoZIhvcNAQELBQAwPjELMAkGA1UEBhMCQ1oxGzAZBgNVBAgMEkhsYXZuaSBtZXN0byBQcmFoYTESMBAGA1UEBwwJSGx1Ym9jZXB5MB4XDTIyMDYyMjE0NTE1OFoXDTIzMDYyMjE0NTE1OFowPjELMAkGA1UEBhMCQ1oxGzAZBgNVBAgMEkhsYXZuaSBtZXN0byBQcmFoYTESMBAGA1UEBwwJSGx1Ym9jZXB5MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyTZ8rfZAzjtgfat1jq6s37Xtx8lITNta4fqAhRlhsfg+bpP3nd7K8kB7U4AV37FZPjM/wbIJOdJ+dR5nkqPCPDJ7q3e9nYX1D5t2rwloCNH1lyDpA5osBV+ohgsR27dOEeKqfN+u16Ev1S7PS+h8MXIOPjuSvjh7/lVZp6jDZwj0MjwJVbTALlkuO9vJj0FnOzmpekWFz/o+/dfnJiZuto97hl0H1O4uBs19uduqpW3T3HQVMNjRMiftGyqlkpE0MSnbPT59xyKpzx/KN1zV57c/QEOOZfOArlRhocKQF7dPqTAS/AN69mv3QAQs6CDlcp6GktHocHvdXsrLXGQ2kQIDAQABo3QwcjAdBgNVHQ4EFgQUaQNqw0xnagL/rud9VntFrRvN2cUwHwYDVR0jBBgwFoAUaQNqw0xnagL/rud9VntFrRvN2cUwDgYDVR0PAQH/BAQDAgWgMCAGA1UdJQEB/wQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjANBgkqhkiG9w0BAQsFAAOCAQEANuxV/KL25wCTDqzCj3GOwheRv/iZmFq0p29i4W5Q8au19BUtCI9FNeB+9mAS95YV7U/QhigfjfYNV/B4B/SbXs4ttkeHI7QRDR0TCsxJHhrYmXolkyIHVjo0bmH87ekccSZjaTUvMcryqd1vIzBXmHpDQeYgQ78XJPPA3liulPLmmftb4lSYrZ3kP2E743O2GRaPnkE9K2fjf5vk8trzkeb2mn1qP/tywlcFud+bMdXyp9OI38WD/FxgE1NYrz83RXJH6J5kGZ+Am9tGEZHiAwtBvAa0n8YmkL3h4HI74YIKPtOeOuU4YX17tOqzWru2QNCxZFCE0uC7guX/SwtWjQ=="
	rules := []RequestRules{{
		UUID: "40f084cd-ddc1-11ec-82b0-34cff65c6ee3",
	}}
	request := Request{
		Certificate: certificate,
		Rules:       rules,
	}

	response, _ := tService.ComplianceCheck("x509", request)
	if len(response.Rules) < 0 {
		t.Error("Validation Failed")
	}
}

func TestService_ComplianceCheckNok(t *testing.T) {
	certificate := "MIIDfjCCAmagAwIBAgIUFanTqP1qmACQDMYlbsvJhyvjvIYwDQYJKoZIhvcNAQELBQAwPjELMAkGA1UEBhMCQ1oxGzAZBgNVBAgMEkhsYXZuaSBtZXN0byBQcmFoYTESMBAGA1UEBwwJSGx1Ym9jZXB5MB4XDTIyMDYyMjE0NTE1OFoXDTIzMDYyMjE0NTE1OFowPjELMAkGA1UEBhMCQ1oxGzAZBgNVBAgMEkhsYXZuaSBtZXN0byBQcmFoYTESMBAGA1UEBwwJSGx1Ym9jZXB5MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyTZ8rfZAzjtgfat1jq6s37Xtx8lITNta4fqAhRlhsfg+bpP3nd7K8kB7U4AV37FZPjM/wbIJOdJ+dR5nkqPCPDJ7q3e9nYX1D5t2rwloCNH1lyDpA5osBV+ohgsR27dOEeKqfN+u16Ev1S7PS+h8MXIOPjuSvjh7/lVZp6jDZwj0MjwJVbTALlkuO9vJj0FnOzmpekWFz/o+/dfnJiZuto97hl0H1O4uBs19uduqpW3T3HQVMNjRMiftGyqlkpE0MSnbPT59xyKpzx/KN1zV57c/QEOOZfOArlRhocKQF7dPqTAS/AN69mv3QAQs6CDlcp6GktHocHvdXsrLXGQ2kQIDAQABo3QwcjAdBgNVHQ4EFgQUaQNqw0xnagL/rud9VntFrRvN2cUwHwYDVR0jBBgwFoAUaQNqw0xnagL/rud9VntFrRvN2cUwDgYDVR0PAQH/BAQDAgWgMCAGA1UdJQEB/wQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjANBgkqhkiG9w0BAQsFAAOCAQEANuxV/KL25wCTDqzCj3GOwheRv/iZmFq0p29i4W5Q8au19BUtCI9FNeB+9mAS95YV7U/QhigfjfYNV/B4B/SbXs4ttkeHI7QRDR0TCsxJHhrYmXolkyIHVjo0bmH87ekccSZjaTUvMcryqd1vIzBXmHpDQeYgQ78XJPPA3liulPLmmftb4lSYrZ3kP2E743O2GRaPnkE9K2fjf5vk8trzkeb2mn1qP/tywlcFud+bMdXyp9OI38WD/FxgE1NYrz83RXJH6J5kGZ+Am9tGEZHiAwtBvAa0n8YmkL3h4HI74YIKPtOeOuU4YX17tOqzWru2QNCxZFCE0uC7guX/SwtWjQ=="
	rules := []RequestRules{{
		UUID: "40f084cd-ddc1-11ec-82b0-34cff65c6ee3",
	}}
	request := Request{
		Certificate: certificate,
		Rules:       rules,
	}

	response, _ := tService.ComplianceCheck("x509", request)
	if response.Status != NOK {
		t.Error("Validation Failed")
	}
}

func TestService_ComplianceCheckAttributes(t *testing.T) {
	certificate := "MIIDfjCCAmagAwIBAgIUFanTqP1qmACQDMYlbsvJhyvjvIYwDQYJKoZIhvcNAQELBQAwPjELMAkGA1UEBhMCQ1oxGzAZBgNVBAgMEkhsYXZuaSBtZXN0byBQcmFoYTESMBAGA1UEBwwJSGx1Ym9jZXB5MB4XDTIyMDYyMjE0NTE1OFoXDTIzMDYyMjE0NTE1OFowPjELMAkGA1UEBhMCQ1oxGzAZBgNVBAgMEkhsYXZuaSBtZXN0byBQcmFoYTESMBAGA1UEBwwJSGx1Ym9jZXB5MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyTZ8rfZAzjtgfat1jq6s37Xtx8lITNta4fqAhRlhsfg+bpP3nd7K8kB7U4AV37FZPjM/wbIJOdJ+dR5nkqPCPDJ7q3e9nYX1D5t2rwloCNH1lyDpA5osBV+ohgsR27dOEeKqfN+u16Ev1S7PS+h8MXIOPjuSvjh7/lVZp6jDZwj0MjwJVbTALlkuO9vJj0FnOzmpekWFz/o+/dfnJiZuto97hl0H1O4uBs19uduqpW3T3HQVMNjRMiftGyqlkpE0MSnbPT59xyKpzx/KN1zV57c/QEOOZfOArlRhocKQF7dPqTAS/AN69mv3QAQs6CDlcp6GktHocHvdXsrLXGQ2kQIDAQABo3QwcjAdBgNVHQ4EFgQUaQNqw0xnagL/rud9VntFrRvN2cUwHwYDVR0jBBgwFoAUaQNqw0xnagL/rud9VntFrRvN2cUwDgYDVR0PAQH/BAQDAgWgMCAGA1UdJQEB/wQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjANBgkqhkiG9w0BAQsFAAOCAQEANuxV/KL25wCTDqzCj3GOwheRv/iZmFq0p29i4W5Q8au19BUtCI9FNeB+9mAS95YV7U/QhigfjfYNV/B4B/SbXs4ttkeHI7QRDR0TCsxJHhrYmXolkyIHVjo0bmH87ekccSZjaTUvMcryqd1vIzBXmHpDQeYgQ78XJPPA3liulPLmmftb4lSYrZ3kP2E743O2GRaPnkE9K2fjf5vk8trzkeb2mn1qP/tywlcFud+bMdXyp9OI38WD/FxgE1NYrz83RXJH6J5kGZ+Am9tGEZHiAwtBvAa0n8YmkL3h4HI74YIKPtOeOuU4YX17tOqzWru2QNCxZFCE0uC7guX/SwtWjQ=="
	attrs := []attributes.RequestAttributes{attributes.RequestAttributes{
		Name:  "condition",
		UUID:  "7ed00782-e706-11ec-8fea-0242ac120002",
		Value: attributes.BaseAttributeContent{Value: "Equals"}},
		attributes.RequestAttributes{
			Name:  "length",
			UUID:  "7ed00886-e706-11ec-8fea-0242ac120002",
			Value: attributes.BaseAttributeContent{Value: 2048},
		}}
	rules := []RequestRules{{
		UUID:       "7ed00480-e706-11ec-8fea-0242ac120002",
		Attributes: attrs,
	}}
	request := Request{
		Certificate: certificate,
		Rules:       rules,
	}

	tService.ComplianceCheck("x509", request)
}
