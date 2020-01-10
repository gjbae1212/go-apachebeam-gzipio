package gzipio

import (
	"bufio"
	"compress/gzip"
	"context"
	"fmt"
	"strings"

	"github.com/apache/beam/sdks/go/pkg/beam"
	"github.com/apache/beam/sdks/go/pkg/beam/io/filesystem"
)

// Read is reading from glob
func Read(s beam.Scope, glob string) beam.PCollection {
	s = s.Scope("gzipio.Read")
	// check if path is valid.
	filesystem.ValidateScheme(glob)

	in := beam.Create(s, glob)
	files := beam.ParDo(s, extractFilesFn, in)
	return beam.ParDo(s, readFilesFn, files)
}

func extractFilesFn(ctx context.Context, glob string, emit func(string)) error {
	if strings.TrimSpace(glob) == "" {
		return fmt.Errorf("[err] empty glob")
	}

	fs, err := filesystem.New(ctx, glob)
	if err != nil {
		return err
	}
	defer fs.Close()

	files, err := fs.List(ctx, glob)
	if err != nil {
		return err
	}
	for _, filename := range files {
		emit(filename)
	}
	return nil
}

func readFilesFn(ctx context.Context, filename string, emit func(string)) error {
	fs, err := filesystem.New(ctx, filename)
	if err != nil {
		return err
	}
	defer fs.Close()

	fd, err := fs.OpenRead(ctx, filename)
	if err != nil {
		return err
	}
	defer fd.Close()

	gs, err := gzip.NewReader(fd)
	if err != nil {
		return err
	}
	defer gs.Close()

	scanner := bufio.NewScanner(gs)
	for scanner.Scan() {
		emit(scanner.Text())
	}
	return scanner.Err()
}

type writeFileFn struct {
	Filename string `json:"filename"`
}

// Write writes data to filename with gzip compress
func Write(s beam.Scope, filename string, col beam.PCollection) {
	s = s.Scope("gzipio.Write")

	// check if path is valid.
	filesystem.ValidateScheme(filename)

	pre := beam.AddFixedKey(s, col)
	post := beam.GroupByKey(s, pre)
	beam.ParDo0(s, &writeFileFn{Filename: filename}, post)
}

// ProcessElement is a boilerplate method using process data.
func (w *writeFileFn) ProcessElement(ctx context.Context, key int, values func(*string) bool) error {
	fs, err := filesystem.New(ctx, w.Filename)
	if err != nil {
		return err
	}
	defer fs.Close()

	fd, err := fs.OpenWrite(ctx, w.Filename)
	if err != nil {
		return err
	}
	defer fd.Close()

	gz := gzip.NewWriter(fd)
	defer gz.Close()

	buf := bufio.NewWriterSize(gz, 1<<20) // use 1MB buffer

	var line string
	for values(&line) {
		if _, err := buf.WriteString(line); err != nil {
			return err
		}
		if _, err := buf.Write([]byte{'\n'}); err != nil {
			return err
		}
	}

	if err := buf.Flush(); err != nil {
		return err
	}
	return nil
}
