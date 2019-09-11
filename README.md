# Automated search for music
Select randomly a range of albums listed by Guts of Darkness 

## Usage
### Dev
`mpsyt //$(go run main.go)`

### Everyday
`go build`

`sudo cp gutsOfDarkness /usr/bin`

`mpsyt //$(gutsOfDarkness)`

## TODO
Considering the number of time we play with indexes, I'm pretty sure there is an off-by-one issue somewhere

Tests :))

Split data recovery into another pkg

Handle no result/single pagination
