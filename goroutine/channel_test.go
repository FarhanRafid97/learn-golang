package sync

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)


	go func(){
		time.Sleep(3 * time.Second)
		channel <- "farhan rafid syauqi"
	}()
	data := <- channel
	fmt.Println(data)
	close(channel)
}

func GiveMeResponse(channel chan string){
	time.Sleep(2 * time.Second)
	channel <- "Farhan rafid syauqi"
}

func TestChanelParams(t *testing.T){
	channel := make(chan string)


	go GiveMeResponse(channel)
	data := <- channel
	fmt.Println(data)
	close(channel)
}


func OnlyIn(channel chan<- string){
	time.Sleep(2 * time.Second)
	channel <- "farhan rafid syauqi"
	
	
}
func GiveMemyName(channel <-chan string){
	data := <- channel
	fmt.Println(data)
	
}
func TestChanelMyName(t *testing.T){
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go GiveMemyName(channel)
	time.Sleep(2* time.Second)
	

}

func TestBuffered(t *testing.T){
	channel := make(chan string,2)
	
	defer close(channel)

	channel <- "farhan"
	channel <- "rafid"
	fmt.Println("selesai")
	fmt.Println(<- channel)
	fmt.Println(<- channel)

}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)


	go func (){
		for i:= 0;i< 10 ;i++{
			channel <- "Perulangan ke" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel{
		fmt.Println("Meneriama data",data)
	}
	fmt.Println("Selesai")

}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go OnlyIn(channel1)
	go OnlyIn(channel2)

	counter := 0
	for {
		select{
		case data := <- channel1:
			fmt.Println("data dari channel 1", data)
			counter++
		
		case data := <- channel2:
			fmt.Println("data dari channel 2", data)
			counter++
		}
		if counter == 2{
			break
		}
	}

}

func TestDefaultChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go OnlyIn(channel1)
	go OnlyIn(channel2)

	counter := 0
	for {
		select{
		case data := <- channel1:
			fmt.Println("data dari channel 1", data)
			counter++
		
		case data := <- channel2:
			fmt.Println("data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu")
		} 
		if counter == 2{
			break
		}
	}

}

func TestMutex(t *testing.T) {
 x := 0
 var mutex sync.Mutex
 for i := 1; i<= 1000;i++{
	go func(){ 
		for j:=1;j <= 100;j++{
			mutex.Lock()
			x = x + 1
			mutex.Unlock()
		}
	}()
 }
 time.Sleep(5 * time.Second)
 fmt.Println(x)
}

type BankAccount struct{
	RWMutex sync.RWMutex
	Balance int
}

func(account *BankAccount) addBalance(amount int){
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func ( acount *BankAccount) GetBalance()int{
	acount.RWMutex.RLock()
	balance := acount.Balance
	acount.RWMutex.RUnlock()
	return balance

}

func TestRWMutex(t *testing.T) {
	account := *&BankAccount{}

	for i := 0;i < 100;i++{
		go func(){
				for j := 0;j< 100;j++{
					account.addBalance(1)
					fmt.Println( account.GetBalance())
				}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance",account.Balance)

}

type UserBalance struct{
	sync.Mutex
	User string
	Balance int
}

func (user *UserBalance) Lock(){
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock(){
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int){
	user.Balance = user.Balance + amount
}

func Transaksi(user1 *UserBalance, user2 *UserBalance,amount int){
 	user1.Lock()
 	fmt.Println("Lock user 1", user1.User)
	user1.Change(-amount)
time.Sleep(1 * time.Second)

	user2.Lock()
 	fmt.Println("Lock user 2", user2.User)
	user2.Change(amount)


	time.Sleep(1 * time.Second)
	user1.Unlock()
	user2.Unlock()
}

func TestUserBalanace(t *testing.T) {
	user1 := UserBalance{
		User: "farhan",
		Balance: 100000,
	}
	user2 := UserBalance{
		User: "Syauqi",
		Balance: 100000,
	}

	go Transaksi(&user1,&user2,10000)
	time.Sleep(5 * time.Second)
	fmt.Println("User" , user1.User, "Balanace", user1.Balance)
	fmt.Println("User" , user2.User, "Balanace", user2.Balance)

}