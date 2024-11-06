package filter

import "context"


type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops{
	return &Ops{repo}
}

func (o *Ops)CreateFiletr(ctx context.Context, f *FilterSet) error{
	// validation check 
	// minprice < max price
	// price>0
	//
	// => convert persain to english numbers

	err := o.repo.Insert(ctx, f)
	if err!=nil{
		return ErrCreateFilter
	}
	return nil
}