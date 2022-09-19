package kata

func NbYear(p0 int, percent float64, aug int, p int) int {
  var population = int(float64(p0) + float64(p0) * (float64(percent)*0.01) + float64(aug))
  var years int
  for i := 2 ; i >= 0 ; i++ {
    population = int(float64(population) + float64(population) * (float64(percent)*0.01) + float64(aug))
    if population < p {
      continue
    } else {
      years = i
      break
    }
  }
  return years
}
