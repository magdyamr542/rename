## Rename utility
Simple command to rename a prefix in a list of files
Usage: `ren --replace login --with WhatIsUp ./`
Example with these files:
```
total 32
drwxr-xr-x   5 apple  staff   160B Aug 14 16:16 .
drwxr-xr-x  13 apple  staff   416B Jul 29 20:57 ..
-rw-r--r--   1 apple  staff   6.3K Aug 14 16:10 login.component.html
-rw-r--r--   1 apple  staff   1.2K Aug 14 16:10 login.component.spec.ts
-rw-r--r--   1 apple  staff   1.3K Aug 14 16:10 login.component.ts
```
After running `ren --replace login --with simpleLogin ./`
```
total 32
drwxr-xr-x   5 apple  staff   160B Aug 14 16:17 .
drwxr-xr-x  13 apple  staff   416B Jul 29 20:57 ..
-rw-r--r--   1 apple  staff   6.3K Aug 14 16:10 simpleLogin.component.html
-rw-r--r--   1 apple  staff   1.2K Aug 14 16:10 simpleLogin.component.spec.ts
-rw-r--r--   1 apple  staff   1.3K Aug 14 16:10 simpleLogin.component.ts
```
