package maps

import "errors"


type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
  definition, err := d[word]

  if !err{
    return "", errors.New("Could not find the definition")
  }

  return definition, nil
}


func (d Dictionary) Add(word, definition string) {
  d[word] = definition
}


// write update method


// Write delete method
//Go has a built-in function delete that works on maps. It takes two arguments. The first is the map and the second is the key to be removed.
