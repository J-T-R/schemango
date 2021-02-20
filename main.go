package main

func main() {
	runner := Runner{}

	// Here we could add a default config getter
	schemas := make(map[string][]byte)

	// Could also we a default config to load
	subscriptions := make(map[string]Address)

	// Port should be from config
	// Although should it just be a full address from config?
	runner.Port = 9000
	runner.Schemas = schemas
	runner.Subscriptions = subscriptions

	runner.runAPI()
}
