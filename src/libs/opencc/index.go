package opencc

import (
	_ "statik"
	"github.com/rakyll/statik/fs"
	"github.com/stevenyao/go-opencc"
	"io"
	"os"
)

func initDepends(fn string) error {
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		sfs, err := fs.New()
		if err != nil {
			return err
		}
		fin, err := sfs.Open("/opencc/" + fn)
		if err != nil {
			return err
		}
		defer fin.Close()

		fout, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
		defer fout.Close()

		_, err = io.Copy(fout, fin)
		if err != nil {
			return err
		}

	}
    return nil
}

func ToTraditional(s string) (string, error) {
    config := `s2t.json`
    for _, fn := range []string{config, `libopencc.dll`, `STPhrases.ocd`, `STCharacters.ocd`} {
        err := initDepends(fn)
        if err != nil {
            return "", err
        }
    }
	return opencc.Convert(s, config), nil
}

func ToSimplified(s string) (string, error) {
    config := `t2s.json`
    for _, fn := range []string{config, `libopencc.dll`, `TSPhrases.ocd`, `TSCharacters.ocd`} {
        err := initDepends(fn)
        if err != nil {
            return "", err
        }
    }
	return opencc.Convert(s, config), nil
}
