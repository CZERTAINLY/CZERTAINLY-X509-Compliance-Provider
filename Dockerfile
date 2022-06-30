#GO lang buld the docket container

#Using Go version 1.18.2
FROM golang:1.18.2
WORKDIR /app

#Copy the modules and the checksum of the modules
COPY go.mod ./
COPY go.sum ./

#Download the go modules
RUN go mod download
COPY *.go ./

#Run build
RUN go build -o /CZERTAINLY-X509-Compliance-Provider

##
## Deploy
##

WORKDIR /app
COPY --from=build /CZERTAINLY-X509-Compliance-Provider /CZERTAINLY-X509-Compliance-Provider
EXPOSE 8080
ENTRYPOINT ["/docker-gs-ping"]