## cmd cli

add github token

## just download 

```
go run main.go --download=true --owner=NubeIO  --repo=rubix-service --dest=../bin  --asset=rubix-service --arch=amd64 --tag=latest --token=12
```
## just install and existing download

```
go run main.go --unzip=true  --dest=../bin --target=rubix-bios-app --existing-path=../bin --existing-asset=rubix-service-1.19.2-5bdd773f.amd64.zip
```


```
export GITHUB_TOKEN=your-token
```

## examples

## get releases

```
go run main.go list  --owner=NubeIO --repo=rubix-service  --per-page=3
```

## get repo info

if no `tag` is provided it will use tag `latest`

```
go run main.go info  --owner=NubeIO --repo=rubix-service 
```

```
go run main.go info  --owner=NubeIO --repo=rubix-service --tag=v1.18.0
```

### download

download a build with tag

```
go run main.go --owner=NubeIO  --repo=rubix-service --dest=../bin  --asset=rubix-service --arch=amd64 --tag=v0.0.1
```

if no `tag` is provided it will use tag `latest`

```
go run main.go --owner=NubeIO  --repo=rubix-service --dest=../bin  --asset=rubix-service --arch=amd64 --tag=latest
```

## manual install

this is meant to be used if the user already has a downloaded version of the asset (zip) on their PC

if `--dont-delete=false` is false then the zip will not be deleted once the installation is completed, set to `true` to
do a cleanup after the installation is done

```
go run main.go manual --manual-asset=rubix-service-1.19.0-eb71da61.amd64.zip --manual-path=/home/aidan  --dest=../bin  --dont-delete=false
```

## make the dir name have the version number

`--version-in-target=true`

will make the dir look like

```
├── data
│    └── rubix-bios-app
│       └── v1.5.2
│          └── rubix-bios

```

```
go run main.go --owner=NubeIO --repo=rubix-bios --dest=../bin --target=rubix-bios-app --asset=rubix-bios --arch=amd64 --tag=v1.5.2 --version-in-target=true
```
