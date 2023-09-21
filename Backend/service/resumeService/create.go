package resumeService

import (
	"context"
	"resume_backend/models"
	"resume_backend/pkg/ctl"
	"resume_backend/pkg/utils"
	"resume_backend/repositories"
)

func (s *ResumeService) CreateResume(ctx context.Context, req *models.CreateResumeReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.LogrusObj.Info(ctx)
		return
	}
	user, err := repositories.NewUserDao(ctx).FindUserByUserId(uint(u.Id))
	if err != nil {
		utils.LogrusObj.Info(ctx)
		return
	}
	resume := &models.Resume{
		UserId:            user.ID,
		Title:             *&req.Title,
		BaseInfo:          *&req.BaseInfo,
		CareerTarget:      *&req.CareerTarget,
		EducationInfo:     *&req.EducationInfo,
		ProjectExperience: *&req.ProjectExperience,
		SkillInfo:         *&req.SkillInfo,
		WorkExperience:    *&req.WorkExperience,
	}
	err = repositories.NewResumeDao(ctx).CreateResume(resume)
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}
	return ctl.RespSuccess(), nil
}
func (s *ResumeService) UpdateResume(ctx context.Context, req *models.UpdateResumeReq) (resp interface{}, err error) {
	// 拿到用户信息
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.LogrusObj.Info(ctx)
		return
	}
	err = repositories.NewResumeDao(ctx).UpdateResume(u.Id, req)
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}
	return ctl.RespSuccess(), nil
}

func (s *ResumeService) DeleteResume(ctx context.Context, req *models.DeleteResumeReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.LogrusObj.Info(ctx)
		return
	}
	err = repositories.NewResumeDao(ctx).DeleteResume(u.Id, req)
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}
	return ctl.RespSuccess(), nil
}
