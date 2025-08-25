package pkgutils

import (
	"errors"

	"advent/utils/collections"

	"golang.org/x/tools/go/packages"
)

func PackageByName(name string) (*packages.Package, error) {

	// must do recursive search - Load itself does NOT fail if it looks for packages that do not exist
	pk, e := packages.Load(&packages.Config{Mode: packages.NeedName | packages.NeedTypes}, "./...")
	if e != nil {
		return nil, e
	}

	// now filter the results
	pk = collections.Filter(pk, func (p *packages.Package) bool { return p.Name == name })
	if len(pk) == 0 {
		return nil, errors.New("package missing")
	}

	return pk[0], nil
}
