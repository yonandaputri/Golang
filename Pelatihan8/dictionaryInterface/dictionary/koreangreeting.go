package dictionary

type KoreaDict struct{}

// disini pointer jadi di app.go harus kirim alamat pake &
func (kr *KoreaDict) GetMorningGreeting() string {
	return "좋은 아침"
}
