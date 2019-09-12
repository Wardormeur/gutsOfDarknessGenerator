# Automated search for music
Select randomly a range of albums listed by Guts of Darkness 

## Usage
### Dev
`mpsyt //$(go run main.go)`

### Everyday
#### Setup

`go build`

`sudo cp gutsOfDarkness /usr/bin`

### Run

`mpsyt //$(gutsOfDarkness)`

### Options
--year: filter on this specific year only, default none

--minRank: filter albums from this rank onwards the maximum, default 5 (0-6)

### Examples

`gutsOfDarkness` will return albums from rank 5 to 6 from any year

`gutsOfDarkness --year=2018` will return albums from rank 5 to 6 from 2018

`gutsOfDarkness --minRank=6` will return albums of rank 6 from any year 

## Return values
### When result found
Artist - Album

Exit(0)

### When no result
Exit(1)

## TODO
Tests :))

Split data recovery into another pkg
