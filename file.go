package structs

import (
	"crypto/rand"
	"fmt"
	"github.com/kazoup/structs/content"
	"github.com/kazoup/structs/intmap"
	"github.com/kazoup/structs/metadata"
	"github.com/kazoup/structs/permissions"
	"mime"
	"os"
	"path/filepath"
	//"strconv"
	"strings"
	"time"
)

// File model
type File struct {
	ExistsOnDisk    bool                    `json:"exists_on_disk"`
	ID              string                  `json:"_id"`
	ArchiveComplete bool                    `json:"archive_complete"`
	FirstSeen       time.Time               `json:"first_seen"`
	IDB64           string                  `json:"id_b64"`
	LastSeen        time.Time               `json:"last_seen"`
	Content         content.Content         `json:"content"`
	Metadata        metadata.Metadata       `json:"metadata"`
	Permissions     permissions.Permissions `json:"permissions"`
}

// LocalFile model
type LocalFile struct {
	Type string
	Path string
	Info os.FileInfo
}

// NewFileFromLocal file constructor
func NewFileFromLocal(lf *LocalFile) *File {
	return &File{
		ExistsOnDisk: true,
		//ID:              "/" + lf.Path + ":" + strconv.FormatInt(lf.Info.ModTime().Unix(), 10),
		ID:              pseudo_uuid(),
		ArchiveComplete: false,
		FirstSeen:       time.Now(),
		Content:         content.Content{},
		Metadata: metadata.Metadata{
			Mimetype:     mime.TypeByExtension(filepath.Ext(lf.Info.Name())),
			DirpathSplit: pathToIntmap(lf.Path),
			Extension:    filepath.Ext(lf.Info.Name()),
			Created:      lf.Info.ModTime(),
			Modified:     lf.Info.ModTime(),
			Filename:     lf.Info.Name(),
			Dirpath:      filepath.Dir(lf.Path),
			Accessed:     lf.Info.ModTime(),
			Fullpath:     lf.Path,
			Sharepath:    filepath.VolumeName(lf.Path),
			Size:         lf.Info.Size(),
		},
		Permissions: permissions.Permissions{},
	}
}

func pathToIntmap(path string) intmap.Intmap {
	results := make(intmap.Intmap)
	dir := filepath.Dir(path)
	parts := strings.Split(dir, "/")
	for k, v := range parts {
		if k == 0 {
			results[k] = "/" + v

		} else {
			results[k] = filepath.Join(results[k-1], v)
		}
	}
	return results
}

func pseudo_uuid() (uuid string) {

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return
}
