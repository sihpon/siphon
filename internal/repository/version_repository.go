package repository

import (
	"context"
	"errors"

	v1 "github.com/siphon/siphon/api/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type VersionRepository interface {
	CreateOrUpdate(ctx context.Context, version *v1.Version) error
	Find(ctx context.Context, id string) (*v1.Version, error)
	All(ctx context.Context) ([]*v1.Version, error)
	Delete(ctx context.Context, id string) error
}

type VersionRepositoryImpl struct {
	Client client.Client
}

func (v *VersionRepositoryImpl) CreateOrUpdate(ctx context.Context, version *v1.Version) error {
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

	siphon.Spec.Versions = append(siphon.Spec.Versions, *version)

	if notFounded {
		err = v.Client.Create(ctx, siphon)
	} else {
		err = v.Client.Update(ctx, siphon)
	}

	return err
}

func (v *VersionRepositoryImpl) All(ctx context.Context) ([]*v1.Version, error) {
	siphon := &v1.Siphon{}
	err := v.Client.Get(ctx, client.ObjectKey{Namespace: "default", Name: "siphon"}, siphon)
	if err != nil {
		return nil, err
	}

	var versions []*v1.Version
	for _, version := range siphon.Spec.Versions {
		versions = append(versions, &version)
	}

	return versions, nil
}

func (v *VersionRepositoryImpl) Find(ctx context.Context, id string) (*v1.Version, error) {
	siphon := &v1.Siphon{}
	err := v.Client.Get(ctx, client.ObjectKey{Namespace: "default", Name: "siphon"}, siphon)
	if err != nil {
		return nil, err
	}

	for _, version := range siphon.Spec.Versions {
		if version.ID == id {
			return &version, nil
		}
	}

	return nil, errors.New("version not found")
}

func (v *VersionRepositoryImpl) Delete(ctx context.Context, id string) error {
	siphon := &v1.Siphon{}
	err := v.Client.Get(ctx, client.ObjectKey{Namespace: "default", Name: "siphon"}, siphon)
	if err != nil {
		return err
	}

	for _, version := range siphon.Spec.Versions {
		if version.ID == id {
			version = v1.Version{}
		}
	}

	return v.Client.Update(ctx, siphon)
}
