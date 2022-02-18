package model

type ProductData struct {
	ID          int
	ProductName string
	Description string
	Category    string
	UserName    string
	UserID      int
}

type Config struct {
	PartitionCount int      `json:"partition_count"`
	ConsumerCount  int      `json:"consumer_count"`
	TopicName      string   `json:"topic_name"`
	InputPath      string   `json:"input_path"`
	OutputPath     string   `json:"output_path"`
	InputFile      string   `json:"input_file"`
	Host           []string `json:"host"`
}
