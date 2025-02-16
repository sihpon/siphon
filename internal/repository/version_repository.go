package repository

import (
	"context"
	"errors"

	v1 "github.com/siphon/siphon/api/v1"
	"github.com/siphon/siphon/internal/model"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type VersionRepository interface {
	CreateOrUpdate(ctx context.Context, version *model.Version) error
	Find(ctx context.Context, id string) (*model.Version, error)
	All(ctx context.Context) ([]*model.Version, error)
	Delete(ctx context.Context, id string) error
}

type VersionRepositoryImpl struct {
	Client client.Client
}

func (v *VersionRepositoryImpl) CreateOrUpdate(ctx context.Context, version *model.Version) error {
	siphon := &v1.Siphon{}
	err := v.Client.Get(ctx, client.ObjectKey{Namespace: "default", Name: "siphon"}, siphon)
	if client.IgnoreNotFound(err) != nil {
		return err
	}

	notFounded := false
	if apierrors.IsNotFound(err) {
		siphon = &v1.Siphon{ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "siphon"}}
		notFounded = true
	}

	siphon.Spec.Versions[version.Version] = version.Description

	if notFounded {
		err = v.Client.Create(ctx, siphon)
	} else {
		err = v.Client.Update(ctx, siphon)
	}

	return err
}

func (v *VersionRepositoryImpl) All(ctx context.Context) ([]*model.Version, error) {
	siphon := &v1.Siphon{}
	err := v.Client.Get(ctx, client.ObjectKey{Namespace: "default", Name: "siphon"}, siphon)
	if err != nil {
		return nil, err
	}

	var versions []*model.Version
	for version, description := range siphon.Spec.Versions {
		versions = append(versions, &model.Version{
			Version:     version,
			Description: description,
		})
	}

	return versions, nil
}

func (v *VersionRepositoryImpl) Find(ctx context.Context, id string) (*model.Version, error) {
	siphon := &v1.Siphon{}
	err := v.Client.Get(ctx, client.ObjectKey{Namespace: "default", Name: "siphon"}, siphon)
	if err != nil {
		return nil, err
	}

	description, ok := siphon.Spec.Versions[id]
	if !ok {
		return nil, errors.New("version not found")
	}

	return &model.Version{
		Version:     id,
		Description: description,
	}, nil
}

func (v *VersionRepositoryImpl) Delete(ctx context.Context, id string) error {
	siphon := &v1.Siphon{}
	err := v.Client.Get(ctx, client.ObjectKey{Namespace: "default", Name: "siphon"}, siphon)
	if err != nil {
		return err
	}

	delete(siphon.Spec.Versions, id)

	return v.Client.Update(ctx, siphon)
}
