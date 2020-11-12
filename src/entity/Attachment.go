package entity

// Attachment is part of a tweet or comment
type Attachment struct {
	ID           int
	Path         string // (/idOfTheUser/uuidOfTheFile) --> /124312478/844e7a29-d8ee-44b4-94ba-adc519912cd2
	MimeType     string // video/mp4 or image/png or video/mkv
	Extension    string // .mp4 or .png or .mkv
	OriginalName string // giffycat
}

func (a Attachment) getOriginalFileName() string {
	return a.OriginalName + a.Extension
}
