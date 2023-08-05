# ldm 
version:0.0.3
ldm=landmarks
you can usr a function to get landmarks with no other libs
like this:
```
# pip install ldm
import  ldm
from skimage import io

imagepath="closed_eye/10.jfif"
ldmer=ldm.LDM()
img=io.imread(imagepath)
ldl,txt=ldmer.landmarks(img)
print txt
for ld in ldl:
    print ld['nose']

```  
20180524
anjiang
