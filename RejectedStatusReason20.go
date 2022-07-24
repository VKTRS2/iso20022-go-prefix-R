package iso20022

// Specifies reasons for the rejected status.
type RejectedStatusReason20 struct {

	// Specifies the reason why the instruction or instruction cancellation has been rejected.
	ReasonCode *RejectedReason19Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectedStatusReason20) AddReasonCode() *RejectedReason19Choice {
	r.ReasonCode = new(RejectedReason19Choice)
	return r.ReasonCode
}

func (r *RejectedStatusReason20) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
