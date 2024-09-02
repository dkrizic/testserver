package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// name,ratings,price,imgURL,camera,display,battery,storage,ram,processor,android_version
type Entry struct {
	Name           string
	Ratings        string
	Price          string
	ImgURL         string
	Camera         string
	Display        string
	Battery        string
	Storage        string
	RAM            string
	Processor      string
	AndroidVersion string
}

func main() {
	// load a CSV into memory
	f, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}
	fmt.Printf("-- +goose Up\n")
	entries := []Entry{}
	for i, record := range records {
		if i == 0 {
			for j, field := range record {
				fmt.Printf("insert into tag (id,name) values (%d, '%s');\n", j+1, field)
			}
			continue
		}
		entry := &Entry{
			Name:           record[0],
			Ratings:        record[1],
			Price:          record[2],
			ImgURL:         record[3],
			Camera:         record[4],
			Display:        record[5],
			Battery:        record[6],
			Storage:        record[7],
			RAM:            record[8],
			Processor:      record[9],
			AndroidVersion: record[10],
		}
		entries = append(entries, *entry)

		fmt.Printf("insert into asset (id,name) values (%d, '%s');\n", i, entry.Name)
		fmt.Printf("insert into tagvalue (id, tag_id, asset_id, value) values (%d, %d, %d, '%s');\n", (i-1)*20+1, 1, i, entry.Name)
		fmt.Printf("insert into tagvalue (id, tag_id, asset_id, value) values (%d, %d, %d, '%s');\n", (i-1)*20+2, 1, i, entry.Ratings)
		fmt.Printf("insert into tagvalue (id, tag_id, asset_id, value) values (%d, %d, %d, '%s');\n", (i-1)*20+3, 2, i, entry.Price)
		fmt.Printf("insert into tagvalue (id, tag_id, asset_id, value) values (%d, %d, %d, '%s');\n", (i-1)*20+4, 3, i, entry.ImgURL)
		fmt.Printf("insert into tagvalue (id, tag_id, asset_id, value) values (%d, %d, %d, '%s');\n", (i-1)*20+5, 4, i, entry.Camera)
		fmt.Printf("insert into tagvalue (id, tag_id, asset_id, value) values (%d, %d, %d, '%s');\n", (i-1)*20+6, 5, i, entry.Display)
		fmt.Printf("insert into tagvalue (id, tag_id, asset_id, value) values (%d, %d, %d, '%s');\n", (i-1)*20+7, 6, i, entry.Battery)
		fmt.Printf("insert into tagvalue (id, tag_id, asset_id, value) values (%d, %d, %d, '%s');\n", (i-1)*20+8, 7, i, entry.Storage)
		fmt.Printf("insert into tagvalue (id, tag_id, asset_id, value) values (%d, %d, %d, '%s');\n", (i-1)*20+9, 8, i, entry.RAM)
		fmt.Printf("insert into tagvalue (id, tag_id, asset_id, value) values (%d, %d, %d, '%s');\n", (i-1)*20+10, 9, i, entry.Processor)
		fmt.Printf("insert into tagvalue (id, tag_id, asset_id, value) values (%d, %d, %d, '%s');\n", (i-1)*20+11, 10, i, entry.AndroidVersion)
	}

}
