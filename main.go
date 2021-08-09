package main

import (
	"encoding/json"
	"io"
	"learning-gcs/gcs"
	"os"
)

type City struct {
	Name     string `json:"name"`
	Province string `json:"province"`
}

func main() {

	os.Setenv("BUCKET", "example")
	os.Setenv("FILE_OBJECT", "city.json")

	// create data from struct
	c1 := City{"Bandung", "Jawa Barat"}

	// create object from struct
	gc := gcs.NewGCS(os.Getenv("BUCKET"), os.Getenv("FILE_OBJECT"))

	err := c1.WriteToJson(gc.UploadFile())
	if err != nil {
		panic(err)
	}
	// list object from bucket
	gc.ListFile()

	// rename object from bucket
	gc.RenameFile()

}

func (c City) WriteToJson(w io.WriteCloser) error {

	j, err := json.Marshal(c)
	if err != nil {
		return err
	}

	_, err = w.Write(j)
	if err != nil {
		return err
	}

	return w.Close()

}
