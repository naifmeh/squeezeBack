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
		DeviceName string `json:"deviceName"`
		Token string `json:"token"`
	}



)