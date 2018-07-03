package controllers

import "squeezecnn/models"

type(
	// POST : /device/register
	DeviceResource struct {
		Data models.Device `json:"data"`
	}

	//POST /device/authenticate
	AuthResource struct {
		Data AuthModel `json:"data"`
	}

	AuthDeviceResource struct {
		Data AuthDeviceModel `json:"data"`
	}

	AuthModel struct {
		DeviceName string `json:"deviceName"`
		DeviceMac string `json:"deviceMac"`
	}

	AuthDeviceModel struct {
		Device models.Device `json:"device"`
		Token string `json:"token"`
	}

)

type (
	EmployeeResource struct {
		Data models.Employee `json:"data"`
	}


	AuthEmployeeResource struct {
		Data models.Employee `json"data"`
	}

)

type NetworkResource struct {
	Data models.Network `json:"data"`
}