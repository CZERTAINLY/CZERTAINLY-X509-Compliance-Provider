#!/bin/sh

czertainlyHome="/opt/czertainly"
source ${czertainlyHome}/static-functions

log "INFO" "Launching X.509 Compliance Provider"
./appbin

#exec "$@"