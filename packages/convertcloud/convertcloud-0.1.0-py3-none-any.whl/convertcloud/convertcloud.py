#! /usr/bin/env python

import os
import sys
import struct
import io

import numpy as np

from .formats import Load, Header

class Converter:
    def __init__(self):

        self._rgb = None
        self._rgba = None
        self._decode = None

        self.points = None
        self.fields = None

    def _get_name(self, path):
        """
        Returns basename and extension of path
        """
        return os.path.splitext(os.path.basename(path))

    def load_points(self, path):
        self.points = []
        self.fields = []

        load = Load() 

        print("Reading: ", path)
        name, extension = self._get_name(path)
        if extension == ".pcd":
            self.points, self.fields = load.pcd(path)
        elif extension == ".ply":
            self.points, self.fields = load.ply(path)
        elif extension == ".zdf":
            self.points, self.fields = load.zdf(path)
        elif extension in [".stl", ".STL"]:
            self.points, self.fields = load.stl(path)
        elif extension == ".xyz":
            self.points = load.xyz(path)
        elif extension == ".a3d":
            self.points = load.a3d(path)
        else:
            print("Error: Unknown file extension {}".format(extension))
            sys.exit(1)

        self._decode_points()

        for field in self.fields:
            if field.name == 'red' and self._rgba == None:
                self._rgb = True
            elif field.name == 'alpha':
                self._rgba = True
                self._rgb = False

    def _decode_points(self):
        for num, point in enumerate(self.points):
            if isinstance(point[0], bytes):
                self.points[num] = [val.decode() for val in point]
            else:
                break

        self.points = np.array(self.points).astype("float32")

    def convert(self, path):
        print('Saving point cloud to', path)
        name, extension = self._get_name(path)

        header = self._generate_header(extension)

        with open(path, "wb") as f:
            f.write(header.encode())
            for pt in self.points:
                if self._rgb:
                    f.write("{} {} {} {} {} {}\n".format(\
                            pt[0], pt[1], pt[2],\
                            int(pt[3]), int(pt[4]), int(pt[5])).encode())

                elif self._rgba:
                    f.write("{} {} {} {} {} {} {}\n".format(\
                            pt[0], pt[1], pt[2],\
                            int(pt[3]), int(pt[4]), int(pt[5]), int(pt[6])).encode())

                else:
                    f.write("{} {} {}\n".format(pt[0], pt[1], pt[2]).encode())

    def _generate_header(self, extension):

        header_gen = Header(len(self.points), self.fields, self._rgb, self._rgba)

        if extension == ".ply":
            header = header_gen.ply()

        elif extension == ".pcd":
            header = header_gen.pcd()

        elif extension in [".xyz", ".a3d"]:
            header = ''

        else:
            print("Error: Can't convert to {}".format(extension))
            sys.exit(1)

        return header

def main():
    if len(sys.argv) != 3:
        print("usage: converter <original.format1> <converted.format2>")
        print("formats supported: .ply, .pcd, .xyz, .zdf")
        sys.exit(1)

    c = Converter()

    c.load_points(sys.argv[1])
    c.convert(sys.argv[2])

if __name__ == "__main__":
    main()
