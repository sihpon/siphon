package version

import (
	"context"
	"fmt"
	"log"

	v1 "github.com/siphon/siphon/api/v1"
	versionv1 "github.com/siphon/siphon/generated/version/v1"
	"gorm.io/gorm"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	log.Println("Creating version")
	siphon := v1.SiphonSpec{
		Foo: "fooooo",
		Versions: map[string]string{
			"v1.2.0": "1.2.0 開発",
		},
		Labels: map[string]string{
			"dev": "普通の開発環境",
		},
	}

	deploymentStore := &v1.Siphon{ObjectMeta: metav1.ObjectMeta{
		Name:      "siphon",
		Namespace: "default",
	}}
	deploymentStore.Spec = siphon

	if err := v.Client.Create(ctx, deploymentStore); err != nil {
		log.Println(err)
		return &versionv1.CreateResponse{}, err
	}

	return &versionv1.CreateResponse{}, nil
}

// Delete implements Versioner.
func (v *VersionService) Delete(ctx context.Context, request *versionv1.DeleteRequest) (*versionv1.DeleteResponse, error) {
	panic("unimplemented")
}

// Get implements Versioner.
func (v *VersionService) Get(ctx context.Context, request *versionv1.GetRequest) (*versionv1.GetResponse, error) {
	siphon := &v1.Siphon{}
	if err := v.Client.Get(ctx, client.ObjectKey{Namespace: "default", Name: "siphon"}, siphon); err != nil {
		return nil, err
	}
	version, ok := siphon.Spec.Versions[request.Id]
	if !ok {
		return nil, fmt.Errorf("version %s not found", request.Id)
	}

	return &versionv1.GetResponse{
		Version: &versionv1.Version{
			Id:          request.Id,
			Description: version,
		},
	}, nil
}

// List implements Versioner.
func (v *VersionService) List(ctx context.Context, request *versionv1.ListRequest) (*versionv1.ListResponse, error) {
	siphon := &v1.Siphon{}
	if err := v.Client.Get(ctx, client.ObjectKey{Namespace: "default", Name: "siphon"}, siphon); err != nil {
		log.Println(err)
		return nil, err
	}

	response := make([]*versionv1.Version, 0, len(siphon.Spec.Versions))
	for key, value := range siphon.Spec.Versions {
		response = append(response, &versionv1.Version{
			Id:   key,
			Name: value,
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
