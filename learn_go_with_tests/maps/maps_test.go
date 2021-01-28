package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known key", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertString(t, got, want)
	})
	t.Run("unknown key", func(t *testing.T) {
		_, got := dictionary.Search("abaaba")
		want := KeyNotFoundError
		assertError(t, got, want)
	})

}
func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "6"
		value := "blaze"
		dictionary.Add(key, value)
		assertDefinition(t, dictionary, key, value)
	})
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "6"
		value := "blaze"
		dictionary.Add(key, value)
		err := dictionary.Add(key, "blazeee")
		assertError(t, err, AlreadyExistError)
		assertDefinition(t, dictionary, key, value)
	})
}
func TestUpdate(t *testing.T) {
	t.Run("update the message", func(t *testing.T) {
		key := "6"
		value := "kaltsit"
		dictionary := Dictionary{key: value}
		newValue := "chen"

		dictionary.Update(key, newValue)
		assertDefinition(t, dictionary, key, newValue)
	})
}

func assertString(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}
func assertDefinition(t *testing.T, dic Dictionary, key, value string) {
	t.Helper()
	got, err := dic.Search(key)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if value != got {
		t.Errorf("got '%s' want '%s'", got, value)
	}
}
