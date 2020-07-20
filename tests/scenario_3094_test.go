// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/driver"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/validation"
)

func TestScenario3094(t *testing.T) {
	spec.ParallelTest(t, &driver.Scenario{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Id: "3094",
		Vars: func(scenario *driver.Scenario) map[string]interface{} {
			return map[string]interface{}{
				"Password":           "Z3VhbmxpeXVhbm1pbWExMjMhQCM=",
				"ChargeType":         "Month",
				"CreateCPU":          1,
				"CreateMem":          1024,
				"NewPassword":        "Z3VhbmxpeXVhbm1pbWExMjMhQCM=",
				"Name":               "uhost-basic-api-N-Normal-LocalDisk-2",
				"ImageID":            "#{u_get_image_resource($Region,$Zone)}",
				"BootSize":           20,
				"BootType":           "LOCAL_NORMAL",
				"DiskSize":           20,
				"DiskType":           "LOCAL_NORMAL",
				"BootBackup":         "NONE",
				"DiskBackup":         "NONE",
				"UHostType":          "N2",
				"MinimalCpuPlatform": "Intel/Haswell",
				"MachineType":        "N",
				"Region":             "cn-bj2",
				"Zone":               "cn-bj2-02",
			}
		},
		Owners: []string{"maggie.an@ucloud.cn"},
		Title:  "Intel/Haswell-主机功能-Auto-N-LOCAL_NORMAL-LOCAL_NORMAL-2",
		Steps: []*driver.Step{
			testStep3094DescribeImage01,
			testStep3094CreateUHostInstance02,
			testStep3094DescribeUHostInstance03,
			testStep3094StopUHostInstance04,
			testStep3094DescribeUHostInstance05,
			testStep3094ResetUHostInstancePassword06,
			testStep3094DescribeUHostInstance07,
			testStep3094ReinstallUHostInstance08,
			testStep3094DescribeUHostInstance09,
			testStep3094PoweroffUHostInstance10,
			testStep3094DescribeUHostInstance11,
			testStep3094TerminateUHostInstance12,
		},
	})
}

var testStep3094DescribeImage01 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewDescribeImageRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":      step.Scenario.GetVar("Zone"),
			"Region":    step.Scenario.GetVar("Region"),
			"OsType":    "Linux",
			"ImageType": "Base",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeImage(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("ImageID", step.Must(utils.GetValue(resp, "ImageSet.0.ImageId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DescribeImageResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取镜像列表",
	FastFail:      true,
}

var testStep3094CreateUHostInstance02 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewCreateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":               step.Scenario.GetVar("Zone"),
			"UHostType":          step.Scenario.GetVar("UHostType"),
			"TimemachineFeature": "No",
			"Tag":                "Default",
			"Region":             step.Scenario.GetVar("Region"),
			"Quantity":           1,
			"Password":           "VXFhNzg5VGVzdCFAIyQ7LA==",
			"NetCapability":      "Normal",
			"Name":               step.Scenario.GetVar("Name"),
			"MinimalCpuPlatform": step.Scenario.GetVar("MinimalCpuPlatform"),
			"Memory":             step.Scenario.GetVar("CreateMem"),
			"MachineType":        step.Scenario.GetVar("MachineType"),
			"LoginMode":          "Password",
			"ImageId":            step.Scenario.GetVar("ImageID"),
			"HotplugFeature":     "false",
			"Disks": []map[string]interface{}{
				{
					"BackupType": step.Scenario.GetVar("BootBackup"),
					"IsBoot":     "True",
					"Size":       step.Scenario.GetVar("BootSize"),
					"Type":       step.Scenario.GetVar("BootType"),
				},
				{
					"BackupType": step.Scenario.GetVar("DiskBackup"),
					"IsBoot":     "False",
					"Size":       step.Scenario.GetVar("DiskSize"),
					"Type":       step.Scenario.GetVar("DiskType"),
				},
			},
			"ChargeType": step.Scenario.GetVar("ChargeType"),
			"CPU":        step.Scenario.GetVar("CreateCPU"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateUHostInstance(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("hostId", step.Must(utils.GetValue(resp, "UHostIds.0")))
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
	Title:         "创建云主机",
	FastFail:      true,
}

var testStep3094DescribeUHostInstance03 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewDescribeUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone": step.Scenario.GetVar("Zone"),
			"UHostIds": []interface{}{
				step.Scenario.GetVar("hostId"),
			},
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeUHostInstance(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.CPU", step.Scenario.GetVar("CreateCPU"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.Memory", step.Scenario.GetVar("CreateMem"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.UHostId", step.Scenario.GetVar("hostId"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.Name", step.Scenario.GetVar("Name"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.BasicImageId", step.Scenario.GetVar("ImageID"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.State", "Running", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.0.IsBoot", "True", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.1.IsBoot", "False", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.0.DiskType", step.Scenario.GetVar("BootType"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.1.DiskType", step.Scenario.GetVar("DiskType"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.0.Size", step.Scenario.GetVar("BootSize"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.1.Size", step.Scenario.GetVar("DiskSize"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.MachineType", step.Scenario.GetVar("MachineType"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.BootDiskState", "Normal", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    200,
	RetryInterval: 30 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep3094StopUHostInstance04 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewStopUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("hostId"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.StopUHostInstance(req)
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
	MaxRetries:    5,
	RetryInterval: 30 * time.Second,
	Title:         "关闭主机",
	FastFail:      true,
}

var testStep3094DescribeUHostInstance05 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewDescribeUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone": step.Scenario.GetVar("Zone"),
			"UHostIds": []interface{}{
				step.Scenario.GetVar("hostId"),
			},
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeUHostInstance(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.State", "Stopped", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    30,
	RetryInterval: 10 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep3094ResetUHostInstancePassword06 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewResetUHostInstancePasswordRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":     step.Scenario.GetVar("Zone"),
			"UHostId":  step.Scenario.GetVar("hostId"),
			"Region":   step.Scenario.GetVar("Region"),
			"Password": step.Scenario.GetVar("NewPassword"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.ResetUHostInstancePassword(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "ResetUHostInstancePasswordResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "重置主机密码",
	FastFail:      true,
}

var testStep3094DescribeUHostInstance07 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewDescribeUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone": step.Scenario.GetVar("Zone"),
			"UHostIds": []interface{}{
				step.Scenario.GetVar("hostId"),
			},
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeUHostInstance(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.State", "Stopped", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(20) * time.Second,
	MaxRetries:    30,
	RetryInterval: 20 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep3094ReinstallUHostInstance08 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewReinstallUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":     step.Scenario.GetVar("Zone"),
			"UHostId":  step.Scenario.GetVar("hostId"),
			"Region":   step.Scenario.GetVar("Region"),
			"Password": step.Scenario.GetVar("NewPassword"),
			"ImageId":  step.Scenario.GetVar("ImageID"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.ReinstallUHostInstance(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "ReinstallUHostInstanceResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "重装系统",
	FastFail:      true,
}

var testStep3094DescribeUHostInstance09 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewDescribeUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone": step.Scenario.GetVar("Zone"),
			"UHostIds": []interface{}{
				step.Scenario.GetVar("hostId"),
			},
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeUHostInstance(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.CPU", step.Scenario.GetVar("CreateCPU"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.Memory", step.Scenario.GetVar("CreateMem"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.UHostId", step.Scenario.GetVar("hostId"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.Name", step.Scenario.GetVar("Name"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.BasicImageId", step.Scenario.GetVar("ImageID"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.State", "Running", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.0.IsBoot", "True", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.1.IsBoot", "False", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.0.DiskType", step.Scenario.GetVar("BootType"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.1.DiskType", step.Scenario.GetVar("DiskType"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.0.Size", step.Scenario.GetVar("BootSize"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.1.Size", step.Scenario.GetVar("DiskSize"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.MachineType", step.Scenario.GetVar("MachineType"), "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    30,
	RetryInterval: 30 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep3094PoweroffUHostInstance10 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewPoweroffUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("hostId"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.PoweroffUHostInstance(req)
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
	MaxRetries:    5,
	RetryInterval: 30 * time.Second,
	Title:         "模拟主机掉电",
	FastFail:      true,
}

var testStep3094DescribeUHostInstance11 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewDescribeUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone": step.Scenario.GetVar("Zone"),
			"UHostIds": []interface{}{
				step.Scenario.GetVar("hostId"),
			},
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeUHostInstance(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.State", "Stopped", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    30,
	RetryInterval: 10 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep3094TerminateUHostInstance12 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewTerminateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("hostId"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.TerminateUHostInstance(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除云主机",
	FastFail:      true,
}