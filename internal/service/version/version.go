package version

import (
	"context"

	versionv1 "github.com/siphon/siphon/generated/version/v1"
	"github.com/siphon/siphon/internal/model"
	"github.com/siphon/siphon/internal/repository"
	"gorm.io/gorm"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type VersionService struct {
	db     *gorm.DB
	Client client.Client
	Scheme *runtime.Scheme
}

func NewVersionService(db *gorm.DB, client client.Client, schema *runtime.Scheme) *VersionService {
	return &VersionService{
		db:     db,
		Client: client,
		Scheme: schema,
	}
}

var _ Versioner = &VersionService{}

// Create implements Versioner.
func (v *VersionService) Create(ctx context.Context, request *versionv1.CreateRequest) (*versionv1.CreateResponse, error) {
	repo := &repository.VersionRepositoryImpl{
		Client: v.Client,
	}

	err := repo.CreateOrUpdate(ctx, &model.Version{
		Version:     request.Version.Id,
		Description: request.Version.Description,
	})

	if err != nil {
		return nil, err
	}

	return &versionv1.CreateResponse{}, nil
}

// Delete implements Versioner.
func (v *VersionService) Delete(ctx context.Context, request *versionv1.DeleteRequest) (*versionv1.DeleteResponse, error) {
	repo := &repository.VersionRepositoryImpl{
		Client: v.Client,
	}

	err := repo.Delete(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return &versionv1.DeleteResponse{}, nil
}

// Get implements Versioner.
func (v *VersionService) Get(ctx context.Context, request *versionv1.GetRequest) (*versionv1.GetResponse, error) {
	repo := &repository.VersionRepositoryImpl{
		Client: v.Client,
	}

	version, err := repo.Find(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return &versionv1.GetResponse{
		Version: &versionv1.Version{
			Id:          version.Version,
			Description: version.Description,
		},
	}, nil
}

// List implements Versioner.
func (v *VersionService) List(ctx context.Context, request *versionv1.ListRequest) (*versionv1.ListResponse, error) {
	repo := &repository.VersionRepositoryImpl{
		Client: v.Client,
	}

	versions, err := repo.All(ctx)
	if err != nil {
		return nil, err
	}

	response := make([]*versionv1.Version, 0, len(versions))
	for _, version := range versions {
		response = append(response, &versionv1.Version{
			Id:          version.Version,
			Description: version.Description,
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
