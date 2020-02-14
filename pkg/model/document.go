package model

import "time"

type Document struct {
	ID string `json:"id"`

	Content string `json:"content"`

	Lodestone DocLodestone `json:"lodestone"`

	// File information/attributes
	File DocFile `json:"file"`

	// Document storage location (and thumbnail storage)
	Storage DocStorage `json:"storage"`

	// Document metadata extracted from document via tika
	Meta DocMeta `json:"meta"`
}

type DocLodestone struct {
	ProcessorVersion string   `json:"processor_version"`
	Title            string   `json:"title"`
	Tags             []string `json:"tags"`
	Bookmark         bool     `json:"bookmark"`
}

type DocFile struct {
	ContentType  string    `json:"content_type"`
	FileName     string    `json:"filename"`
	Extension    string    `json:"extension"` //does not include .
	Filesize     int64     `json:"filesize"`
	IndexedChars int64     `json:"indexed_chars"`
	IndexedDate  time.Time `json:"indexed_date,omitempty"`
	Created      time.Time `json:"created,omitempty"`
	LastModified time.Time `json:"last_modified,omitempty"`
	LastAccessed time.Time `json:"last_accessed,omitempty"`
	Checksum     string    `json:"checksum"`

	Group string `json:"group"`
	Owner string `json:"owner"`
}

type DocStorage struct {
	Bucket      string `json:"bucket"`       //bucket name, does not include "/" char
	Path        string `json:"path"`         //key, does not include "/" prefix
	ThumbBucket string `json:"thumb_bucket"` //bucket name, does not include "/" char
	ThumbPath   string `json:"thumb_path"`   //key, does not include "/" prefix
}

type DocMeta struct {
	Author      string    `json:"author"`
	CreatedDate time.Time `json:"created"`
	SavedDate   time.Time `json:"saved"`
	Date        time.Time `json:"date,omitempty"`

	Keywords    []string `json:"keywords"`
	Title       string   `json:"title"`
	Language    string   `json:"language"`
	Format      string   `json:"format"`
	Identifier  string   `json:"identifier"`
	Contributor string   `json:"contributor"`
	Modifier    string   `json:"modifier"`
	CreatorTool string   `json:"creator_tool"`
	Publisher   string   `json:"publisher"`
	Relation    string   `json:"relation"`
	Rights      string   `json:"rights"`
	Source      string   `json:"source"`
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Latitude    string   `json:"latitude"`
	Longitude   string   `json:"longitude"`
	Altitude    string   `json:"altitude"`
	Rating      byte     `json:"rating"`
	Comments    string   `json:"comments"`
	Pages       string   `json:"pages"`
}
