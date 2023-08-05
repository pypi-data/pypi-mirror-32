# ldm 
version:0.0.4
ldm=landmarks
you can use a function to get landmarks and face  with no other libs
like this:
```
import  ldm
from skimage import io

imagepath="closed_eye/10.jfif"
ldmer=ldm.LDM()
img=io.imread(imagepath)
ldl,facel,txt=ldmer.landmarks(img)
print txt
for ld in ldl:
    print 10*'-'
    print 'nose:'
    print ld['nose']
for face in facel:
    print 10*'-'
    print 'face:'
    print face

```  
20180524
anjiang
