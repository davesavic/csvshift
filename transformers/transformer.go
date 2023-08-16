package transformers

type Transformer interface {
	Apply(row map[string]interface{})
}

type RowTransformer struct {
	Transformers []Transformer
}
