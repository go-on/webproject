package main

import (
	"encoding/json"

	"gopkg.in/go-on/cdncache.v1"
	. "gopkg.in/go-on/lib.v2/html"
	"gopkg.in/go-on/lib.v2/internal/bootstrap/bs3"
	"gopkg.in/go-on/lib.v2/internal/bootstrap/bs3/bs3menu"
	"gopkg.in/go-on/lib.v2/types"
	// "gopkg.in/go-on/lib.v2/html/h"
	// . "gopkg.in/go-on/lib.v2/html/tag"
	"net/http"

	"gopkg.in/go-on/lib.v2/internal/menu"
	"gopkg.in/go-on/lib.v2/internal/menu/menuhandler"
)

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

func main() {
	/*
		m := &menu.Node{
			Edges: []*menu.Node{

				&menu.Node{
					Edges: []*menu.Node{
						&menu.Node{
							Edges: []*menu.Node{
								&menu.Node{
									Leaf: menu.Item("american english", "/en_us"),
								},
								&menu.Node{
									Leaf: menu.Item("british english", "/en_gb"),
								},
							},
							Leaf: menu.Item("english", "/english"),
						},
						&menu.Node{Leaf: menu.Item("---", "")},
						&menu.Node{Leaf: menu.Item("french", "/fr")},
					},
					Leaf: menu.Item("Languages", "/languages"),
				},
				&menu.Node{
					Edges: []*menu.Node{
						&menu.Node{Leaf: menu.Item("USA", "/usa")},
						&menu.Node{Leaf: menu.Item("Brazil", "_")},
						&menu.Node{
							Edges: []*menu.Node{
								&menu.Node{Leaf: menu.Item("UK", "#uk")},
								&menu.Node{Leaf: menu.Item("France", "#france")},
							},
							Leaf: menu.Item("Europe", "/europe"),
						},
					},
					Leaf: menu.Item("Countries", "/countries"),
				},

				&menu.Node{
					Leaf: menu.Item("Currencies", "/currencies"),
				},
			},
		}
	*/
	//data, _ := json.Marshal(m)

	//fmt.Printf("%s", data)
	Menu := &menu.Node{}
	err := json.Unmarshal([]byte(menuJson), &Menu)

	if err != nil {
		panic(err.Error())
	}

	cdn := cdncache.CDN("/cdn-cache/")

	http.Handle("/",
		HTML5(
			HEAD(
				CharsetUtf8(),
				CssHref(cdn(bs3.CDN_3_1_1_min)),
				CssHref(cdn(bs3.CDN_3_1_1_theme_min)),
				HttpEquiv("X-UA-Compatible", "IE=edge"),
				Viewport("width=device-width, initial-scale=1"),
				TITLE("Example with bootsrap"),
			),

			BODY(
				NAV(bs3.Navbar, bs3.Navbar_default,
					DIV(
						bs3.Container_fluid,
						DIV(bs3.Navbar_header, SPAN(bs3.Navbar_brand, "NavBar Menu @ 0-1")),
						menuhandler.NewStatic(Menu, 1, bs3menu.NavBar()),
					),
				),

				DIV(
					bs3.Container_fluid,
					menuhandler.NewStatic(Menu, 2, bs3menu.Breadcrumb()),
				),

				DIV(
					bs3.Container_fluid,

					DIV(bs3.Col_sm_6,
						DIV(bs3.Panel, bs3.Panel_default,
							DIV(bs3.Panel_heading, "Dropdown Buttons"),
							DIV(bs3.Panel_body,
								DIV(bs3.Btn_toolbar,
									DIV(
										bs3.Btn_group,
										menuhandler.NewStaticSub(Menu, 0, 0, bs3menu.DropdownButton(bs3.Btn_default, "", "Category")),
										menuhandler.NewStaticSub(Menu, 0, 0, bs3menu.Dropdown()),
									),
									DIV(
										bs3.Btn_group,
										menuhandler.NewStaticSub(Menu, 1, 1, bs3menu.Button(bs3.Btn_success, "%s", "»")),
										menuhandler.NewStaticSub(Menu, 1, 1, bs3menu.DropdownButton(bs3.Btn_success, "", "")),
										menuhandler.NewStaticSub(Menu, 1, 1, bs3menu.Dropdown()),
									),
									DIV(
										bs3.Btn_group,
										menuhandler.NewStaticSub(Menu, 2, 2, bs3menu.Button(bs3.Btn_warning, "%s", "»")),
										menuhandler.NewStaticSub(Menu, 2, 2, bs3menu.DropdownButton(bs3.Btn_warning, "", "")),
										menuhandler.NewStaticSub(Menu, 2, 2, bs3menu.Dropdown()),
									),
								),
							),
						),
					),

					DIV(bs3.Col_sm_3,
						DIV(bs3.Panel, bs3.Panel_default,
							DIV(bs3.Panel_heading, "Pills stacked @ 1-2"),
							DIV(bs3.Panel_body,
								menuhandler.NewStaticSub(Menu, 1, 2, bs3menu.Pills(true, bs3.Nav_stacked)),
							),
						),
					),

					DIV(bs3.Col_sm_3,
						DIV(bs3.Panel, bs3.Panel_default,
							DIV(bs3.Panel_heading, "Listgroup @ 2"),
							menuhandler.NewStaticSub(Menu, 2, 2, bs3menu.ListGroup()),
						),
					),

					DIV(bs3.Col_sm_6,
						DIV(bs3.Panel, bs3.Panel_default,
							DIV(bs3.Panel_heading, "Tabs @ 2"),
							DIV(bs3.Panel_body,
								menuhandler.NewStaticSub(Menu, 2, 2, bs3menu.Tabs(true, true)),
								DIV(bs3.Tab_content,
									DIV(bs3.Tab_pane, types.Id("uk"),
										AHref("http://en.wikipedia.org/wiki/United_Kingdom", "From Wikipedia:"),
										CITE(
											`The United Kingdom of Great Britain and Northern Ireland,`+
												`commonly known as the United Kingdom (UK) or Britain /ˈbrɪ.tən/, `+
												`is a sovereign state located off the north-western coast of `+
												`continental Europe.`,
										),
									),
									DIV(bs3.Tab_pane, types.Id("france"),
										AHref("http://en.wikipedia.org/wiki/France", "From Wikipedia:"),
										CITE(
											`France (UK: /frɑːns/; US: Listeni/fræns/; French: [fʁɑ̃s], `+
												`officially the French Republic (French: République française [ʁepyblik fʁɑ̃sɛz]), `+
												`is a sovereign country in Western Europe that includes overseas `+
												`regions and territories.`,
										),
									),
								),
							),
						),
					),
				),
				JsSrc(cdn("//code.jquery.com/jquery-1.11.0.min.js")),
				JsSrc(cdn(bs3.CDN_3_1_1_js_min)),
			),
		),
	)
	http.ListenAndServe(":8585", nil)
}
