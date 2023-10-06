package schemaregistry

type SchemaResponse struct {
	Subject    string       `json:"subject"`
	ID         int          `json:"id,omitempty"`
	Version    int          `json:"version"`
	Schema     string       `json:"schema"`
	SchemaType string       `json:"schemaType"`
	References []*Reference `json:"references"`
}

type SchemaRequest struct {
	Schema     string       `json:"schema"`
	SchemaType string       `json:"schemaType,omitempty"`
	References []*Reference `json:"references,omitempty"`
}

type Reference struct {
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Version int    `json:"version"`
}

type SubjectVersion struct {
	Subject string `json:"subject"`
	Version int    `json:"version"`
}

type idOnlyJSON struct {
	ID int `json:"id"`
}

type isCompatibleJSON struct {
	IsCompatible bool `json:"is_compatible"`
}
