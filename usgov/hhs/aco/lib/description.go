package lib

import (
  "os"
  "encoding/csv"
  "bitbucket.org/gocodo/bloomsource"
  "bitbucket.org/gocodo/bloomsource/helpers"
)

type Description struct {}

var sourceUris = map[string]string {
  "usgov.hhs.aco_contacts": "https://s3.amazonaws.com/gocodo/usgov/hhs/aco/aco_contacts.csv",
  "usgov.hhs.aco_puf": "https://s3.amazonaws.com/gocodo/usgov/hhs/aco/aco_puf.csv",
  "usgov.hhs.aco_year1": "https://s3.amazonaws.com/gocodo/usgov/hhs/aco/aco_year1.csv",
  "usgov.hhs.aco": "https://s3.amazonaws.com/gocodo/usgov/hhs/aco/aco_names.csv",
  "usgov.hhs.aco_benchmarks": "https://s3.amazonaws.com/gocodo/usgov/hhs/aco/aco_benchmarks.csv",
  "usgov.hhs.aco_measures": "https://s3.amazonaws.com/gocodo/usgov/hhs/aco/aco_measures.csv",
  "usgov.hhs.aco_benchmark_years": "https://s3.amazonaws.com/gocodo/usgov/hhs/aco/aco_benchmark_years.csv",
  "usgov.hhs.aco_benchmark_measures": "https://s3.amazonaws.com/gocodo/usgov/hhs/aco/aco_benchmark_measures.csv",
}

func (d *Description) Available() ([]bloomsource.Source, error) {
  return []bloomsource.Source{
    bloomsource.Source{
      Name: "usgov.hhs.aco_year1",
      Version: "2015-00",
    },
    bloomsource.Source{
      Name: "usgov.hhs.aco_puf",
      Version: "2015-00",
    },
    bloomsource.Source{
      Name: "usgov.hhs.aco_contacts",
      Version: "2015-00",
    },
    bloomsource.Source{
      Name: "usgov.hhs.aco",
      Version: "2015-00",
    },
    bloomsource.Source{
      Name: "usgov.hhs.aco_benchmarks",
      Version: "2015-00",
    },
    bloomsource.Source{
      Name: "usgov.hhs.aco_measures",
      Version: "2015-00",
    },
    bloomsource.Source{
      Name: "usgov.hhs.aco_benchmark_years",
      Version: "2015-00",
    },
    bloomsource.Source{
      Name: "usgov.hhs.aco_benchmark_measures",
      Version: "2015-00",
    },
  }, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  uri := sourceUris[sourceName]

  downloader := bloomsource.NewDownloader("data/", nil)
  path, err := downloader.Fetch(uri)
  if err != nil {
    return nil, err
  }

  reader, err := os.Open(path)
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

func (d *Description) Reader(source bloomsource.Source) (bloomsource.ValueReader, error) {
  uri := sourceUris[source.Name]

  downloader := bloomsource.NewDownloader("data/", nil)
  path, err := downloader.Fetch(uri)
  if err != nil {
    return nil, err
  }

  reader, err := os.Open(path)
  if err != nil {
    return nil, err
  }

  zipReader := helpers.NewCsvReader(reader)

  return zipReader, nil
}