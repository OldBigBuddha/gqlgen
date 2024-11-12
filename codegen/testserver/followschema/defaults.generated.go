// Code generated by github.com/OldBigBuddha/gqlgen, DO NOT EDIT.

package followschema

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/OldBigBuddha/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type MutationResolver interface {
	DefaultInput(ctx context.Context, input DefaultInput) (*DefaultParametersMirror, error)
	OverrideValueViaInput(ctx context.Context, input FieldsOrderInput) (*FieldsOrderPayload, error)
	UpdateSomething(ctx context.Context, input SpecialInput) (string, error)
	UpdatePtrToPtr(ctx context.Context, input UpdatePtrToPtrOuter) (*PtrToPtrOuter, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_Mutation_defaultInput_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	arg0, err := ec.field_Mutation_defaultInput_argsInput(ctx, rawArgs)
	if err != nil {
		return nil, err
	}
	args["input"] = arg0
	return args, nil
}
func (ec *executionContext) field_Mutation_defaultInput_argsInput(
	ctx context.Context,
	rawArgs map[string]interface{},
) (DefaultInput, error) {
	// We won't call the directive if the argument is null.
	// Set call_argument_directives_with_null to true to call directives
	// even if the argument is null.
	_, ok := rawArgs["input"]
	if !ok {
		var zeroVal DefaultInput
		return zeroVal, nil
	}

	ctx = graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
	if tmp, ok := rawArgs["input"]; ok {
		return ec.unmarshalNDefaultInput2githubᚗcomᚋOldBigBuddhaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐDefaultInput(ctx, tmp)
	}

	var zeroVal DefaultInput
	return zeroVal, nil
}

func (ec *executionContext) field_Mutation_overrideValueViaInput_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	arg0, err := ec.field_Mutation_overrideValueViaInput_argsInput(ctx, rawArgs)
	if err != nil {
		return nil, err
	}
	args["input"] = arg0
	return args, nil
}
func (ec *executionContext) field_Mutation_overrideValueViaInput_argsInput(
	ctx context.Context,
	rawArgs map[string]interface{},
) (FieldsOrderInput, error) {
	// We won't call the directive if the argument is null.
	// Set call_argument_directives_with_null to true to call directives
	// even if the argument is null.
	_, ok := rawArgs["input"]
	if !ok {
		var zeroVal FieldsOrderInput
		return zeroVal, nil
	}

	ctx = graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
	if tmp, ok := rawArgs["input"]; ok {
		return ec.unmarshalNFieldsOrderInput2githubᚗcomᚋOldBigBuddhaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐFieldsOrderInput(ctx, tmp)
	}

	var zeroVal FieldsOrderInput
	return zeroVal, nil
}

func (ec *executionContext) field_Mutation_updatePtrToPtr_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	arg0, err := ec.field_Mutation_updatePtrToPtr_argsInput(ctx, rawArgs)
	if err != nil {
		return nil, err
	}
	args["input"] = arg0
	return args, nil
}
func (ec *executionContext) field_Mutation_updatePtrToPtr_argsInput(
	ctx context.Context,
	rawArgs map[string]interface{},
) (UpdatePtrToPtrOuter, error) {
	// We won't call the directive if the argument is null.
	// Set call_argument_directives_with_null to true to call directives
	// even if the argument is null.
	_, ok := rawArgs["input"]
	if !ok {
		var zeroVal UpdatePtrToPtrOuter
		return zeroVal, nil
	}

	ctx = graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
	if tmp, ok := rawArgs["input"]; ok {
		return ec.unmarshalNUpdatePtrToPtrOuter2githubᚗcomᚋOldBigBuddhaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐUpdatePtrToPtrOuter(ctx, tmp)
	}

	var zeroVal UpdatePtrToPtrOuter
	return zeroVal, nil
}

func (ec *executionContext) field_Mutation_updateSomething_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	arg0, err := ec.field_Mutation_updateSomething_argsInput(ctx, rawArgs)
	if err != nil {
		return nil, err
	}
	args["input"] = arg0
	return args, nil
}
func (ec *executionContext) field_Mutation_updateSomething_argsInput(
	ctx context.Context,
	rawArgs map[string]interface{},
) (SpecialInput, error) {
	// We won't call the directive if the argument is null.
	// Set call_argument_directives_with_null to true to call directives
	// even if the argument is null.
	_, ok := rawArgs["input"]
	if !ok {
		var zeroVal SpecialInput
		return zeroVal, nil
	}

	ctx = graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
	if tmp, ok := rawArgs["input"]; ok {
		return ec.unmarshalNSpecialInput2githubᚗcomᚋOldBigBuddhaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐSpecialInput(ctx, tmp)
	}

	var zeroVal SpecialInput
	return zeroVal, nil
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _DefaultParametersMirror_falsyBoolean(ctx context.Context, field graphql.CollectedField, obj *DefaultParametersMirror) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_DefaultParametersMirror_falsyBoolean(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.FalsyBoolean, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*bool)
	fc.Result = res
	return ec.marshalOBoolean2ᚖbool(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_DefaultParametersMirror_falsyBoolean(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "DefaultParametersMirror",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _DefaultParametersMirror_truthyBoolean(ctx context.Context, field graphql.CollectedField, obj *DefaultParametersMirror) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_DefaultParametersMirror_truthyBoolean(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.TruthyBoolean, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*bool)
	fc.Result = res
	return ec.marshalOBoolean2ᚖbool(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_DefaultParametersMirror_truthyBoolean(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "DefaultParametersMirror",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_defaultInput(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_defaultInput(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().DefaultInput(rctx, fc.Args["input"].(DefaultInput))
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*DefaultParametersMirror)
	fc.Result = res
	return ec.marshalNDefaultParametersMirror2ᚖgithubᚗcomᚋOldBigBuddhaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐDefaultParametersMirror(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_defaultInput(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "falsyBoolean":
				return ec.fieldContext_DefaultParametersMirror_falsyBoolean(ctx, field)
			case "truthyBoolean":
				return ec.fieldContext_DefaultParametersMirror_truthyBoolean(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type DefaultParametersMirror", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_defaultInput_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_overrideValueViaInput(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_overrideValueViaInput(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().OverrideValueViaInput(rctx, fc.Args["input"].(FieldsOrderInput))
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*FieldsOrderPayload)
	fc.Result = res
	return ec.marshalNFieldsOrderPayload2ᚖgithubᚗcomᚋOldBigBuddhaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐFieldsOrderPayload(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_overrideValueViaInput(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "firstFieldValue":
				return ec.fieldContext_FieldsOrderPayload_firstFieldValue(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type FieldsOrderPayload", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_overrideValueViaInput_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_updateSomething(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_updateSomething(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().UpdateSomething(rctx, fc.Args["input"].(SpecialInput))
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_updateSomething(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_updateSomething_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_updatePtrToPtr(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_updatePtrToPtr(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().UpdatePtrToPtr(rctx, fc.Args["input"].(UpdatePtrToPtrOuter))
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*PtrToPtrOuter)
	fc.Result = res
	return ec.marshalNPtrToPtrOuter2ᚖgithubᚗcomᚋOldBigBuddhaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐPtrToPtrOuter(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_updatePtrToPtr(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "name":
				return ec.fieldContext_PtrToPtrOuter_name(ctx, field)
			case "inner":
				return ec.fieldContext_PtrToPtrOuter_inner(ctx, field)
			case "stupidInner":
				return ec.fieldContext_PtrToPtrOuter_stupidInner(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type PtrToPtrOuter", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_updatePtrToPtr_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputDefaultInput(ctx context.Context, obj interface{}) (DefaultInput, error) {
	var it DefaultInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	if _, present := asMap["falsyBoolean"]; !present {
		asMap["falsyBoolean"] = false
	}
	if _, present := asMap["truthyBoolean"]; !present {
		asMap["truthyBoolean"] = true
	}

	fieldsInOrder := [...]string{"falsyBoolean", "truthyBoolean"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "falsyBoolean":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("falsyBoolean"))
			data, err := ec.unmarshalOBoolean2ᚖbool(ctx, v)
			if err != nil {
				return it, err
			}
			it.FalsyBoolean = data
		case "truthyBoolean":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("truthyBoolean"))
			data, err := ec.unmarshalOBoolean2ᚖbool(ctx, v)
			if err != nil {
				return it, err
			}
			it.TruthyBoolean = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var defaultParametersMirrorImplementors = []string{"DefaultParametersMirror"}

func (ec *executionContext) _DefaultParametersMirror(ctx context.Context, sel ast.SelectionSet, obj *DefaultParametersMirror) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, defaultParametersMirrorImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("DefaultParametersMirror")
		case "falsyBoolean":
			out.Values[i] = ec._DefaultParametersMirror_falsyBoolean(ctx, field, obj)
		case "truthyBoolean":
			out.Values[i] = ec._DefaultParametersMirror_truthyBoolean(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var mutationImplementors = []string{"Mutation"}

func (ec *executionContext) _Mutation(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, mutationImplementors)
	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "Mutation",
	})

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		innerCtx := graphql.WithRootFieldContext(ctx, &graphql.RootFieldContext{
			Object: field.Name,
			Field:  field,
		})

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Mutation")
		case "defaultInput":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_defaultInput(ctx, field)
			})
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "overrideValueViaInput":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_overrideValueViaInput(ctx, field)
			})
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "updateSomething":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_updateSomething(ctx, field)
			})
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "updatePtrToPtr":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_updatePtrToPtr(ctx, field)
			})
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNDefaultInput2githubᚗcomᚋOldBigBuddhaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐDefaultInput(ctx context.Context, v interface{}) (DefaultInput, error) {
	res, err := ec.unmarshalInputDefaultInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNDefaultParametersMirror2githubᚗcomᚋOldBigBuddhaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐDefaultParametersMirror(ctx context.Context, sel ast.SelectionSet, v DefaultParametersMirror) graphql.Marshaler {
	return ec._DefaultParametersMirror(ctx, sel, &v)
}

func (ec *executionContext) marshalNDefaultParametersMirror2ᚖgithubᚗcomᚋOldBigBuddhaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐDefaultParametersMirror(ctx context.Context, sel ast.SelectionSet, v *DefaultParametersMirror) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._DefaultParametersMirror(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
