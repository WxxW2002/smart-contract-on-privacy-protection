package models

type EncryptedDocument struct {
    DocID       string            `json:"docID"`
    Content     []byte            `json:"content"`
    Keywords    []string          `json:"keywords"`
    Owner       string            `json:"owner"`
    Timestamp   int64             `json:"timestamp"`
}

type SearchIndex struct {
    Keyword     string            `json:"keyword"`
    DocIDs      []string          `json:"docIDs"`
}