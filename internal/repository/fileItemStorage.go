package pkg

import (
	"encoding/json"
	"io"
	"os"
)

type FileItemStorage struct {
	filePath string
}

func NewFileItemRepository() FileItemStorage {
	return FileItemStorage{filePath: "storage.json"}
}

func (r *FileItemStorage) GetItems() ([]Item, error) {
	jsonFile, err := os.Open(r.filePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(io.Reader(jsonFile))
	if err != nil {
		return nil, err
	}
	var items []Item
	err = json.Unmarshal(byteValue, &items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *FileItemStorage) StoreItemsState(items []Item) error {
	file, err := os.Create(r.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(items)
}
