<<<<<<< HEAD
import sys
=======
>>>>>>> b6a5a1e (changed SConstruct)
with open("VERSION.txt", "r") as f:
    print("current version: " + f.read())
import os
SCONS_WRITE = 1
match os.environ.get("SCONS_WRITE"):
    case "0" | "":
        SCONS_WRITE = 0
<<<<<<< HEAD
if SCONS_WRITE: sys.exit()
version=input("Enter version # (none keeps original version): \n")
if not version: sys.exit() 
with open("VERSION.txt", "w") as f:
    f.write(version)
=======
version=input("Enter version # (none keeps original version): \n")
if bool(version) & SCONS_WRITE:
    with open("VERSION.txt", "w") as f:
        f.write(version)
>>>>>>> b6a5a1e (changed SConstruct)
