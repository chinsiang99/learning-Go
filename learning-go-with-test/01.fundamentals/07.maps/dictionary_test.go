package maps

import "testing"

// func TestSearch(t *testing.T) {
// 	assertBalance := func(t testing.TB, got string, want string) {
// 		t.Helper()
// 		// got := wallet.Balance().String()

// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}
// 	}

// 	t.Run("deposit successfully", func(t *testing.T) {
// 		wallet := Wallet{}

// 		wallet.Deposit(Bitcoin(10))

// 		got := wallet.Balance().String()
// 		want := Bitcoin(10).String()

// 		assertBalance(t, got, want)
// 	})
// }

func TestSearch(t *testing.T) {
	// dictionary := Dictionary{"test": "this is just a test"}

	// got := dictionary.Search("test")
	// want := "this is just a test"

	// assertStrings(t, got, want)

	t.Run("Dictionary search functionality", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		searchTestCases := []struct {
			name       string
			dictionary Dictionary
			searchWord string
			wantString string
			expectErr  error
		}{
			{name: "search known word", dictionary: dictionary, searchWord: "test", wantString: "this is just a test", expectErr: nil},
			{name: "search unknown word", dictionary: dictionary, searchWord: "unknown", wantString: "", expectErr: ErrorNotFoundWord},
		}

		for _, tc := range searchTestCases {
			t.Run(tc.name, func(t *testing.T) {
				// since they dont share state, we can run it parallelly
				t.Parallel()

				got, err := tc.dictionary.Search(tc.searchWord)

				assertError(t, err, tc.expectErr)

				assertStrings(t, got, tc.wantString)
			})
		}
	})

	t.Run("Dictionary add functionality", func(t *testing.T) {
		dictionary := Dictionary{"test": "testValue"}
		searchTestCases := []struct {
			name       string
			dictionary Dictionary
			addWord    string
			value      string
			wantString string
			expectErr  error
		}{
			{name: "add a new word", dictionary: dictionary, addWord: "testing new word", value: "value is here", wantString: "value is here", expectErr: nil},
			{name: "aadd existing word", dictionary: dictionary, addWord: "test", value: "hello there", wantString: "testValue", expectErr: ErrorWordExists},
		}

		for _, tc := range searchTestCases {
			t.Run(tc.name, func(t *testing.T) {
				// since they dont share state, we can run it parallelly
				t.Parallel()

				err := tc.dictionary.Add(tc.addWord, tc.value)

				assertError(t, err, tc.expectErr)

				got, err := tc.dictionary.Search(tc.addWord)

				assertError(t, err, err)
				assertStrings(t, got, tc.wantString)
			})
		}
	})

	t.Run("Dictionary update functionality", func(t *testing.T) {
		dictionary := Dictionary{"test": "testValue"}
		searchTestCases := []struct {
			name       string
			dictionary Dictionary
			updateWord string
			value      string
			wantString string
			expectErr  error
		}{
			{name: "update existing word", dictionary: dictionary, updateWord: "test", value: "value is here", wantString: "value is here", expectErr: nil},
			{name: "update non-existent word", dictionary: dictionary, updateWord: "test2", value: "value is here", wantString: "", expectErr: ErrorNotFoundWord},
		}

		for _, tc := range searchTestCases {
			t.Run(tc.name, func(t *testing.T) {
				// since they dont share state, we can run it parallelly
				t.Parallel()

				err := tc.dictionary.Update(tc.updateWord, tc.value)

				assertError(t, err, tc.expectErr)

				got, err := tc.dictionary.Search(tc.updateWord)

				assertError(t, err, err)
				assertStrings(t, got, tc.wantString)
			})
		}
	})

	t.Run("Dictionary delete functionality", func(t *testing.T) {
		dictionary := Dictionary{"test": "testValue"}
		searchTestCases := []struct {
			name       string
			dictionary Dictionary
			deleteWord string
			expectErr  error
			wantString string
		}{
			{name: "delete existing word", dictionary: dictionary, deleteWord: "test", expectErr: nil, wantString: ""},
			{name: "delete non-existent word", dictionary: dictionary, deleteWord: "test2", expectErr: ErrorNotFoundWord, wantString: ""},
		}

		for _, tc := range searchTestCases {
			t.Run(tc.name, func(t *testing.T) {
				// since they dont share state, we can run it parallelly
				t.Parallel()

				err := tc.dictionary.Delete(tc.deleteWord)

				assertError(t, err, tc.expectErr)

				got, err := tc.dictionary.Search(tc.deleteWord)

				assertError(t, err, err)
				assertStrings(t, got, tc.wantString)
			})
		}
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q, want %q", got, want)
	}
}
