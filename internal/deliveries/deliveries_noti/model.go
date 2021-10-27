package deliveries_noti

type Destination struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

// Where data from
type Source struct {
	Name          string `json:"name"`
	ServerAddress string `json:"server_address"`
	DatabaseName  string `json:"database_name"`
	TableName     string `json:"table_name"`
	DataKey       string `json:"data_key"`
}

type Expression struct {
	Name                string `json:"name"`
	SourceID            string `json:"source_id"`
	Condition           string `json:"condition"`
	DestinationID       string `json:"destination_id"`
	Interval            string `json:"interval"`
	MessageInCondition  string `json:"message_in_condition"`
	MessageOutCondition string `json:"message_out_condition"`
}
