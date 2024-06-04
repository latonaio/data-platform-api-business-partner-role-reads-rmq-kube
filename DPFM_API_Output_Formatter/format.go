package dpfm_api_output_formatter

import (
	"data-platform-api-business-partner-role-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToBusinessPartnerRole(rows *sql.Rows) (*[]BusinessPartnerRole, error) {
	defer rows.Close()
	businessPartnerRole := make([]BusinessPartnerRole, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.BusinessPartnerRole{}

		err := rows.Scan(
			&pm.BusinessPartnerRole,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &businessPartnerRole, nil
		}

		data := pm
		businessPartnerRole = append(businessPartnerRole, BusinessPartnerRole{
			BusinessPartnerRole:	data.BusinessPartnerRole,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &businessPartnerRole, nil
}

func ConvertToText(rows *sql.Rows) (*[]Text, error) {
	defer rows.Close()
	text := make([]Text, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Text{}

		err := rows.Scan(
			&pm.BusinessPartnerRole,
			&pm.Language,
			&pm.BusinessPartnerRoleName,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &text, err
		}

		data := pm
		text = append(text, Text{
			BusinessPartnerRole:    	data.BusinessPartnerRole,
			Language:          			data.Language,
			BusinessPartnerRoleName:	data.BusinessPartnerRoleName,
			CreationDate:				data.CreationDate,
			LastChangeDate:				data.LastChangeDate,
			IsMarkedForDeletion:		data.IsMarkedForDeletion,
		})
	}

	return &text, nil
}
