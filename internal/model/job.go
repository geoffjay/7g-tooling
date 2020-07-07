package model

type Job struct {
	BaseModelSeq
	Name  *string
	Tasks []Task `gorm:"many2many:job_tasks;association_autoupdate:false;association_autocreate:false"`
}
