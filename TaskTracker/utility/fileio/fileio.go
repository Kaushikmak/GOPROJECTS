package fileio

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/kaushikmak/go-projects/TaskTracker/models"
)



const DIR_NAME = ".task-tracker"
const FILE_NAME = "data.json"


func EnsureStorage() (string,error) {
    home,err := os.UserHomeDir()
    if err != nil {
        return "",err
    }
    dir := filepath.Join(home,DIR_NAME)
    if err := os.MkdirAll(dir,0755); err != nil {
        return "",err
    }

    path := filepath.Join(dir,FILE_NAME)

    if _,err := os.Stat(path); os.IsNotExist(err){
        f,err := os.OpenFile(path,os.O_CREATE,0644)
        if err != nil {
            return "",nil
        }
        f.Close()
    }

    return path,nil
}

func Load(path string) ([]models.Task,error) {
    data,err := os.ReadFile(path)
    if err != nil {
        return nil,err
    }
    if len(data) == 0{
        return []models.Task{},nil
    }
    var tasks []models.Task
    if err := json.Unmarshal(data,&tasks); err != nil {
        return nil,err
    }
    return tasks,nil
}

func Save(path string,tasks []models.Task) error {
    data,err := json.MarshalIndent(tasks,""," ")
    if err != nil {
        return err
    }
    tmp := path + ".tmp"

    if err := os.WriteFile(tmp,data,0644); err != nil {
        return nil
    }

    return os.Rename(tmp,path)
}























