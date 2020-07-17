package ladon

type ProviderImageApp struct {
	Image      string `firestore:"image" json:"image"`
	ProviderId int    `firestore:"provider_id" json:"provider_id"`
	Url        string `firestore:"url" json:"url"`
}
