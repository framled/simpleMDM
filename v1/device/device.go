package device

import "time"

type Attribute struct {
    Name                          string    `json:"name,omitempty"`
    LastSeenAt                    time.Time `json:"last_seen_at,omitempty"`
    Status                        string    `json:"status,omitempty"`
    DeviceName                    string    `json:"device_name,omitempty"`
    OsVersion                     string    `json:"os_version,omitempty"`
    BuildVersion                  string    `json:"build_version,omitempty"`
    ModelName                     string    `json:"model_name,omitempty"`
    Model                         string    `json:"model,omitempty"`
    ProductName                   string    `json:"product_name,omitempty"`
    UniqueIdentifier              string    `json:"unique_identifier,omitempty"`
    SerialNumber                  string    `json:"serial_number,omitempty"`
    IMEI                          string    `json:"imei,omitempty"`
    Meid                          string    `json:"meid,omitempty"`
    DeviceCapacity                float32   `json:"device_capacity,omitempty"`
    AvailableDeviceCapacity       float32   `json:"available_device_capacity,omitempty"`
    BatteryLevel                  string    `json:"battery_level,omitempty"`
    ModemFirmwareVersion          string    `json:"modem_firmware_version,omitempty"`
    IccId                         string    `json:"iccid,omitempty"`
    BluetoothMac                  string    `json:"bluetooth_mac,omitempty"`
    WifiMac                       string    `json:"wifi_mac,omitempty"`
    CurrentCarrierNetwork         string    `json:"current_carrier_network,omitempty"`
    SimCarrierNetwork             string    `json:"sim_carrier_network,omitempty"`
    SubscriberCarrierNetwork      string    `json:"subscriber_carrier_network,omitempty"`
    CarrierSettingsVersion        string    `json:"carrier_settings_version,omitempty"`
    PhoneNumber                   string    `json:"phone_number,omitempty"`
    VoiceRoamingEnabled           bool      `json:"voice_roaming_enabled,omitempty"`
    DataRoamingEnabled            bool      `json:"data_roaming_enabled,omitempty"`
    IsRoaming                     bool      `json:"is_roaming,omitempty"`
    SubscriberMcc                 string    `json:"subscriber_mcc,omitempty"`
    SimMNC                        string    `json:"simmnc,omitempty"`
    CurrentMCC                    string    `json:"current_mcc,omitempty"`
    CurrentMNC                    string    `json:"current_mnc,omitempty"`
    HardWareEncryptionCaps        int       `json:"hardware_encryption_caps,omitempty"`
    PassCodePresent               bool      `json:"passcode_present,omitempty"`
    PassCodeCompliant             bool      `json:"passcode_compliant,omitempty"`
    PassCodeCompliantWithProfiles bool      `json:"passcode_compliant_with_profiles,omitempty"`
    SubscriberMNC                 string    `json:"subscriber_mnc,omitempty"`
    SimMCC                        string    `json:"simmcc,omitempty"`
    IsSupervised                  bool      `json:"is_supervised,omitempty"`
    IsDeviceLocatorServiceEnabled bool      `json:"is_device_locator_service_enabled,omitempty"`
    IsDoNotDisturbInEffect        bool      `json:"is_do_not_disturb_in_effect,omitempty"`
    PersonalHotspotEnabled        bool      `json:"personal_hotspot_enabled,omitempty"`
    ItunesStoreAccountIsActive    bool      `json:"itunes_store_account_is_active,omitempty"`
    CellularTechnology            int       `json:"cellular_technology,omitempty"`
    LastCloudBackupDate           string    `json:"last_cloud_backup_date,omitempty"`
    IsActivationLockEnabled       bool      `json:"is_activation_lock_enabled,omitempty"`
    IsCloudBackupEnabled          bool      `json:"is_cloud_backup_enabled,omitempty"`
    LocationLatitude              string    `json:"location_latitude,omitempty"`
    LocationLongitude             string    `json:"location_longitude,omitempty"`
    LocationAccuracy              int       `json:"location_accuracy,omitempty"`
    LocationUpdatedAt             time.Time `json:"location_updated_at,omitempty"`
}

type Relationships struct {
    DeviceGroup Group
}

type Group struct {
    Data struct {
        Id   int    `json:"id"`
        Type string `json:"type"`
    } `json:"data"`
}

type ResponseMany struct {
    Data    *[]Device `json:"data,omitempty"`
    HasMore bool      `json:"has_more,omitempty"`
}

type ResponseOne struct {
    Data *Device `json:"data"`
}

type Device struct {
    Type          string         `json:"type,omitempty"`
    Id            int            `json:"id,omitempty"`
    Attributes    *Attribute     `json:"attributes,omitempty"`
    Relationships *Relationships `json:"relationships,omitempty"`
}
