package maps

type Dictionary map[string]string

const (
	KeyNotFoundError  = DictionaryErr("not found")
	AlreadyExistError = DictionaryErr("we already have this key")
	KeyNotExistError  = DictionaryErr("not exist")
)

type DictionaryErr string

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", KeyNotFoundError
	}
	return value, nil
}
func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)
	switch err {
	case KeyNotFoundError:
		d[key] = value

	case nil:
		return AlreadyExistError

	default:
		return err
	}
	return nil
}
func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case KeyNotFoundError:
		return KeyNotExistError
	case nil:
		d[key] = value
	default:
		return err
	}
	return nil
}
func (d Dictionary) Delete(key string) {
	delete(d, key)
}
func (e DictionaryErr) Error() string {
	return string(e)
}
