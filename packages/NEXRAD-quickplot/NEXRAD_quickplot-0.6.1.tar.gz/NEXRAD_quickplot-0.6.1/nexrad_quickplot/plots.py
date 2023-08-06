from pathlib import Path
import xarray
from typing import Dict, Any
from matplotlib.pyplot import figure, draw, fignum_exists
import matplotlib.ticker as mticker
import matplotlib.dates as mdates
#
import cartopy
from cartopy.mpl.gridliner import LONGITUDE_FORMATTER, LATITUDE_FORMATTER
# WGS84 is the default, just calling it out explicity so somene doesn't wonder.
GREF = cartopy.crs.PlateCarree()  # globe=cartopy.crs.Globe(ellipse='WGS84')

labels = [[-117.1625, 32.715, 'San Diego'],
          [-87.9073, 41.9742, 'KORD'],
          [-90.3755, 38.7503, 'KSUS'],
          [-97.040443, 32.897480, 'KDFW'],
          [-104.6731667, 39.8616667, 'KDEN'],
          [-111.1502604, 45.7772358, 'KBZN'],
          [-106.6082622, 35.0389316, 'KABQ']
          ]


def overlay2d(img: xarray.DataArray, ofn: Path=None, mlp: Dict[str, Any]=None,
              verbose: bool=False) -> dict:
    """plot NEXRAD reflectivity on map coordinates"""
    def _savemap(ofn, fg):
        if ofn is not None:
            ofn = Path(ofn).expanduser()
            print('saving Nexrad map:', ofn, end='\r')
            fg.savefig(ofn, bbox_inches='tight')

    if mlp is not None and fignum_exists(mlp['fg'].number):
        mlp['himg'].set_data(img)
        mlp['ht'].set_text(img.filename.name)
        draw()
        _savemap(ofn, mlp['fg'])
        return mlp

    fg = figure(figsize=(15, 10))

    ax = fg.gca(projection=GREF)

    ht = ax.set_title(img.filename.name)

    ax.add_feature(cartopy.feature.COASTLINE, linewidth=0.5, linestyle=':')
    ax.add_feature(cartopy.feature.NaturalEarthFeature('cultural', 'admin_1_states_provinces',
                                                       '50m',
                                                       linestyle=':', linewidth=0.5, edgecolor='grey', facecolor='none'))

    if verbose:
        for l in labels:
            ax.plot(l[0], l[1], 'bo', markersize=7, transform=GREF)
            ax.annotate(l[2], xy=(l[0], l[1]), xytext=(3, 3), textcoords='offset points')

    himg = ax.imshow(img, origin='upper',
                     extent=[img.lon[0], img.lon[-1], img.lat[0], img.lat[-1]],
                     transform=GREF)
# %% grid lines and labels
    gl = ax.gridlines(crs=GREF, draw_labels=True,
                      linewidth=1, color='gray', alpha=0.5, linestyle='--')
    gl.xlabels_top = False
    gl.ylabels_left = False
    gl.xformatter = LONGITUDE_FORMATTER
    gl.yformatter = LATITUDE_FORMATTER
    # optional
    gl.xlocator = mticker.FixedLocator(range(-140, -40, 20))
    gl.ylocator = mticker.FixedLocator(range(20, 54, 4))

    draw()
    _savemap(ofn, fg)

    mlp = {'fg': fg, 'himg': himg, 'ht': ht}

    return mlp


def keogram(keo: xarray.DataArray, ofn: Path=None):
    """stack a single lat or lon index"""

# %%
    fg = figure(figsize=(15, 10))
    ax = fg.gca()

    tlim = mdates.date2num(keo.time[[0, -1]].values)

    ax.imshow(keo.values, origin='upper',
              aspect='auto',  # crucial for time-based imshow()
              extent=[tlim[0], tlim[1], keo.lon[0].item(), keo.lon[-1].item()])

    ax.xaxis_date()
    ax.xaxis.set_major_formatter(mdates.DateFormatter('%H:%M:%S'))
    fg.autofmt_xdate()

    ax.set_xlabel('Time [UTC]')
    ax.set_ylabel('Longitude [deg.]')
    ax.set_title(f'NEXRAD Keogram: cut at lat={keo.lat}\n'
                 f'{keo.time.values[0]} to {keo.time.values[-1]}')

    draw()
# %%
    if ofn is not None:
        ofn = Path(ofn).expanduser()
        print('saving keogram to', ofn)
        fg.savefig(ofn, bbox_inches='tight')
