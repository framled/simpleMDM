package device

import (
    "github.com/stretchr/testify/assert"
    "net/http"
    "net/http/httptest"
    "testing"
)

const apiKey = "Dummy api key"
const device = `{
    "type":"device",
    "id":123,
    "attributes":{
        "name":"*",
        "last_seen_at":"2019-11-14T18:31:45.000Z",
        "status":"enrolled",
        "device_name":"test",
        "os_version":"12.4",
        "build_version":"16G77",
        "model_name":"iPad 9.7-Inch 6th Gen",
        "model":"MRJN2CI",
        "product_name":"iPad7,5",
        "unique_identifier":"*",
        "serial_number":"*",
        "imei":null,
        "meid":null,
        "device_capacity":26.19,
        "available_device_capacity":25.1312,
        "battery_level":"83%",
        "modem_firmware_version":null,
        "iccid":null,
        "bluetooth_mac":"*",
        "ethernet_macs":[],
        "wifi_mac":"*",
        "current_carrier_network":null,
        "sim_carrier_network":null,
        "subscriber_carrier_network":null,
        "carrier_settings_version":null,
        "phone_number":null,
        "voice_roaming_enabled":false,
        "data_roaming_enabled":false,
        "is_roaming":false,
        "subscriber_mcc":null,
        "subscriber_mnc":null,
        "simmnc":null,
        "current_mcc":null,
        "current_mnc":null,
        "hardware_encryption_caps":3,
        "passcode_present":false,
        "passcode_compliant":true,
        "passcode_compliant_with_profiles":true,
        "is_supervised":true,
        "is_dep_enrollment":false,
        "is_user_approved_enrollment":null,
        "is_device_locator_service_enabled":false,
        "is_do_not_disturb_in_effect":false,
        "personal_hotspot_enabled":null,
        "itunes_store_account_is_active":false,
        "cellular_technology":0,
        "last_cloud_backup_date":null,
        "is_activation_lock_enabled":false,
        "is_cloud_backup_enabled":false,
        "filevault_enabled":false,
        "filevault_recovery_key":null,
        "firmware_password_enabled":false,
        "firmware_password":null,
        "location_latitude":null,
        "location_longitude":null,
        "location_accuracy":null,
        "location_updated_at":null
    },
    "relationships":{
        "device_group":{
            "data":{
                "type":"device_group",
                "id":0
            }
        }
    }
}`
const list = `{"data":[` + device +`], "has_more":false}`

func TestGetById(t *testing.T) {
    handler := func (w http.ResponseWriter, r *http.Request) {
        assert.Equal(t, apiKey, r.Header.Get("authorization"))
        w.Write([]byte(device))
    }
    testServer := httptest.NewServer(http.HandlerFunc(handler))
    defer testServer.Close()

    mdm := NewMDM(apiKey, testServer.URL)
    _, err := mdm.getById("dummyId")
    if err != nil {
        t.FailNow()
    }
}

func TestFindAll(t *testing.T) {
    handler := func (w http.ResponseWriter, r *http.Request) {
        assert.Equal(t, apiKey, r.Header.Get("authorization"))
        w.Write([]byte(list))
    }
    testServer := httptest.NewServer(http.HandlerFunc(handler))
    defer testServer.Close()

    mdm := NewMDM(apiKey, testServer.URL)
    _, err := mdm.findAll(0, 0)
    if err != nil {
        t.FailNow()
    }
}
