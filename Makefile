main:
	go fmt
	go mod tidy
	go mod vendor
	go mod verify
	read -p "Enter version # (none keeps original version): " version
	if [ "$version" != "" ]; then
		echo "$version" > ../VERSION.TXT
	fi
	exit 0

