package models

type Asset struct {
	ID             string `json:"ID"`
	Color          string `json:"Color"`
	Size           int    `json:"Size"`
	Owner          string `json:"Owner"`
	AppraisedValue int    `json:"AppraisedValue"`
	EncryptedColor string `json:"EncryptedColor"`
	EncryptedOwner string `json:"EncryptedOwner"`
}

type AssetPrivateDetails struct {
	ID             string `json:"ID"`
	EncryptedColor string `json:"EncryptedColor"`
	EncryptedOwner string `json:"EncryptedOwner"`
}

func (a *Asset) ToAssetPrivateDetails() *AssetPrivateDetails {
	return &AssetPrivateDetails{
		ID:             a.ID,
		EncryptedColor: a.EncryptedColor,
		EncryptedOwner: a.EncryptedOwner,
	}
}
