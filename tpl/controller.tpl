package controllers

import "wstmart/models"

type <%@#&%> struct {
	Base
}

func (this *<%@#&%>) Index()  {
	
	if this.Ctx.Input.IsPost() {
		queryParam := new(models.<%@#&%>QueryParam)
		queryParam.Limit, _ = this.GetInt("limit", 10)
		page, _ := this.GetInt("page", 1)
		queryParam.Offset = (page - 1) * queryParam.Limit
		
		m := models.New<%@#&%>()
		
		data, count := m.PageQuery(queryParam)
		
		this.JsonOut(data, len(data))
	}

	this.Data["data"] = data 
}

func (this *<%@#&%>) Get() {
	id, _ := this.GetInt(":id", 0)
	m := models.New<%@#&%>()
	data := m.GetById(id)
	this.Data["data"] = data
	this.TplName = "<%@#&%>/detail.html"
}

func (this *<%@#&%>) Post() {
	mid, _ := this.GetInt(":id", 0)
	m := models.New<%@#&%>()
	id,err := m.Add(m)
	code := 0
	if err != nil {
		code = 99
	}
	this.JsonOut(m, id, code, err)
}

func (this *<%@#&%>) Put() {
	id, _ := this.GetInt(":id", 0)
	filters := make(map[string]interface{})
	
	m := models.New<%@#&%>()
	num, err := m.Update(filters, data)
	code := 0
	if err != nil {
		code = 99
	} 
	this.JsonOut(data, num, code, err)
}

func (this *<%@#&%>) Delete () {
	id, _ := this.GetInt(":id", 0)
	m := models.New<%@#&%>()
	num, err := m.Delete(id)
	code := 0 
	
	if err != nil {
		code = 99
	}

	this.JsonOut(id, num, code, err)

}