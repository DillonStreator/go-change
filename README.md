# go-change

Change calculation that can be run via cli or as an api

### CLI

```sh
./go-change -h
Usage of ./go-change:
  -a	run as api (shorthand)
  -api
    	run as api
  -d string
    	the drawer (shorthand)
  -drawer string
    	the drawer
  -j string
    	json input with keys paid, owed, drawer (shorthand)
  -json string
    	json input with keys paid, owed, drawer
  -o int
    	amount that is owed (shorthand)
  -owed int
    	amount that is owed
  -p int
    	amount that is paid (shorthand)
  -paid int
    	amount that is paid
```

### API

Run the binary with the `a` flag to start as an api\
`./go-change -a`

or with a custom port\
`PORT=4141 ./go-change -a`

POST `/`\
body
```json
{
    "drawer": [1,2,3],
    "owed": 10,
    "paid": 14
}
```