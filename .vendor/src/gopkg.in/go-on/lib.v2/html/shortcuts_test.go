package html

import (
	"testing"
)

func TestShortCuts(t *testing.T) {

	tests := map[string]string{
		AHref("#", "t").String():                        `<a href="#">t</a>`,
		JsSrc("j.js").String():                          `<script type="text/javascript" src="j.js"></script>`,
		CssHref("c.css").String():                       `<link rel="stylesheet" type="text/css" href="c.css" />`,
		ImgSrc("i.jpg").String():                        `<img src="i.jpg" />`,
		LabelFor("x", "y").String():                     `<label for="x">y</label>`,
		Charset("iso 8859-1").String():                  `<meta charset="iso 8859-1" />`,
		CharsetUtf8().String():                          `<meta charset="utf-8" />`,
		HttpEquivUtf8().String():                        `<meta http-equiv="Content-Type" content="text/html;charset=utf-8" />`,
		HttpEquiv("Content-Type", "text/html").String(): `<meta http-equiv="Content-Type" content="text/html" />`,
		Viewport("x").String():                          `<meta name="viewport" content="x" />`,
		FormGet("x", "y").String():                      `<form method="get" action="x">y</form>`,
		FormPost("x", "y").String():                     `<form method="post" action="x">y</form>`,
		FormPostMultipart("x", "y").String():            `<form method="post" action="x" enctype="multipart/form-data">y</form>`,
		FormPut("x", "y").String():                      `<form method="post" action="x"><input name="_method" value="PUT" type="hidden" />y</form>`,
		FormPatch("x", "y").String():                    `<form method="post" action="x"><input name="_method" value="PATCH" type="hidden" />y</form>`,
		FormDelete("x", "y").String():                   `<form method="post" action="x"><input name="_method" value="DELETE" type="hidden" />y</form>`,
		InputHidden("x").String():                       `<input type="hidden" name="x" />`,
		InputSubmit("x").String():                       `<input type="submit" name="x" />`,
		InputText("x").String():                         `<input type="text" name="x" />`,
		InputButton("x").String():                       `<input type="button" name="x" />`,
		InputPassword("x").String():                     `<input type="password" name="x" />`,
		InputFile("x").String():                         `<input type="file" name="x" />`,
		InputRadio("x").String():                        `<input type="radio" name="x" />`,
		InputCheckbox("x").String():                     `<input type="checkbox" name="x" />`,
	}

	for got, expected := range tests {
		if got != expected {
			t.Errorf("got: %#v expected: %#v", got, expected)
		}
	}

}
