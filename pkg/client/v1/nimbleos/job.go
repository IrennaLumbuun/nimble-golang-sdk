// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Job - Jobs are operations in progress in the system.
// Export JobFields for advance operations like search filter etc.
var JobFields *JobStringFields

func init() {
	CompletionTimefield := "completion_time"
	CreationTimefield := "creation_time"
	CurrentPhasefield := "current_phase"
	CurrentPhaseDescriptionfield := "current_phase_description"
	Descriptionfield := "description"
	IDfield := "id"
	Namefield := "name"
	LastModifiedfield := "last_modified"
	ObjectIdfield := "object_id"
	OpTypefield := "op_type"
	Typefield := "type"
	ParentJobIdfield := "parent_job_id"
	PercentCompletefield := "percent_complete"
	Requestfield := "request"
	Responsefield := "response"
	Statefield := "state"
	Resultfield := "result"
	TotalPhasesfield := "total_phases"

	JobFields = &JobStringFields{
		CompletionTime:          &CompletionTimefield,
		CreationTime:            &CreationTimefield,
		CurrentPhase:            &CurrentPhasefield,
		CurrentPhaseDescription: &CurrentPhaseDescriptionfield,
		Description:             &Descriptionfield,
		ID:                      &IDfield,
		Name:                    &Namefield,
		LastModified:            &LastModifiedfield,
		ObjectId:                &ObjectIdfield,
		OpType:                  &OpTypefield,
		Type:                    &Typefield,
		ParentJobId:             &ParentJobIdfield,
		PercentComplete:         &PercentCompletefield,
		Request:                 &Requestfield,
		Response:                &Responsefield,
		State:                   &Statefield,
		Result:                  &Resultfield,
		TotalPhases:             &TotalPhasesfield,
	}
}

type Job struct {
	// CompletionTime - Completion time of the job.
	CompletionTime *int64 `json:"completion_time,omitempty"`
	// CreationTime - Time when this job was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// CurrentPhase - Phase number of the job in progress.
	CurrentPhase *int64 `json:"current_phase,omitempty"`
	// CurrentPhaseDescription - Description of the current phase of the job.
	CurrentPhaseDescription *string `json:"current_phase_description,omitempty"`
	// Description - Description of the job.
	Description *string `json:"description,omitempty"`
	// ID - Identifier for job.
	ID *string `json:"id,omitempty"`
	// Name - Name of the job.
	Name *string `json:"name,omitempty"`
	// LastModified - Time of the last update from the job.
	LastModified *int64 `json:"last_modified,omitempty"`
	// ObjectId - Identifier for object being acted upon.
	ObjectId *string `json:"object_id,omitempty"`
	// OpType - Type of operation.
	OpType *NsOperationType `json:"op_type,omitempty"`
	// Type - Job type.
	Type *NsJobType `json:"type,omitempty"`
	// ParentJobId - Identifier of parent job.
	ParentJobId *string `json:"parent_job_id,omitempty"`
	// PercentComplete - Progress of the job as a percentage.
	PercentComplete *int64 `json:"percent_complete,omitempty"`
	// Request - Original request that the job is responsible for.
	Request *NsRequest `json:"request,omitempty"`
	// Response - Response from the operation as the job executes.
	Response *NsResponse `json:"response,omitempty"`
	// State - Status of the job.
	State *NsJobStatus `json:"state,omitempty"`
	// Result - Result of the job.
	Result *NsJobResult `json:"result,omitempty"`
	// TotalPhases - Total number of phases of the job.
	TotalPhases *int64 `json:"total_phases,omitempty"`
}

// Struct for JobFields
type JobStringFields struct {
	CompletionTime          *string
	CreationTime            *string
	CurrentPhase            *string
	CurrentPhaseDescription *string
	Description             *string
	ID                      *string
	Name                    *string
	LastModified            *string
	ObjectId                *string
	OpType                  *string
	Type                    *string
	ParentJobId             *string
	PercentComplete         *string
	Request                 *string
	Response                *string
	State                   *string
	Result                  *string
	TotalPhases             *string
}
