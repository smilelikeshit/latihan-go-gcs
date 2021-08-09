package gcs

import (
	"context"
	"fmt"
	"io"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type gcs struct {
	clientgcs  *storage.Client
	bucketname string
	filename   string
}

func NewGCS(newbucketname, newfile string) *gcs {

	ctx := context.Background()

	// Creates a client.
	newclient, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return &gcs{
		clientgcs:  newclient,
		bucketname: newbucketname,
		filename:   newfile,
	}
}

func (g *gcs) UploadFile() io.WriteCloser {

	ctx := context.Background()
	return g.clientgcs.Bucket(g.bucketname).Object(g.filename).NewWriter(ctx)

}

func (g *gcs) ListFile() {

	it := g.clientgcs.Bucket(g.bucketname).Objects(context.Background(), nil)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(attrs.Name)
	}

}

func (g *gcs) RoleFilePublic(filename string) {

	acl := g.clientgcs.Bucket(g.bucketname).Object(filename).ACL()
	if err := acl.Set(context.Background(), storage.AllUsers, storage.RoleReader); err != nil {
		fmt.Println(err)

	}

	fmt.Printf("bucketname %s , filename %s has successfully to public ACL", g.bucketname, filename)

}

func (g *gcs) RenameFile() string {

	ctx := context.Background()

	dstName := g.filename + "-rename"
	src := g.clientgcs.Bucket(g.bucketname).Object(g.filename)
	dst := g.clientgcs.Bucket(g.bucketname).Object(dstName)

	if _, err := dst.CopierFrom(src).Run(ctx); err != nil {
		return dstName
	}
	if err := src.Delete(ctx); err != nil {
		return dstName
	}
	fmt.Printf("Blob %v moved to %v.\n", g.filename, dstName)

	return dstName

}
