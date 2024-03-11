pngloss
=======

Lossily compress your PNG images with pngloss. The program reads the original
PNG file, modifies the pixels to make them more compressible, and writes a new
file with the extension -loss.png.

The compression technique relies on making small adjustments to pixel colors.
It works best on true-color images with a wide variety of colors, like
photographs or computer generated graphics with realistic lighting. It does
not do a good job on paletted images or images with large areas of flat color.

### Heritage

The lossy compression in pngloss is based on an algorithm in Michael Vinther's
graphics editor, [Image Analyzer](http://meesoft.logicnet.dk/Analyzer/).

The command line tool is based on Kornel Lesiński's PNG quantization tool,
[pngquant](https://pngquant.org/). Additionally, pngloss includes his work to
port the lossy compression algorithm from Go to C as part of his PNG
compression suite, [ImageOptim](https://imageoptim.com/).

William MacKay brought these pieces together and made improvements to the
original compression algorithm, including:
- Diffuse color error using Sierra dithering.
- Instead of using a static, evenly spaced symbol table (e.g. 0, 20, 40...),
pngloss chooses symbols which are more frequent in the original file.
- Instead of using only PNG's "average" filter, pngloss tries all five PNG
filters for each row, keeping whichever has the best quality/size ratio.
- Explicitly tell libpng which filter to use for each row instead of hoping
its adaptive filter selection algorithm will choose the correct one.

### Installation

    git clone https://github.com/foobaz/pngloss.git
    cd pngloss
    make
    sudo make install

There is no configure script. The only dependency is libpng. The makefile installs the binary to `/usr/local/bin/pngloss` and the man page to `/usr/local/share/man/man1/pngloss.1`.

### Synopsis

`pngloss [options] <file> [<file>...]`

### Options

`-s`, `--strength`
How much quality to sacrifice, from 0 to 85 (default 19). Strength 0 is
lossless and does not modify the pixel data, although it may convert
colorspace or strip PNG chunks. The maximum is 85 because the algorithm
needs at least three symbols to perform well and 256 / 3 ~= 85. One symbol
near zero is used for the majority of pixels, plus one positive and one
negative to insert when color error builds up enough that they are needed.

`-b`, `--bleed`
Color bleed divider, from 1 to 32767 (default 2). A divider of 1
propagates all of the error from quantization to neighboring pixels, which
improves visual quality but also increases filesize. The default of 2
propagates half (1/2) of the error, which is usually a good tradeoff.

`-v`, `--verbose`
Verbose - print additional information about compression.

`-q`, `--quiet`
Quiet - don't print information about compression (the default).

`-f`, `--force`
Force - overwrite existing output image.

`--no-force`
Don't overwrite existing output image - overrides an earlier "force" argument.

`--ext`
Specify filename extension. Defaults to "-loss.png". Use `-f --ext .png` to
overwrite original files in-place if the original has the extension ".png".

`--skip-if-larger`
Don't write compressed image if it's larger than the original.

`-o`, `--output`
Output filename. When this option is given only one input file is accepted.

`--strip`
Remove unnecessary chunks (metadata) from input file when writing output.

`-V`, `--version`
Print version number.

`-h`, `--help`
Display usage information.

### Examples
| Original | -s 20 | -s 40 |
| :------: | :---: | :---: |
| ![David, original](http://frammish.org/pngloss/david.png) | ![David, s20](http://frammish.org/pngloss/david-s20.png) | ![David, s40](http://frammish.org/pngloss/david-s40.png) |
| 18kB | 7kB (36%) | 5kB (23%) |

| Original | -s 20 | -s 40 |
| :------: | :---: | :---: |
| ![Lena, original](http://frammish.org/pngloss/lena.png) | ![Lena, s20](http://frammish.org/pngloss/lena-s20.png) | ![Lena, s40](http://frammish.org/pngloss/lena-s40.png) |
| 475kB | 65kB (13%) | 35kB (7%) |

| Original | -s 20 | -s 40 |
| :------: | :---: | :---: |
| ![Tenko, original](http://frammish.org/pngloss/tenko.png) | ![Tenko, s20](http://frammish.org/pngloss/tenko-s20.png) | ![Tenko, s40](http://frammish.org/pngloss/tenko-s40.png) |
| 234kB | 47kB (20%) | 30kB (13%) |

See origin samples in [pngloss-samples](https://github.com/ImageProcessing-ElectronicPublications/pngloss-samples).

