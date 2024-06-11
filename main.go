package main;

import (
	"context"
	"fmt"
  "time"

	client "github.com/celestiaorg/celestia-openrpc"
	"github.com/celestiaorg/celestia-openrpc/types/blob"
	"github.com/celestiaorg/celestia-openrpc/types/share"
)


func main() {
  err := SubmitBlob(context.Background(), "ws://localhost:26658", "")
  if err != nil {
    panic(err)
  }
}

// SubmitBlob submits a blob containing "Hello, World!" to the 0xDEADBEEF namespace. It uses the default signer on the running node.
func SubmitBlob(ctx context.Context, url string, token string) error {
	client, err := client.NewClient(ctx, url, token)
	if err != nil {
		return err
	}

	// let's post to 0xDEADBEEF namespace
	namespace, err := share.NewBlobNamespaceV0([]byte{0xDE, 0xAD, 0xBE, 0xEF})
	if err != nil {
		return err
	}

	// create a blob
	helloWorldBlob, err := blob.NewBlobV0(namespace, []byte("Hello, World!"))
	if err != nil {
		return err
	}

  for i := range(10) {
    // submit the blob to the network
    go func() {
      height, err := client.Blob.Submit(ctx, []*blob.Blob{helloWorldBlob}, blob.DefaultGasPrice())
      if err != nil {
        fmt.Printf("Error on blob %d submit: %v\n", i, err)
      }

      fmt.Printf("Blob %d was included at height %d\n", i, height)
    }()
  }

  time.Sleep(time.Minute * 5)
	return nil
}
