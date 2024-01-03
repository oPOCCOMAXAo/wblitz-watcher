package du

func HasPermissions(userPermissions, checkPermissions int64) bool {
	return userPermissions&checkPermissions == checkPermissions
}
