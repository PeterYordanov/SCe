package core

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows"
)

func IsRunningAsAdmin() bool {
	var sid *windows.SID

	// Create a well-known SID for the Administrators group.
	err := windows.AllocateAndInitializeSid(
		&windows.SECURITY_NT_AUTHORITY,
		2,
		windows.SECURITY_BUILTIN_DOMAIN_RID,
		windows.DOMAIN_ALIAS_RID_ADMINS,
		0, 0, 0, 0, 0, 0,
		&sid,
	)
	if err != nil {
		fmt.Println("Failed to initialize SID:", err)
		os.Exit(1)
	}
	defer windows.FreeSid(sid)

	// Open the current process token.
	var token windows.Token

	//TODO: Deprecated
	_, err = windows.OpenCurrentProcessToken()
	if err != nil {
		fmt.Println("Failed to open process token:", err)
		os.Exit(1)
	}
	defer token.Close()

	// Check whether the token includes the Administrators SID.
	isAdmin, err := token.IsMember(sid)
	if err != nil {
		fmt.Println("Failed to check token membership:", err)
		os.Exit(1)
	}
	return isAdmin
}
