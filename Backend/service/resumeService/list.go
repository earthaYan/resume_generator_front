package resumeService

import (
	"context"
	"resume_backend/models"
	"resume_backend/pkg/ctl"
	"resume_backend/pkg/utils"
	"resume_backend/repositories"
)

func (s *ResumeService) ListResume(ctx context.Context, req *models.ListResumeReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.LogrusObj.Info(ctx)
		return
	}
	resumes, total, err := repositories.NewResumeDao(ctx).ListResume(req.Start, req.Limit, u.Id)
	tRespList := make([]*models.ListResumeResp, 0)
	for _, v := range resumes {
		tRespList = append(tRespList, &models.ListResumeResp{
			Id:        int64(v.ID),
			Title:     v.Title,
			Author:    u.UserName,
			CreatedAt: v.CreatedAt.Unix(),
			UpdatedAt: v.UpdatedAt.Unix(),
		})
	}
	return ctl.ResponseList(tRespList, total), nil
}
func (s *ResumeService) DetailResume(ctx context.Context, req *models.DetailReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.LogrusObj.Info(ctx)
		return
	}
	resume, err := repositories.NewResumeDao(ctx).FindResumeByResumeIdAndUserId(req.ResumeId, u.Id)

	if err != nil {
		utils.LogrusObj.Info(ctx)
		return
	}
	var BaseInfo models.ResumeBasicInfo = resume.BaseInfo
	var CareerTarget models.ResumeCareerTarget = resume.CareerTarget
	var EducationInfo []models.ResumeEducation = resume.EducationInfo
	var ProjectExperience []models.ResumeProjectExperience = resume.ProjectExperience
	var SkillInfo []models.ResumeSkill = resume.SkillInfo
	var WorkExperience []models.ResumeWorkExperience = resume.WorkExperience
	v := &models.DetailResp{
		ID:                resume.ID,
		Title:             resume.Title,
		BaseInfo:          BaseInfo,
		CareerTarget:      CareerTarget,
		EducationInfo:     EducationInfo,
		ProjectExperience: ProjectExperience,
		SkillInfo:         SkillInfo,
		WorkExperience:    WorkExperience,
	}
	return ctl.RespSuccessWithData(v), nil
}
