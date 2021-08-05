# Methods

> For our purposes, an *object* is simply a value or variable that has methods, and a *method* is a function associated with a particular type.

A *method* is declared with a variant of the ordinary function declaration in which an extra parameter appears before the function name. The parameter attaches the function to the type of that parameter.

```golang
func (p Point) Distance(q Point) float64 {
  return math.Hypot(q.X - p.X, q.Y - p.Y)
}

p := Point{1, 2}
q := Point{4, 6}

p.Distance(q))

distanceFromP := p.Distance // method value
distanceFromP(q)
```

If you need to update a variable, or an argument is so large that we wish to avoid copying it, we must use a pointer.

```golang
func (p *Point) ScaleBy(factor float64) {
  p.X *= factor
  p.Y *= factor
}
```
