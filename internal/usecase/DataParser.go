package usecase

import (
	"encoding/csv"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/ujunglangit-id/redpanda-kafka-bench/internal/model"
	"os"
)

type DataParser struct {
	cfg *model.Config
}

func New(cfg *model.Config) *DataParser {
	return &DataParser{
		cfg: cfg,
	}
}

func (d *DataParser) ParseMessage(record *kgo.Record, idx int) {
	var msg model.ProductData
	err := json.Unmarshal(record.Value, &msg)
	if err != nil {
		log.Errorf("[parsemsg] error unmarshal incoming msg , err %#v", err)
		return
	}

	csvFile, err := os.Create("employee.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvFile.Close()

}

func (d *DataParser) writeToCsv(msg model.ProductData) {
	//check if file is exist

	//append to file

}

func addcol(fname string, column []string) error {
	// read the file
	f, err := os.Open(fname)
	if err != nil {
		return err
	}
	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	if err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}

	// add column
	l := len(lines)
	if len(column) < l {
		l = len(column)
	}
	for i := 0; i < l; i++ {
		lines[i] = append(lines[i], column[i])
	}

	// write the file
	f, err = os.Create(fname)
	if err != nil {
		return err
	}
	w := csv.NewWriter(f)
	if err = w.WriteAll(lines); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}
