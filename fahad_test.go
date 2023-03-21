package fahad

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertDataAkreditas(t *testing.T) {
	dbname := "infokomdb"
	dataakreditas:= DataAkreditas{
		ID:      primitive.NewObjectID(),
		Perguruan_Tinggi:    "ULBI",
		Program_studi:   "Teknik Informatika",
		Peringkat: "B",
		Status: "Masih Berlaku",
	}
	insertedID := InsertDataAkreditas(dbname, dataakreditas)
	if insertedID == nil {
		t.Error("Failed to insert user")
	}
}

func TestInsertDataProgramStudi(t *testing.T) {
	dbname := "infokomdb"
	dataprogramstudi := DataProgramStudi{
		ID:       primitive.NewObjectID(),
		Program_studi: "Teknik Informatika",
		Program:  "Sarjana Terapan",
	}
	insertedID := InsertDataProgramStudi(dbname, dataprogramstudi)
	if insertedID == nil {
		t.Error("Failed to insert surat")
	}
}


func TestInsertProfile(t *testing.T) {
	dbname := "infokomdb"
	profile := Profile{
		ID:     primitive.NewObjectID(),
		Isi_satu: "kampus merdeka",
	}
	insertedID := InsertProfile(dbname, profile)
	if insertedID == nil {
		t.Error("Failed to insert lokasi")
	}
}
