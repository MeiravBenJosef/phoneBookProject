package utills

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

//A struct to contain pagination feature params
type mongoPaginate struct {
	limit int64
	page int64
 }

 //A constructor for the mongoPaginate struct
func NewMongoPaginate(limit, page int) *mongoPaginate {
	return &mongoPaginate{
	   limit: int64(limit),
	   page:  int64(page),
	}
 }

 //GetPaginatedOpts will return the correct oprtions fields with pagination feature params to send mongo db.
 func (mp *mongoPaginate) GetPaginatedOpts() *options.FindOptions {
	l := mp.limit
	skip := mp.page*mp.limit - mp.limit
	fOpt := options.FindOptions{Limit: &l, Skip: &skip}
 
	return &fOpt
 }