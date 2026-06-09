package proposal

type ProposalFactoryInterface interface {
	CreateProposal(clientID string) (*Proposal, error)
}

type ProposalFactory struct {
	productRepository any
}

func NewProposalFactory(productRepository any) ProposalFactoryInterface {
	return &ProposalFactory{productRepository: productRepository}
}

func (f *ProposalFactory) CreateProposal(clientID string) (*Proposal, error) {
	return &Proposal{
		ClientID: clientID,
	}, nil
}
