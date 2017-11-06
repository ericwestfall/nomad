package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	cafile  = "../helper/tlsutil/testdata/ca.pem"
	foocert = "../helper/tlsutil/testdata/nomad-foo.pem"
	fookey  = "../helper/tlsutil/testdata/nomad-foo-key.pem"
)

func TestTLSConfig_Equals_False_NilExisting(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	const (
		cafile  = "../helper/tlsutil/testdata/ca.pem"
		foocert = "../helper/tlsutil/testdata/nomad-foo.pem"
		fookey  = "../helper/tlsutil/testdata/nomad-foo-key.pem"
	)

	existingConfig := &TLSConfig{}

	newConfig := &TLSConfig{
		EnableHTTP:           true,
		EnableRPC:            true,
		VerifyServerHostname: true,
		CAFile:               cafile,
		CertFile:             foocert,
		KeyFile:              fookey,
	}

	assert.False(existingConfig.Equals(newConfig))
}

func TestTLSConfig_Equals_False_NilNew(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	const (
		cafile  = "../helper/tlsutil/testdata/ca.pem"
		foocert = "../helper/tlsutil/testdata/nomad-foo.pem"
		fookey  = "../helper/tlsutil/testdata/nomad-foo-key.pem"
	)

	existingConfig := &TLSConfig{
		EnableHTTP:           true,
		EnableRPC:            true,
		VerifyServerHostname: true,
		CAFile:               cafile,
		CertFile:             foocert,
		KeyFile:              fookey,
	}

	assert.False(existingConfig.Equals(nil))
}

func TestTLSConfig_Equals_False(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	const (
		cafile  = "../helper/tlsutil/testdata/ca.pem"
		foocert = "../helper/tlsutil/testdata/nomad-foo.pem"
		fookey  = "../helper/tlsutil/testdata/nomad-foo-key.pem"
	)

	newConfig := &TLSConfig{
		EnableHTTP:           true,
		EnableRPC:            true,
		VerifyServerHostname: true,
		CAFile:               "bar.pem",
		CertFile:             "nomad-bar.pem",
		KeyFile:              "nomad-bar-key.pem",
	}

	existingConfig := &TLSConfig{
		EnableHTTP:           true,
		EnableRPC:            true,
		VerifyServerHostname: true,
		CAFile:               cafile,
		CertFile:             foocert,
		KeyFile:              fookey,
	}

	assert.False(existingConfig.Equals(newConfig))
}

func TestTLSConfig_Equals_True_WhenEqual(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	newConfig := &TLSConfig{
		EnableHTTP:           true,
		EnableRPC:            true,
		VerifyServerHostname: true,
		CAFile:               cafile,
		CertFile:             foocert,
		KeyFile:              fookey,
	}

	existingConfig := &TLSConfig{
		EnableHTTP:           true,
		EnableRPC:            true,
		VerifyServerHostname: true,
		CAFile:               cafile,
		CertFile:             foocert,
		KeyFile:              fookey,
	}

	assert.True(existingConfig.Equals(newConfig))
}

func TestTLSConfig_IsEmpty(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	empty := &TLSConfig{}
	assert.True(empty.IsEmpty())

	notEmpty := &TLSConfig{
		EnableHTTP:           true,
		EnableRPC:            true,
		VerifyServerHostname: true,
		CAFile:               cafile,
		CertFile:             foocert,
		KeyFile:              fookey,
	}
	assert.False(notEmpty.IsEmpty())
}
