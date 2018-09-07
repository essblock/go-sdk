package esstoml

import "log"

// ExampleGetTOML gets the ess.toml file for coins.asia
func ExampleClient_GetEssToml() {
	_, err := DefaultClient.GetEssToml("coins.asia")
	if err != nil {
		log.Fatal(err)
	}
}
