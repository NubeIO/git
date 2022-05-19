package git

import "errors"

func MakeAssetOptions(asset AssetOptions, manual manualInstall) (*AssetOptions, error) {
	if asset.Repo == "" {
		return nil, errors.New("require asset")
	}

	return &AssetOptions{
		Owner:         asset.Owner,
		Repo:          asset.Repo,
		Tag:           asset.Tag,
		OS:            asset.OS,
		OSAlias:       asset.OSAlias,
		Arch:          asset.Arch,
		ArchAlias:     asset.ArchAlias,
		DestPath:      asset.DestPath,
		Target:        asset.Target,
		manualInstall: manual,
	}, nil
}
