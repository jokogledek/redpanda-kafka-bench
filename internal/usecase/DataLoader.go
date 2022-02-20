package usecase

import (
	"encoding/csv"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/ujunglangit-id/redpanda-kafka-bench/internal/model"
	"io"
	"os"
)

func (d *DataParser) LoadCSV() (err error) {
	var (
		f *os.File
	)
	// open file
	f, err = os.Open(d.cfg.InputPath + d.cfg.InputFile)
	if err != nil {
		log.Errorf("[dataloader] error open csv file at %s%s, %#v", d.cfg.InputPath, d.cfg.InputFile, err)
		return
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	for {
		var outData model.ProductData
		rec := []string{}
		rec, err = csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Errorf("[dataloader] error read csv file, %#v", err)
			return
		}
		// do something with read line
		log.Infof("data : %v", rec)
		outData, err = d.ExtractField(rec)
		if err != nil {
			log.Errorf("[dataloader] error read csv file, %#v", err)
			continue
		}
		d.repo.PublishMsg(outData)
	}
	return
}

func (d *DataParser) ExtractField(data []string) (out model.ProductData, err error) {
	if len(data) == 10 {
		out = model.ProductData{
			ID:          data[0],
			ProductName: data[7],
			Description: data[1],
			Category:    data[8],
			UserName:    data[2],
			UserID:      data[3],
		}
	} else {
		err = errors.New("invalid length")
	}
	return
}
