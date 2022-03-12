// bitbucket.org/gotamer/cfg is a json configuration package
//     * You define your config items in your application as a struct.
//     * A json config file template from your struct is saved on the first run
//     * You can save runtime modifications to the config
// See doc.go for an example
package cfg

import (
	"encoding/json"
	"io/ioutil"
)

// Load gets your config from the json file,
// and fills your struct with the option
func Load(filename string, o interface{}) error {
	b, err := ioutil.ReadFile(filename)
	if err == nil {
		err = json.Unmarshal(b, &o)
	}
	return err
}

// Save will save your struct to the given filename,
// this is a good way to create a json template
func Save(filename string, o interface{}) error {
	j, err := json.MarshalIndent(&o, "", "\t")
	if err == nil {
		err = ioutil.WriteFile(filename, j, 0660)
	}
	return err
}
