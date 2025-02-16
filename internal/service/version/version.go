package version

import (
	"context"

	v1 "github.com/siphon/siphon/api/v1"
	versionv1 "github.com/siphon/siphon/generated/version/v1"
	"github.com/siphon/siphon/internal/container"
	"github.com/siphon/siphon/internal/repository"
)

type VersionService struct {
	container container.Containerer
}

func NewVersionService(container container.Containerer) *VersionService {
	return &VersionService{
		container: container,
	}
}

var _ Versioner = &VersionService{}

// Create implements Versioner.
func (v *VersionService) Create(ctx context.Context, request *versionv1.CreateRequest) (*versionv1.CreateResponse, error) {
	repo := &repository.VersionRepositoryImpl{
		Client: v.container.KubeClient(),
	}

	err := repo.CreateOrUpdate(ctx, &v1.Version{
		ID: request.Version.Id,
	})

	if err != nil {
		return nil, err
	}

	return &versionv1.CreateResponse{}, nil
}

// Delete implements Versioner.
func (v *VersionService) Delete(ctx context.Context, request *versionv1.DeleteRequest) (*versionv1.DeleteResponse, error) {
	repo := &repository.VersionRepositoryImpl{
		Client: v.container.KubeClient(),
	}

	err := repo.Delete(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return &versionv1.DeleteResponse{}, nil
}

// Get implements Versioner.
func (v *VersionService) Get(ctx context.Context, request *versionv1.GetRequest) (*versionv1.GetResponse, error) {
	repo := &repository.VersionRepositoryImpl{Client: v.container.KubeClient()}

	version, err := repo.Find(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return &versionv1.GetResponse{
		Version: &versionv1.Version{
			Id: version.ID,
		},
	}, nil
}

// List implements Versioner.
func (v *VersionService) List(ctx context.Context, request *versionv1.ListRequest) (*versionv1.ListResponse, error) {
	repo := &repository.VersionRepositoryImpl{
		Client: v.container.KubeClient(),
	}

	versions, err := repo.All(ctx)
	if err != nil {
		return nil, err
	}

	response := make([]*versionv1.Version, 0, len(versions))
	for _, version := range versions {
		response = append(response, &versionv1.Version{
			Id: version.ID,
		})
	}

	return &versionv1.ListResponse{
		Versions: response,
	}, nil
}

// Update implements Versioner.
func (v *VersionService) Update(ctx context.Context, request *versionv1.UpdateRequest) (*versionv1.UpdateResponse, error) {
	panic("unimplemented")
}
