package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"runtime"
	"strconv"
	"time"
)

var CityToProvinces = []map[string]string{
	//"北京":   "北京市",
	{"泉州": "福建"},
	{"福州": "福建"},
	{"漳州": "福建"},
	{"厦门": "福建"},

	{"太原*": "山西"},
	{"太原": "山西"},
	{"大同": "山西"},

	{"武汉": "湖北"},
	{"宜昌": "湖北"},
	{"孝感": "湖北"},
	{"荆州": "湖北"},

	{"济南": "山东"},
	{"泰安": "山东"},

	{"广州": "广东"},
	{"珠海": "广东"},

	{"天津": "天津"},

	{"呼和浩特": "呼和浩特"},

	{"包头": "包头"},

	{"沈阳": "沈阳"},

	{"上海": "上海"},

	{"昆明": "昆明"},

	{"长沙": "长沙"},

	{"重庆": "重庆"},

	{"哈尔滨": "哈尔滨"},

	{"南昌": "南昌"},

	{"保定": "保定"},

	{"吴忠": "吴忠"},

	{"银川": "银川"},

	//"泾阳":   "陕西省",
	{"西宁": "西宁"},

	{"昆明*": "昆明*"},
}

type A struct {
	name string
	Sun  B
}

type B struct {
	name string
}

func min(v ...string) string {
	return ""
}

func MyMin(map1 map[string]string) string {
	var valueSlice []string
	for _, value := range map1 {
		valueSlice = append(valueSlice, value)
	}
	return min(valueSlice...)
}

func main() {

	path := os.Getenv("PATH")
	fmt.Println("$PATH:", path)

	myMap := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	MyMin(myMap)

	//startDateStr := "2023-03-25 00:00:00"
	//endDateStr := "2023-04-09 00:00:01"

	// 获取当前时间的纳秒级时间戳
	nanoTimestamp := time.Now().UnixNano()

	// 将纳秒级时间戳转换为毫秒级时间戳
	milliTimestamp := nanoTimestamp / int64(time.Millisecond)

	println(milliTimestamp)

	return

	num := 3.14159
	formatted := strconv.FormatFloat(num, 'f', 2, 64)
	fmt.Println(formatted) // 输出: 3.14

	return

	encryptPassword, err := bcrypt.GenerateFromPassword([]byte("23434"), bcrypt.DefaultCost)
	if err != nil {
		println("生成密码错误:", err.Error())
	}
	println("生成密码成功:", string(encryptPassword))
	return

	var a A = A{
		name: "结构1",
		Sun: B{
			name: "结构2",
		},
	}

	fmt.Printf("%v", a.name)
	return

	layout := "2006-01-02 15:04:05" // 根据你的时间格式更改
	t, _ := time.Parse(layout, "2023-02-28 16:08:07")

	//now := time.Now()

	//year, month, _ := t.Date()

	lt := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	lt.Format("2006-01-02")

	println(lt.Format("2006-01-02"))
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("Allocated memory: %v bytes\n", mem.Alloc)
	fmt.Printf("Total memory: %v bytes\n", mem.TotalAlloc)
	fmt.Printf("Heap memory: %v bytes\n", mem.HeapAlloc)
	//Alloc 表示当前已分配但未释放的内存量。
	//TotalAlloc 表示总共已分配的内存量。
	//HeapAlloc 表示当前堆内存中已分配但未释放的内存量。
	//注意：MemStats 结构体中还包含了许多其他的内存指标，可以根据需要进行查看。

}
