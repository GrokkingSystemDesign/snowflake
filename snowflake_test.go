package snowflake

import (
	"testing"
)

func TestNewMachine_NormalID(t *testing.T) {
	_, err := NewMachine(16)
	if err != nil {
		t.Fatalf("error when creating machine: %s", err.Error())
	}
}

func TestNewMachine_ErrorID(t *testing.T) {
	_, err := NewMachine(1200)
	if err == nil {
		t.Fatalf("should raise an error but not")
	}
}

func TestPrintAll(t *testing.T) {
	m, err := NewMachine(0)
	if err != nil {
		t.Fatalf("error creating machine, %s", err)
	}
	id, err := m.Generate()
	if err != nil {
		t.Fatalf("error generating id")
	}
	t.Logf("Int64: %#v", id.Int64())
	t.Logf("String: %#v", id.String())
	t.Logf("Base64: %#v", id.Base64())
	t.Logf("Bytes: %#v", id.Bytes())
}

func BenchmarkGenerate(b *testing.B) {

	m, _ := NewMachine(0)

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, _ = m.Generate()
	}
}
