package mapper

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/pijalu/go.hands.two/frinsultdata/model"
	"github.com/pijalu/go.hands.two/frinsultproto"
)

// buildFrinsults builds simple test case objects
func buildFrinsults() (m *model.Frinsult, p *frinsultproto.Frinsult) {
	m = &model.Frinsult{
		Model: gorm.Model{ID: 42},
		Text:  "text",
		Score: 23,
	}

	p = &frinsultproto.Frinsult{
		ID:    42,
		Text:  "text",
		Score: 23,
	}

	return
}

func TestProto2Db(t *testing.T) {
	expected, input := buildFrinsults()
	actual := Proto2Db(input)
	if expected.ID != actual.ID ||
		expected.Text != actual.Text ||
		expected.Score != actual.Score {
		t.Fatalf("Expected %v but got %v", expected, actual)
	}
}

func TestDb2Proto(t *testing.T) {
	input, expected := buildFrinsults()
	actual := Db2Proto(input)
	if expected.ID != actual.ID ||
		expected.Text != actual.Text ||
		expected.Score != actual.Score {
		t.Fatalf("Expected %v but got %v", expected, actual)
	}
}
