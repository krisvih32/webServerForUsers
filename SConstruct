import os
os.system("""
go fmt
go mod tidy
go mod vendor
go mod verify
""")
with open("VERSION.txt", "r") as f:
    print("current version: " + f.read())
version=input("Enter version # (none keeps original version): \n")
if version:
    with open("VERSION.txt", "w") as f:
        f.write(version)