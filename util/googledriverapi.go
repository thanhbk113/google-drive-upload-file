package util

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"thanhbk113/internal/config"

	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

//serviceAuth: path to file json auth google driver
//filePath: path to file zip

// Upload file to google driver
func UploadFileToGoogleDriver(serviceAuth, filePath string) string {
	SCOPE := drive.DriveScope
	ctx := context.Background()
	srv, err := drive.NewService(ctx, option.WithCredentialsFile(serviceAuth), option.WithScopes(SCOPE))
	if err != nil {
		fmt.Println("err src google driver api:", err)
		return ""
	}

	file, err := os.Open(filePath)
	info, _ := file.Stat()
	if err != nil {
		fmt.Println("err open file:", err)
		return ""
	}

	if err != nil {
		fmt.Println("err open file:", err)
		return ""
	}
	defer file.Close()

	// Create File metadata
	f := &drive.File{Name: info.Name(),
		Parents: []string{config.GetEnv().GOOGLE_DRIVER_FOLDERID}}

	// Create and upload the file
	res, err := srv.Files.
		Create(f).
		Media(file). //context.Background(), file, fileInf.Size(), baseMimeType).
		ProgressUpdater(func(now, size int64) { fmt.Printf("%d, %d\r", now, size) }).
		Do()

	if err != nil {
		fmt.Println("err upload file:", err)
		return ""
	}

	fmt.Printf("New file google driver id: %s\n", res.Id)

	url, err := GeneratePublicLinkFileInGoogleDriver(srv, res.Id)

	if err != nil {
		fmt.Println("err get public link file:", err)
		return ""
	}

	fmt.Printf("Public link file: %s\n", url)
	return url

}

// generate public link file in google driver
func GeneratePublicLinkFileInGoogleDriver(srv *drive.Service, fileID string) (string, error) {
	//add permission role reader
	perm := &drive.Permission{
		Role: "reader",
		Type: "anyone",
	}

	_, err := srv.Permissions.Create(fileID, perm).Do()

	if err != nil {
		return "", err
	}

	//get public link file for viewer
	url := fmt.Sprintf("https://drive.google.com/uc?export=view&id=%s", fileID)

	return url, nil

}

func CountFilesInFolder(serviceAuth, folderID string) (int, error) {
	SCOPE := drive.DriveScope
	ctx := context.Background()
	srv, err := drive.NewService(ctx, option.WithCredentialsFile(serviceAuth), option.WithScopes(SCOPE))
	if err != nil {
		return 0, err
	}

	// Set the query to search for files in the specified folder
	query := fmt.Sprintf("'%s' in parents", folderID)

	// Execute the query to retrieve the list of files in the folder
	files, err := srv.Files.List().Q(query).Do()
	if err != nil {
		return 0, err
	}

	// Count the number of files in the folder
	count := len(files.Files)

	return count, nil
}

// Delete file in google driver
func DeleteFileInGoogleDriver(serviceAuth string) {
	SCOPE := drive.DriveScope
	ctx := context.Background()
	srv, err := drive.NewService(ctx, option.WithCredentialsFile(serviceAuth), option.WithScopes(SCOPE))
	if err != nil {
		fmt.Println("err src google driver api:", err)
	}

	query := fmt.Sprintf("'%s' in parents", config.GetEnv().GOOGLE_DRIVER_FOLDERID)

	files, err := srv.Files.List().Q(query).Do()
	if err != nil {
		fmt.Printf("Unable to retrieve files: %v", err)
	}

	oldestFile := files.Files[len(files.Files)-1]

	err = srv.Files.Delete(oldestFile.Id).Do()
	fmt.Println("Deleted file oldestFile.CreatedTime: ", oldestFile.CreatedTime)
	if err != nil {
		fmt.Printf("Unable to delete file: %v", err)
	}
}

func IsCommandAvailable(command string) bool {
	cmd := exec.Command("which", command)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
