package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/TakeshiHA/test-middleware/functions"
	"github.com/TakeshiHA/test-middleware/models"
	"github.com/TakeshiHA/test-middleware/repository"
	"github.com/labstack/echo/v4"
)

type ClientUsecase struct {
	clientRepository *repository.ClientRepository
	contextTimeout   time.Duration
}

type ClientUsecaseInterface interface {
	CreateClient(ctx context.Context, body *models.Client) (*models.Client, *echo.HTTPError)
	GetClients(ctx context.Context) ([]*models.Client, *echo.HTTPError)
	UpdateClient(ctx context.Context, id string, client *models.Client) (*models.ResponseSuccess, *echo.HTTPError)
	DeleteClientById(ctx context.Context, id string) (*models.ResponseSuccess, *echo.HTTPError)
}

func NewClientUsecase(
	clientRepository *repository.ClientRepository,
	timeout time.Duration,
) *ClientUsecase {
	return &ClientUsecase{
		clientRepository: clientRepository,
		contextTimeout:   timeout,
	}
}

func (clientUCase *ClientUsecase) CreateClient(ctx context.Context, body *models.Client) (*models.Client, *echo.HTTPError) {
	ctx, cancel := context.WithTimeout(ctx, clientUCase.contextTimeout)
	defer cancel()

	if len(body.DNI) != 8 || !functions.ValidateDNI(body.DNI) {
		return nil, echo.NewHTTPError(420, models.ResponseError{Message: "El DNI debe ser 8 dígitos."})
	}

	if len(body.Phone) != 9 || !functions.ValidatePhoneNumber(body.Phone) {
		return nil, echo.NewHTTPError(421, models.ResponseError{Message: "El número de celular debe ser 9 dígitos."})
	}

	if !functions.ValidateEmail(body.Email) {
		return nil, echo.NewHTTPError(422, models.ResponseError{Message: "El correo electrónico es inválido."})
	}

	flagClient, _ := clientUCase.clientRepository.GetClientByDNI(ctx, body.DNI)

	if flagClient != nil {
		return nil, echo.NewHTTPError(422, models.ResponseError{Message: fmt.Sprintf("El cliente con el DNI %s ya se encuentra registado.", body.DNI)})
	}

	client, err := clientUCase.clientRepository.CreateClient(ctx, body)

	if err != nil {
		return nil, echo.NewHTTPError(400, models.ResponseError{Message: "Error! Contáctese con el administrador."})
	}

	return client, nil
}

func (clientUCase *ClientUsecase) GetClients(ctx context.Context) ([]*models.Client, *echo.HTTPError) {
	ctx, cancel := context.WithTimeout(ctx, clientUCase.contextTimeout)
	defer cancel()

	cities, err := clientUCase.clientRepository.GetClients(ctx)

	if err != nil {
		return nil, echo.NewHTTPError(400, models.ResponseError{Message: "Error! Contáctese con el administrador."})
	}

	return cities, nil
}

func (clientUCase *ClientUsecase) UpdateClient(ctx context.Context, id string, body *models.Client) (*models.ResponseSuccess, *echo.HTTPError) {
	ctx, cancel := context.WithTimeout(ctx, clientUCase.contextTimeout)
	defer cancel()

	if len(body.DNI) != 8 || !functions.ValidateDNI(body.DNI) {
		return nil, echo.NewHTTPError(420, models.ResponseError{Message: "El DNI debe ser 8 dígitos."})
	}

	if len(body.Phone) != 9 || !functions.ValidatePhoneNumber(body.Phone) {
		return nil, echo.NewHTTPError(421, models.ResponseError{Message: "El número de celular debe ser 9 dígitos."})
	}

	if !functions.ValidateEmail(body.Email) {
		return nil, echo.NewHTTPError(422, models.ResponseError{Message: "El correo electrónico es inválido."})
	}

	flagClient, _ := clientUCase.clientRepository.GetClientByDNI(ctx, body.DNI)

	if flagClient != nil && flagClient.ID != id {
		return nil, echo.NewHTTPError(422, models.ResponseError{Message: fmt.Sprintf("El cliente con el DNI %s ya se encuentra registado.", body.DNI)})
	}

	isUpdated, err := clientUCase.clientRepository.UpdateClient(ctx, id, body)

	if err != nil && !isUpdated {
		return nil, echo.NewHTTPError(400, models.ResponseError{Message: "Error! Contáctese con el administrador."})
	}

	return &models.ResponseSuccess{Message: "Cliente actualizado correctamente!"}, nil
}

func (clientUCase *ClientUsecase) DeleteClientById(ctx context.Context, id string) (*models.ResponseSuccess, *echo.HTTPError) {
	ctx, cancel := context.WithTimeout(ctx, clientUCase.contextTimeout)
	defer cancel()

	isDeleted, err := clientUCase.clientRepository.DeleteClientById(ctx, id)

	if err != nil && !isDeleted {
		return nil, echo.NewHTTPError(400, models.ResponseError{Message: "Error! Contáctese con el administrador."})
	}

	return &models.ResponseSuccess{Message: "Cliente eliminado correctamente!"}, nil
}

func (clientUCase *ClientUsecase) GetClientById(ctx context.Context, id string) (*models.Client, *echo.HTTPError) {
	ctx, cancel := context.WithTimeout(ctx, clientUCase.contextTimeout)
	defer cancel()

	cities, err := clientUCase.clientRepository.GetClientById(ctx, id)

	if err != nil {
		return nil, echo.NewHTTPError(400, models.ResponseError{Message: "Error! Contáctese con el administrador."})
	}

	return cities, nil
}
