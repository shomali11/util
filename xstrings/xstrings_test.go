package xstrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	assert.True(t, IsEmpty(""))
	assert.False(t, IsEmpty("text"))
	assert.False(t, IsEmpty("	"))
}

func TestIsNotEmpty(t *testing.T) {
	assert.False(t, IsNotEmpty(""))
	assert.True(t, IsNotEmpty("text"))
	assert.True(t, IsNotEmpty("	"))
}

func TestIsBlank(t *testing.T) {
	assert.True(t, IsBlank(""))
	assert.True(t, IsBlank("	"))
	assert.False(t, IsBlank("text"))
}

func TestIsNotBlank(t *testing.T) {
	assert.False(t, IsNotBlank(""))
	assert.False(t, IsNotBlank("	"))
	assert.True(t, IsNotBlank("text"))
}

func TestLeft(t *testing.T) {
	assert.Equal(t, Left("", 5), "     ")
	assert.Equal(t, Left("X", 5), "X    ")
	assert.Equal(t, Left("b\u0301", 3), "b\u0301  ")
	assert.Equal(t, Left("😎⚽", 4), "😎⚽  ")
	assert.Equal(t, Left("Les Mise\u0301rables", 5), "Les Mise\u0301rables")
	assert.Equal(t, Left("ab\u0301cde", 8), "ab\u0301cde   ")
	assert.Equal(t, Left("This `\xc5` is an invalid UTF8 character", 5), "This `\xc5` is an invalid UTF8 character")
	assert.Equal(t, Left("The quick bròwn 狐 jumped over the lazy 犬", 5), "The quick bròwn 狐 jumped over the lazy 犬")
	assert.Equal(t, Left("رائد شوملي", 10), "رائد شوملي")
}

func TestRight(t *testing.T) {
	assert.Equal(t, Right("", 5), "     ")
	assert.Equal(t, Right("X", 5), "    X")
	assert.Equal(t, Right("b\u0301", 3), "  b\u0301")
	assert.Equal(t, Right("😎⚽", 4), "  😎⚽")
	assert.Equal(t, Right("Les Mise\u0301rables", 5), "Les Mise\u0301rables")
	assert.Equal(t, Right("ab\u0301cde", 8), "   ab\u0301cde")
	assert.Equal(t, Right("This `\xc5` is an invalid UTF8 character", 5), "This `\xc5` is an invalid UTF8 character")
	assert.Equal(t, Right("The quick bròwn 狐 jumped over the lazy 犬", 5), "The quick bròwn 狐 jumped over the lazy 犬")
	assert.Equal(t, Right("رائد شوملي", 10), "رائد شوملي")
}

func TestCenter(t *testing.T) {
	assert.Equal(t, Center("", 5), "     ")
	assert.Equal(t, Center("X", 5), "  X  ")
	assert.Equal(t, Center("b\u0301", 3), " b\u0301 ")
	assert.Equal(t, Center("😎⚽", 4), " 😎⚽ ")
	assert.Equal(t, Center("Les Mise\u0301rables", 5), "Les Mise\u0301rables")
	assert.Equal(t, Center("ab\u0301cde", 8), "  ab\u0301cde ")
	assert.Equal(t, Center("This `\xc5` is an invalid UTF8 character", 5), "This `\xc5` is an invalid UTF8 character")
	assert.Equal(t, Center("The quick bròwn 狐 jumped over the lazy 犬", 5), "The quick bròwn 狐 jumped over the lazy 犬")
	assert.Equal(t, Center("رائد شوملي", 10), "رائد شوملي")
}

func TestLength(t *testing.T) {
	assert.Equal(t, Length(""), 0)
	assert.Equal(t, Length("X"), 1)
	assert.Equal(t, Length("b\u0301"), 1)
	assert.Equal(t, Length("😎⚽"), 2)
	assert.Equal(t, Length("Les Mise\u0301rables"), 14)
	assert.Equal(t, Length("ab\u0301cde"), 5)
	assert.Equal(t, Length("This `\xc5` is an invalid UTF8 character"), 37)
	assert.Equal(t, Length("The quick bròwn 狐 jumped over the lazy 犬"), 40)
	assert.Equal(t, Length("رائد شوملي"), 10)
}

func TestReverse(t *testing.T) {
	assert.Equal(t, Reverse(""), "")
	assert.Equal(t, Reverse("X"), "X")
	assert.Equal(t, Reverse("b\u0301"), "b\u0301")
	assert.Equal(t, Reverse("😎⚽"), "⚽😎")
	assert.Equal(t, Reverse("Les Mise\u0301rables"), "selbare\u0301siM seL")
	assert.Equal(t, Reverse("ab\u0301cde"), "edcb\u0301a")
	assert.Equal(t, Reverse("This `\xc5` is an invalid UTF8 character"), "retcarahc 8FTU dilavni na si `�` sihT")
	assert.Equal(t, Reverse("The quick bròwn 狐 jumped over the lazy 犬"), "犬 yzal eht revo depmuj 狐 nwòrb kciuq ehT")
	assert.Equal(t, Reverse("رائد شوملي"), "يلموش دئار")
}
