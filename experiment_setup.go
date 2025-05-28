package pilot

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// ======================================
//        Experiment Setup Struct
// ======================================

type ExperimentSetup struct {
	Name        string
	BaseDir     string            // e.g. "./data/mnist"
	MainURL     string            // e.g. "https://storage.googleapis.com/cvdf-datasets/mnist/"
	Files       map[string]string // key: compressed, value: uncompressed
	DownloadLog []string
}

// ======================================
//        Public Setup Method
// ======================================

func (e *ExperimentSetup) Init() error {
	if err := os.MkdirAll(e.BaseDir, os.ModePerm); err != nil {
		return err
	}

	for compressed, uncompressed := range e.Files {
		cPath := filepath.Join(e.BaseDir, compressed)
		uPath := filepath.Join(e.BaseDir, uncompressed)

		if _, err := os.Stat(uPath); os.IsNotExist(err) {
			if _, err := os.Stat(cPath); os.IsNotExist(err) {
				fmt.Printf("📥 Downloading %s...\n", compressed)
				if err := downloadFile(e.MainURL+compressed, cPath); err != nil {
					return err
				}
				e.DownloadLog = append(e.DownloadLog, compressed)
			}
			fmt.Printf("📦 Unzipping %s...\n", compressed)
			if err := unzipFile(cPath, uPath); err != nil {
				return err
			}
		}
	}

	return nil
}

// ======================================
//         Internal Helpers
// ======================================

func downloadFile(url, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func unzipFile(src, dest string) error {
	fSrc, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fSrc.Close()

	gzReader, err := gzip.NewReader(fSrc)
	if err != nil {
		return err
	}
	defer gzReader.Close()

	fDest, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer fDest.Close()

	_, err = io.Copy(fDest, gzReader)
	return err
}
