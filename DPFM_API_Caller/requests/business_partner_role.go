package requests

type BusinessPartnerRole struct {
	BusinessPartnerRole	string	`json:"BusinessPartnerRole"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
