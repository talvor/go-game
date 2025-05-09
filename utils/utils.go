package utils

import "path/filepath"

func GetAssetsDirectory(pathSegments ...string) string {
	segments := append([]string{"assets"}, pathSegments...)
	assetDir, err := filepath.Abs(filepath.Join(segments...))
	if err != nil {
		panic(err)
	}

	return assetDir
}
