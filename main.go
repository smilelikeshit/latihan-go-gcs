package main

import (
	"encoding/json"
	"io"
	"learning-gcs/gcs"
	"os"
	"time"
)

type City struct {
	Name     string `json:"name"`
	Province string `json:"province"`
}

func main() {

	now := time.Now()

	// create file by date format
	timeformat, _ := time.Parse(time.RFC822, "02 Jan 06 15:04 MST")

	os.Setenv("BUCKET", "example")
	os.Setenv("FILE_OBJECT", now.Format(timeformat.String()))

	// dummy create data from struct implement io.Writer
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
	z := gc.RenameFile()

	// set acl to public
	gc.RoleFilePublic(z)

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
