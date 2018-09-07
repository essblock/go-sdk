package config

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountIDValidator(t *testing.T) {
	var val struct {
		Empty        string `valid:"ess_accountid"`
		NotSTRKey    string `valid:"ess_accountid"`
		NotAccountID string `valid:"ess_accountid"`
		Valid        string `valid:"ess_accountid"`
		WrongType    int    `valid:"ess_accountid"`
	}

	val.NotSTRKey = "hello"
	val.NotAccountID = "SA5MATAU4RNJDKCTIC6VVSYSGB7MFFBVU3OKWOA5K67S62EYB5ESKLTV"
	val.Valid = "GBXS6WTZNRS7LOGHM3SCMAJD6M6JCXB3GATXECCZ3C5NJ3PVSZ23PEWX"
	val.WrongType = 100

	// run the validation
	ok, err := govalidator.ValidateStruct(val)
	require.False(t, ok)
	require.Error(t, err)

	fields := govalidator.ErrorsByField(err)

	// ensure valid is not in the invalid map
	_, ok = fields["Valid"]
	assert.False(t, ok)

	_, ok = fields["Empty"]
	assert.True(t, ok, "Empty is not an invalid field")

	_, ok = fields["NotSTRKey"]
	assert.True(t, ok, "NotSTRKey is not an invalid field")

	_, ok = fields["NotAccountID"]
	assert.True(t, ok, "NotAccountID is not an invalid field")

	_, ok = fields["WrongType"]
	assert.True(t, ok, "WrongType is not an invalid field")
}

func TestSeedValidator(t *testing.T) {
	var val struct {
		Empty     string `valid:"ess_seed"`
		NotSTRKey string `valid:"ess_seed"`
		NotSeed   string `valid:"ess_seed"`
		Valid     string `valid:"ess_seed"`
		WrongType int    `valid:"ess_seed"`
	}

	val.NotSTRKey = "hello"
	val.NotSeed = "GBXS6WTZNRS7LOGHM3SCMAJD6M6JCXB3GATXECCZ3C5NJ3PVSZ23PEWX"
	val.Valid = "SA5MATAU4RNJDKCTIC6VVSYSGB7MFFBVU3OKWOA5K67S62EYB5ESKLTV"
	val.WrongType = 100

	// run the validation
	ok, err := govalidator.ValidateStruct(val)
	require.False(t, ok)
	require.Error(t, err)

	fields := govalidator.ErrorsByField(err)

	// ensure valid is not in the invalid map
	_, ok = fields["Valid"]
	assert.False(t, ok)

	_, ok = fields["Empty"]
	assert.True(t, ok, "Empty is not an invalid field")

	_, ok = fields["NotSTRKey"]
	assert.True(t, ok, "NotSTRKey is not an invalid field")

	_, ok = fields["NotSeed"]
	assert.True(t, ok, "NotSeed is not an invalid field")

	_, ok = fields["WrongType"]
	assert.True(t, ok, "WrongType is not an invalid field")
}

func TestUndecoded(t *testing.T) {
	var val struct {
		Test string `toml:"test" valid:"optional"`
		TLS  struct {
			CertificateFile string `toml:"certificate-file" valid:"required"`
			PrivateKeyFile  string `toml:"private-key-file" valid:"required"`
		} `valid:"optional"`
	}

	// Notice _ in certificate_file
	toml := `test="abc"
[tls]
certificate_file="hello"
private-key-file="world"`

	err := decode(toml, &val)
	require.Error(t, err)
	assert.Equal(t, "Unknown fields: [tls.certificate_file]", err.Error())
}

func TestCorrect(t *testing.T) {
	var val struct {
		Test string `toml:"test" valid:"optional"`
		TLS  struct {
			CertificateFile string `toml:"certificate-file" valid:"required"`
			PrivateKeyFile  string `toml:"private-key-file" valid:"required"`
		} `valid:"optional"`
	}

	// Notice _ in certificate_file
	toml := `test="abc"
[tls]
certificate-file="hello"
private-key-file="world"`

	err := decode(toml, &val)
	require.NoError(t, err)
}
