package configuration

import("flag")

func NewCLIParser() (string, int) {
  year := flag.String("year", "", "Year of release")
  minRank := flag.Int("minRank", 0, "Lowest rank, default 5")
  flag.Parse()
  return *year, *minRank
}
