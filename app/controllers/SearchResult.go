package controllers

import (
	"github.com/revel/revel"
	"revelSample/app/models"
	"revelSample/app/dbconnector"
)

func (c App) SearchResult(name string) revel.Result {

	// modelsを作成し、入力値をセット
	inputData := new(models.SearchInput)
	inputData.SetInput(name)

	// 社員情報を検索
	results, err:= dbconnector.Search(name)
	if err != nil {
		c.RenderError(err)
	}

	return c.Render(inputData,results)
}
