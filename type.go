package fahad

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DataAkreditas struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Perguruan_Tinggi string             `bson:"perguruan_tinggi" json:"perguruan_tinggi"`
	Program_studi  string             `bson:"program_studi" json:"program_studi"`
	Peringkat  string             `bson:"peringkat" json:"peringkat"`
	Status  string             `bson:"status" json:"status"`
}
type DataProgramStudi struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Program_studi string             `bson:"program_studi" json:"program_studi"`
	Program   string          `bson:"program" json:"program"`
}

type Profile struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Isi_satu string             `bson:"isi_satu,omitempty" json:"isi_satu,omitempty"`

}