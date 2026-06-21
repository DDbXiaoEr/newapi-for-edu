#!/bin/bash
set -e

DOMAIN="${DOMAIN:-localhost}"
CERT_DIR="/etc/nginx/certs"
CERT_FILE="${CERT_DIR}/server.crt"
KEY_FILE="${CERT_DIR}/server.key"

if [ -f "${CERT_FILE}" ] && [ -f "${KEY_FILE}" ]; then
    echo "Certificates already exist, skipping generation."
    exit 0
fi

mkdir -p "${CERT_DIR}"

openssl req -x509 -nodes -days 3650 -newkey rsa:2048 \
    -keyout "${KEY_FILE}" \
    -out "${CERT_FILE}" \
    -subj "/CN=${DOMAIN}/O=newapiedu/C=CN" \
    -addext "subjectAltName=DNS:${DOMAIN},DNS:*.${DOMAIN},IP:127.0.0.1"

chmod 644 "${CERT_FILE}"
chmod 600 "${KEY_FILE}"

echo "Self-signed certificate generated for domain: ${DOMAIN}"
