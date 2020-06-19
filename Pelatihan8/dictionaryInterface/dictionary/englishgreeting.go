package dictionary

type EnglishDict struct{}

// disini pointer jadi di app.go harus kirim alamat pake &
func (en *EnglishDict) GetMorningGreeting() string {
	return "Hello, good morning"
}
