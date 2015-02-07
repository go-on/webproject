package html

import (
	"testing"
)

func TestAttributes(t *testing.T) {

	tests := map[string]string{
		Rel_("x").String():        `rel="x"`,
		Href_("x").String():       `href="x"`,
		Type_("x").String():       `type="x"`,
		Media_("x").String():      `media="x"`,
		Src_("x").String():        `src="x"`,
		Name_("x").String():       `name="x"`,
		Content_("x").String():    `content="x"`,
		Charset_("x").String():    `charset="x"`,
		CharsetUtf8_().String():   `charset="utf-8"`,
		Lang_("x").String():       `lang="x"`,
		Role_("x").String():       `role="x"`,
		Value_("x").String():      `value="x"`,
		Alt_("x").String():        `alt="x"`,
		Title_("x").String():      `title="x"`,
		Method_("x").String():     `method="x"`,
		Action_("x").String():     `action="x"`,
		For_("x").String():        `for="x"`,
		Width_("x").String():      `width="x"`,
		Height_("x").String():     `height="x"`,
		OnSubmit_("x").String():   `onsubmit="javascript:x"`,
		OnClick_("x").String():    `onclick="javascript:x"`,
		Enctype_("x").String():    `enctype="x"`,
		Target_("x").String():     `target="x"`,
		DataToggle_("x").String(): `data-toggle="x"`,
		DataTarget_("x").String(): `data-target="x"`,
		DataId_("x").String():     `data-id="x"`,
		Style_("x").String():      `style="x"`,
		About_("x").String():      `about="x"`,
		TypeOf_("x").String():     `typeof="x"`,
		Property_("x").String():   `property="x"`,
	}

	for got, expected := range tests {
		if got != expected {
			t.Errorf("got: %#v expected: %#v", got, expected)
		}
	}

}
