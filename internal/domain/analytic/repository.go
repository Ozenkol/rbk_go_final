package analytic

type AnalyticRepositoryInterface interface {
	Create(analytic *Analytic) (*Analytic, error)
	GetByID(id string) (*Analytic, error)
	Update(analytic *Analytic) (*Analytic, error)
	Delete(id string) error
	List() ([]*Analytic, error)
}
