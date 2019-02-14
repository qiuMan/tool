package models

type AdsQueryParam struct {
	BaseQueryParam
}

type Ads struct {
	Id		int		`orm:"pk;column(adId)"`
	PositionType		string		`orm:"column(positionType)"`
	AdPositionId		int		`orm:"column(adPositionId)"`
	AdName		string		`orm:"column(adName)"`
	AdStartDate		string		`orm:"column(adStartDate)"`
	AdSort		int		`orm:"column(adSort)"`
	AdClickNum		int		`orm:"column(adClickNum)"`
	DataFlag		string		`orm:"column(dataFlag)"`
	CreateTime		string		`orm:"column(createTime)"`
	AdFile		string		`orm:"column(adFile)"`
	AdURL		string		`orm:"column(adURL)"`
	AdEndDate		string		`orm:"column(adEndDate)"`
} 

func NewAds() *Ads {
	obj := new(Ads)
	return obj
}

 func (this *Ads) TableName() string {
    return GetTableName("Ads")
}

 func (this *Ads) ListQuery(Params *AdsQueryParam) []*Ads  {
	table := this.TableName()
	query := Db(table)

	total, _ := query.Count()
	data := make([]*Ads, 0)
	
	if total > 0 {
		query.Limit(Params.Limit, Params.Offset).All(&data)
	}

	return data, total
 }

 func (this *Ads) GetById(Id int) Ads {
	var Ads Ads
	table := this.TableName()
	Db(table).Filter("adId", Id).Filter("dataFlag", 1).One(&Ads)
	return Ads
 }

 func (this *Ads) Add(data *Ads) (int64, error) {
	table := this.TableName()
	db := Db(table)
	i, _ := db.PrepareInsert()
	id, err := i.Insert(data)	 
	
	return id, err
 }

func (this *Ads) Update(filter map[string]interface{}, data map[string]interface{}) (int64, error) {
	table := this.TableName()
	db := Db(table)
	for key, value := range filter {
		db = db.Filter(key, value)
	}
	
	num, err := db.Update(data)
	return num, err
}

func (this *Ads) Delete(id int) (int64, error){
	table := this.TableName()
	db := Db(table)
	num, err := db.Filter("adId", id).Delete()
	return num, err
}
