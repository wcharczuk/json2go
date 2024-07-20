package autogen

// AutoGenerated is an autogenerated type from JSON.
type AutoGenerated struct {
	Report struct {
		AssetReportID  string  `json:"asset_report_id,omitempty"`
		ClientReportID string  `json:"client_report_id,omitempty"`
		DateGenerated  string  `json:"date_generated,omitempty"`
		DaysRequested  float64 `json:"days_requested,omitempty"`
		Items          []struct {
			Accounts []struct {
				AccountID string `json:"account_id,omitempty"`
				Balances  struct {
					Available              float64 `json:"available,omitempty"`
					Current                float64 `json:"current,omitempty"`
					IsoCurrencyCode        string  `json:"iso_currency_code,omitempty"`
					Limit                  any     `json:"limit,omitempty"`
					MarginLoanAmount       any     `json:"margin_loan_amount,omitempty"`
					UnofficialCurrencyCode any     `json:"unofficial_currency_code,omitempty"`
				} `json:"balances,omitempty"`
				DaysAvailable      float64 `json:"days_available,omitempty"`
				HistoricalBalances []struct {
					Current                float64 `json:"current,omitempty"`
					Date                   string  `json:"date,omitempty"`
					IsoCurrencyCode        string  `json:"iso_currency_code,omitempty"`
					UnofficialCurrencyCode any     `json:"unofficial_currency_code,omitempty"`
				} `json:"historical_balances,omitempty"`
				Mask         string `json:"mask,omitempty"`
				Name         string `json:"name,omitempty"`
				OfficialName string `json:"official_name,omitempty"`
				Owners       []struct {
					Addresses []struct {
						Data struct {
							City       string `json:"city,omitempty"`
							Country    string `json:"country,omitempty"`
							PostalCode string `json:"postal_code,omitempty"`
							Region     string `json:"region,omitempty"`
							Street     string `json:"street,omitempty"`
						} `json:"data,omitempty"`
						Primary bool `json:"primary,omitempty"`
					} `json:"addresses,omitempty"`
					Emails       []AutoGenerated_Report_Items_Accounts_Owners_Emails `json:"emails,omitempty"`
					Names        []string                                            `json:"names,omitempty"`
					PhoneNumbers []AutoGenerated_Report_Items_Accounts_Owners_Emails `json:"phone_numbers,omitempty"`
				} `json:"owners,omitempty"`
				OwnershipType any    `json:"ownership_type,omitempty"`
				Subtype       string `json:"subtype,omitempty"`
				Transactions  []struct {
					AccountID              string  `json:"account_id,omitempty"`
					Amount                 float64 `json:"amount,omitempty"`
					Date                   string  `json:"date,omitempty"`
					IsoCurrencyCode        string  `json:"iso_currency_code,omitempty"`
					OriginalDescription    string  `json:"original_description,omitempty"`
					Pending                bool    `json:"pending,omitempty"`
					TransactionID          string  `json:"transaction_id,omitempty"`
					UnofficialCurrencyCode any     `json:"unofficial_currency_code,omitempty"`
				} `json:"transactions,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"accounts,omitempty"`
			DateLastUpdated string `json:"date_last_updated,omitempty"`
			InstitutionID   string `json:"institution_id,omitempty"`
			InstitutionName string `json:"institution_name,omitempty"`
			ItemID          string `json:"item_id,omitempty"`
		} `json:"items,omitempty"`
		User struct {
			ClientUserID string `json:"client_user_id,omitempty"`
			Email        string `json:"email,omitempty"`
			FirstName    string `json:"first_name,omitempty"`
			LastName     string `json:"last_name,omitempty"`
			MiddleName   string `json:"middle_name,omitempty"`
			PhoneNumber  string `json:"phone_number,omitempty"`
			Ssn          string `json:"ssn,omitempty"`
		} `json:"user,omitempty"`
	} `json:"report,omitempty"`
	RequestID string `json:"request_id,omitempty"`
	Warnings  []any  `json:"warnings,omitempty"`
}

// AutoGenerated_Report_Items_Accounts_Owners_Emails is an aliasesd type from JSON.
type AutoGenerated_Report_Items_Accounts_Owners_Emails struct {
	Data    string `json:"data,omitempty"`
	Primary bool   `json:"primary,omitempty"`
	Type    string `json:"type,omitempty"`
}
