with open("VERSION.txt", "r") as f:
    print("current version: " + f.read())
import os
SCONS_WRITE = 1
match os.environ.get("SCONS_WRITE"):
    case "0" | "":
        SCONS_WRITE = 0
version=input("Enter version # (none keeps original version): \n")
if bool(version) & SCONS_WRITE:
    with open("VERSION.txt", "w") as f:
        f.write(version)