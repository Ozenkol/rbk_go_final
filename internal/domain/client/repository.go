package client

type ClientRepositoryInterface interface {
	Create(client *Client) (*Client, error)
	GetByID(id string) (*Client, error)
	Update(client *Client) (*Client, error)
	Delete(id string) error
	List() ([]*Client, error)
}
