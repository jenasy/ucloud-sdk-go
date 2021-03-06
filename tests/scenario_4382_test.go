// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/services/ulb"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/driver"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/validation"
)

func TestScenario4382(t *testing.T) {
	spec.ParallelTest(t, &driver.Scenario{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Id: "4382",
		Vars: func(scenario *driver.Scenario) map[string]interface{} {
			return map[string]interface{}{
				"Region": "cn-bj2",
				"Zone":   "cn-bj2-02",
			}
		},
		Owners: []string{"li.wei@ucloud.cn"},
		Title:  "内网-ulb4自动化回归-枚举参数验证-06",
		Steps: []*driver.Step{
			testStep4382CreateULB01,
			testStep4382CreateVServer02,
			testStep4382DescribeVServer03,
			testStep4382CreateVServer04,
			testStep4382DescribeVServer05,
			testStep4382CreateVServer06,
			testStep4382DescribeVServer07,
			testStep4382DeleteULB08,
		},
	})
}

var testStep4382CreateULB01 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewCreateULBRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ULBName":   "测试-inner",
			"Tag":       "Default",
			"Region":    step.Scenario.GetVar("Region"),
			"OuterMode": "No",
			"InnerMode": "Yes",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateULB(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("ULBId_inner", step.Must(utils.GetValue(resp, "ULBId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "创建负载均衡",
	FastFail:      false,
}

var testStep4382CreateVServer02 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewCreateVServerRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VServerName":     "inner_tcp",
			"ULBId":           step.Scenario.GetVar("ULBId_inner"),
			"Region":          step.Scenario.GetVar("Region"),
			"Protocol":        "TCP",
			"PersistenceType": "None",
			"Method":          "ConsistentHashPort",
			"ListenType":      "PacketsTransmit",
			"FrontendPort":    1024,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateVServer(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "创建VServer",
	FastFail:      false,
}

var testStep4382DescribeVServer03 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewDescribeVServerRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ULBId":  step.Scenario.GetVar("ULBId_inner"),
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeVServer(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "获取VServer信息",
	FastFail:      false,
}

var testStep4382CreateVServer04 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewCreateVServerRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VServerName":     "inner_udp",
			"ULBId":           step.Scenario.GetVar("ULBId_inner"),
			"Region":          step.Scenario.GetVar("Region"),
			"Protocol":        "UDP",
			"PersistenceType": "None",
			"MonitorType":     "Ping",
			"Method":          "ConsistentHashPort",
			"ListenType":      "PacketsTransmit",
			"FrontendPort":    1024,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateVServer(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "创建VServer",
	FastFail:      false,
}

var testStep4382DescribeVServer05 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewDescribeVServerRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ULBId":  step.Scenario.GetVar("ULBId_inner"),
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeVServer(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "获取VServer信息",
	FastFail:      false,
}

var testStep4382CreateVServer06 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewCreateVServerRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VServerName":     "inner_udp",
			"ULBId":           step.Scenario.GetVar("ULBId_inner"),
			"Region":          step.Scenario.GetVar("Region"),
			"Protocol":        "UDP",
			"PersistenceType": "None",
			"MonitorType":     "Port",
			"Method":          "ConsistentHashPort",
			"ListenType":      "PacketsTransmit",
			"FrontendPort":    1026,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateVServer(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "创建VServer",
	FastFail:      false,
}

var testStep4382DescribeVServer07 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewDescribeVServerRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ULBId":  step.Scenario.GetVar("ULBId_inner"),
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeVServer(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取VServer信息",
	FastFail:      false,
}

var testStep4382DeleteULB08 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewDeleteULBRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ULBId":  step.Scenario.GetVar("ULBId_inner"),
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteULB(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "删除负载均衡",
	FastFail:      false,
}
