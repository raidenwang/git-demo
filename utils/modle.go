package utils

import (
	"fmt"
	"os/exec"
)

type User struct {
	Name     string
	password int64
}
type FamilyAccount struct {
	input   float64
	output  float64
	balance float64
	details string
	note    string
	loop    bool
	key     string
	choice  int
}

func NewUser() *User {
	return &User{
		Name:     "",
		password: 123456,
	}
}
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		input:   0.0,
		output:  0.0,
		balance: 0.0,
		details: "收支\t\t账户余额\t\t本次收支金额说明：",
		note:    "",
		loop:    true,
		key:     "",
		choice:  0,
	}
}
func (user *User) RevisePassword() {
	fmt.Println("欢迎使用皮卡丘版家庭收支记账软件")
	fmt.Println("请输入用户名:")
	fmt.Scanln(&user.Name)
	fmt.Printf("欢迎用户%v\n", user.Name)
	fmt.Printf("默认密码是%v\n", user.password)
	fmt.Println("你是否要修改密码 Yes/No")
	var choice string
	fmt.Scanln(&choice)
	if choice == "Yes" {
		fmt.Println("请输入新密码：")
		fmt.Scanln(&user.password)
		fmt.Println("密码修改成功")
		cmd := exec.Command("CLS")
		cmd.Run()
	} else {
		return
	}
}
func (user *User) GetPassword() {
	fmt.Println("请输入密码：")
	var password int64
	var icount int = 0
	fmt.Scanln(&password)
	for {
		if icount == 3 {
			fmt.Println("请过一会儿在试")
			return
		}
		if password == user.password {
			fmt.Println("密码正确")
			break
		} else {
			fmt.Println("密码错误")
			icount++
			fmt.Printf("你还剩下%d次机会\n", 3-icount)
		}
	}
}
func (this *FamilyAccount) ShowDetails() {
	fmt.Println("-----------------当前收支详情--------------")
	if this.balance == 0 {
		fmt.Println("家庭账户余额为0")
	} else {
		fmt.Println(this.details)
	}
}
func (this *FamilyAccount) Income() {
	fmt.Println("本次收入金额")
	fmt.Scanln(&this.input)
	this.balance += this.input
	fmt.Println("本次收入来源")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入%v\t账户余额%v\t\t%v", this.input, this.balance, this.note)
}
func (this *FamilyAccount) Outcome() {
	fmt.Println("本次支出金额")
	fmt.Scanln(&this.output)
	if this.output > this.balance {
		fmt.Println("余额不足")
		return
	} else {
		this.balance -= this.output
		fmt.Println("本次支出去向")
		fmt.Scanln(&this.note)
		this.details += fmt.Sprintf("\n支出%v\t账户余额%v\t\t%v", this.output, this.balance, this.note)
	}
}
func (this *FamilyAccount) Exit() {
	fmt.Println("你确定要退出吗？Yes/No")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Yes" || this.key == "No" {
			break
		} else {
			fmt.Println("您的输入有误，请重新输入")
		}

	}
	if this.key == "Yes" {
		this.loop = false
		fmt.Println("已退出家庭收支记账软件")
	}
}
func (this *FamilyAccount) MainMenu(user *User) {
	user.RevisePassword()
	fmt.Println("欢迎使用皮卡丘版家庭收支记账软件")
	user.GetPassword()
	fmt.Printf("欢迎用户%v\n", user.Name)
	for {
		fmt.Println("-----------家庭收支记账软件---------------")
		fmt.Println("----------- 1 收支明细-------------------")
		fmt.Println("----------- 2 登记收入-------------------")
		fmt.Println("----------- 3 登记支出-------------------")
		fmt.Println("----------- 4 退出   -------------------")
		fmt.Print("请选择（1-4）：")
		fmt.Scanln(&this.choice)
		switch this.choice {
		case 1:
			fmt.Println("----------- 1 收支明细-------------------")
			this.ShowDetails()
		case 2:
			fmt.Println("----------- 2 登记收入-------------------")
			this.Income()
		case 3:
			fmt.Println("----------- 3 登记支出-------------------")
			this.Outcome()
		case 4:
			fmt.Println("----------- 4 退出   -------------------")
			this.Exit()
		default:
			fmt.Println("请选择正确选项")
		}
		if !this.loop {
			break
		}
	}
}
