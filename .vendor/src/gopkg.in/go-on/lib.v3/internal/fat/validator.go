package fat

import "fmt"

type Validater interface {
	Validate(*Field) error
}

type ValidaterFunc func(*Field) error

func (vf ValidaterFunc) Validate(f *Field) error {
	return vf(f)
}

func Validaters(vs ...Validater) func(f *Field) (errs []error) {
	return func(f *Field) (errs []error) {
		errs = []error{}
		for _, v := range vs {
			err := v.Validate(f)
			if err != nil {
				errs = append(errs, err)
			}
		}
		return
	}
}

var StringMustNotBeEmpty = ValidaterFunc(func(f *Field) (err error) {
	// fmt.Println("corporation validator called")
	if f.String() == "" {
		return fmt.Errorf("must not be empty")
	}
	return
})
