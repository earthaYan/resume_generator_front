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
