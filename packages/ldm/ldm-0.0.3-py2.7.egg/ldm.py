import os
import dlib
from skimage import io
import numpy as np

class LDM:
    def __init__(self):
        self.predictor_path=self.get_model() 
        self.detector=dlib.get_frontal_face_detector()
        self.predictor=dlib.shape_predictor(self.predictor_path)
   
    def get_part_landmarks(self,shape,start_index,end_index):
        '''
        {   
            dxRange jaw;       // [0 , 16]
            IdxRange rightBrow; // [17, 21]
            IdxRange leftBrow;  // [22, 26]
            IdxRange nose;      // [27, 35]
            IdxRange rightEye;  // [36, 41]
        IdxRange leftEye;   // [42, 47]
        IdxRange mouth;     // [48, 59]
        IdxRange mouth2;    // [60, 67]
        }
        make the shape to be a dict that can easy get  part like jaw,brow,nose,eye,mouth
        '''
        jaw=[]
        for i in range(start_index,end_index):
            #print i+1,shape.part(i)
            jaw.append(np.array((shape.part(i).x,shape.part(i).y)))
        return jaw
    
    def landmark_list(self,img):
        # get all 68 landmarks from the img
        # return the list ldl of the landmark dict ld
        dets = self.detector(img, 1)
        #print("Number of faces detected: {}".format(len(dets)))
        ldl=[]
        for k,d in enumerate(dets):
            shape=self.predictor(img,d)
            #print k,d
            ld={'help':'jaw,right_brow,left_brow,nose,right_eye,left_eye,mouth,mouth2'}
            ld['jaw']=self.get_part_landmarks(shape,0,17)
            ld['right_brow']=self.get_part_landmarks(shape,17,22)
            ld['left_brow']=self.get_part_landmarks(shape,22,27)
            ld['nose']=self.get_part_landmarks(shape,27,36)
            ld['right_eye']=self.get_part_landmarks(shape,36,42)
            ld['left_eye']=self.get_part_landmarks(shape,42,48)
            ld['mouth']=self.get_part_landmarks(shape,48,59)
            ld['mouth2']=self.get_part_landmarks(shape,60,67)
            ldl.append(ld)
        #for ld in ldl:
        #    print ld['right_eye']
            #print np.array(ld)
        return ldl
    
    def get_model(self):
        predictor_path="landmarks_68.dat"
        url='http://dlib.net/files/shape_predictor_68_face_landmarks.dat.bz2'
        if not os.path.exists(predictor_path):
            print os.path.exists("%s.bz2"%(predictor_path))
            if not os.path.exists("%s.bz2"%(predictor_path)):
                os.system('wget -O %s.bz2 %s'%(predictor_path,url))   
            os.system('bunzip2 %s.bz2'%(predictor_path))   
        return predictor_path
    
    def landmarks(self,img):
       
        #predictor_path=get_model() 
        #detector=dlib.get_frontal_face_detector()
        #predictor=dlib.shape_predictor(predictor_path)
        ldl=self.landmark_list(img)
        helptxt='dict[0]_item:jaw,right_brow,left_brow,nose,right_eye,left_eye,mouth,mouth2'
        return ldl,helptxt+',model_paht='+self.predictor_path
