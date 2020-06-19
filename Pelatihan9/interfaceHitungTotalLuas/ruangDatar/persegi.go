package ruangDatar

type Persegi struct {
	Sisi float64
}

func (p Persegi) HitungLuas() float64 {
	return p.Sisi * p.Sisi
}
