package maps

type Dictionary map[string]string

func (d Dictionary) Search(searchkey string) string {
	return d[searchkey]
}
