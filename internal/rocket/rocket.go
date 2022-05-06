//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/misua/go-grpc-svc/internal/rocket Store

package rocket

import "context"

//NOTE Rocket  - should contain the definition of our Rocket
type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}

// NOTE Store - defines the interface we expect
// our database implementation to follow
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

//NOTE  Service- responsible for updating rocket inventory
type Service struct {
	Store Store
}

func New(store Store) Service {
	return Service{
		Store: store,
	}
}

//NOTE GetRocketByID - retries a rocket based on the id from the store
func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, err
	}
	//NOTE if it manages to get rocket just return nil
	return rkt, nil

}

//NOTE - insert a new rocket into the store
func (s Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

//NOTE  DeleteRocket - deletes a rocket from our inventory
func (s Service) DeleteRocket(id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}
