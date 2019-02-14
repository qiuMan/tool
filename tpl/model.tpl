package models

type <%@#&%>QueryParam struct {
	BaseQueryParam
}

<%@#struct%>

func New<%@#&%>() *<%@#&%> {
	obj := new(<%@#&%>)
	return obj
}

 func (this *<%@#&%>) TableName() string {
    return GetTableName("<%@#&%>")
}

 func (this *<%@#&%>) ListQuery(Params *<%@#&%>QueryParam) []*<%@#&%>  {
	table := this.TableName()
	query := Db(table)

	total, _ := query.Count()
	data := make([]*<%@#&%>, 0)
	
	if total > 0 {
		query.Limit(Params.Limit, Params.Offset).All(&data)
	}

	return data, total
 }

 func (this *<%@#&%>) GetById(Id int) <%@#&%> {
	var <%@#&%> <%@#&%>
	table := this.TableName()
	Db(table).Filter("<%@#ID%>", Id).Filter("dataFlag", 1).One(&<%@#&%>)
	return <%@#&%>
 }

 func (this *<%@#&%>) Add(data *<%@#&%>) (int64, error) {
	table := this.TableName()
	db := Db(table)
	i, _ := db.PrepareInsert()
	id, err := i.Insert(data)	 
	
	return id, err
 }

func (this *<%@#&%>) Update(filter map[string]interface{}, data map[string]interface{}) (int64, error) {
	table := this.TableName()
	db := Db(table)
	for key, value := range filter {
		db = db.Filter(key, value)
	}
	
	num, err := db.Update(data)
	return num, err
}

func (this *<%@#&%>) Delete(id int) (int64, error){
	table := this.TableName()
	db := Db(table)
	num, err := db.Filter("<%@#ID%>", id).Delete()
	return num, err
}