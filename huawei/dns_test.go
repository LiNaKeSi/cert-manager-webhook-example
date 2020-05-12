package huawei

import (
	"testing"
)

// const (
// 	appKey    = "app key"
// 	appSecret = "app secret"
// 	appDomain = "linakesi.com"
// )

func TestPing(t *testing.T) {
	c, err := NewHuaweiDNSClient(appKey, appSecret, appDomain)
	if err != nil {
		t.Fatal(err)
	}
	rs, err := c.List("")
	if err != nil {
		t.Fatal(err)
	}
	if len(rs) == 0 {
		t.Fatal("Can't find any dns record")
	}
}

func TestAddDel(t *testing.T) {
	c, err := NewHuaweiDNSClient(appKey, appSecret, appDomain)
	if err != nil {
		t.Fatal(err)
	}
	tname := "testsnyh.linakesi.com"
	err = c.AddDomainRecord(tname, "TXT", "123")
	if err != nil {
		t.Fatal(err)
	}
	c.DeleteDomainRecord(tname, "TXT")
}
