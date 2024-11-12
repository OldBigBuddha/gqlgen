// Code generated by github.com/OldBigBuddha/gqlgen, DO NOT EDIT.

package graph

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/OldBigBuddha/gqlgen/plugin/federation/fedruntime"
)

var (
	ErrUnknownType  = errors.New("unknown type")
	ErrTypeNotFound = errors.New("type not found")
)

func (ec *executionContext) __resolve__service(ctx context.Context) (fedruntime.Service, error) {
	if ec.DisableIntrospection {
		return fedruntime.Service{}, errors.New("federated introspection disabled")
	}

	var sdl []string

	for _, src := range sources {
		if src.BuiltIn {
			continue
		}
		sdl = append(sdl, src.Input)
	}

	return fedruntime.Service{
		SDL: strings.Join(sdl, "\n"),
	}, nil
}

func (ec *executionContext) __resolve_entities(ctx context.Context, representations []map[string]interface{}) []fedruntime.Entity {
	list := make([]fedruntime.Entity, len(representations))

	repsMap := ec.buildRepresentationGroups(ctx, representations)

	switch len(repsMap) {
	case 0:
		return list
	case 1:
		for typeName, reps := range repsMap {
			ec.resolveEntityGroup(ctx, typeName, reps, list)
		}
		return list
	default:
		var g sync.WaitGroup
		g.Add(len(repsMap))
		for typeName, reps := range repsMap {
			go func(typeName string, reps []EntityWithIndex) {
				ec.resolveEntityGroup(ctx, typeName, reps, list)
				g.Done()
			}(typeName, reps)
		}
		g.Wait()
		return list
	}
}

type EntityWithIndex struct {
	// The index in the original representation array
	index  int
	entity EntityRepresentation
}

// EntityRepresentation is the JSON representation of an entity sent by the Router
// used as the inputs for us to resolve.
//
// We make it a map because we know the top level JSON is always an object.
type EntityRepresentation map[string]any

// We group entities by typename so that we can parallelize their resolution.
// This is particularly helpful when there are entity groups in multi mode.
func (ec *executionContext) buildRepresentationGroups(
	ctx context.Context,
	representations []map[string]any,
) map[string][]EntityWithIndex {
	repsMap := make(map[string][]EntityWithIndex)
	for i, rep := range representations {
		typeName, ok := rep["__typename"].(string)
		if !ok {
			// If there is no __typename, we just skip the representation;
			// we just won't be resolving these unknown types.
			ec.Error(ctx, errors.New("__typename must be an existing string"))
			continue
		}

		repsMap[typeName] = append(repsMap[typeName], EntityWithIndex{
			index:  i,
			entity: rep,
		})
	}

	return repsMap
}

func (ec *executionContext) resolveEntityGroup(
	ctx context.Context,
	typeName string,
	reps []EntityWithIndex,
	list []fedruntime.Entity,
) {
	if isMulti(typeName) {
		err := ec.resolveManyEntities(ctx, typeName, reps, list)
		if err != nil {
			ec.Error(ctx, err)
		}
	} else {
		// if there are multiple entities to resolve, parallelize (similar to
		// graphql.FieldSet.Dispatch)
		var e sync.WaitGroup
		e.Add(len(reps))
		for i, rep := range reps {
			i, rep := i, rep
			go func(i int, rep EntityWithIndex) {
				entity, err := ec.resolveEntity(ctx, typeName, rep.entity)
				if err != nil {
					ec.Error(ctx, err)
				} else {
					list[rep.index] = entity
				}
				e.Done()
			}(i, rep)
		}
		e.Wait()
	}
}

func isMulti(typeName string) bool {
	switch typeName {
	default:
		return false
	}
}

func (ec *executionContext) resolveEntity(
	ctx context.Context,
	typeName string,
	rep EntityRepresentation,
) (e fedruntime.Entity, err error) {
	// we need to do our own panic handling, because we may be called in a
	// goroutine, where the usual panic handling can't catch us
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
		}
	}()

	switch typeName {
	case "EmailHost":
		resolverName, err := entityResolverNameForEmailHost(ctx, rep)
		if err != nil {
			return nil, fmt.Errorf(`finding resolver for Entity "EmailHost": %w`, err)
		}
		switch resolverName {

		case "findEmailHostByID":
			id0, err := ec.unmarshalNString2string(ctx, rep["id"])
			if err != nil {
				return nil, fmt.Errorf(`unmarshalling param 0 for findEmailHostByID(): %w`, err)
			}
			entity, err := ec.resolvers.Entity().FindEmailHostByID(ctx, id0)
			if err != nil {
				return nil, fmt.Errorf(`resolving Entity "EmailHost": %w`, err)
			}

			return entity, nil
		}
	case "User":
		resolverName, err := entityResolverNameForUser(ctx, rep)
		if err != nil {
			return nil, fmt.Errorf(`finding resolver for Entity "User": %w`, err)
		}
		switch resolverName {

		case "findUserByID":
			id0, err := ec.unmarshalNID2string(ctx, rep["id"])
			if err != nil {
				return nil, fmt.Errorf(`unmarshalling param 0 for findUserByID(): %w`, err)
			}
			entity, err := ec.resolvers.Entity().FindUserByID(ctx, id0)
			if err != nil {
				return nil, fmt.Errorf(`resolving Entity "User": %w`, err)
			}

			return entity, nil
		}

	}
	return nil, fmt.Errorf("%w: %s", ErrUnknownType, typeName)
}

func (ec *executionContext) resolveManyEntities(
	ctx context.Context,
	typeName string,
	reps []EntityWithIndex,
	list []fedruntime.Entity,
) (err error) {
	// we need to do our own panic handling, because we may be called in a
	// goroutine, where the usual panic handling can't catch us
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
		}
	}()

	switch typeName {

	default:
		return errors.New("unknown type: " + typeName)
	}
}

func entityResolverNameForEmailHost(ctx context.Context, rep EntityRepresentation) (string, error) {
	for {
		var (
			m   EntityRepresentation
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["id"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findEmailHostByID", nil
	}
	return "", fmt.Errorf("%w for EmailHost", ErrTypeNotFound)
}

func entityResolverNameForUser(ctx context.Context, rep EntityRepresentation) (string, error) {
	for {
		var (
			m   EntityRepresentation
			val interface{}
			ok  bool
		)
		_ = val
		// if all of the KeyFields values for this resolver are null,
		// we shouldn't use use it
		allNull := true
		m = rep
		val, ok = m["id"]
		if !ok {
			break
		}
		if allNull {
			allNull = val == nil
		}
		if allNull {
			break
		}
		return "findUserByID", nil
	}
	return "", fmt.Errorf("%w for User", ErrTypeNotFound)
}
