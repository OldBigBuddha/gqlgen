// Code generated by github.com/OldBigBuddha/gqlgen, DO NOT EDIT.

package followschema

import (
	"context"

	"github.com/OldBigBuddha/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputInputWithEnumValue(ctx context.Context, obj interface{}) (InputWithEnumValue, error) {
	var it InputWithEnumValue
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"enum"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "enum":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("enum"))
			data, err := ec.unmarshalNEnumTest2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐEnumTest(ctx, v)
			if err != nil {
				return it, err
			}
			it.Enum = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNEnumTest2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐEnumTest(ctx context.Context, v interface{}) (EnumTest, error) {
	var res EnumTest
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNEnumTest2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐEnumTest(ctx context.Context, sel ast.SelectionSet, v EnumTest) graphql.Marshaler {
	return v
}

func (ec *executionContext) unmarshalOInputWithEnumValue2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐInputWithEnumValue(ctx context.Context, v interface{}) (*InputWithEnumValue, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputInputWithEnumValue(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

// endregion ***************************** type.gotpl *****************************
