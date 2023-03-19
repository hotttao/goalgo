package algo

type Product struct {
	Name       string
	Attributes map[string]string
}

func (p *Product) AddAttr(attr, value string) {
	if p.Attributes == nil {
		p.Attributes = make(map[string]string)
	}
	p.Attributes[attr] = value
}

