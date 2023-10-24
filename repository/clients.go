package repository

import (
	"context"

	"github.com/TakeshiHA/test-middleware/database"
	"github.com/TakeshiHA/test-middleware/models"
)

type ClientRepository struct {
	clientCollection *database.ClientCollection
}

type ClientRepositoryInterface interface {
	CreateClient(ctx context.Context, body *models.Client) (*models.Client, error)
	GetClients(ctx context.Context) ([]*models.Client, error)
	UpdateClient(ctx context.Context, id string, client *models.Client) (bool, error)
	DeleteClientById(ctx context.Context, id string) (bool, error)
	GetClientById(ctx context.Context, id string) (*models.Client, error)
}

func NewClientRepository() *ClientRepository {
	return &ClientRepository{database.GetClientCollection()}
}

func (clientRepo *ClientRepository) CreateClient(ctx context.Context, body *models.Client) (*models.Client, error) {
	cities, err := clientRepo.clientCollection.CreateClient(ctx, body)
	return cities, err
}

func (clientRepo *ClientRepository) GetClients(ctx context.Context) ([]*models.Client, error) {
	cities, err := clientRepo.clientCollection.GetClients(ctx)
	return cities, err
}

func (clientRepo *ClientRepository) GetClientByDNI(ctx context.Context, dni string) (*models.Client, error) {
	cities, err := clientRepo.clientCollection.GetClientByDNI(ctx, dni)
	return cities, err
}

func (clientRepo *ClientRepository) GetClientById(ctx context.Context, id string) (*models.Client, error) {
	cities, err := clientRepo.clientCollection.GetClientById(ctx, id)
	return cities, err
}

func (clientRepo *ClientRepository) UpdateClient(ctx context.Context, id string, client *models.Client) (bool, error) {
	cities, err := clientRepo.clientCollection.UpdateClient(ctx, id, client)
	return cities, err
}

func (clientRepo *ClientRepository) DeleteClientById(ctx context.Context, id string) (bool, error) {
	cities, err := clientRepo.clientCollection.DeleteClientById(ctx, id)
	return cities, err
}
