package cookie

import (
	"net/url"
	"io/ioutil"
	"strings"
	"net/http"
)

func AccessCookie ()(cookie string){
    v := url.Values{}
	v.Set("__VIEWSTATE", "/wEPDwULLTEwMDA0NDY4OTNkZNs/q7P6bnuSJlfMjnry7zi/8yUw")
	v.Set("txtUserName", "fb")
	v.Set("txtPass", "654321")

	body := ioutil.NopCloser(strings.NewReader(v.Encode()))
	datalen := string(len([]rune(v.Encode())))
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://10.188.151.44/login.aspx", body)

	req.Header.Set("Accept", "application/x-ms-application, image/jpeg, application/xaml+xml, image/gif, image/pjpeg, application/x-ms-xbap, application/vnd.ms-excel, application/vnd.ms-powerpoint, application/msword, */*")
	req.Header.Set("Referer", "http://10.188.151.44/Pages/Report/WellMonthlyRecord.aspx")
	req.Header.Set("Accept-Language", "zh-CN")
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/7.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; InfoPath.2)")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Content-Length", string(datalen))
	req.Header.Set("Connection", "Keep-Alive")
	//req.Header.Set("Cookie", cookie)
	req.Host = "10.188.151.44"
	req.Header.Set("Pragma", "no-cache")

	resp, _ := client.Do(req)
	cookies := resp.Cookies()[0].Raw
	s := strings.Split(cookies, ";")[0]
	return s
}
