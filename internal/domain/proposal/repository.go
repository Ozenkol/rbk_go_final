package proposal

type ProposalRepositoryInterface interface {
	Create(proposal *Proposal) (*Proposal, error)
	GetByID(id string) (*Proposal, error)
	Update(proposal *Proposal) (*Proposal, error)
	Delete(id string) error
	List() ([]*Proposal, error)
}
