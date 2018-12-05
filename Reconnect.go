package main


import (
"fmt"
	"time"
	"os/exec"
	"net/http"
	"io/ioutil"
	"strings"
	"encoding/json"
	"os"
)

type Account struct {
	Username string
	Password string
	QueryString string
}


func main() {


	for
	{
		time.Sleep(1000000000)
		netWorkStatus := NetWorkStatus()
		//fmt.Println(netWorkStatus)
		if netWorkStatus == true {

			continue
		}else {
			fmt.Println("reconnecting....")
			data, _ := ioutil.ReadFile(getCurrentPath()+"account.json")
			account :=Account{}
			errJson := json.Unmarshal(data,&account)
			if errJson!=nil{
				fmt.Println("username and password marshal erro:",errJson)
			}

			url := "http://211.87.158.84/eportal/InterFace.do?method=login"
			fmt.Println(account.Username,"pass:",account.Password)
			payload := strings.NewReader("userId="+account.Username+"&password="+
				account.Password+"&service=internet&queryString="+
				account.QueryString+"&operatorPwd=&operatorUserId=&validcode=&passwordEncrypt=false&undefined=")

			req, _ := http.NewRequest("POST", url, payload)

			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Add("cache-control", "no-cache")
			req.Header.Add("Postman-Token", "99f9e080-8a59-4b57-a3c1-918dfd652c88")

			res, _ := http.DefaultClient.Do(req)

			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)

			fmt.Println(res)
			fmt.Println(string(body))
		}
	}

}
func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	checkErr(err)
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func NetWorkStatus() bool {

	fmt.Println("checking networking")
	status := ExecCommand("ping www.baidu.com -c 5")
	if(len(status)<10){
		return false
	}else {
		fmt.Println("networking ok")
		return true
	}

}
func ExecCommand(strCommand string)(string){
	cmd := exec.Command("/bin/bash", "-c", strCommand)


	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil{
		fmt.Println("Execute failed when Start:" + err.Error())
		return ""
	}

	out_bytes, _ := ioutil.ReadAll(stdout)
	stdout.Close()

	if err := cmd.Wait(); err != nil {
		fmt.Println("Execute failed when Wait:" + err.Error())
		return ""
	}
	return string(out_bytes)
}





