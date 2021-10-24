package deliveries_noti

type Destination struct {
	Name           string `json:"name"`
	Type           string `json:"type"`
	Identification string `jsob:"identification"`
}

type Source struct {
	Name          string `json:"name"`
	ServerAddress string `json:"server_address"`
	DatabaseName  string `jsob:"database_name"`
	TableName     string `jsob:"table_name"`
	DataKey       string `jsob:"data_key"`
	DataType      string `jsob:"data_type"`
	DataInterval  string `jsob:"data_interval"`
}

type Expression struct {
	SourceID      string `json:"source_id"`
	Condition     string `json:"condition"`
	DestinationID string `json:"destination_id"`
}
