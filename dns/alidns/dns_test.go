package alidns

import (
	"github.com/cloudlibz/gocloud/aliauth"
	"testing"
)

func init() {
	aliauth.LoadConfig()
}

func TestCreateDns(t *testing.T) {
	var aliDNS Alidns
	createDNS := map[string]interface{}{
		"DomainName": "oddcn.cn",
		"RR":         "gocloud.test",
		"Type":       "A",
		"Value":      "202.106.0.20",
		"TTL":        600,
	}
	_, err := aliDNS.CreateDns(createDNS)
	if err != nil {
		t.Errorf("CreateDns Test Fail: %s", err)
		return
	}
	t.Logf("Ali DNS is created successfully.")
}

func TestCreateDNSBuilder(t *testing.T) {
	var aliDNS Alidns
	createDNS, err := NewCreateDNSBuilder().
		DomainName("oddcn.cn").
		RR("gocloud.test").
		Type("A").
		Value("202.106.0.20").
		TTL(600).
		Build()
	_, err = aliDNS.CreateDns(createDNS)
	if err != nil {
		t.Errorf("CreateDns Test Fail: %s", err)
		return
	}
	t.Logf("Ali DNS is created successfully.")
}

func TestDeleteDns(t *testing.T) {
	var aliDNS Alidns
	deleteDNS := map[string]interface{}{
		"RecordId": "9999985",
	}
	_, err := aliDNS.DeleteDns(deleteDNS)
	if err != nil {
		t.Errorf("DeleteDns Test Fail: %s", err)
		return
	}
	t.Logf("Ali DNS is deleted successfully.")
}

func TestDeleteDNSBuilder(t *testing.T) {
	var aliDNS Alidns
	deleteDNS, err := NewDeleteDNSBuilder().RecordId("9999985").Build()
	_, err = aliDNS.DeleteDns(deleteDNS)
	if err != nil {
		t.Errorf("DeleteDns Test Fail: %s", err)
		return
	}
	t.Logf("Ali DNS is deleted successfully.")
}

func TestListDns(t *testing.T) {
	var aliDNS Alidns
	list := map[string]interface{}{
		"DomainName":   "oddcn.cn",
		"PageNumber":   1,
		"PageSize":     20,
		"RRKeyWord":    "www",
		"TypeKeyWord":  "MX",
		"ValueKeyWord": "com",
	}
	resp, err := aliDNS.ListDns(list)
	if err != nil {
		t.Errorf("ListDns Test Fail: %s", err)
		return
	}
	t.Logf("Ali DNS is listed successfully.")
	t.Logf("%s", resp.(map[string]interface{})["body"])
}

func TestListDNSBuilder(t *testing.T) {
	var aliDNS Alidns
	list, err := NewListDNSBuilder().
		DomainName("oddcn.cn").
		PageNumber(1).
		PageSize(20).
		Build()
	resp, err := aliDNS.ListDns(list)
	if err != nil {
		t.Errorf("ListDns Test Fail: %s", err)
		return
	}
	t.Logf("Ali DNS is listed successfully.")
	t.Logf("%s", resp.(map[string]interface{})["body"])
}

func TestParseListDnsResp(t *testing.T) {
	var aliDNS Alidns
	list := map[string]interface{}{
		"DomainName": "oddcn.cn",
		"PageNumber": 1,
		"PageSize":   20,
	}
	resp, err := aliDNS.ListDns(list)
	if err != nil {
		t.Errorf("ListDns Test Fail: %s", err)
		return
	}
	listDnsResp, err := ParseListDnsResp(resp)
	if err != nil {
		t.Errorf("ListDns Test Fail: %s", err)
	}
	for _, value := range listDnsResp.DomainRecords.Record {
		t.Logf("%+v\n", value)
	}
}
