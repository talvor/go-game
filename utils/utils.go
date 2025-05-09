package utils

import "path/filepath"

func GetAssetsDirectory(assetType string) string {
	assetDir, err := filepath.Abs(filepath.Join("assets", assetType))
	if err != nil {
		panic(err)
	}

	return assetDir
}
