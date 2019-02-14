package models

type UsersQueryParam struct {
	BaseQueryParam
}

type Users struct {
	Id		int		`orm:"pk;column(userId)"`
	PayPwd		string		`orm:"column(payPwd)"`
	True_name		string		`orm:"column(true_name)"`
	UserType		string		`orm:"column(userType)"`
	UserName		string		`orm:"column(userName)"`
	UserFrom		string		`orm:"column(userFrom)"`
	LockMoney		string		`orm:"column(lockMoney)"`
	DataFlag		string		`orm:"column(dataFlag)"`
	CreateTime		string		`orm:"column(createTime)"`
	LoginName		string		`orm:"column(loginName)"`
	Brithday		string		`orm:"column(brithday)"`
	UserPhoto		string		`orm:"column(userPhoto)"`
	UserPhone		string		`orm:"column(userPhone)"`
	LastTime		string		`orm:"column(lastTime)"`
	LoginPwd		string		`orm:"column(loginPwd)"`
	UserEmail		string		`orm:"column(userEmail)"`
	UserTotalScore		int		`orm:"column(userTotalScore)"`
	LastIP		string		`orm:"column(lastIP)"`
	UserMoney		string		`orm:"column(userMoney)"`
	LoginSecret		int		`orm:"column(loginSecret)"`
	UserSex		string		`orm:"column(userSex)"`
	TrueName		string		`orm:"column(trueName)"`
	UserQQ		string		`orm:"column(userQQ)"`
	UserScore		int		`orm:"column(userScore)"`
	UserStatus		string		`orm:"column(userStatus)"`
	Table		string		`orm:"column(table)"`
} 

func NewUsers() *Users {
	obj := new(Users)
	return obj
}

 func (this *Users) TableName() string {
    return GetTableName("Users")
}

 func (this *Users) ListQuery(Params *UsersQueryParam) []*Users  {
	table := this.TableName()
	query := Db(table)

	total, _ := query.Count()
	data := make([]*Users, 0)
	
	if total > 0 {
		query.Limit(Params.Limit, Params.Offset).All(&data)
	}

	return data, total
 }

 func (this *Users) GetById(Id int) Users {
	var Users Users
	table := this.TableName()
	Db(table).Filter("userId", Id).Filter("dataFlag", 1).One(&Users)
	return Users
 }

 func (this *Users) Add(data *Users) (int64, error) {
	table := this.TableName()
	db := Db(table)
	i, _ := db.PrepareInsert()
	id, err := i.Insert(data)	 
	
	return id, err
 }

func (this *Users) Update(filter map[string]interface{}, data map[string]interface{}) (int64, error) {
	table := this.TableName()
	db := Db(table)
	for key, value := range filter {
		db = db.Filter(key, value)
	}
	
	num, err := db.Update(data)
	return num, err
}

func (this *Users) Delete(id int) (int64, error){
	table := this.TableName()
	db := Db(table)
	num, err := db.Filter("userId", id).Delete()
	return num, err
}
