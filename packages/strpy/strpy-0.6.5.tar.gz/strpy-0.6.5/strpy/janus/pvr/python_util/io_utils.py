""" Utility functions related to input and output
"""
import numpy as np
from PIL import Image
import os
import geometry_utils
# tifffile module can be imported from either tifffile or skimage.external.tifffile
has_tifffile = True
try:
    import tifffile
except ImportError:
    try:
        import skimage.external.tifffile as tifffile
    except ImportError:
        has_tifffile = False


def read_token(file_obj, tok=None, ignore_char=None):
    """ return a generator that seperates file based on whitespace, or optionally, tok """
    for line in file_obj:
        for token in line.split(tok):
            if len(token) == 0 or (ignore_char is not None and token[0] == ignore_char):
                # ignore rest of line
                break
            yield token


def read_list(filename):
    """ read a list of strings from file, one per line """
    fd = open(filename,'r')
    lines = []
    for line in fd:
        lines.append(line.strip())
    return lines


def write_list(thelist, filename):
    """ write each element to a seperate line in the file """
    try:
        fd = open(filename,'w')
    except IOError:
        print('Error opening file ' + filename)
        return []
    for list_el in thelist:
        fd.write(str(list_el) + '\n')
    return


def read_vector_float(filename):
    """ read each float into a single vector """
    elements = []
    lines = read_list(filename)
    for line in lines:
        elements_str = line.split()
        for s in elements_str:
            elements.append(float(s))
    vec = np.array(elements)
    return vec


def read_vectors_float(filename):
    """ read each line as a seperate vector of floats """
    lines = read_list(filename)
    vectors = []
    for line in lines:
        elements_str = line.split()
        elements = []
        for s in elements_str:
            elements.append(float(s))
        if (len(elements) > 0):
            vec = np.array(elements)
            vectors.append(vec)
    return vectors


def write_vectors_float(vector_list, filename):
    """ write each vector of floats to a seperate line """
    str_list = []
    for v in vector_list:
        v_str = ''
        for x in v:
            v_str = v_str + str(x) + ' '
        str_list.append(v_str)
    write_list(str_list, filename)
    return str_list


def imread(filename):
    """ read the image to a numpy array """
    _, ext = os.path.splitext(filename)
    if has_tifffile and (ext == '.tiff' or ext == '.tif'):
        # if image is tiff, use tifffile module
        img_np = tifffile.imread(filename)
    else:
        img = Image.open(filename)
        # workaround for 16 bit images
        if img.mode == 'I;16':
            img_np = np.array(img.getdata(), dtype=np.uint16).reshape(img.size[::-1])
        else:
            img_np = np.array(img)
    return img_np


def imwrite(img, filename):
    """ write the numpy array as an image """
    _, ext = os.path.splitext(filename)
    is_multiplane = len(img.shape) > 2
    if has_tifffile and (ext == '.tiff' or ext == '.tif') and is_multiplane:
        # if image is tiff, use tifffile module
        tifffile.imsave(filename, img)
    else:
        pilImg = Image.fromarray(img)
        if pilImg.mode == 'L':
            pilImg.convert('I')  # convert to 32 bit signed mode
        pilImg.save(filename)
    return


def imwrite_byte(img, vmin, vmax, filename):
    """ write the 2-d numpy array as an image, scale to byte range first """
    img_byte = np.uint8(np.zeros_like(img))
    img_norm = (img - vmin)/(vmax-vmin)
    img_norm = img_norm.clip(0.0, 1.0)
    img_byte[:] = img_norm * 255
    imwrite(img_byte, filename)


# remove directory and extension from filename
def filename_base(filename):
    """ remove the directory and extension from the filename """
    (_, filename_wext) = os.path.split(filename)
    (base, _) = os.path.splitext(filename_wext)
    return base


def read_vector(vec_string):
    """ read the individual floats from a string """
    elements_str = vec_string.split()
    elements = []
    for s in elements_str:
        elements.append(float(s))
    vec = np.array(elements)
    return vec


def read_matrix_lines(row_strings):
    """ read the individual matrix elements from a list of strings (one per row) """
    rows = []
    for line in row_strings:
        elements_str = line.split()
        elements = []
        for s in elements_str:
            elements.append(float(s))
        if (len(elements) > 0):
            vec = np.array(elements)
            rows.append(vec)
    M = np.array(rows)
    return M


def read_matrix(filename):
    """ read a matrix from a file. assumes each row is on a new line """
    return np.loadtxt(filename)


def write_matrix(M, filename):
    """ write out 1D or 2D numpy array M as ascii text file, one row per line """
    np.savetxt(filename, M)


def read_camera_KRT(filename):
    """ read a KRT camera from text file """
    lines = read_list(filename)
    # remove any empty lines
    lines = [line for line in lines if line]
    K = read_matrix_lines(lines[0:3])
    R = read_matrix_lines(lines[3:6])
    T = read_vector(lines[6])
    return K, R, T


def write_camera_KRT(K,R,T, filename):
    """ write a KRT camera from text file """
    if K.shape != (3,3):
        raise Exception('K matrix should be 3x3')
    if R.shape != (3,3):
        raise Exception('R matrix should be 3x3')
    if T.shape != (3,):
        raise Exception('T vector should have length 3')
    with open(filename,'w') as fd:
        for row in K:
            fd.write('%f %f %f\n' % (row[0],row[1],row[2]))
        for row in R:
            fd.write('%f %f %f\n' % (row[0],row[1],row[2]))
        fd.write('%f %f %f\n' % (T[0],T[1],T[2]))


def read_bundler_file(filename):
    """ read an output file from the 'bundler' program.  Return the cameras and points """
    lines = read_list(filename)
    # first line is comment
    # second line has number of cameras, number of pts
    cam_pts_str = lines[1].split()
    num_cams = int(cam_pts_str[0])
    num_pts = int(cam_pts_str[1])
    intrinsics = []
    Rs = []
    Ts = []
    lines_per_cam = 5
    for i in range(num_cams):
        start_line = lines_per_cam * i + 2
        intrinsic = read_vector(lines[start_line])
        R = read_matrix_lines(lines[start_line + 1: start_line + 4])
        T = read_vector(lines[start_line + 4])
        intrinsics.append(intrinsic)
        Rs.append(R)
        Ts.append(T)
    pts = []
    pts_start_line = lines_per_cam * num_cams + 2
    lines_per_pt = 3
    for i in range(num_pts):
        start_line = pts_start_line + lines_per_pt*i
        p = read_vector(lines[start_line])
        pts.append(p)
    return intrinsics, Rs, Ts, pts


def read_vsfm_nvm_file(filename):
    """ read an output file from the "VisualSFM" program in .nvm format """
    try:
        fd = open(filename,'r')
    except IOError:
        print('Error opening file ' + filename)
        return None
    # first line should contain version string and optionally, fixed calibration info
    magic_string = 'NVM_V3'
    magic_string_R9T = 'NVM_V3_R9T'
    first_line_toks = fd.readline().split()
    if len(first_line_toks) == 0:
        print('Error: Expecting first token in file to be ' + magic_string)
        return None
    format_R9T = False
    if first_line_toks[0] == magic_string_R9T:
        format_R9T = True
    elif first_line_toks[0] != magic_string:
        print('Error: Expecting first token in file to be ' + magic_string)
        return None
    if len(first_line_toks) > 1 and first_line_toks[1] == 'FixedK':
        # file has fixed calibration info
        # skip for now
        print('WARNING: skipping read of fixed calibration info')

    # from here on out, read file on token at a time
    tokgen = read_token(fd, ignore_char='#')

    num_cameras = int(next(tokgen))
    print('%d cameras' % num_cameras)
    img_fnames = []
    fs = []
    Rs = []
    Ts = []
    for c in range(num_cameras):
        fname = next(tokgen)
        f = float(next(tokgen))
        if format_R9T:
            rvals = np.zeros(9)
            for ri in range(9):
                rvals[ri] = float(next(tokgen))
            R = rvals.reshape((3,3))
            T = np.zeros(3)
            for ti in range(3):
                T[ti] = float(next(tokgen))
            cam_center = np.dot(-R.transpose(),T)
        else:
            q = np.zeros(4)
            for qi in range(4):
                q[qi] = float(next(tokgen))
            R = geometry_utils.quaternion_to_matrix(q)
            cam_center = np.zeros(3)
            for ci in range(3):
                cam_center[ci] = float(next(tokgen))
        dist_coef = float(next(tokgen))
        if (dist_coef != 0.0):
            print('WARNING: ignoring nonzero distortion coefficent for camera %d' % c)

        T = np.dot(-R, cam_center)

        img_fnames.append(fname)
        fs.append(f)
        Rs.append(R)
        Ts.append(T)
        # read '0' as end of camera
        if next(tokgen) != '0':
            print('Error: expecting \'0\' delimiter and end of camera %d section' % c)
            #return None
    num_points = int(next(tokgen))
    print('%d points' % num_points)
    pts = []
    colors = []
    for _p in range(num_points):
        pt = np.zeros(3)
        for pti in range(3):
            pt[pti] = float(next(tokgen))
        pts.append(pt)
        rgb = np.zeros(3, 'uint8')
        for rgbi in range(3):
            rgb[rgbi] = np.uint8(next(tokgen))
        colors.append(rgb)
        num_measurements = int(next(tokgen))
        for _m in range(num_measurements):
            # ignore measurement info for now
            # image index, feature index, x, y
            for _mm in range(4):
                next(tokgen)

    return img_fnames, fs, Rs, Ts, pts, colors

# from http://stackoverflow.com/questions/260273/most-efficient-way-to-search-the-last-x-lines-of-a-file-in-python.
# surprisingly, this was not the accepted answer.
def reversed_lines(file):
    """ generate the lines of file in reverse order """
    part = ''
    for block in reversed_blocks(file):
        for c in reversed(block):
            if c == '\n' and part:
                yield part[::-1]
                part = ''
            part += c
    if part: yield part[::-1]

def reversed_blocks(file, blocksize=4096):
    """ generate blocks of file's contents in reverse order """
    file.seek(0, os.SEEK_END)
    here = file.tell()
    while 0 < here:
        delta = min(blocksize, here)
        here -= delta
        file.seek(here, os.SEEK_SET)
        yield file.read(delta)


def write_ply(filename, verts, faces=None, vert_colors=None, vert_uv=None, vert_normals=None):
    """ Save a mesh in the Stanford ply format
        verts should be a Nx3 array of points
        faces should be a Fx3 array of integer vertex indices
    """
    num_verts = verts.shape[0]
    if verts.shape[1] != 3:
        raise Exception('verts matrix should be Nx3')
    with open(filename,'w') as fd:
        fd.write('ply\n')
        fd.write('format ascii 1.0\n')
        fd.write('element vertex ' + str(num_verts) + '\n')
        fd.write('property float x\n')
        fd.write('property float y\n')
        fd.write('property float z\n')
        if vert_colors is not None:
            if vert_colors.shape != (num_verts,3):
                raise Exception('vert_colors matrix should be Nx3')
            if vert_colors.dtype != np.uint8:
                raise Exception('vert_colors dtype should be uint8')
            fd.write('property uint8 red\n')
            fd.write('property uint8 green\n')
            fd.write('property uint8 blue\n')
        if vert_normals is not None:
            if vert_normals.shape != (num_verts,3):
                raise Exception('vert_normals matrix should be Nx3')
            fd.write('property float nx\n')
            fd.write('property float ny\n')
            fd.write('property float nz\n')
        if vert_uv is not None:
            if vert_uv.shape != (num_verts,2):
                raise Exception('vert_uv matrix should be Nx2')
            fd.write('property float s\n')
            fd.write('property float t\n')
        if faces is not None:
            if faces.shape[1] != 3:
                raise Exception('faces matrix should be Nx3')
            num_faces = faces.shape[0]
            fd.write('element face ' + str(num_faces) + '\n')
            fd.write('property list uchar int vertex_indices\n')
        fd.write('end_header\n')
        for i in range(num_verts):
            fd.write('%0.4f %0.4f %0.4f' % (verts[i,0], verts[i,1], verts[i,2]))
            if vert_colors is not None:
                fd.write(' %d %d %d' % (vert_colors[i,0], vert_colors[i,1], vert_colors[i,2]))
            if vert_normals is not None:
                fd.write(' %f %f %f' % (vert_normals[i,0], vert_normals[i,1], vert_normals[i,2]))
            if vert_uv is not None:
                fd.write(' %0.4f %0.4f' % (vert_uv[i,0], vert_uv[i,1]))
            fd.write('\n')
        if faces is not None:
            for i in range(num_faces):
                fd.write('3 %d %d %d\n' % (faces[i,0], faces[i,1], faces[i,2]))
