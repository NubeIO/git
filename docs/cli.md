## cmd cli

these docs are for using the linux or windows command line

- cd into the cmd dir
- once in the main dir there is access to different protocols, or you can for example add/remove a host

```
cd cmd
```

## main app

### hosts

add a new a host

````
go run main.go hosts --new=true --name=test --ip=192.178.12.1
````

update a host ip (will up by the host name)

