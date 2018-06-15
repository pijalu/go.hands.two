package handler

import (
	"context"

	"github.com/pijalu/go.hands.two/frinsultdata/mapper"
	"github.com/pijalu/go.hands.two/frinsultdata/repository"
	"github.com/pijalu/go.hands.two/frinsultproto"
	"github.com/pkg/errors"
)

// FrinsultHandler is the RPC implementation
type FrinsultHandler struct{}

// GetFrinsultByID finds an insult by id
func (f *FrinsultHandler) GetFrinsultByID(c context.Context, req *frinsultproto.ByIDRequest, rep *frinsultproto.Frinsult) error {
	ID := uint(req.GetID())
	if ID == 0 {
		return errors.New("invalid ID")
	}

	dbInsult, err := repository.GetFrinsultByID(ID)
	if err != nil {
		return err
	}

	*rep = *mapper.Db2Proto(dbInsult)
	return nil
}

// DeleteFrinsultByID deletes an insults by id
func (f *FrinsultHandler) DeleteFrinsultByID(c context.Context, req *frinsultproto.ByIDRequest, _ *frinsultproto.Void) error {
	ID := uint(req.GetID())
	if ID == 0 {
		return errors.New("invalid ID")
	}

	return repository.DeleteFrinsultByID(ID)
}

// UpdateFrinsult updates an insult
func (f *FrinsultHandler) UpdateFrinsult(c context.Context, req *frinsultproto.Frinsult, _ *frinsultproto.Void) error {
	ID := uint(req.GetID())
	if ID == 0 {
		return errors.New("invalid ID")
	}

	return repository.UpdateFrinsult(
		mapper.Proto2Db(req))
}

// InsertFrinsult updates an insult
func (f *FrinsultHandler) InsertFrinsult(c context.Context, req *frinsultproto.Frinsult, rep *frinsultproto.Frinsult) error {
	fi, err := repository.InsertFrinsult(
		mapper.Proto2Db(req))
	if err != nil {
		return err
	}

	*rep = *mapper.Db2Proto(fi)
	return nil
}

// VoteFrinsultByID update score of a given insult
func (f *FrinsultHandler) VoteFrinsultByID(c context.Context, v *frinsultproto.VoteRequest, _ *frinsultproto.Void) error {
	ID := uint(v.GetID())
	if ID == 0 {
		return errors.New("invalid ID")
	}

	return repository.VoteForFrinsult(uint(v.GetID()), int(v.GetVote()))
}

// GetFrinsults returns a list of items
func (f *FrinsultHandler) GetFrinsults(ctx context.Context, in *frinsultproto.Void, out *frinsultproto.Frinsults) error {
	insults, err := repository.GetFrinsults()
	if err != nil {
		return err
	}

	out.Insults = make([]*frinsultproto.Frinsult, 0, len(insults))

	for _, fri := range insults {
		out.Insults = append(out.Insults, mapper.Db2Proto(&fri))
	}

	return nil
}
