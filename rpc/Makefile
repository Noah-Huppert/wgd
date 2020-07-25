.PHONY: certs-clean

# Protocol buffers source files.
PB_SRC_NAME ?= rpc
PB_SRC_FILE ?= ${PB_SRC_NAME}.proto

# Generated Go output files.
PB_GO_OUT_DIR ?= .
PB_GO_OUT ?= ${PB_GO_OUT_DIR}/${PB_SRC_NAME}.pb.go

# Directories in which generated protocol buffers files will be placed.
PB_SERVER_DIR ?= ../server/rpc
PB_CLIENT_DIR ?= ../client/rpc

# Generate GRPC Go bindings.
generate: ${PB_GO_OUT}
${PB_GO_OUT}: ${PB_SRC_FILE}
	# Generate files
	protoc --go_out=plugins=grpc:${PB_GO_OUT_DIR} ${PB_SRC_FILE}

	# Copy generated files into server and client directories
	if [ ! -d ${PB_SERVER_DIR} ]; then mkdir ${PB_SERVER_DIR}; fi
	if [ ! -d ${PB_CLIENT_DIR} ]; then mkdir ${PB_CLIENT_DIR}; fi
	cp ${PB_GO_OUT} ${PB_SERVER_DIR}/${PB_SRC_NAME}.pb.go
	cp ${PB_GO_OUT} ${PB_CLIENT_DIR}/${PB_SRC_NAME}.pb.go


# Certificates generated using commands from: https://itnext.io/practical-guide-to-securing-grpc-connections-with-go-and-tls-part-1-f63058e9d6d1

# Prefix applied to all certificate files
CERTS_PREFIX ?= dev-

CERTS_COUNTRY = 00
CERTS_STATE = 00
CERTS_ORG = 00

# Root certificate files.
CERTS_ROOT_CA_KEY ?= ${CERTS_PREFIX}ca.key
CERTS_ROOT_CA_CERT ?= ${CERTS_PREFIX}ca.cert

CERTS_SERVER_CONFIG ?= certs-server.conf

CERTS_SERVER_KEY ?= ${CERTS_PREFIX}server.key
CERTS_SERVER_CSR ?= ${CERTS_PREFIX}server.csr
CERTS_SERVER_CERT ?= ${CERTS_PREFIX}server.pem

# Generate authentication certificates.
certs: ${CERTS_ROOT_CA_KEY} ${CERTS_ROOT_CA_CERT} ${CERTS_SERVER_KEY} ${CERTS_SERVER_CSR} ${CERTS_SERVER_CERT}

# Delete all certificate files.
certs-clean:
	rm *.srl || true
	rm ${CERTS_ROOT_CA_KEY} || true
	rm ${CERTS_ROOT_CA_CERT} || true
	rm ${CERTS_SERVER_KEY} || true
	rm ${CERTS_SERVER_CSR} || true
	rm ${CERTS_SERVER_CERT} || true

# Generate the root CA signing key.
${CERTS_ROOT_CA_KEY}:
	openssl genrsa -out ${CERTS_ROOT_CA_KEY} 4096

# Generate the root CA certificate.
${CERTS_ROOT_CA_CERT}:
	openssl req \
		-new \
		-x509 \
		-key ${CERTS_ROOT_CA_KEY} \
		-sha256 \
		-subj "/C=${CERTS_COUNTRY}/ST=${CERTS_STATE}/O=${CERTS_ORG}" \
		-days 365 \
		-out ${CERTS_ROOT_CA_CERT}

# Generate server's key.
${CERTS_SERVER_KEY}:
	openssl genrsa -out ${CERTS_SERVER_KEY} 4096

# Generate server's certificate signing request
${CERTS_SERVER_CSR}:
	openssl req \
		-new \
		-key ${CERTS_SERVER_KEY} \
		-out ${CERTS_SERVER_CSR} \
		-config ${CERTS_SERVER_CONFIG}

# Generate server's certificate
${CERTS_SERVER_CERT}:
	openssl x509 \
		-req -in ${CERTS_SERVER_CSR} \
		-CA ${CERTS_ROOT_CA_CERT} \
		-CAkey ${CERTS_ROOT_CA_KEY} \
		-CAcreateserial \
		-out ${CERTS_SERVER_CERT} \
		-days 365 \
		-sha256 \
		-extfile ${CERTS_SERVER_CONFIG} \
		-extensions req_ext
