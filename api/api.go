package api

type APIConfig struct {
	REST    *RESTConfig    `json:"rest,omitempty"`
	GraphQL *GraphQLConfig `json:"graphql,omitempty"`
	Guard   *GuardConfig   `json:"guard,omitempty"`
}

type RESTConfig struct {
	Description string                     `json:"description,omitempty"`
	Path        string                     `json:"path"`
	Aliases     map[string]RESTAliasConfig `json:"aliases"`
}

type RESTAliasConfig struct {
	Description string                 `json:"description,omitempty"`
	Action      string                 `json:"action"`
	Params      map[string]interface{} `json:"params,omitempty"`
}

type GraphQLConfig struct {
	Description string                            `json:"description,omitempty"`
	TypeDefs    string                            `json:"typeDefs"`

	// TODO: can suggest more detailed types
	Resolvers   map[string]map[string]interface{} `json:"resolvers"`
}

type GuardConfig map[string]string
