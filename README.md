# password

CLI/Docker container to generate strong passwords.

# features
- super secure defaults
- customizable
- 12 factor
- no HTTP API
- no JavaScript frontend

# installation
## cli
```sh

```
## docker
```sh
docker pull admiralobvious/password
```
## golang
```sh
go get github.com/admiralobvious/password
```

# usage
## cli
```sh
password generate
```
##docker
```sh
docker run --rm -it admiralobvious/password generate
```

# options
```
  -r, --allow-repeat           allows characters to repeat
      --digits string          list of permitted digits (default "0123456789")
  -h, --help                   help for generate
  -l, --length int             total number of characters in the password (default 64)
      --lower-letters string   list of permitted lowercase letters (default "abcdefghijklmnopqrstuvwxyz")
  -u, --no-upper               excludes uppercase letters
  -d, --num-digits int         number of digits to include in the result (default 10)
  -s, --num-symbols int        number of symbols to include in the result (default 10)
      --symbols string         list of permitted symbols (default "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./")
      --upper-letters string   list of permitted uppercase letters (default "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
```

# credits
[@sethvargo](https://github.com/sethvargo) for [go-password](https://github.com/sethvargo/go-password)
