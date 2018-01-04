package download

import (
	"databse/xml2txt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	//"fmt"

	"fmt"
)

func Download(ddstation string, well []string, cookie string, view string) {
	for i := 0; i < len(well); i++ {
		v := url.Values{}
		v.Set("__EVENTTARGET", "")
		v.Set("__EVENTARGUMENT", "")
		v.Set("__LASTFOCUS", "")
		v.Set("__VIEWSTATE", view)
		v.Set("ddlStation", ddstation)
		v.Set("ddlWells", well[i])
		v.Set("tbStartDate", "2006-01-01")
		v.Set("tbEndDate", "2017-12-21")
		v.Set("btnExport", "导  出")

		body := ioutil.NopCloser(strings.NewReader(v.Encode()))
		datalen := string(len([]rune(v.Encode())))
		client := &http.Client{}
		req, _ := http.NewRequest("POST", "http://10.188.151.44/Pages/Report/WellMonthlyRecord.aspx", body)

		req.Header.Set("Accept", "application/x-ms-application, image/jpeg, application/xaml+xml, image/gif, image/pjpeg, application/x-ms-xbap, application/vnd.ms-excel, application/vnd.ms-powerpoint, application/msword, */*")
		req.Header.Set("Referer", "http://10.188.151.44/Pages/Report/WellMonthlyRecord.aspx")
		req.Header.Set("Accept-Language", "zh-CN")
		req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/7.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; InfoPath.2)")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Accept-Encoding", "gzip, deflate")
		req.Header.Set("Content-Length", string(datalen))
		req.Header.Set("Connection", "Keep-Alive")
		req.Header.Set("Cookie", cookie)
		req.Host = "10.188.151.44"
		req.Header.Set("Pragma", "no-cache")
		resp, _ := client.Do(req)
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data))
		xml2txt.Xml2txt(data)
		fmt.Println("%s has been downloaded", well[i])
	}
}
