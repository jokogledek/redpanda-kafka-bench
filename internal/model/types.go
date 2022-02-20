package model

type ProductData struct {
	ID          string `json:"id"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	UserName    string `json:"user_name"`
	UserID      string `json:"user_id"`
}

type Config struct {
	PartitionCount int      `json:"partition_count"`
	ConsumerCount  int      `json:"consumer_count"`
	TopicName      string   `json:"topic_name"`
	ConsumerGroup  string   `json:"consumer_group"`
	InputPath      string   `json:"input_path"`
	OutputPath     string   `json:"output_path"`
	InputFile      string   `json:"input_file"`
	Host           []string `json:"host"`
}
