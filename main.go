package main

import (
	"go-crud-demo/Dao"
)

func main() {
	Dao.InitDB()
	
	InitRouter()
	//go Controller.NewUser(*db, *r)

	// 插入单条数据
	//user := User{Name: "陈海潇", Age: "21", Message: "654"}
	//db.Create(&user)
	// 插入多条数据
	//users := []User{
	//	{Name: "陈海潇1", Age: 21, Message: "6542"},
	//	{Name: "陈海潇2", Age: 213, Message: "6541"},
	//	{Name: "陈海潇3", Age: 214, Message: "6544"},
	//}
	//db.Create(&users)
	//db.CreateInBatches(users, 100) // 设置分批添加数据，每批100条数据
	//for _, user := range users {
	//	fmt.Println(user.ID)
	//}
	// 查询单条数据
	//user := User{}
	//db.First(&user, 1) // 通过ID进行查询一条信息
	//db.First(&user, "name=?", "陈海潇1") // 条件查询
	//fmt.Println(user)
	// 查询多条数据
	//users := []User{}
	//db.Find(&users) // 查询所有信息
	//db.Find(&users, []int{1, 2, 3})   // 默认使用ID主键进行查询
	//db.Find(&users, "name=?", "陈海潇1") // 条件查询
	//fmt.Println(users)
	// 修改数据
	//user := User{}
	//var user User
	//db.Model(&user).Where("name=?", "陈海潇2").Update("Age", 999) // 条件更新单个字段
	//db.Model(&user).Where("name=?", "陈海潇1").Updates(User{Name: "陈海潇", Age: 111}) // 条件更新多个字段
	// 删除数据
	//user := User{}
	//db.Where("name=?", "陈海潇3").Delete(&user)

	// 接口
	//r := gin.Default()
	//
	//// mapper层
	//// 测试
	//r.GET("/", func(context *gin.Context) {
	//	context.JSON(200, gin.H{
	//		"message": "请求成功",
	//	})
	//})
	//// 注册
	//r.POST("/douyin/user/register/", func(c *gin.Context) {
	//	var user User
	//	user.Name = c.Query("username")
	//	user.Password = c.Query("password")
	//	fmt.Println(user)
	//	db.Create(&user)
	//	if user.ID != 0 {
	//		c.JSON(200, gin.H{
	//			"code":    200,
	//			"message": "注册成功",
	//			"data":    user,
	//		})
	//	} else {
	//		c.JSON(400, gin.H{
	//			"code":    400,
	//			"message": "注册失败",
	//			"data":    gin.H{},
	//		})
	//	}
	//})
	//r.GET("/douyin/user/", func(c *gin.Context) {
	//	var user User
	//	result := db.Where("name=?", "106703658@qq.com").Where("password=?", "123456").Find(&user)
	//	if result.RowsAffected == 0 {
	//		c.JSON(400, gin.H{
	//			"code":    400,
	//			"message": "登陆失败",
	//			"data":    gin.H{},
	//		})
	//	} else {
	//		c.JSON(200, gin.H{
	//			"code":    200,
	//			"message": "登陆成功",
	//			"data":    user,
	//		})
	//	}
	//})
	//// 登陆
	//r.POST("/douyin/user/login/", func(c *gin.Context) {
	//	var user User
	//	username := c.Query("username")
	//	password := c.Query("password")
	//	result := db.Where("name=?", username).Where("password=?", password).Find(&user)
	//	fmt.Println(user)
	//	if result.RowsAffected == 0 {
	//		c.JSON(400, gin.H{
	//			"code":    400,
	//			"message": "登陆失败",
	//			"data":    gin.H{},
	//		})
	//	} else {
	//		c.JSON(200, gin.H{
	//			"code":    200,
	//			"message": "登陆成功",
	//			"data":    user,
	//		})
	//	}
	//})
	//// 查询用户
	//// 通过用户ID进行查找用户
	//r.POST("/user/findById", func(c *gin.Context) {
	//	var user User
	//	if err := c.ShouldBindJSON(&user); err != nil {
	//		c.JSON(200, gin.H{
	//			"code":    400,
	//			"message": "查询失败",
	//			"data":    gin.H{},
	//		})
	//	} else {
	//		result := db.Find(&user, user.ID)
	//		if result.RowsAffected == 0 {
	//			c.JSON(200, gin.H{
	//				"code":    400,
	//				"message": "不存在该用户",
	//				"data":    gin.H{},
	//			})
	//		} else {
	//			c.JSON(200, gin.H{
	//				"code":    200,
	//				"message": "查询成功",
	//				"data":    user,
	//			})
	//		}
	//	}
	//})
	//// 通过用户名进行查找用户
	//r.POST("/user/findByName", func(c *gin.Context) {
	//	var user User
	//	var users []User
	//	if err := c.ShouldBindJSON(user); err != nil {
	//		c.JSON(200, gin.H{
	//			"code":    400,
	//			"message": "查询失败",
	//			"data":    err,
	//		})
	//	} else {
	//		result := db.Find(&users, "name=?", user.Name)
	//		if result.RowsAffected == 0 {
	//			c.JSON(200, gin.H{
	//				"code":    400,
	//				"message": "不存在用户" + user.Name,
	//				"data":    gin.H{},
	//			})
	//		} else {
	//			c.JSON(200, gin.H{
	//				"code":    200,
	//				"message": "查询成功",
	//				"data":    users,
	//			})
	//		}
	//	}
	//})
	// 修改用户信息
	//r.POST()
	// 设置端口号
}
