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
