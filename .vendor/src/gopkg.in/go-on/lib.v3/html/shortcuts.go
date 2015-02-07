package html

import (
	. "gopkg.in/go-on/lib.v3/html/internal/element"
	"gopkg.in/go-on/lib.v3/types"
)

func JsSrc(url string, objects ...interface{}) *Element {
	params := []interface{}{Type_("text/javascript"), Src_(url)}
	return SCRIPT(append(params, objects...)...)
}

func CssHref(url string, objects ...interface{}) *Element {
	params := []interface{}{Rel_("stylesheet"), Type_("text/css"), Href_(url)}
	return LINK(append(params, objects...)...)
}

func FormGet(action string, objects ...interface{}) *Element {
	params := []interface{}{Method_("get"), Action_(action)}
	return FORM(append(params, objects...)...)
}

func FormPost(action string, objects ...interface{}) *Element {
	params := []interface{}{Method_("post"), Action_(action)}
	return FORM(append(params, objects...)...)
}

//t.Add(Attr("enctype", "multipart/form-data", "method", "post"))
func FormPostMultipart(action string, objects ...interface{}) *Element {
	//t.Add(Attr("enctype", "multipart/form-data", "method", "post"))
	params := []interface{}{Method_("post"), Action_(action), MultiPart_}
	return FORM(append(params, objects...)...)
}

func FormPut(action string, objects ...interface{}) *Element {
	params := []interface{}{
		Method_("post"),
		Action_(action),
		INPUT(
			Name_("_method"),
			Value_("PUT"),
			Type_("hidden"))}
	return FORM(append(params, objects...)...)
}

func FormPatch(action string, objects ...interface{}) *Element {
	params := []interface{}{
		Method_("post"),
		Action_(action),
		INPUT(
			Name_("_method"),
			Value_("PATCH"),
			Type_("hidden"))}
	return FORM(append(params, objects...)...)
}

func FormDelete(action string, objects ...interface{}) *Element {
	params := []interface{}{
		Method_("post"),
		Action_(action),
		INPUT(
			Name_("_method"),
			Value_("DELETE"),
			Type_("hidden"))}
	return FORM(append(params, objects...)...)
}

func inputType(typ string, name string, objects ...interface{}) *Element {
	params := []interface{}{
		Type_(typ),
		Name_(name),
	}
	return INPUT(append(params, objects...)...)
}

func InputHidden(name string, objects ...interface{}) *Element {
	return inputType("hidden", name, objects...)
}

func InputSubmit(name string, objects ...interface{}) *Element {
	return inputType("submit", name, objects...)
}

func InputText(name string, objects ...interface{}) *Element {
	return inputType("text", name, objects...)
}

func InputButton(name string, objects ...interface{}) *Element {
	return inputType("button", name, objects...)
}

func InputPassword(name string, objects ...interface{}) *Element {
	return inputType("password", name, objects...)
}

func InputRadio(name string, objects ...interface{}) *Element {
	return inputType("radio", name, objects...)
}

func InputCheckbox(name string, objects ...interface{}) *Element {
	return inputType("checkbox", name, objects...)
}

func InputFile(name string, objects ...interface{}) *Element {
	return inputType("file", name, objects...)
}

func AHref(url string, objects ...interface{}) *Element {
	params := []interface{}{Href_(url)}
	return A(append(params, objects...)...)
}

func ImgSrc(src string, objects ...interface{}) *Element {
	params := []interface{}{Src_(src)}
	return IMG(append(params, objects...)...)
}

func LabelFor(for_ string, objects ...interface{}) *Element {
	params := []interface{}{For_(for_)}
	return LABEL(append(params, objects...)...)
}

func Charset(charset string, objects ...interface{}) *Element {
	params := []interface{}{types.Attribute{"charset", charset}}
	return META(append(params, objects...)...)
}

func CharsetUtf8(objects ...interface{}) *Element {
	return Charset("utf-8", objects...)
}

func HttpEquiv(http_equiv string, content string, objects ...interface{}) *Element {
	params := []interface{}{
		types.Attribute{"http-equiv", http_equiv},
		types.Attribute{"content", content},
	}
	return META(append(params, objects...)...)
}

func HttpEquivUtf8(objects ...interface{}) *Element {
	return HttpEquiv("Content-Type", "text/html;charset=utf-8", objects...)
}

func Viewport(content string, objects ...interface{}) *Element {
	params := []interface{}{
		types.Attribute{"name", "viewport"},
		types.Attribute{"content", content},
	}
	return META(append(params, objects...)...)
}

func Canonical(url string, objects ...interface{}) *Element {
	params := []interface{}{
		Rel_("canonical"),
		Href_(url),
	}
	return LINK(append(params, objects...)...)
}
