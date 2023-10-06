package schemaregistry

type Fetcher interface {
	FetchSchema(*SubjectVersion) (*SchemaResponse, error)
	FetchLatestVersion(string) (int, error)
	FetchSchemaByID(int) (*SchemaResponse, error)
	FetchVersions(string) ([]int, error)
}
