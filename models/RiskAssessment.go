package models


import (
	"encoding/json"
	"gorm.io/gorm"
	"ametory-crud/requests"

	"time"


)

type RiskAssessment struct {
	Base
	RiskId string `gorm:"type:char(36);NOT NULL" json:"risk_id"`
	Likelihood string `gorm:"type:risk_assessment_likelihood_enum;NOT NULL" json:"likelihood"`
	Impact string `gorm:"type:risk_assessment_impact_enum;NOT NULL" json:"impact"`
	RiskScore int `gorm:"type:int;NOT NULL" json:"risk_score"`
	AssessedBy string `gorm:"type:char(36);NOT NULL" json:"assessed_by"`
	AssessedAt time.Time `gorm:"type:timestamp ;DEFAULT CURRENT_TIMESTAMP" json:"assessed_at"`
	}

func init() {
	RegisterModel(&RiskAssessment{})
}

func (p *RiskAssessment) BeforeCreate(tx *gorm.DB) error {
	p.ID = generateUUID()
	return nil
}

func (p RiskAssessment) MarshalJSON() ([]byte, error) {
	return json.Marshal(requests.RiskAssessmentResponse{
		ID:       p.ID,
		RiskId: p.RiskId,
		Likelihood: p.Likelihood,
		Impact: p.Impact,
		RiskScore: p.RiskScore,
		AssessedBy: p.AssessedBy,
		AssessedAt: p.AssessedAt,
		})
}

type RiskAssessmentResp struct {
	Pagination 	PaginationResponse 	`json:"pagination"`
	Data		[]RiskAssessment 	`json:"data"`
	Message 	string 				`json:"msg"`
}

type RiskAssessmentSingleResp struct {
	Data		RiskAssessment 	`json:"data"`
	Message 	string 				`json:"msg"`
}

func (p *RiskAssessment) UnmarshalJSON(data []byte) error {
	var req requests.RiskAssessmentRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return err
	}
	p.RiskId = req.RiskId
	p.Likelihood = req.Likelihood
	p.Impact = req.Impact
	p.RiskScore = req.RiskScore
	p.AssessedBy = req.AssessedBy
	p.AssessedAt = req.AssessedAt
	
	return nil
}
