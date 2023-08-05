"""
Annulus mask
============

Create an antialiased annulus mask using PIL image commands.
"""

from PIL import Image, ImageDraw
import numpy

def annular_mask_antialiased(shape, center, inner_radius, outer_radius,
                             background_value=0.0, mask_value=1.0,
                             oversampling=8):
    # type: (Tuple[int, int], Tuple[float, float], float, float, float, float, int)  -> numpy.ndarray
    """
    Takes the following:

    * *shape* tuple: (x, y) - this is the size of the output image
    * *center* tuple: (x, y)
    * *inner_radius*: float
    * *outer_radius*: float
    * *background_value*: float (the image is initialized to this value)
    * *mask_value*: float (the annulus is drawn with this value)
    * *oversampling*: int (the mask is drawn on a canvas this many times bigger
      than the final size, then resampled down to give smoother edges)
    """
    # Create a 32-bit float image
    intermediate_shape = (shape[0]*int(oversampling), shape[1]*int(oversampling))
    im = Image.new('F', intermediate_shape, color=background_value)

    # Making a handle to the drawing tool
    draw = ImageDraw.Draw(im)

    # Have to scale everything in the problem by the oversampling
    outer_radius_r = outer_radius * oversampling
    inner_radius_r = inner_radius * oversampling
    center_r = (center[0] * oversampling, center[1] * oversampling)


    # Calculate bounding box for outer circle
    x_outer_min = center_r[0] - outer_radius_r
    x_outer_max = center_r[0] + outer_radius_r
    y_outer_min = center_r[1] - outer_radius_r
    y_outer_max = center_r[1] + outer_radius_r
    outer_bbox = [x_outer_min, y_outer_min, x_outer_max, y_outer_max]

    # Calculate bounding box for inner circle
    x_inner_min = center_r[0] - inner_radius_r
    x_inner_max = center_r[0] + inner_radius_r
    y_inner_min = center_r[1] - inner_radius_r
    y_inner_max = center_r[1] + inner_radius_r
    inner_bbox = [x_inner_min, y_inner_min, x_inner_max, y_inner_max]

    # Draw the circles:  outer one first
    draw.ellipse(outer_bbox, fill=mask_value)

    # Now overlay the inner circle
    draw.ellipse(inner_bbox, fill=background_value)

    # Now bring it back to size, with antialiasing
    #im.thumbnail(shape, Image.ANTIALIAS)
    # This produced artifacts - output.max() was > mask_value by 10% or more!

    # Using numpy reshape instead (rebinning) - see Scipy cookbook
    output = numpy.asarray(im)
    output = output.reshape(shape[0], oversampling, shape[1], oversampling).mean(1).mean(2)
    return output

def pie_bounding_box(center, radius, start_angle, end_angle):
    # assumes Euclidean angles (CCW) - make sure to specify pieslice with negative angles (CW)
    points = [
        center,
        [center[0] + numpy.cos(start_angle)*radius, center[1] + numpy.sin(start_angle)*radius],
        [center[0] + numpy.cos(end_angle)*radius, center[1] + numpy.sin(end_angle)*radius]        
    ]
    pangle = numpy.pi/2.0 # angle of the cardinal points
    for cardinal in numpy.arange(pangle*numpy.ceil(start_angle/pangle), end_angle, pangle):
        points.append([center[0] + numpy.cos(cardinal)*radius, center[1] + numpy.sin(cardinal)*radius])
        
    points = numpy.array(points, dtype='float')
    xmin = points[:,0].min()
    xmax = points[:,0].max()
    ymin = points[:,1].min()
    ymax = points[:,1].max()
    return [xmin, ymin, xmax, ymax]
    
def sector_cut_antialiased(shape, center, inner_radius, outer_radius, start_angle=0.0, end_angle=numpy.pi,
                mirror=True, background_value=0.0, mask_value=1.0,
                oversampling=8):
    # type: (Tuple[int, int], Tuple[float, float], float, float, float, float, int)  -> numpy.ndarray
    """
    Takes the following:

    * *shape* tuple: (x, y) - this is the size of the output image
    * *center* tuple: (x, y)
    * *inner_radius*: float
    * *outer_radius*: float
    * *start_angle*: float (radians) start of sector cut range
    * *end_angle*: float (radians) end of sector cut range (with defaults, covers full circle)
    * *mirror*: bool (take cut on both sides of origin)
    * *background_value*: float (the image is initialized to this value)
    * *mask_value*: float (the annulus is drawn with this value)
    * *oversampling*: int (the mask is drawn on a canvas this many times bigger
      than the final size, then resampled down to give smoother edges)
    """
    d_start = -numpy.degrees(start_angle)
    d_end = -numpy.degrees(end_angle)
    
    # Create a 32-bit float image
    intermediate_shape = (shape[0]*int(oversampling), shape[1]*int(oversampling))
    im = Image.new('F', intermediate_shape, color=background_value)

    # Making a handle to the drawing tool
    draw = ImageDraw.Draw(im)

    # Have to scale everything in the problem by the oversampling
    outer_radius_r = outer_radius * oversampling
    inner_radius_r = inner_radius * oversampling
    center_r = (center[0] * oversampling, center[1] * oversampling)

    # Calculate bounding box for outer circle
    x_outer_min = center_r[0] - outer_radius_r
    x_outer_max = center_r[0] + outer_radius_r
    y_outer_min = center_r[1] - outer_radius_r
    y_outer_max = center_r[1] + outer_radius_r
    outer_bbox = [x_outer_min, y_outer_min, x_outer_max, y_outer_max]
    
    # draw pie slices
    draw.pieslice(outer_bbox, d_end, d_start, fill=mask_value)
    if mirror:
        draw.pieslice(outer_bbox, d_end-180.0, d_start-180.0,  fill=mask_value)
    
    # Calculate bounding box for inner circle
    x_inner_min = center_r[0] - inner_radius_r
    x_inner_max = center_r[0] + inner_radius_r
    y_inner_min = center_r[1] - inner_radius_r
    y_inner_max = center_r[1] + inner_radius_r
    inner_bbox = [x_inner_min, y_inner_min, x_inner_max, y_inner_max]
    
    # Now overlay the inner circle
    draw.ellipse(inner_bbox, fill=background_value)
    
    # Now bring it back to size, with antialiasing
    #im.thumbnail(shape, Image.ANTIALIAS)
    # This produced artifacts - output.max() was > mask_value by 10% or more!

    # Using numpy reshape instead (rebinning) - see Scipy cookbook
    output = numpy.asarray(im)
    output = output.reshape(shape[0], oversampling, shape[1], oversampling).mean(1).mean(2)
    return output

def test_sector_cut_antialiased(shape, center, inner_radius, outer_radius, start_angle=0.0, end_angle=numpy.pi,
                mirror=True, background_value=0.0, mask_value=1.0,
                oversampling=8):
    # type: (Tuple[int, int], Tuple[float, float], float, float, float, float, int)  -> numpy.ndarray
    """
    Takes the following:

    * *shape* tuple: (x, y) - this is the size of the output image
    * *center* tuple: (x, y)
    * *inner_radius*: float
    * *outer_radius*: float
    * *start_angle*: float (radians) start of sector cut range
    * *end_angle*: float (radians) end of sector cut range (with defaults, covers full circle)
    * *background_value*: float (the image is initialized to this value)
    * *mask_value*: float (the annulus is drawn with this value)
    * *oversampling*: int (the mask is drawn on a canvas this many times bigger
      than the final size, then resampled down to give smoother edges)
    """
    d_start = -numpy.degrees(start_angle)
    d_end = -numpy.degrees(end_angle)
    
    # Create a 32-bit float image
    intermediate_shape = (shape[0]*int(oversampling), shape[1]*int(oversampling))
    im = Image.new('F', intermediate_shape, color=background_value)

    # Making a handle to the drawing tool
    draw = ImageDraw.Draw(im)

    # Have to scale everything in the problem by the oversampling
    outer_radius_r = outer_radius * oversampling
    inner_radius_r = inner_radius * oversampling
    center_r = (center[0] * oversampling, center[1] * oversampling)

    # Calculate bounding box for outer circle
    x_outer_min = center_r[0] - outer_radius_r
    x_outer_max = center_r[0] + outer_radius_r
    y_outer_min = center_r[1] - outer_radius_r
    y_outer_max = center_r[1] + outer_radius_r
    outer_bbox = [x_outer_min, y_outer_min, x_outer_max, y_outer_max]
    draw.pieslice(outer_bbox, d_end, d_start, fill=mask_value)
    if mirror:
        draw.pieslice(outer_bbox, d_end-180.0, d_start-180.0,  fill=mask_value)
    
    # Calculate bounding box for inner circle
    x_inner_min = center_r[0] - inner_radius_r
    x_inner_max = center_r[0] + inner_radius_r
    y_inner_min = center_r[1] - inner_radius_r
    y_inner_max = center_r[1] + inner_radius_r
    inner_bbox = [x_inner_min, y_inner_min, x_inner_max, y_inner_max]
    
    # Now overlay the inner circle
    draw.ellipse(inner_bbox, fill=background_value)
    im.show()
    
def test():
    test_sector_cut_antialiased(
        (128,128), 
        (64, 60), 
        10, 
        20, 
        start_angle=numpy.pi/4, 
        end_angle=numpy.pi/2,
        mirror=False, 
        background_value=0.0, 
        mask_value=100.0,
        oversampling=8
    )
