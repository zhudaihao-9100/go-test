package homework

import (
	"fmt"
	"go-test/testutils"
	"testing"

	"gorm.io/gorm"
)

// 只出现一次数字 测试用例
func TestSingleNumber(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
	}

	for _, item := range tests {
		got := SingleNumber(item.input)
		if got != item.expect {
			t.Errorf("输入%v，预期%d，得到%d", item.input, item.expect, got)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	// 回文数测试用例
	tests := []struct {
		input  int
		expect bool
	}{
		{121, true},
		{-121, false},
		{0, true},
		{12321, true},
		{123, false},
	}

	for _, item := range tests {

		got := IsPalindrome(item.input)

		if got != item.expect {
			t.Errorf("s输入: %d ，预期: %v ，实际得到: %v", item.input, item.expect, got)
		}
	}

}

func TestIsValidParentheses(t *testing.T) {
	tests := []struct {
		input  string
		expect bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"([])", true},
		{"(]", false},
		{"{[(])}", false}, //错误嵌套
		{"(", false},      //左括号多余
		{")", false},      //开头右括号
		{"((()))", true},
	}

	for _, item := range tests {
		got := IsValidParentheses(item.input)
		if got != item.expect {
			t.Errorf("输入字符串：%s，预期：%v，实际得到：%v", item.input, item.expect, got)
		}
	}

}

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		input  []string
		expect string
	}{
		{[]string{"flott", "floii", "flff"}, "fl"},
		{[]string{"mcctt", "mccii", "mcoff"}, "mc"},
	}

	for _, item := range tests {
		got := LongestCommonPrefix(item.input)
		if got != item.expect {
			t.Errorf("输入%v 预期:%s 实际:%s", item.input, item.expect, got)
		}

	}

}

func TestPlusOne(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
		{[]int{8, 9, 9}, []int{9, 0, 0}},
		{[]int{8, 8, 9}, []int{8, 9, 0}},
	}

	for _, item := range tests {
		got := PlusOne(item.input)

		match := true

		if len(got) != len(item.expect) {
			match = false
		} else {
			for i := range got {

				if got[i] != item.expect[i] {
					match = false
				}

			}
		}

		if !match {
			t.Errorf("输入%v 预期%v 实际%v", item.input, item.expect, got)
		}
	}

}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		expect []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, item := range tests {
		got := TwoSum2(item.nums, item.target)
		match := false
		//两种顺序都算正确 [a,b] 或 [b,a]
		if (got[0] == item.expect[0] && got[1] == item.expect[1]) ||
			(got[0] == item.expect[1] && got[1] == item.expect[0]) {
			match = true
		}
		if !match {
			t.Errorf("nums:%v target:%d 预期%v 实际%v",
				item.nums, item.target, item.expect, got)
		}
	}
}

//=============== 进阶gorm  ========================

// 用户
type User struct {
	ID   uint
	Name string

	Posts []Post //一对多

	PostCount int64 // 用户文章数量统计，仅存储，不写钩子
}

// 文章
type Post struct {
	ID    uint
	Title string

	UserID   uint
	User     User      //preload 预加载 预留
	Comments []Comment //一对多： 一篇文章多条评论

	CommentStat string // 评论状态：正常 / 无评论
}

// 评价
type Comment struct {
	ID      uint
	Content string

	PostID uint
	Post   Post //preload 预加载 预留
}

// 定义中间结构体 接受统计结果
type PostCommentCount struct {
	Post

	CommentCount int64
}

//  题目1：模型定义

func TestGorm(t *testing.T) {

	t.Helper()

	db := testutils.NewTestDB(t, "user.db")

	//建表
	if err := db.AutoMigrate(&User{}, &Post{}, &Comment{}); err != nil {
		t.Fatalf("auto migrate: %v", err)
	}

	users := []User{
		{
			Name: "张三",
		},
		{
			Name: "李四",
		},
	}

	err := db.Create(&users).Error

	if err != nil {
		t.Fatalf("create user err: %v", err)
	}

	posts := []Post{
		{
			Title:  "Go 入门",
			UserID: users[0].ID,
		},
		{
			Title:  "GORM 入门",
			UserID: users[0].ID,
		},
		{
			Title:  "Go 协程goroutine 入门",
			UserID: users[1].ID,
		},
	}

	if err := db.Create(&posts).Error; err != nil {
		t.Fatalf("create post err: %v", err)
	}

	t.Logf("=======post Create")

	comments := []Comment{
		{Content: "写的不错", PostID: posts[0].ID},
		{Content: "学到了", PostID: posts[0].ID}, // Go入门 共2条

		{Content: "1很好", PostID: posts[1].ID},
		{Content: "2很好", PostID: posts[1].ID},
		{Content: "3很好", PostID: posts[1].ID},
		{Content: "4很好", PostID: posts[1].ID},
		{Content: "5很好", PostID: posts[1].ID}, // GORM学习 共5条（最多）

		{Content: "GMP讲解很棒", PostID: posts[2].ID}, // Go协程 1条
	}

	if err := db.Create(&comments).Error; err != nil {
		t.Fatalf("create comment err %v", err)
	}

	u, err := QueryUserPostsAndComments(db, users[1].ID)

	if err != nil {
		t.Fatalf("query err: %v", err)
	}
	t.Logf("查询结果：%+v", u)

	p, err := QueryMostCommentPost(db)

	if err != nil {
		t.Fatalf("query err: %v", err)
	}
	t.Logf("查询结果2：%+v", p)

	//获取一条评论
	var delComment Comment
	err = db.Model(&Comment{}).Where("post_id = ?", posts[0].ID).First(&delComment).Error

	if err != nil {
		t.Fatal(err)
	}

	//删除评论实体
	err = db.Delete(&delComment).Error
	if err != nil {
		t.Fatal(err)
	}

	//删除
	err = db.Delete(&comments[3]).Error

	t.Logf("已删除评论 id=%d", delComment.ID)
}

// 查询 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息
func QueryUserPostsAndComments(db *gorm.DB, userId uint) (*User, error) {
	var user User

	//嵌套预加载
	err := db.Preload("Posts.Comments").First(&user, userId).Error

	if err != nil {
		return nil, err
	}

	return &user, nil

}

// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
func QueryMostCommentPost(db *gorm.DB) (*Post, error) {
	var res PostCommentCount

	err := db.
		//查询主表
		Model(&Post{}).
		//Select 作用 指定查询返回那些列
		//posts.*    posts表全部字段
		//COUNT(comments.id)   其中COUNT 是SQL固定关键字聚合函数  统计行数 ；统计当前分组里 comments.id 不为null一共有多少行
		//AS  是SQL 固定关键字  计算结果起别名
		Select("posts.*, COUNT(comments.id) AS comment_count").

		// LEFT JOIN 做链接 固定写法  以左边表的数据为基准，左边全部保留，右边匹配不到就填null
		// comments 代表要关联的副表(右表) 数据库副表名称
		// ON 固定  后面写两张表的匹配条件
		//comments.post_id   里面的comments  副表；post_id评论表里的外键字段，存的文章ID
		// =  等于，匹配条件
		//posts.id  里面的posts 主表， id 主表里面的id
		Joins("LEFT JOIN comments ON comments.post_id = posts.id"). //Joins 是表连接，把多张表联合查询

		//按文章id分组
		Group("posts.id").
		//Order排序； 按字段 comment_count 排序 ； DESC 降序
		Order("comment_count DESC").
		//限制查询返回的行数 只取1条记录
		Limit(1).
		//聚会查询（对一批数据做计算，汇总统计，不是单纯把原始行拿出来）
		//聚会查询推荐scan
		Scan(&res).Error

	if err != nil {
		return nil, err
	}

	if res.Post.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	//预加载，把评论填充进Post.Comments
	// Preload  作用  预加载关联数据，解决GORM 一对多/一对一 关联 查询
	// 查询主表数据同时，额外把关联的子数据查询出来，填充到结构体的关联切片/对象字段
	//Preload 参数 是结构体里面的字段名

	//First 取查询结果第一条记录 按主键排序
	// First  参数 1：接收对象 （必须传指针），参数2 主键值
	err = db.Preload("Comments").First(&res.Post, res.Post.ID).Error
	if err != nil {
		return nil, err
	}
	return &res.Post, nil
}

// Post 的 BeforeCreate 钩子（指针接收者，必须用 tx）

// BeforeCreate 创建文章前触发
func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {

	fmt.Printf("=======post Create===BeforeCreate ")
	//用户文章计数+1  使用expr 数据库层面自增 并发安全
	err = tx.
		Model(&User{}).
		Where("id = ?", p.UserID).
		Update("post_count", gorm.Expr("post_count + ?", 1)).Error

	return err

}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
// Comment 的 BeforeDelete 钩子  CommentStat
func (c *Comment) BeforeDelete(tx *gorm.DB) (err error) {
	var remainCnt int64

	err = tx.
		//查询表  Comment
		Model(&Comment{}).
		//条件 post_id = 对应id
		//AND  并且两个条件同时满足
		//<>   SQL 语句里的不等于
		Where("post_id = ? AND <> ?", c.PostID, c.ID).
		Count(&remainCnt).Error

	if remainCnt == 0 {
		err = tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_stat", "无评论").Error
	}

	return err

}
