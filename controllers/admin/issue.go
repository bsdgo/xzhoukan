package admin

import (
	"fmt"
	"strings"
	"time"
	"xzhoukan.com/models"
)

type IssueController struct {
	BaseController
}

func (this *IssueController) Add() {
	this.Data["Current"] = this.getTime()
	this.display()
}

func (this *IssueController) DoAdd() {
	var title string = strings.TrimSpace(this.GetString("title"))
	var subTitle string = strings.TrimSpace(this.GetString("sub_title"))
	issueNo, _ := this.GetInt16("issue_num")
	publishDate, _ := time.Parse("2006-01-02", this.GetString("publish_date"))
	fmt.Println(publishDate)
	fmt.Println(this.GetString("publish_date"))
	var desc string = strings.TrimSpace(this.GetString("desc"))

	var issue models.Issue
	issue.Title = title
	issue.SubTitle = subTitle
	issue.IssueNum = issueNo
	issue.PublishDate = publishDate
	issue.Desc = desc

	err := issue.Add()
	if err != nil {
		fmt.Println(err)
	}

	//this.Redirect("/admin/issue/list", 302)
	this.EnableRender = false
}
