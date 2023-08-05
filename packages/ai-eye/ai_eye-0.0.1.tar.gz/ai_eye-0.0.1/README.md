# ldm 
version:0.0.2
ldm=landmarks
you can usr a function to get landmarks with no other libs
like this:
```
# pip install gldm
from gldm import landmarks
from skimage import io

imagepath="closed_eye/10.jfif"

img=io.imread(imagepath)
ldl,txt=landmarks(img)
print txt
for ld in ldl:
    print ld['nose'] 
```  
20180524
anjiang
