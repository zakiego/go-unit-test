package helper

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// only run once every helper
func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkipOnLinux(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Can not run on linux")
	}
}

func TestHelloWorld(t *testing.T) {

	result := HelloWorld("Zaki")

	s := strings.Split(result, " ")

	require.Equal(t, "Hello", s[0], "The first word must be 'Hello'")
	require.Equal(t, "Zaki", s[1], "Output name must be 'Zaki'")
}

func TestHelloWorldMoreThanOneWord(t *testing.T) {
	result := HelloWorld("Zaki Terbang Ke Angkasa")

	s := strings.Split(result, " ")
	sJoin := strings.Join(s[1:], " ")

	require.Equal(t, "Hello", s[0], "The first word must be 'Hello'")
	require.Equal(t, "Zaki Terbang Ke Angkasa", sJoin, "Output name must be 'Zaki Terbang Ke Angkasa")
}

func TestSubTest(t *testing.T) {
	t.Run("Zaki", func(t *testing.T) {
		result := HelloWorld("Zaki")
		require.Equal(t, "Hello Zaki", result, "Result must be 'Hello Zaki'")
	})

	t.Run("Agus", func(t *testing.T) {
		result := HelloWorld("Agus")
		require.Equal(t, "Hello Agus", result, "Result must be 'Hello Agus'")
	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Zaki",
			request:  "Zaki",
			expected: "Hello Zaki",
		},
		{
			name:     "Muhammmad Zaki",
			request:  "Muhammmad Zaki",
			expected: "Hello Muhammmad Zaki",
		},
	}

	for _, test := range tests {
		result := HelloWorld(test.request)
		require.Equal(t, test.expected, result)
	}
}
