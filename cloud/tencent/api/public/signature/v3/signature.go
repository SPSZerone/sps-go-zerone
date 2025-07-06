package v3

import (
	"encoding/hex"
	"fmt"

	spshmac "github.com/SPSZerone/sps-go-zerone/algorithm/crypto/hmac"
	spsauth "github.com/SPSZerone/sps-go-zerone/net/http/authorization"
)

// Signature API
//
// see [DOC]
//
// [DOC]: https://cloud.tencent.com/document/api/1312/48202
func Signature(config *Config, canonicalRequest *CanonicalRequest) string {
	// # build base
	timestamp, date, credentialScope := config.Build()

	// # build canonical request string
	hashedCanonicalRequest := canonicalRequest.Build()

	// # build string to sign
	string2sign := fmt.Sprintf("%s\n%d\n%s\n%s", config.Algorithm, timestamp, credentialScope, hashedCanonicalRequest)

	// # sign string
	secretDate := spshmac.Sha256Str(date, fmt.Sprintf("TC3%s", config.SecretKey))
	secretService := spshmac.Sha256Str(config.Service, secretDate)
	secretSigning := spshmac.Sha256Str("tc3_request", secretService)
	signature := hex.EncodeToString(spshmac.Sha256Bytes(string2sign, secretSigning))

	// # build authorization
	credentials := fmt.Sprintf(
		"Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		config.SecretId, credentialScope,
		canonicalRequest.SignedHeaders,
		signature,
	)
	return spsauth.Format(config.Algorithm, credentials)
}
