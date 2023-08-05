# ai_face 
version:0.0.1

analyze the face artribution in the image
like this:
```
# pip install ldm
import ldm
# pip install ai_eye
from ai_face import *
from skimage import io

imagepath="closed_eye/10.jfif"
ldmer=ldm.LDM()
img=io.imread(imagepath)
ldl,facel,helptxt=ldmer.landmarks(img)
print helptxt
print 'face_num:'
print face_num(facel)
print 'face_avre_rate:'
print face_area_rate(img,facel)
print 'face_center_degree:'
print face_center_degree(img,facel)
print 'face_direction:'
print face_direction(facel)
print 'face_feature:'
print face_feature(img,facel)
print 'face_compare:'
feature1=face_feature(img,facel)
feature2=face_feature(img,facel)
print compare2facefeature(feature1,feature2)

```  
20180524
anjiang
