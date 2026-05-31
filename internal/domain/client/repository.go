package client

type ClientRepositoryInterface interface {
	Save(client *Client) error
	FindByID(id string) (*Client, error)
	FindAll() ([]*Client, error)
	Update(client *Client) error
	Delete(id string) error
}