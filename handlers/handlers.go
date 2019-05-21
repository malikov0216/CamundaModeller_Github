package handlers

import (
	conf "../config"
	"../models"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"os"
	"strings"
	"time"
)

func userPass (encodedData string) (string, string) {
	decodedString, _ := base64.StdEncoding.DecodeString(encodedData[6:])
	loginData := string(decodedString)
	slicedUserPass := strings.Split(loginData, ":")
	return slicedUserPass[0], slicedUserPass[1] //User, Pass
}
// @Summary Camunda Modeller
// @Description Serves changing in diagram and upload/update it in GitHub
// @Accept  multipart/form-data
// @Produce  json
// @Param deployment-name formData string false "Repository name"
// @Param diagram_1.bpmn formData file false "Camunda Modeller"
// @Success 200 {string} string "answer"
// @Header 200 {string} string "Header"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Security BasicAuth
// @Router /deployment/create [post]
func CamundaModeller (c *gin.Context) {
	// UUID for the name of folder
	u1, _ := uuid.NewV4()


	errorToPrint := c.Errors.ByType(gin.ErrorTypePublic).Last()
	//Initiate config
	config, err := conf.LoadConfiguration("./config/config.json")
	if err != nil {
		c.JSON(401, errorToPrint.Meta)
		return
	}

	//Get Endpoit (repository name) from Camunda
	endpoint := c.PostForm("deployment-name")

	//Get Username, Password from Camunda
	username, password := userPass(c.Request.Header.Get("Authorization"))

	var (
		directory = u1.String()
		url = config.RepUrl + endpoint //URL FROM CONFIG
		fileName = config.FileName
	)

	msg := new(models.Errors)
	//Cloning from repository
	if err := gitClone(username, password, directory, url, plumbing.ReferenceName(config.BranchName)); err != nil {
		msg.Message = err.Error()
		c.JSON(404, msg.Message)
		return
	}

	//Change cloned file to file that we got from Camunda
	a, err  := c.FormFile("diagram_1.bpmn")
	if err != nil {
		msg.Message = err.Error()
		c.JSON(404, msg.Message)
		return
	}
	if err := c.SaveUploadedFile(a, directory + "/" + fileName); err != nil {
		msg.Message = err.Error()
		c.JSON(404, msg.Message)
		return
	}

	//Committing
	if err := gitCommit(username, directory, fileName); err != nil {
		msg.Message = err.Error()
		c.JSON(404, msg.Message)
		return
	}
	//Pushing
	if err := gitPush(username, password, directory); err != nil {
		msg.Message = err.Error()
		c.JSON(404, msg.Message)
		return
	}

	//Deleting files
	if err = deleteFile(directory); err != nil {
		msg.Message = err.Error()
		c.JSON(404, msg.Message)
		return
	}
}

// Clone the given repository to the given directory
func gitClone (username, password, directory, url string, branch plumbing.ReferenceName) error{
	Info("git clone %s %s", username, directory)

	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: username,
			Password: password,
		},
		ReferenceName: "refs/heads/" + branch,
		URL:      url,
		Progress: os.Stdout,
	})
	if err != nil {
		return err
	}
	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	if err != nil {
		return err
	}
	// ... retrieving the commit object
	_, err = r.CommitObject(ref.Hash())
	if err != nil {
		return err
	}

	return err
}
func gitCommit(username, directory, fileName string) error{
	// Opens an already existing repository.
	r, err := git.PlainOpen(directory)
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	//git add
	Info("git add " + fileName)
	_, err = w.Add(fileName)
	if err != nil {
		return err
	}

	// We can verify the current status of the worktree using the method Status.
	Info("git status --porcelain")
	_, err = w.Status()
	if err != nil {
		return err
	}
	Info("git commit -m \"example go-git commit\"")
	commit, err := w.Commit("commiting camunda file", &git.CommitOptions{
		Author: &object.Signature{
			Name:  username,
			Email: "asd@homebank.kz",
			When:  time.Now(),
		},
	})

	if err != nil {
		return err
	}

	// Prints the current HEAD to verify that all worked well.
	Info("git show -s")
	_, err = r.CommitObject(commit)
	if err != nil {
		return err
	}
	return err
}
func gitPush(username, password, directory string) error{
	path := directory
	r, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	Info("git push")
	// push using default options
	err = r.Push(&git.PushOptions{
		Auth:&http.BasicAuth{
			Username:username,
			Password:password,
		},

	})
	if err != nil {
		return err
	}
	return err
}
func deleteFile(directory string) error{
	err := os.RemoveAll(directory)
	return err
}