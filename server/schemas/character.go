package schemas

// The juicy stuff

type CharacterResponse struct {
	Name        string    `json:"name"`
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Sources     *[]string `json:"sources,omitempty"` // Sources are optional
	Tags        []string  `json:"tags"`

	Image EntityImage `json:"img"`

	Gallery []struct {
		EntityImage
		Source  string       `json:"source"`
		AddedBy *AddedByUser `json:"added_by"`
	} `json:"gallery"`

	Attributes struct {
		Gender GenderAlias `json:"gender"`
		Traits []string    `json:"traits"`

		Franchise struct {
			Name   string               `json:"name"`
			Year   int                  `json:"year"`
			Medium FranchiseMediumAlias `json:"medium"`

			Holder []struct {
				Name string `json:"name"`
				EntityRef
			} `json:"holder"`
		} `json:"franchise"`

		Quotes []struct {
			Quote   string       `json:"quote"`
			AddedBy *AddedByUser `json:"added_by"`
		} `json:"quotes"`

		Aliases *[]struct {
			Type string  `json:"type"`
			Name string  `json:"name"`
			Lang *string `json:"lang,omitempty"`
		} `json:"aliases,omitempty"`

		VoiceActor []struct {
			Name string
			As   string
			EntityRef
		} `json:"voice_actor"`

		Others []any `json:"others"`
	} `json:"attributes"`

	AddedBy AddedByUser `json:"added_by"`
}
