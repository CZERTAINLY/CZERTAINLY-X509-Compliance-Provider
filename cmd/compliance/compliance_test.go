package compliance

import (
	"CZERTAINLY-X509-Compliance-Provider/cmd/attributes"
	"CZERTAINLY-X509-Compliance-Provider/cmd/rules"
	"testing"

	"go.uber.org/zap"
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
	if len(response.Rules) == 0 {
		t.Error("Validation Failed")
	}
}

func TestService_ComplianceCheckNa(t *testing.T) {
	certificate := "MIIDfjCCAmagAwIBAgIUFanTqP1qmACQDMYlbsvJhyvjvIYwDQYJKoZIhvcNAQELBQAwPjELMAkGA1UEBhMCQ1oxGzAZBgNVBAgMEkhsYXZuaSBtZXN0byBQcmFoYTESMBAGA1UEBwwJSGx1Ym9jZXB5MB4XDTIyMDYyMjE0NTE1OFoXDTIzMDYyMjE0NTE1OFowPjELMAkGA1UEBhMCQ1oxGzAZBgNVBAgMEkhsYXZuaSBtZXN0byBQcmFoYTESMBAGA1UEBwwJSGx1Ym9jZXB5MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyTZ8rfZAzjtgfat1jq6s37Xtx8lITNta4fqAhRlhsfg+bpP3nd7K8kB7U4AV37FZPjM/wbIJOdJ+dR5nkqPCPDJ7q3e9nYX1D5t2rwloCNH1lyDpA5osBV+ohgsR27dOEeKqfN+u16Ev1S7PS+h8MXIOPjuSvjh7/lVZp6jDZwj0MjwJVbTALlkuO9vJj0FnOzmpekWFz/o+/dfnJiZuto97hl0H1O4uBs19uduqpW3T3HQVMNjRMiftGyqlkpE0MSnbPT59xyKpzx/KN1zV57c/QEOOZfOArlRhocKQF7dPqTAS/AN69mv3QAQs6CDlcp6GktHocHvdXsrLXGQ2kQIDAQABo3QwcjAdBgNVHQ4EFgQUaQNqw0xnagL/rud9VntFrRvN2cUwHwYDVR0jBBgwFoAUaQNqw0xnagL/rud9VntFrRvN2cUwDgYDVR0PAQH/BAQDAgWgMCAGA1UdJQEB/wQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjANBgkqhkiG9w0BAQsFAAOCAQEANuxV/KL25wCTDqzCj3GOwheRv/iZmFq0p29i4W5Q8au19BUtCI9FNeB+9mAS95YV7U/QhigfjfYNV/B4B/SbXs4ttkeHI7QRDR0TCsxJHhrYmXolkyIHVjo0bmH87ekccSZjaTUvMcryqd1vIzBXmHpDQeYgQ78XJPPA3liulPLmmftb4lSYrZ3kP2E743O2GRaPnkE9K2fjf5vk8trzkeb2mn1qP/tywlcFud+bMdXyp9OI38WD/FxgE1NYrz83RXJH6J5kGZ+Am9tGEZHiAwtBvAa0n8YmkL3h4HI74YIKPtOeOuU4YX17tOqzWru2QNCxZFCE0uC7guX/SwtWjQ=="
	rules := []RequestRules{{
		UUID: "40f084cd-ddc1-11ec-82b0-34cff65c6ee3",
	}}
	request := Request{
		Certificate: certificate,
		Rules:       rules,
	}

	response, _ := tService.ComplianceCheck("x509", request)
	if response.Status != NA {
		t.Error("Validation Failed")
	}
}

func TestService_ComplianceCheckNok(t *testing.T) {
	certificate := "MIIDKDCCAhCgAwIBAgIUUodCaUBN/2ZPTFLjlC+nnDNS7gMwDQYJKoZIhvcNAQELBQAwEzERMA8GA1UEAwwIdGVzdHg1MDkwHhcNMjIwNjMwMTAyMjAyWhcNMjUwNjI5MTAyMjAyWjATMREwDwYDVQQDDAh0ZXN0eDUwOTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKCGkhhZjpFZVvuvKaJ/w1dNeRNNZCmpW2BzaYtsrEpSVyRVRUo4fWf0M2xwiirR663jUpoyh/wGEg0mzs/HG4CWCyBEJuvtCAG0PmyF+QrVn0QQFu1V+0zG3OvaF/nXpxnsuTzKMuwf/DKGKMSLZVlJjeemcNLYW7lWciYXI3IoVasBBYkVJCiWevwW4I9guvMUx+ani+kyGMzpGVtDmJaBFJL/TR4Z752+UFgdZpjOq9FRAct2G0gS+9zDXFaC9g4xAYWIVtXiGACW9koTA7qT/mcFzAeJEKHTv381Ak/dnLM5qhlEbJU73OpVwSOjAJQdDaYr4+GToQ1bnYg8kHUCAwEAAaN0MHIwHQYDVR0OBBYEFENDlZGYNla0JMDX5frcTl3zgyQAMB8GA1UdIwQYMBaAFENDlZGYNla0JMDX5frcTl3zgyQAMA4GA1UdDwEB/wQEAwIFoDAgBgNVHSUBAf8EFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwDQYJKoZIhvcNAQELBQADggEBACmJqJFATkP/Nj+YJ3+/EDagpr+Fu8KYqvh0O5wcooyEl6iFOPdbnAr0SxUjECifJ/Lt+R3HOB8c5UQPcI0dFD0P5kFe37ZGKW/R28F97omtAMf3dxb4r9uINHsho1a0MXEAKfIT16/c+M7hZgU1h1TIv6Qg6O1OyQZd8aj1ZmiGxA/ULFbQ6BsNSLyc6M8nY9BAT8XNgzbgvJVbcmsLOSMV++mozplTAQSFfePsV943p6HfrxhQmOJyzP8IEmZoatzCllG265AJXtIhgjzD15v+4r+EJsGGPZd9VA/z+XNf4lAt6TqJVG1uU4wU1w0KCqvft1D+cGTaX1D80UUsy0M="
	rules := []RequestRules{{
		UUID: "40f0ac56-ddc1-11ec-9825-34cff65c6ee3",
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
	attrs := []attributes.RequestAttributes{{
		Name:    "condition",
		UUID:    "7ed00782-e706-11ec-8fea-0242ac120002",
		Content: attributes.BaseAttributeContent{Value: "Equals"}},
		{
			Name:    "length",
			UUID:    "7ed00886-e706-11ec-8fea-0242ac120002",
			Content: attributes.BaseAttributeContent{Value: 2048},
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
