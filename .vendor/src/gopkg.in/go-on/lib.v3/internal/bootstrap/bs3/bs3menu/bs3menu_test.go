package bs3menu

import (
	"testing"
)

/*
menuhandler.NewStatic(Menu, 1, bs3menu.NavBar()),

Menu := &menu.Node{}
	err := json.Unmarshal([]byte(menuJson), &Menu)

	if err != nil {
		panic(err.Error())
	}
var menuJson = `
{
  "Subs": [
    { "Text": "Languages", "Path": "/languages",
      "Subs": [
        { "Text": "english", "Path": "/english",
          "Subs": [
            { "Text": "american english", "Path": "/en_us"},
            { "Text": "british english", "Path": "/en_gb" }
          ]
        },
        {"Text": "---"},
        {"Text": "french", "Path": "/fr"}
      ]
    },
    { "Text": "Countries", "Path": "/countries",
      "Subs": [
        { "Text": "USA", "Path": "/usa" },
        { "Text": "Brazil", "Path": "_" },
        { "Text": "Europe", "Path": "/europe",
          "Subs": [
            { "Text": "UK", "Path": "#uk" },
            { "Text": "France", "Path": "#france"}
          ]
        }
      ]
    },
    { "Text": "Currencies", "Path": "/currencies" }
  ]
}
`
*/

func TestX(t *testing.T) {

}
