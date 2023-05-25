package endpoints

import (
	"context"
	"errors"

	"github.com/steefanSS/watermark-service/internal"
)

type InterfaceSet Set

//Separated logic a bit more, this are the methods that are going to be used. The flow goes like this. Service defines concrete models, MakeEndpoint factory functions, wrap the types and form and point
//Finally the new type InterfaceSet of type Set is created and with it, concrete methods to be used to get the job done
//(didn't make sense to me to mix factory functions and methods in the same file, it's easily confusing for anyone who doesn't have experience with golang and in general structured programming lang like go)

func (s *InterfaceSet) Get(ctx context.Context, filters ...internal.Filter) ([]internal.Document, error) {
	resp, err := s.GetEndpoint(ctx, GetRequest{Filters: filters})

	if err != nil {
		return []internal.Document{}, err
	}

	getResp := resp.(GetResponse)

	if getResp.Err != "" {
		return []internal.Document{}, errors.New(getResp.Err)
	}

	return getResp.Document, nil
}
