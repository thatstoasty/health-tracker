package load

import (
	"encoding/csv"
	"os"
	"log"

    "github.com/thatstoasty/health-tracker/shared/queries"
)

func GetRecordsFromFile(path string) [][]string {
    // open file
    f, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }

    // remember to close the file at the end of the program
    defer f.Close()

    // read csv values using csv.Reader
    csvReader := csv.NewReader(f)
    data, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

	return data
}

func CreateCompositionList(data [][]string) []queries.SubmitCompositionParams {
    var compositionList []queries.SubmitCompositionParams
    for i, line := range data {
        if i > 0 { // omit header line
            var rec queries.SubmitCompositionParams
            for j, field := range line {
                if j == 0 {
                    rec.SubmittedOn = field
                } else if j == 1 {
                    rec.Weight = field
                } else if j == 2 {
                    rec.Bodyfat = field
                }
            }
            compositionList = append(compositionList, rec)
        }
    }
    return compositionList
}
