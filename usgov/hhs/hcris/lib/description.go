package lib

import (
  "io"
  "regexp"
  "encoding/csv"
  "github.com/bloomapi/dataloading"
  "github.com/bloomapi/dataloading/helpers"
)

type Description struct {}

func (d *Description) Available() ([]dataloading.Source, error) {
  return []dataloading.Source{
    dataloading.Source{
      Name: "usgov.hhs.hcris_ime_gme2013",
      Version: "2013-00",
    },
  }, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  fileMatch := regexp.MustCompile(`IME_GME2013.CSV$`)
  reader, err := getFileReader("http://downloads.cms.gov/files/hcris/hosp10-reports.zip", fileMatch)
  if err != nil {
    return nil, err
  }

  csvReader := csv.NewReader(reader)
  if err != nil {
    return nil, err
  }

  columns, err := csvReader.Read()
  if err != nil {
    return nil, err
  }

  return columns, nil
}

func getFileReader(uri string, zipPattern *regexp.Regexp) (io.Reader, error) {
  downloader := dataloading.NewDownloader("data/", nil)
  path, err := downloader.Fetch(uri)
  if err != nil {
    return nil, err
  }

  reader, err := helpers.OpenExtractZipReader(path, zipPattern)
  if err != nil {
    return nil, err
  }

  return reader, nil
}

func (d *Description) Reader(source dataloading.Source) (dataloading.ValueReader, error) {
  fileMatch := regexp.MustCompile(`IME_GME2013.CSV$`)
  reader, err := getFileReader("http://downloads.cms.gov/files/hcris/hosp10-reports.zip", fileMatch)
  if err != nil {
    return nil, err
  }

  zipReader := helpers.NewCsvReader(reader)

  return zipReader, nil
}