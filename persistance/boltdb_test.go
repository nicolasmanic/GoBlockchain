package persistance

import (
	"os"
	"testing"
)

func TestSaveBlockMetadata(t *testing.T) {
	dbManager := newDBManager("test.db")
	defer dbManager.CloseDb()
	err := dbManager.SaveBlockMetadata([]byte("1234"), &BlockMetadata{1, "/test/testblock1.dat"})
	if err != nil {
		t.Error("Could not save block to DB")
	}
}

func TestRetrieveBlockPathByHash(t *testing.T) {
	dbManager := newDBManager("test.db")
	defer dbManager.CloseDb()
	returnedPath, err := dbManager.RetrieveBlockPathByHash([]byte("1234"))

	if err != nil {
		t.Errorf("Could not retrieve block path error message: %v", err)
	}
	if "/test/testblock1.dat" != returnedPath {
		t.Error("Could not retrieve block path from DB")
	}
}

func TestLastUsedHashDB(t *testing.T) {
	dbManager := newDBManager("test.db")
	defer dbManager.CloseDb()
	lastHash := dbManager.lastUsedHash()

	if "1234" != string(lastHash[:]) {
		t.Error("Could not retrieve lash blocks hash from DB")

	}
}

func tearDown() {
	os.Remove("test.db")
}

func TestMain(m *testing.M) {
	retCode := m.Run()
	tearDown()
	os.Exit(retCode)
}
