package esstoml

import "net/http"

// EssTomlMaxSize is the maximum size of ess.toml file
const EssTomlMaxSize = 5 * 1024

// WellKnownPath represents the url path at which the ess.toml file should
// exist to conform to the federation protocol.
const WellKnownPath = "/.well-known/ess.toml"

// DefaultClient is a default client using the default parameters
var DefaultClient = &Client{HTTP: http.DefaultClient}

// Client represents a client that is capable of resolving a ess.toml file
// using the internet.
type Client struct {
	// HTTP is the http client used when resolving a ess.toml file
	HTTP HTTP

	// UseHTTP forces the client to resolve against servers using plain HTTP.
	// Useful for debugging.
	UseHTTP bool
}

type ClientInterface interface {
	GetEssToml(domain string) (*Response, error)
	GetEssTomlByAddress(addy string) (*Response, error)
}

// HTTP represents the http client that a stellertoml resolver uses to make http
// requests.
type HTTP interface {
	Get(url string) (*http.Response, error)
}

// Response represents the results of successfully resolving a ess.toml file
type Response struct {
	AuthServer       string `toml:"AUTH_SERVER"`
	FederationServer string `toml:"FEDERATION_SERVER"`
	EncryptionKey    string `toml:"ENCRYPTION_KEY"`
	SigningKey       string `toml:"SIGNING_KEY"`
}

// GetEssToml returns ess.toml file for a given domain
func GetEssToml(domain string) (*Response, error) {
	return DefaultClient.GetEssToml(domain)
}

// GetEssTomlByAddress returns ess.toml file of a domain fetched from a
// given address
func GetEssTomlByAddress(addy string) (*Response, error) {
	return DefaultClient.GetEssTomlByAddress(addy)
}
