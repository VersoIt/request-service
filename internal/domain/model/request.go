package model

import "time"

type Request struct {
	ID        int64
	Type      RequestType
	Payload   []byte
	Status    RequestStatus
	CreatedAt time.Time
}

type Requests []Request

type RequestType string

const (
	TypeCertificate         RequestType = "certificate"
	TypePassportApplication             = "passport_application"
	TypeTexDeclaration                  = "tex_declaration"
	TypeLicenseRequest                  = "license_request"
	TypeSocialBenefit                   = "social_benefit"
	TypeRegistrationUpdate              = "registration_update"
	TypeUnknown                         = "unknown"
)

type RequestStatus string

const (
	StatusPending    RequestStatus = "pending"
	StatusProcessing               = "processing"
	StatusSuccess                  = "success"
	StatusFailed                   = "failed"
	StatusUnknown                  = "unknown"
)
