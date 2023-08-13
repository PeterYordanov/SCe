package capturer

/*
	Windows:
		key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall`, registry.ENUMERATE_SUB_KEYS|registry.QUERY_VALUE)
		if err != nil {
			log.Fatal(err)
		}
		defer key.Close()

		subkeys, err := key.ReadSubKeyNames(-1)
		if err != nil {
			log.Fatal(err)
		}
		for _, sk := range subkeys {
			subkey, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\`+sk, registry.QUERY_VALUE)
			if err != nil {
				continue
			}
			name, _, err := subkey.GetStringValue("DisplayName")
			if err == nil {
				fmt.Println(name)
			}
			subkey.Close()
		}
		"golang.org/x/sys/windows/registry"


	dpkg:
		cmd := exec.Command("dpkg", "--list")
		output, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(output))
*/
