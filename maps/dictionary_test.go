package maps

import "testing"


func TestSearch(t *testing.T) {
  
  t.Run("Know words", func(t *testing.T){
    dictionary := Dictionary{"test":"this is a test"}
    got, _:= dictionary.Search("test")
    want := "this is a test"
    
    AssertStrings(t, got, want)
 })

  t.Run("Unkown words", func(t *testing.T) {
    dict := Dictionary{"test": "this is a test"}
    _, err := dict.Search("Unkown")

    if err == nil {
			t.Fatal("expected to get an error.")
		}

  })
}


func TestAdd(t *testing.T) {
  dictionary := Dictionary{}
  dictionary.Add("test", "this is a test")

  want := "this is a test"
  got, err := dictionary.Search("test")
  if err != nil {
    t.Fatal("Should find added word:",err)
  }

  AssertStrings(t, got, want)
}


func AssertStrings(t *testing.T, got string, want string) {
  if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
