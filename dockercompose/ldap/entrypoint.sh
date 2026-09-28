#!/bin/bash
set -e

CUSTOM_LDIF_DIR="/container/service/slapd/assets/config/bootstrap/ldif/custom"
BOOTSTRAP_FILE="/bootstrap/init-data.ldif"

if [ -f "$BOOTSTRAP_FILE" ]; then
  mkdir -p "$CUSTOM_LDIF_DIR"
  cp "$BOOTSTRAP_FILE" "$CUSTOM_LDIF_DIR/50-bootstrap.ldif"
fi

exec /container/tool/run "$@"
