package constants

import (
    "encoding/json"
    "io/ioutil"
    "errors"
    "log"
)

var errorLookup map[int]error 

func init() {

    type ErrorJson struct {
        Code int `json:"code"`
        Error string `json:"error"`
    }

    var errs []ErrorJson

    raw, err := ioutil.ReadFile("./data/errors.json")
    if err != nil {
        log.Fatal(err.Error())
    }

    json.Unmarshal(raw, &errs)

    for _, v := range errs {

        err := errors.New(v.Error)

        errorLookup[v.Code] = err
    }
}

func Error (code int) error {
    if errorLookup[code] != nil {
        return errorLookup[code]
    }

    return nil

}
