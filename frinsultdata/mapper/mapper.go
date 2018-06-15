package mapper

import (
	"github.com/jinzhu/gorm"
	"github.com/pijalu/go.hands.two/frinsultdata/model"
	"github.com/pijalu/go.hands.two/frinsultproto"
)

// Proto2Db convert a proto frinsult to a model
func Proto2Db(i *frinsultproto.Frinsult) *model.Frinsult {
	return &model.Frinsult{
		Model: gorm.Model{ID: uint(i.ID)},
		Text:  i.Text,
		Score: int(i.Score),
	}
}

// Db2Proto convert a proto frinsult to a model
func Db2Proto(i *model.Frinsult) *frinsultproto.Frinsult {
	return &frinsultproto.Frinsult{
		ID:    int64(i.ID),
		Text:  i.Text,
		Score: int64(i.Score),
	}
}
