package version

import (
	"context"

	versionv1 "github.com/siphon/siphon/generated/version/v1"
	"github.com/siphon/siphon/internal/model"
	"gorm.io/gorm"
)

type VersionService struct {
	db *gorm.DB
}

func NewVersionService(db *gorm.DB) *VersionService {
	return &VersionService{db: db}
}

var _ Versioner = &VersionService{}

// Create implements Versioner.
func (v *VersionService) Create(ctx context.Context, request *versionv1.CreateRequest) (*versionv1.CreateResponse, error) {
	panic("unimplemented")
}

// Delete implements Versioner.
func (v *VersionService) Delete(ctx context.Context, request *versionv1.DeleteRequest) (*versionv1.DeleteResponse, error) {
	panic("unimplemented")
}

// Get implements Versioner.
func (v *VersionService) Get(ctx context.Context, request *versionv1.GetRequest) (*versionv1.GetResponse, error) {
	panic("unimplemented")
}

// List implements Versioner.
func (v *VersionService) List(ctx context.Context, request *versionv1.ListRequest) (*versionv1.ListResponse, error) {
	versions := []*model.Version{}
	if err := v.db.Find(&versions).Error; err != nil {
		return nil, err
	}

	response := make([]*versionv1.Version, 0, len(versions))
	for _, version := range versions {
		response = append(response, &versionv1.Version{
			Id:          version.ID,
			Name:        version.Name,
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
