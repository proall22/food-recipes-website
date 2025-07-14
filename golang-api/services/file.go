package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

type FileService struct {
	uploadDir string
}

func NewFileService() *FileService {
	uploadDir := os.Getenv("UPLOAD_DIR")
	if uploadDir == "" {
		uploadDir = "./uploads"
	}

	// Create upload directory if it doesn't exist
	os.MkdirAll(uploadDir, 0755)
	os.MkdirAll(filepath.Join(uploadDir, "images"), 0755)
	os.MkdirAll(filepath.Join(uploadDir, "recipes"), 0755)
	os.MkdirAll(filepath.Join(uploadDir, "avatars"), 0755)

	return &FileService{
		uploadDir: uploadDir,
	}
}

type UploadResult struct {
	URL      string `json:"url"`
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
}

func (s *FileService) UploadImage(file *multipart.FileHeader, category string) (*UploadResult, error) {
	// Validate file type
	if !s.isValidImageType(file.Filename) {
		return nil, fmt.Errorf("invalid file type. Only JPEG, PNG, and WebP are allowed")
	}

	// Validate file size (max 10MB)
	if file.Size > 10*1024*1024 {
		return nil, fmt.Errorf("file size too large. Maximum 10MB allowed")
	}

	// Generate unique filename
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%s_%d%s", uuid.New().String(), time.Now().Unix(), ext)

	// Determine upload path based on category
	var uploadPath string
	switch category {
	case "avatar":
		uploadPath = filepath.Join(s.uploadDir, "avatars", filename)
	case "recipe":
		uploadPath = filepath.Join(s.uploadDir, "recipes", filename)
	default:
		uploadPath = filepath.Join(s.uploadDir, "images", filename)
	}

	// Open uploaded file
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	// Create destination file
	dst, err := os.Create(uploadPath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	// Copy file content
	size, err := io.Copy(dst, src)
	if err != nil {
		return nil, err
	}

	// Generate URL
	relativePath := strings.TrimPrefix(uploadPath, s.uploadDir)
	url := fmt.Sprintf("/uploads%s", relativePath)

	return &UploadResult{
		URL:      url,
		Filename: filename,
		Size:     size,
	}, nil
}

func (s *FileService) UploadMultipleImages(files []*multipart.FileHeader, category string) ([]*UploadResult, error) {
	var results []*UploadResult

	for _, file := range files {
		result, err := s.UploadImage(file, category)
		if err != nil {
			return nil, fmt.Errorf("failed to upload %s: %v", file.Filename, err)
		}
		results = append(results, result)
	}

	return results, nil
}

func (s *FileService) DeleteImage(filename string) error {
	// Find file in all possible directories
	possiblePaths := []string{
		filepath.Join(s.uploadDir, "images", filename),
		filepath.Join(s.uploadDir, "recipes", filename),
		filepath.Join(s.uploadDir, "avatars", filename),
	}

	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			return os.Remove(path)
		}
	}

	return fmt.Errorf("file not found: %s", filename)
}

func (s *FileService) isValidImageType(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	validExts := []string{".jpg", ".jpeg", ".png", ".webp"}

	for _, validExt := range validExts {
		if ext == validExt {
			return true
		}
	}

	return false
}
