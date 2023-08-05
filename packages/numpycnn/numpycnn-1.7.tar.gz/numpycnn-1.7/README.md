Convolutional neural network implementation using NumPy.
Just three layers are created which are convolution (conv for short), ReLU, and max pooling.
The major steps involved are as follows:

1. Reading the input image.

2. Preparing filters.

3. Conv layer: Convolving each filter with the input image.

4. ReLU layer: Applying ReLU activation function on the feature maps (output of conv layer).

5. Max Pooling layer: Applying the pooling operation on the output of ReLU layer.

6. Stacking conv, ReLU, and max pooling layers



The project is tested using Python 3.5.2 installed inside Anaconda 4.2.0 (64-bit)
NumPy version used is 1.14.0



The file named **example.py** is an example of using the project.

The code starts by reading an input image. That image can be either single or multi-dimensional image.
In this examplel, an input gray is used and this is why it is required to ensure the image is already gray.



The filters of the first conv layer are prepared according to the input image dimensions. The filter is created by specifying the following:

1) Number of filters.

2) Size of first dimension.

3) Size of second dimension.
4) Size of third dimension and so on.


Because the previous image is just gray, then the filter will have just width and height and no depth. That is why it is created by specifying just three numbers (number of filters, width, and height). 

The code can still work with RGb images. The only difference is using filters of similar shape to the image. If the image is RGB and not converted to gray, then the filter will be created by specifying 4 numbers (number of filters, width, height, and number of channels).
------------------------------



For more details, refer to the article describing this project which is titled "Building Convolutional Neural Network using NumPy from Scratch". It is available in these links:

LinkedIn: https://www.linkedin.com/pulse/building-convolutional-neural-network-using-numpy-from-ahmed-gad/

KDnuggets: https://www.kdnuggets.com/2018/04/building-convolutional-neural-network-numpy-scratch.html
It is also translated into Chinese: http://m.aliyun.com/yunqi/articles/585741
------------------------------





Contact the author:

KDnuggets: https://www.kdnuggets.com/author/ahmed-gad
LinkedIn: https://www.linkedin.com/in/ahmedfgad

Facebook: https://www.facebook.com/ahmed.f.gadd

ahmed.f.gad@gmail.com
ahmed.fawzy@ci.menofia.edu.eg