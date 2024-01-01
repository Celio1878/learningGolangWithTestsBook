package dictionary

import (
	"errors"
	"testing"
)

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})

	t.Run("add word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "wondering"
		definition := "desire to know something, curious"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("add existing word", func(t *testing.T) {
		word := "rather than"
		definition := "instead of"

		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("update word", func(t *testing.T) {
		word := "make"
		definition := "build something"
		dictionary := Dictionary{word: definition}

		newDefinition := "create something"
		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("Updating word not found", func(t *testing.T) {
		word := "unknown"
		definition := "unknown"

		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})

	t.Run("delete word", func(t *testing.T) {
		word := "delete"
		dictionary := Dictionary{word: "delete"}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		if !errors.Is(err, ErrNotFound) {
			t.Errorf("Expected %q to be deleted", word)
		}
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, definition)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if !errors.Is(got, want) {
		t.Errorf("got error %q want %q", got, want)
	}
}
