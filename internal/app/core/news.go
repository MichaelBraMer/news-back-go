package core

type News struct {
    ID string `bson:"_id,omitempty" json:"id"`
    Title string `bson:"title" json:"title"`
    Paragraph string `bson:"paragraph" json:"paragraph"`
}