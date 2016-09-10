package models

type QuotaParam struct {
	Quota string
	Volume string
	Percent string
}

func CreateQuota(v *QuotaParam) (string, error) {
	_, err := Gluster("volume", "quota", v.Volume, "enable")
	if (err != nil) {
		return "", err
	}
	_, err = Gluster("volume", "quota", v.Volume, "limit-usage","/",v.Quota+"GB")
	if (err != nil) {
		return "", err
	}
	result, err := Gluster("volume", "quota", v.Volume, "default-soft-limit",v.Percent)
	if (err != nil) {
		return "", err
	}
	return result, nil
}


func ChangeQuota(v *QuotaParam) (string, error) {
	result, err := Gluster("volume", "quota", v.Volume, "limit-usage","/",v.Quota+"GB")
	if (err != nil) {
		return "", err
	}
	return result, nil
}


func RemoveQuota(volume string) (string, error) {
	result, err := Gluster("volume", "quota", volume, "remove","/")
	if (err != nil) {
		return "", err
	}
	return result, nil
}

func QueryQuota(volume string) (string, error) {
	result, err := Gluster("volume", "quota", volume, "list")
	if (err != nil) {
		return "", err
	}
	return result, nil
}