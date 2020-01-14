package dictionary

//Err type
type Err string

func (e Err) Error() string {
	return string(e)
}

const (
	//ErrNotFound when definition not found in dictionary
	ErrNotFound = Err("could not find the word you were looking for")
	//ErrWordExisted word existed error
	ErrWordExisted = Err("word existed")
	//ErrWordDoesNotExisted word not existed for update
	ErrWordDoesNotExisted = Err("cannot update word because it does not exist")
)

//Dictionary type
type Dictionary map[string]string

//Search text in dictionary
func (dict Dictionary) Search(word string) (string, error) {
	definition, ok := dict[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

//Add word to dictionary
func (dict Dictionary) Add(word, definition string) error {
	_, err := dict.Search(word)
	switch err {
	case ErrNotFound:
		dict[word] = definition
	case nil:
		return ErrWordExisted
	default:
		return err
	}
	return nil
}

//Update existed word in dictionary
func (dict Dictionary) Update(word, newDefinition string) error {
	_, err := dict.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExisted
	case nil:
		dict[word] = newDefinition
	default:
		return err
	}
	return nil
}

//Delete a word in dictionary
func (dict Dictionary) Delete(word string) {
	delete(dict, word)
}
