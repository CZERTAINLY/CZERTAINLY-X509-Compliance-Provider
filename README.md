# CZERTAINLY X509 Compliance Provider

> This repository is part of the commercial open-source project CZERTAINLY, but the connector is available under subscription. You can find more information about the project at [CZERTAINLY](https://github.com/3KeyCompany/CZERTAINLY) repository, including the contribution guide.

X509 Compliance `Connector` is the implementation of the following `Function Groups` and `Kinds`:

| Function Group | Kind |
| --- | --- |
| `Compliance Provider` | `x509` |

X509 Compliance Provider is the implementation of compliance check and management for x509 certificates that are managed by CZERTAINLY. This `Connector` performs compliance check for the certificates of type x509. List of items to be considered to determine the compliance of a certificate is administrated by the rules and groups.

### Rules

Rules in the X509 Compliance Provider describes the condition that should be applied to the certificate for determination of the Compliance Status. Rules are the individual objects that contribute to the overall compliance status of the certificate

### Groups

Groups are the logical grouping of the rules organized by some baseline similarities. A group may contain two or more rules and when a compliance profile is added with the group, all the rules in the group will be applied to the certificate compliance determination


X509 Compliance Provider allows you to perform the following operations:
- Check compliance of x509 Certificate

## Database requirements

This `Connector` does not require any database as it does not store any information

## Short Process Description

Compliance of the certificate is calculated by the use of `Compliance Profiles`. `Compliance Profiles` are the entities that holds the list of rules and groups to be considered for the compliance calculation. These objects holds rules and groups from different `Connectors`, apply them accordingly and compute the status based on the rule validations.

X509 Compliance Provider consumes ZLint for some rules and groups. To know more about ZLint, refer to [ZLint](https://github.com/zmap/zlint)


To know more about the `Core`, refer to [CZERTAINLY Core](https://github.com/3KeyCompany/CZERTAINLY-Core)


## Interfaces

X509 Compliance `Connector` implements `Compliance Provider` interfaces. To learn more about the interfaces and end points, refer to the [CZERTAINLY Interfaces](https://github.com/3KeyCompany/CZERTAINLY-Interfaces).

For more information, please refer to the [CZERTAINLY documentation](https://docs.czertainly.com).

## Docker container

X509 Compliance `Connector` is provided as a Docker container. Use the `3keycompany/czertainly-x509-compliance-provider:tagname` to pull the required image from the repository. It can be configured using the following environment variables:

| Variable | Description | Required | Default value |
| --- | --- | --- | --- |
| `SERVER_PORT` | Port where the service is exposed | No | 8080 |
| `LOG_LEVEL` | Logging level for the service | No | INFO |