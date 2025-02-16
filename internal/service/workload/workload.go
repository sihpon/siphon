package workload

import (
	"context"
	"fmt"
	"strings"

	v1 "github.com/siphon/siphon/api/v1"
	workloadv1 "github.com/siphon/siphon/generated/workload/v1"
	"github.com/siphon/siphon/internal/model"
	"github.com/siphon/siphon/internal/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type WorkloadService struct {
	db     *gorm.DB
	Client client.Client
}

func NewWorkloadService(
	db *gorm.DB,
	client client.Client,
) *WorkloadService {
	return &WorkloadService{
		db:     db,
		Client: client,
	}
}

var _ Workloader = &WorkloadService{}

func (s *WorkloadService) List(ctx context.Context, req *workloadv1.ListRequest) (*workloadv1.ListResponse, error) {
	repo := &repository.WorkloadRepositoryImpl{Client: s.Client}
	workloads, err := repo.All(ctx)
	if err != nil {
		return nil, err
	}

	response := make([]*workloadv1.Workload, 0, len(workloads))
	for _, workload := range workloads {
		response = append(response, &workloadv1.Workload{
			Id:          workload.Spec.ID,
			Name:        workload.Spec.Name,
			Description: workload.Spec.Description,
			Version:     workload.Spec.Version,
			CreatedAt:   timestamppb.New(workload.Spec.CreatedAt.Time),
			UpdatedAt:   timestamppb.New(workload.Spec.UpdatedAt.Time),
		})
	}

	return &workloadv1.ListResponse{
		Workloads: response,
	}, nil
}

func (s *WorkloadService) Get(ctx context.Context, request *workloadv1.GetRequest) (*workloadv1.GetResponse, error) {
	workload := &v1.SiphonManagedWorkload{}
	if err := s.Client.Get(ctx, client.ObjectKey{Namespace: "default", Name: request.Id}, workload); err != nil {
		return nil, err
	}

	response := &workloadv1.Workload{
		Id:          workload.Spec.ID,
		Name:        workload.Spec.Name,
		Description: workload.Spec.Description,
		Version:     workload.Spec.Version,
		CreatedAt:   timestamppb.New(workload.Spec.CreatedAt.Time),
		UpdatedAt:   timestamppb.New(workload.Spec.UpdatedAt.Time),
	}

	return &workloadv1.GetResponse{
		Workload: response,
	}, nil
}

func (s *WorkloadService) Create(ctx context.Context, request *workloadv1.CreateRequest) (*workloadv1.CreateResponse, error) {
	id := fmt.Sprintf("%s-%s", request.VersionId, request.Name)
	namespace := strings.Replace(id, ".", "-", -1)

	workload := v1.SiphonManagedWorkload{
		ObjectMeta: metav1.ObjectMeta{
			Name:      namespace,
			Namespace: "default",
		},
		Spec: v1.SiphonManagedWorkloadSpec{
			ID:          namespace,
			Namespace:   namespace,
			Name:        request.Name,
			Version:     request.VersionId,
			Description: request.Description,
			CreatedAt:   metav1.Now(),
			UpdatedAt:   metav1.Now(),
		},
	}

	repo := &repository.WorkloadRepositoryImpl{Client: s.Client}
	if err := repo.Create(ctx, &workload); err != nil {
		return nil, err
	}

	return &workloadv1.CreateResponse{
		Workload: &workloadv1.Workload{
			Id:          workload.Spec.ID,
			Name:        workload.Spec.Name,
			Description: workload.Spec.Description,
			Version:     workload.Spec.Version,
			CreatedAt:   timestamppb.New(workload.Spec.CreatedAt.Time),
			UpdatedAt:   timestamppb.New(workload.Spec.UpdatedAt.Time),
		},
	}, nil
}

func (s *WorkloadService) Update(ctx context.Context, request *workloadv1.UpdateRequest) (*workloadv1.UpdateResponse, error) {
	workload := &model.Workload{}
	if err := s.db.First(&workload, request.Workload.Id).Error; err != nil {
		return nil, err
	}

	workload.Name = request.Workload.Name
	workload.Description = request.Workload.Description
	workload.Version = request.Workload.Version

	if err := s.db.Save(&workload).Error; err != nil {
		return nil, err
	}

	return &workloadv1.UpdateResponse{
		Workload: request.Workload,
	}, nil
}

func (s *WorkloadService) Delete(ctx context.Context, request *workloadv1.DeleteRequest) (*workloadv1.DeleteResponse, error) {
	if err := s.Client.Delete(ctx, &v1.SiphonManagedWorkload{ObjectMeta: metav1.ObjectMeta{Name: request.Id, Namespace: "default"}}); err != nil {
		return nil, err
	}

	return &workloadv1.DeleteResponse{}, nil
}
