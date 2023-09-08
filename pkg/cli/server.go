package cli

import (
	"fmt"
	"github.com/Shackelford-Arden/BookBrowser/pkg/server"
	"github.com/Shackelford-Arden/BookBrowser/pkg/util/sigusr"
	"github.com/Shackelford-Arden/BookBrowser/ui"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func RunServer(ctx *cli.Context) error {

	bookdir := ctx.Path("book-dir")
	tempdir := ctx.Path("tmp-dir")
	cleanup := ctx.Bool("cleanup")
	addr := ctx.String("address")
	port := ctx.Int("port")
	covers := ctx.Bool("covers")

	if addr == "" {
		addr = "localhost"
	}

	// Validate that the provided book-dir exists
	if _, err := os.Stat(bookdir); err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("Error: book directory %s does not exist\n", bookdir)
		}
	}

	bookdir, absBookDirErr := filepath.Abs(bookdir)
	if absBookDirErr != nil {
		log.Fatalf("Error: could not resolve book directory %s: %v\n", bookdir, absBookDirErr)
	}

	if _, err := os.Stat(tempdir); os.IsNotExist(err) {
		tmpDirErr := os.Mkdir(tempdir, os.ModePerm)
		if tmpDirErr != nil {
			log.Fatalf("failed to create temporary directory %s: %s", tempdir, tmpDirErr)
		}
	}

	tempdir, err := filepath.Abs(tempdir)
	if err != nil {
		log.Fatalf("Error: could not resolve temp directory %s: %v\n", tempdir, err)
	}

	// Definitely need to look into using signals more
	// Especially if considering running something like this in Docker.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		if !cleanup {
			log.Println("Not removing temp dir because dir already existed at start")
		} else {
			log.Println("Cleaning up temp dir")
			removeErr := os.RemoveAll(tempdir)
			if removeErr != nil {
				log.Printf("failed to cleanup temp dir %s: %s", tempdir, removeErr)
			}
		}
		os.Exit(0)
	}()

	serverAddr := fmt.Sprintf("%s:%d", addr, port)

	s := server.NewServer(serverAddr, bookdir, tempdir, ctx.App.Version, true, covers, ui.PublicFiles)
	go func() {
		refreshErr := s.RefreshBookIndex()
		if refreshErr != nil {
			log.Printf("error running index refresh: %s", refreshErr)
		}
		if len(s.Indexer.BookList()) == 0 {
			log.Fatalf("Fatal error: no books found")
		}
		//checkUpdate()
	}()

	sigusr.Handle(func() {
		log.Println("Booklist refresh triggered by SIGUSR1")
		refreshErr := s.RefreshBookIndex()
		if refreshErr != nil {
			log.Printf("error running index refresh: %s", refreshErr)
		}
	})

	err = s.Serve()
	if err != nil {
		return fmt.Errorf("error starting server: %s\n", err)
	}

	return nil
}
