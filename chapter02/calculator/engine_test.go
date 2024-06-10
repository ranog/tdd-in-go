package calculator_test

import (
	"github.com/PacktPublishing/Test-Driven-Development-in-Go/chapter02/calculator"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()

	e := m.Run()

	teardown()

	os.Exit(e)
}

func setup() {
	log.Println("Setting up.")
}

func teardown() {
	log.Println("Tearing down.")
}

func init() {
	log.Println("Init setup.")
}

func TestAdd(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	e := calculator.Engine{}
	actAssert := func(x, y, want float64) {
		got := e.Add(x, y)
		if got != want {
			t.Errorf("Add(%.2f, %.2f) incorrect, got: %.2f, want: %.2f", x, y, got, want)
		}
	}
	t.Run("positive input", func(t *testing.T) {
		x, y := 2.5, 3.5
		want := x + y
		actAssert(x, y, want)
	})
	t.Run("negative input", func(t *testing.T) {
		x, y := -2.5, -3.5
		want := x + y
		actAssert(x, y, want)
	})
}

func BenchmarkAdd(b *testing.B) {
	e := calculator.Engine{}
	for i := 0; i < b.N; i++ {
		e.Add(2.5, 3.5)
	}
}
