package requests

type Text struct {
	BusinessPartnerRole		string  `json:"BusinessPartnerRole"`
	Language          		string  `json:"Language"`
	BusinessPartnerRoleName	string  `json:"BusinessPartnerRoleName"`
	CreationDate			string	`json:"CreationDate"`
	LastChangeDate			string	`json:"LastChangeDate"`
	IsMarkedForDeletion		*bool	`json:"IsMarkedForDeletion"`
}
