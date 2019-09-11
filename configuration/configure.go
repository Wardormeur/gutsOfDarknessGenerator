package configuration

import(
  "math/rand"
  "time"
)

type configuration struct {
  Year string
  MinRank int
  MaxRank int
  Rank int
}

func New(year string, minRank int) *configuration {
  var defaultMinRank int = 5;
  var maxRank int = 6
  if minRank == 0 {
    minRank = defaultMinRank
  }

  conf := &configuration{Year: year, MinRank: minRank, MaxRank: maxRank}
  conf.Rank = conf.GetRank()
  return conf
}

func (c configuration) GetRank() int {
  var rank int = 0
  rand.Seed(time.Now().UnixNano())
  if c.MinRank != c.MaxRank {
    rank = rand.Intn((c.MaxRank + 1) - c.MinRank) + c.MinRank
  } else {
    rank = c.MinRank
  }
  return rank
}
