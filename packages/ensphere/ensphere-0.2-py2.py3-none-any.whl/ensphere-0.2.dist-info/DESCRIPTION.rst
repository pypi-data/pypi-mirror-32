# ensphere

Add XMP info for PanoSphere in the pictures.

## Requirements

Install `exempi` in advance.

On MacOS,

    brew install exempi

## Usage

    usage: ensphere [-h] [--version]
                    [--set UsePanoramaViewer=True,ProjectionType=equirectangular]
                    filenames [filenames ...]

    Add XMP metadata for PanoSphere in the pictures. (version 0.2)

    positional arguments:
      filenames             File names (JPEG files only)

    optional arguments:
      -h, --help            show this help message and exit
      --version, -V         show program's version number and exit
      --set UsePanoramaViewer=True,ProjectionType=equirectangular, -s UsePanoramaViewer=True,ProjectionType=equirectangular
                            Set metadata for PanSpheres. Available variables:
                            ['UsePanoramaViewer', 'CaptureSoftware',
                            'StitchingSoftware', 'ProjectionType',
                            'PoseHeadingDegrees', 'PosePitchDegrees',
                            'PoseRollDegrees', 'InitialViewHeadingDegrees',
                            'InitialViewPitchDegrees', 'InitialViewRollDegrees',
                            'InitialHorizontalFOVDegrees', 'FirstPhotoDate',
                            'LastPhotoDate', 'SourcePhotosCount',
                            'ExposureLockUsed', 'CroppedAreaImageWidthPixels',
                            'CroppedAreaImageHeightPixels', 'FullPanoWidthPixels',
                            'FullPanoHeightPixels', 'CroppedAreaLeftPixels',
                            'CroppedAreaTopPixels', 'InitialCameraDolly']


## Example

To add the default metadata for a equirectangular JPEG image.

    ensphere equirectangular.jpg

For a cropped image, you also have to specify the original image size with `-s` option.

    ensphere -s 'FullPanoWidthPixels=8000,FullPanoHeightPixels=4000,CroppedAreaLeftPixels=0,CroppedAreaTopPixels=0' partial.jpg



