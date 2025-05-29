# Validation - Mitz IT

Standard contracts to create validators instances on top of [ozzo-validation](https://github.com/go-ozzo/ozzo-validation).

## Installation

```bash
go get -u github.com/payly-solucoes-de-pagamentos/golang-validation
```

## Usage

Implement the `IValidator` interface:

```go
package customvalidators

import (
  validation "github.com/payly-solucoes-de-pagamentos/golang-validation"
  ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

type CustomValidator struct {
}

func (validator *CustomValidator) Validate(instance interface{}) error {
  // validation logic goes here...
  
  return err
}

func NewCustomValidator() validation.IValidator {
  return &CustomValidator{}
}
```

Use the validator:

```go
// ...
var validator validation.IValidator
validator = customvalidators.NewCustomValidator()
err := validator.Validate(command)
// ...
```

## Utilities

`golang-validation` supports _Nested_ validations for hierarchical structs, and a __default__ `StringRule` array which states that strings are __required__ and must length between 1 and 200 charcaters.

```go
package customvalidators

import (
  validation "github.com/payly-solucoes-de-pagamentos/golang-validation"
  ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

type CustomValidator struct {
}

func (validator *CustomValidator) Validate(instance interface{}) error {
  nestedStruct := instance.(mystructs.NestedStruct)

  err := ozzo.ValidateStruct(&nestedStruct,
      ozzo.Field(&nestedStruct.Id, ozzo.Required),
      ozzo.Field(&nestedStruct.Name, validation.StringRule...),
      ozzo.Field(&nestedStruct.Data, ozzo.Required),
      validation.Nested(&nestedStruct.Data,
        ozzo.Field(&nestedStruct.Data.Label, validation.StringRule...),
        ozzo.Field(&nestedStruct.Data.Parameters, ozzo.Required),
      ),
    )
  
  return err
}

func NewCustomValidator() validation.IValidator {
  return &CustomValidator{}
}
```
