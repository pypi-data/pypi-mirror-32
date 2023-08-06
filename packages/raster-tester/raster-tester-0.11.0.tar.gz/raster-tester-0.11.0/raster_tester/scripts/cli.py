
import sys

import click
import raster_tester


@click.group()
def cli():
    pass


@cli.command("compare")
@click.argument("input_1", type=click.Path(exists=True))
@click.argument("input_2", type=click.Path(exists=True))
@click.option("--pixel-threshold", "-p", type=int, default=0,
              help='Threshold for pixel diffs [default=0]')
@click.option("--downsample", "-d", type=int, default=1,
              help='Downsample via decimated read for faster comparison, and to handle variation in compression artifacts [default=1]')
@click.option("--upsample", "-u", type=int, default=1,
              help='Upsample to handle variation in compression artifacts [default=1]')
@click.option("--compare-masked", is_flag=True,
              help='Only compare masks + unmasked areas of RGBA rasters')
@click.option("--no-error", is_flag=True,
              help='Compare in non stderr mode: echos "(ok|not ok) - <input_1> is (within|not within) <pixel-threshold> pixels of <input 2>"')
@click.option("--debug", is_flag=True,
              help='Print ascii preview of errors')
@click.option("--flex-mode", is_flag=True,
              help='Allow comparison of masked RGB + RGBA')
def compare(input_1, input_2, pixel_threshold, upsample, downsample,
            no_error, compare_masked, debug, flex_mode):
    raster_tester.compare(
        input_1, input_2, pixel_threshold, upsample, downsample,
        compare_masked, no_error, debug, flex_mode)


@cli.command("isempty")
@click.argument("input_1", type=click.Path(exists=True))
@click.option('--bidx', '-b', default=4,
              help="Bands to blob [default = 4]")
@click.option("--randomize", is_flag=True,
              help='iterate through windows in a psuedorandom fashion')
def isempty(input_1, randomize, bidx):
    empty = raster_tester.is_empty(input_1, randomize, bidx, )
    if empty:
        click.echo('{} is empty'.format(input_1))
    else:
        raise click.ClickException("{} is not empty".format(input_1))


@cli.command("isaligned")
@click.argument('sources', required=True, nargs=-1)
def isaligned(sources):
    aligned, msg = raster_tester.aligned(sources)
    if aligned:
        click.echo("ok: {} are aligned ({})".format(', '.join(sources), msg))
    else:
        raise click.ClickException("not ok: {} are not aligned ({})".format(', '.join(sources), msg))


@cli.command("istiled")
@click.argument('sources', required=True, nargs=-1)
@click.option('--blocksize/--no-blocksize', is_flag=True, default=True,
              help="assert that sources are internally tiled")
def istiled(sources, blocksize):
    result, msg = raster_tester.tiled(sources, blocksize)
    if result:
        click.echo("ok: {} are tiled ({})".format(', '.join(sources), msg))
    else:
        raise click.ClickException("not ok: {} are not all tiled ({})".format(', '.join(sources), msg))


@cli.command('crossesdateline')
@click.argument('input', type=click.Path(exists=True))
def crossesdateline(input):
    result = raster_tester.crosses_dateline(input)
    if result:
        raise click.ClickException('{} crosses dateline; exit 1'.format(input))
    else:
        click.echo('{} does not cross dateline; exit 0'.format(input))
