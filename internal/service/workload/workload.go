package workload

import (
	"context"

	workloadv1 "github.com/siphon/siphon/generated/workload/v1"
	"github.com/siphon/siphon/internal/model"
	"gorm.io/gorm"
)

type WorkloadService struct {
	db *gorm.DB
}

func NewWorkloadService(db *gorm.DB) *WorkloadService {
	return &WorkloadService{db: db}
}

var _ Workloader = &WorkloadService{}

func (s *WorkloadService) List(context.Context, *workloadv1.ListRequest) (*workloadv1.ListResponse, error) {
	workloads := []*model.Workload{}
	if err := s.db.Find(&workloads).Error; err != nil {
		return nil, err
	}

	response := []*workloadv1.Workload{}
	for _, workload := range workloads {
		response = append(response, &workloadv1.Workload{
			Id:          workload.ID,
			Name:        workload.Name,
			Description: workload.Description,
			Version:     workload.Version,
		})
	}

	return &workloadv1.ListResponse{
		Workloads: response,
	}, nil
}

func (s *WorkloadService) Get(ctx context.Context, request *workloadv1.GetRequest) (*workloadv1.GetResponse, error) {
	workload := &model.Workload{}
	if err := s.db.First(&workload, request.Id).Error; err != nil {
		return nil, err
	}

	response := &workloadv1.Workload{
		Id:          workload.ID,
		Name:        workload.Name,
		Description: workload.Description,
		Version:     workload.Version,
	}

	return &workloadv1.GetResponse{
		Workload: response,
	}, nil
}

func (s *WorkloadService) Create(ctx context.Context, request *workloadv1.CreateRequest) (*workloadv1.CreateResponse, error) {
	workload := &model.Workload{
		ID:          request.Workload.Id,
		Name:        request.Workload.Name,
		Description: request.Workload.Description,
		Version:     request.Workload.Version,
	}

	if err := s.db.Create(&workload).Error; err != nil {
		return nil, err
	}

	response := &workloadv1.Workload{
		Id:          workload.ID,
		Name:        workload.Name,
		Description: workload.Description,
		Version:     workload.Version,
	}

	return &workloadv1.CreateResponse{
		Workload: response,
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
	workload := &model.Workload{}
	if err := s.db.First(&workload, request.Id).Error; err != nil {
		return nil, err
	}

	if err := s.db.Delete(&workload).Error; err != nil {
		return nil, err
	}

	return &workloadv1.DeleteResponse{}, nil
}
