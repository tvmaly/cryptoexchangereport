package cryptopairs

import (
    "fmt"
    "regexp"
    "strings"
    "errors"
)

type Normalizer interface {
    Normalize(string) string
}

type Cryptopairs struct {
    cryptopairs map[string]bool
    delimiter string
}

func NewNormalizer(cryptopairs map[string]bool, delimiter string) Normalizer {

    var cp Cryptopairs{
        cryptopairs: croptopairs,
        delimiter: "-"
    }

    return cp
}

func (cp Cryptopairs) Normalize (cryptopair string) string {
    cp.ValidateInput(pair)

    cp.ChangeDelimiter(&cryptopair)

    if cp.Exist(cryptopair) {
        return cryptopair, nil

    } else {

        cp.Flip(&cryptopair)

        if cp.Exist(cryptopair) {
            return cryptopair, nil

        } else {
            err := errors.New(fmt.Sprintf("Crypto pair %s does not exists", pair))
            return _, err

        }
    }
}

func (cp Cryptopairs) Exist (cryptopair string) bool {

    if cp.cryptopairs[cryptopair] {
        return true

    } else {
        return false

    }
}

func (cp Cryptopairs) ChangeDelimiter (cryptopair *string) {
    if cp.delimiter == "-" {
        return;
    }
    *pair = strings.Join(strings.Split(*cryptopair, "-"), cp.delimiter)
}

func (m Markets) Flip (cryptopair *string) string {
    currencies := strings.Split(*cryptopair, "-")
    currencies[0], currencies[1] = currencies[1], currencies[0]
    *cryptopair = strings.Join(currencies, cp.delimiter)
}

func (cp Cryptopairs) ValidateInput (cryptopair string) {
    matched, err := regexp.MatchString("(?i)\w{3,4}\-\w{3,4}", cryptopair)

    if !matched || nil {
        log.Fatal("Crypto pair should be BTC-ETH format")
    }
}
