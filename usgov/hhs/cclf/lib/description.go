package lib

import (
  "path/filepath"
  "os"
  "encoding/csv"
  "regexp"
  "github.com/bloomapi/dataloading"
  "github.com/bloomapi/dataloading/helpers"
  "github.com/bloomapi/datasources/usgov/hhs/cclf/schema"
 	"github.com/spf13/viper"
 	"errors"
)

type Description struct {}

var cclfFileName = regexp.MustCompile(`ACO\.CCLF(\d)\.(D\d\d\d\d\d\d\.T\d\d\d\d\d\d\d)`)

var sourceNames = map[string]string {
	"1": "usgov.hhs.cclf_parta",
	"2": "usgov.hhs.cclf_parta_revenue",
	"3": "usgov.hhs.cclf_parta_proc",
	"4": "usgov.hhs.cclf_parta_dx",
	"5": "usgov.hhs.cclf_partb_phy",
	"6": "usgov.hhs.cclf_partb_dme",
	"7": "usgov.hhs.cclf_partd",
	"8": "usgov.hhs.cclf_beneficiary",
	"9": "usgov.hhs.cclf_bene_xref",
}

// Inverse of sourceNames
var fileIndexes = map[string]string {
	"usgov.hhs.cclf_parta": "1",
	"usgov.hhs.cclf_parta_revenue": "2",
	"usgov.hhs.cclf_parta_proc": "3",
	"usgov.hhs.cclf_parta_dx": "4",
	"usgov.hhs.cclf_partb_phy": "5",
	"usgov.hhs.cclf_partb_dme": "6",
	"usgov.hhs.cclf_partd": "7",
	"usgov.hhs.cclf_beneficiary": "8",
	"usgov.hhs.cclf_bene_xref": "9",
}

var sourceFields = map[string]*schema.Cclf {
	"usgov.hhs.cclf_parta": schema.Parta,
	"usgov.hhs.cclf_parta_revenue": schema.PartaRevenue,
	"usgov.hhs.cclf_parta_proc": schema.PartaProc,
	"usgov.hhs.cclf_parta_dx": schema.PartaDx,
	"usgov.hhs.cclf_partb_phy": schema.PartbPhy,
	"usgov.hhs.cclf_partb_dme": schema.PartbDme,
	"usgov.hhs.cclf_partd": schema.Partd,
	"usgov.hhs.cclf_beneficiary": schema.Beneficiary,
	"usgov.hhs.cclf_bene_xref": schema.BeneXref,
}

func (d *Description) Available() ([]dataloading.Source, error) {
	sources := []dataloading.Source{}

	cclfDir := viper.GetString("cclfDir")

	err := filepath.Walk(cclfDir, func (path string, file os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !file.IsDir() {
				matches := cclfFileName.FindStringSubmatch(file.Name())
				if len(matches) > 0 {
					fileNum := matches[1]
					fileVersion := matches[2]
					sourceName := sourceNames[fileNum]

					if sourceName != "" {
						sources = append(sources, dataloading.Source{
							Name: sourceName,
							Version: fileVersion,
							Action: "upsert",
						})
					}
				}
			}

			return nil
		})

	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(cclfDir + "/uid_to_hicn.csv"); err == nil {
		sources = append(sources, dataloading.Source{
				Name: "usgov.hhs.cclf.uid_to_hicn",
				Version: "2015-01",
				Action: "sync",
			})
	}

	if _, err := os.Stat(cclfDir + "/uid_beneficiaries.csv"); err == nil {
		sources = append(sources, dataloading.Source{
				Name: "usgov.hhs.cclf.uid_beneficiaries",
				Version: "2015-01",
				Action: "sync",
			})
	}

	return sources, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
	cclfDir := viper.GetString("cclfDir")
	var path string
	
	switch sourceName {
	case "usgov.hhs.cclf.uid_to_hicn":
		path = cclfDir + "/uid_to_hicn.csv"
	case "usgov.hhs.cclf.uid_beneficiaries":
		path = cclfDir + "/uid_beneficiaries.csv"
	default:
		fields := sourceFields[sourceName].FieldNames()
		names := make([]string, len(fields))

		for i, field := range fields {
			names[i] = field
		}

		return names, nil
	}

  fileReader, err := os.Open(path)
  if err != nil {
    return nil, err
  }

  csvReader := csv.NewReader(fileReader)
  if err != nil {
    return nil, err
  }

  columns, err := csvReader.Read()
  if err != nil {
    return nil, err
  }

  return columns, nil
}

func (d *Description) Reader(source dataloading.Source) (dataloading.ValueReader, error) {
	cclfDir := viper.GetString("cclfDir")
	var path string

	switch source.Name {
	case "usgov.hhs.cclf.uid_to_hicn":
		path = cclfDir + "/uid_to_hicn.csv"
	case "usgov.hhs.cclf.uid_beneficiaries":
		path = cclfDir + "/uid_beneficiaries.csv"
	default:
		fileIndex := fileIndexes[source.Name]
	  fileRegex := regexp.MustCompile(`ACO\.CCLF` + fileIndex + `\.` + source.Version)

		var sourcePath string

		err := filepath.Walk(cclfDir, func (path string, file os.FileInfo, err error) error {
				if fileRegex.MatchString(file.Name()) {
					sourcePath = path
				}

				return nil
			})

		if err != nil {
			return nil, err
		}

		if sourcePath == "" {
			return nil, errors.New("Source " + source.Name + " with version " + source.Version + " not found.")
		}

		file, err := os.Open(sourcePath)
		if err != nil {
			return nil, err
		}

		return helpers.NewTabReader(file, sourceFields[source.Name].Fields(source.Version)), nil
	}

  fileReader, err := os.Open(path)
  if err != nil {
    return nil, err
  }

  csvReader := helpers.NewCsvReader(fileReader)

  return csvReader, nil
}