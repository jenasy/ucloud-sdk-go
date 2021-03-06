// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/services/uaccount"
	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/driver"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/validation"
)

func TestScenario3611(t *testing.T) {
	spec.ParallelTest(t, &driver.Scenario{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Id: "3611",
		Vars: func(scenario *driver.Scenario) map[string]interface{} {
			return map[string]interface{}{
				"Password":           "Z3VhbmxpeXVhbm1pbWExMjMhQCM=",
				"ImageName":          "ImageTest",
				"TargetImageName":    "ImageCopyTest",
				"TargetRegion":       "cn-bj2",
				"TargetZone":         "cn-bj2-02",
				"myImage":            "#{u_get_image_resource($Region,$Zone)}",
				"BootSize":           20,
				"BootType":           "LOCAL_NORMAL",
				"DiskSize":           20,
				"DiskType":           "LOCAL_NORMAL",
				"BootBackup":         "NONE",
				"DiskBackup":         "NONE",
				"MinimalCpuPlatform": "Intel/Auto",
				"MachineType":        "N",
				"Region":             "cn-bj2",
				"Zone":               "cn-bj2-02",
			}
		},
		Owners: []string{"maggie.an@ucloud.cn"},
		Title:  "主机功能之镜像跨可用区-N-LOCAL_NORMAL-LOCAL_NORMAL-2",
		Steps: []*driver.Step{
			testStep3611DescribeImage01,
			testStep3611CreateUHostInstance02,
			testStep3611DescribeUHostInstance03,
			testStep3611StopUHostInstance04,
			testStep3611DescribeUHostInstance05,
			testStep3611CreateCustomImage06,
			testStep3611DescribeImage07,
			testStep3611CreateUHostInstance08,
			testStep3611DescribeUHostInstance09,
			testStep3611GetProjectList10,
			testStep3611CopyCustomImage11,
			testStep3611DescribeImage12,
			testStep3611CreateUHostInstance13,
			testStep3611DescribeUHostInstance14,
			testStep3611TerminateCustomImage15,
			testStep3611StopUHostInstance16,
			testStep3611StopUHostInstance17,
			testStep3611DescribeUHostInstance18,
			testStep3611TerminateUHostInstance19,
			testStep3611TerminateUHostInstance20,
			testStep3611TerminateCustomImage21,
			testStep3611StopUHostInstance22,
			testStep3611DescribeUHostInstance23,
			testStep3611TerminateUHostInstance24,
		},
	})
}

var testStep3611DescribeImage01 = &driver.Step{
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

		step.Scenario.SetVar("myImage", step.Must(utils.GetValue(resp, "ImageSet.0.ImageId")))
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

var testStep3611CreateUHostInstance02 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewCreateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":               step.Scenario.GetVar("Zone"),
			"Region":             step.Scenario.GetVar("Region"),
			"Password":           "VXFhNzg5VGVzdCFAIyQ7LA==",
			"MinimalCpuPlatform": step.Scenario.GetVar("MinimalCpuPlatform"),
			"Memory":             1024,
			"MachineType":        step.Scenario.GetVar("MachineType"),
			"LoginMode":          "Password",
			"ImageId":            step.Scenario.GetVar("myImage"),
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
			"CPU": 1,
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

var testStep3611DescribeUHostInstance03 = &driver.Step{
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
			validation.Builtins.NewValidator("Action", "DescribeUHostInstanceResponse", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.State", "Running", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.BootDiskState", "Normal", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.0.IsBoot", "True", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.1.IsBoot", "False", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.0.DiskType", step.Scenario.GetVar("BootType"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.1.DiskType", step.Scenario.GetVar("DiskType"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.0.Size", step.Scenario.GetVar("BootSize"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.1.Size", step.Scenario.GetVar("DiskSize"), "str_eq"),
		}
	},
	StartupDelay:  time.Duration(400) * time.Second,
	MaxRetries:    60,
	RetryInterval: 60 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep3611StopUHostInstance04 = &driver.Step{
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
			validation.Builtins.NewValidator("Action", "StopUHostInstanceResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(30) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "关闭主机",
	FastFail:      true,
}

var testStep3611DescribeUHostInstance05 = &driver.Step{
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
			validation.Builtins.NewValidator("Action", "DescribeUHostInstanceResponse", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.State", "Stopped", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.BootDiskState", "Normal", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(30) * time.Second,
	MaxRetries:    30,
	RetryInterval: 30 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep3611CreateCustomImage06 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewCreateCustomImageRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":      step.Scenario.GetVar("Zone"),
			"UHostId":   step.Scenario.GetVar("hostId"),
			"Region":    step.Scenario.GetVar("Region"),
			"ImageName": step.Scenario.GetVar("ImageName"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateCustomImage(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("newImageId", step.Must(utils.GetValue(resp, "ImageId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "创建自制镜像",
	FastFail:      true,
}

var testStep3611DescribeImage07 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewDescribeImageRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"Region":  step.Scenario.GetVar("Region"),
			"ImageId": step.Scenario.GetVar("newImageId"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeImage(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DescribeImageResponse", "str_eq"),
			validation.Builtins.NewValidator("ImageSet.0.State", "Available", "str_eq"),
			validation.Builtins.NewValidator("ImageSet.0.ImageId", step.Scenario.GetVar("newImageId"), "str_eq"),
		}
	},
	StartupDelay:  time.Duration(60) * time.Second,
	MaxRetries:    100,
	RetryInterval: 30 * time.Second,
	Title:         "获取镜像列表",
	FastFail:      true,
}

var testStep3611CreateUHostInstance08 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewCreateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":               step.Scenario.GetVar("Zone"),
			"Region":             step.Scenario.GetVar("Region"),
			"Password":           "VXFhNzg5VGVzdCFAIyQ7LA==",
			"MinimalCpuPlatform": step.Scenario.GetVar("MinimalCpuPlatform"),
			"Memory":             1024,
			"MachineType":        step.Scenario.GetVar("MachineType"),
			"LoginMode":          "Password",
			"ImageId":            step.Scenario.GetVar("myImage"),
			"Disks": []map[string]interface{}{
				{
					"BackupType": "NONE",
					"IsBoot":     "True",
					"Size":       20,
					"Type":       step.Scenario.GetVar("BootType"),
				},
				{
					"BackupType": "NONE",
					"IsBoot":     "False",
					"Size":       20,
					"Type":       step.Scenario.GetVar("DiskType"),
				},
			},
			"CPU": 1,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateUHostInstance(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("hostId2", step.Must(utils.GetValue(resp, "UHostIds.0")))
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

var testStep3611DescribeUHostInstance09 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewDescribeUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"UHostIds": []interface{}{
				step.Scenario.GetVar("hostId2"),
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
			validation.Builtins.NewValidator("Action", "DescribeUHostInstanceResponse", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.State", "Running", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(60) * time.Second,
	MaxRetries:    20,
	RetryInterval: 30 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep3611GetProjectList10 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UAccount")
		if err != nil {
			return nil, err
		}
		client := c.(*uaccount.UAccountClient)

		req := client.NewGetProjectListRequest()
		err = utils.SetRequest(req, map[string]interface{}{})
		if err != nil {
			return nil, err
		}

		resp, err := client.GetProjectList(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("TargetProjectID", step.Must(utils.GetValue(resp, "ProjectSet.1.ProjectId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "GetProjectListResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取帐号下的项目列表",
	FastFail:      true,
}

var testStep3611CopyCustomImage11 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewCopyCustomImageRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":            step.Scenario.GetVar("Zone"),
			"TargetRegion":    step.Scenario.GetVar("TargetRegion"),
			"TargetProjectId": step.Scenario.GetVar("TargetProjectID"),
			"TargetImageName": step.Scenario.GetVar("TargetImageName"),
			"SourceImageId":   step.Scenario.GetVar("newImageId"),
			"Region":          step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CopyCustomImage(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("cpImageId", step.Must(utils.GetValue(resp, "TargetImageId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "复制自制镜像",
	FastFail:      true,
}

var testStep3611DescribeImage12 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewDescribeImageRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region":    step.Scenario.GetVar("TargetRegion"),
			"ProjectId": step.Scenario.GetVar("TargetProjectID"),
			"ImageId":   step.Scenario.GetVar("cpImageId"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeImage(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DescribeImageResponse", "str_eq"),
			validation.Builtins.NewValidator("ImageSet.0.State", "Available", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    120,
	RetryInterval: 60 * time.Second,
	Title:         "获取镜像列表",
	FastFail:      true,
}

var testStep3611CreateUHostInstance13 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewCreateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":               step.Scenario.GetVar("TargetZone"),
			"Region":             step.Scenario.GetVar("TargetRegion"),
			"ProjectId":          step.Scenario.GetVar("TargetProjectID"),
			"Password":           "VXFhNzg5VGVzdCFAIyQ7LA==",
			"MinimalCpuPlatform": step.Scenario.GetVar("MinimalCpuPlatform"),
			"Memory":             1024,
			"MachineType":        step.Scenario.GetVar("MachineType"),
			"LoginMode":          "Password",
			"ImageId":            step.Scenario.GetVar("cpImageId"),
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
			"CPU": 1,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateUHostInstance(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("hostId_new", step.Must(utils.GetValue(resp, "UHostIds.0")))
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

var testStep3611DescribeUHostInstance14 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewDescribeUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone": step.Scenario.GetVar("TargetZone"),
			"UHostIds": []interface{}{
				step.Scenario.GetVar("hostId_new"),
			},
			"Region":    step.Scenario.GetVar("TargetRegion"),
			"ProjectId": step.Scenario.GetVar("TargetProjectID"),
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
			validation.Builtins.NewValidator("Action", "DescribeUHostInstanceResponse", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.State", "Running", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.BootDiskState", "Normal", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    60,
	RetryInterval: 30 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep3611TerminateCustomImage15 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewTerminateCustomImageRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"Region":  step.Scenario.GetVar("Region"),
			"ImageId": step.Scenario.GetVar("newImageId"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.TerminateCustomImage(req)
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
	Title:         "删除自制镜像",
	FastFail:      true,
}

var testStep3611StopUHostInstance16 = &driver.Step{
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
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "关闭主机",
	FastFail:      true,
}

var testStep3611StopUHostInstance17 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewStopUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"UHostId": step.Scenario.GetVar("hostId2"),
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
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "关闭主机",
	FastFail:      true,
}

var testStep3611DescribeUHostInstance18 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewDescribeUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"UHostIds": []interface{}{
				step.Scenario.GetVar("hostId"),
				step.Scenario.GetVar("hostId2"),
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
			validation.Builtins.NewValidator("Action", "DescribeUHostInstanceResponse", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.State", "Stopped", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.1.State", "Stopped", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    30,
	RetryInterval: 10 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep3611TerminateUHostInstance19 = &driver.Step{
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
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除云主机",
	FastFail:      true,
}

var testStep3611TerminateUHostInstance20 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewTerminateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"UHostId": step.Scenario.GetVar("hostId2"),
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
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除云主机",
	FastFail:      true,
}

var testStep3611TerminateCustomImage21 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewTerminateCustomImageRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":      step.Scenario.GetVar("TargetZone"),
			"Region":    step.Scenario.GetVar("TargetRegion"),
			"ProjectId": step.Scenario.GetVar("TargetProjectID"),
			"ImageId":   step.Scenario.GetVar("cpImageId"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.TerminateCustomImage(req)
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
	Title:         "删除自制镜像",
	FastFail:      true,
}

var testStep3611StopUHostInstance22 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewStopUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":      step.Scenario.GetVar("TargetZone"),
			"UHostId":   step.Scenario.GetVar("hostId_new"),
			"Region":    step.Scenario.GetVar("TargetRegion"),
			"ProjectId": step.Scenario.GetVar("TargetProjectID"),
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
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "关闭主机",
	FastFail:      true,
}

var testStep3611DescribeUHostInstance23 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewDescribeUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone": step.Scenario.GetVar("TargetZone"),
			"UHostIds": []interface{}{
				step.Scenario.GetVar("hostId_new"),
			},
			"Region":    step.Scenario.GetVar("TargetRegion"),
			"ProjectId": step.Scenario.GetVar("TargetProjectID"),
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
			validation.Builtins.NewValidator("Action", "DescribeUHostInstanceResponse", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.State", "Stopped", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    30,
	RetryInterval: 10 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep3611TerminateUHostInstance24 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewTerminateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":      step.Scenario.GetVar("TargetZone"),
			"UHostId":   step.Scenario.GetVar("hostId_new"),
			"Region":    step.Scenario.GetVar("TargetRegion"),
			"ProjectId": step.Scenario.GetVar("TargetProjectID"),
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
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除云主机",
	FastFail:      true,
}
