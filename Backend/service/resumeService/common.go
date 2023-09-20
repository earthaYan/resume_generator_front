package resumeService

import "sync"

var ResumeServiceIns *ResumeService
var ResumeServiceOnce sync.Once

type ResumeService struct{}

func GetResumeService() *ResumeService {
	ResumeServiceOnce.Do(func() {
		ResumeServiceIns = &ResumeService{}
	})
	return ResumeServiceIns
}
