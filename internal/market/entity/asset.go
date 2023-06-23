package entity

// agora a ação em si
type Asset struct {
	ID           string
	Name         string
	MarketVolume int // n de ações no mercado
}

// função construtora gerando um novo objeto
func NewAsset(id string, name string, marketVolume int) *Asset {
	return &Asset{
		ID:           id,
		Name:         name,
		MarketVolume: marketVolume,
	}
}
