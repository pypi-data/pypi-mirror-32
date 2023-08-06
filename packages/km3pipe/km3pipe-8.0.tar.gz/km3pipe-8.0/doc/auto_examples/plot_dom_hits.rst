

.. _sphx_glr_auto_examples_plot_dom_hits.py:


==================
DOM hits.
==================

Estimate track/DOM distances using the number of hits per DOM.




.. image:: /auto_examples/images/sphx_glr_plot_dom_hits_001.png
    :align: center


.. rst-class:: sphx-glr-script-out

 Out::

    Loading style definitions from '/Users/tamasgal/Dev/km3pipe/km3pipe/kp-data/stylelib/km3pipe.mplstyle'
    Detector: Parsing the DETX header
    Detector: Reading PMT information...
    Detector: Done.
    km3pipe.io.hdf5.HDF5Pump: Reading group information from '/group_info'.
    Pipeline and module initialisation took 0.030s (CPU 0.020s).
    --------------------------[ Blob     100 ]---------------------------
    --------------------------[ Blob     200 ]---------------------------
    --------------------------[ Blob     300 ]---------------------------
    --------------------------[ Blob     400 ]---------------------------
    --------------------------[ Blob     500 ]---------------------------
    ================================[ . ]================================
             distance  n_hits
    0      319.432073       3
    1       84.027581       3
    2       74.933425       4
    3       82.240334       4
    4      102.219993       2
    5       89.088059       2
    6      398.380855       2
    7       70.289436       3
    8       61.301031       8
    9       52.498108       6
    10      43.992152       5
    11      35.994317       2
    12      28.929146      11
    13      23.647904      16
    14      21.507544      15
    15      23.386937      23
    16      28.501691      34
    17     378.369175       2
    18     113.416443       2
    19      71.812430       2
    20      53.919127       4
    21      45.320313       2
    22      37.170048       3
    23      29.838187       2
    24      24.084136       5
    25      21.231982      18
    26      22.418112       6
    27      27.117701       7
    28      33.899718       7
    29      41.761648       7
    ...           ...     ...
    20888  511.669102       2
    20889   36.874603      14
    20890   31.723224      16
    20891   28.047885      11
    20892   26.470722       7
    20893   30.492830       5
    20894   35.283296       6
    20895   47.709439       3
    20896   54.702576       2
    20897   61.985865       2
    20898   69.468104       2
    20899   84.817689       5
    20900   69.092469       2
    20901   61.235423       8
    20902   53.474098       5
    20903   45.857122      12
    20904   38.470331      14
    20905   31.476205      16
    20906   25.203789      42
    20907   20.332357      25
    20908   18.035324      43
    20909   19.257086      33
    20910   23.454044      22
    20911   29.377599      11
    20912   87.368194       2
    20913   79.212811       5
    20914   71.062168       7
    20915   62.918109       6
    20916   54.783569       4
    20917   46.663528      17

    [20918 rows x 2 columns]
    ============================================================
    500 cycles drained in 11.084468s (CPU 10.136821s). Memory peak: 274.55 MB
      wall  mean: 0.022019s  medi: 0.016253s  min: 0.005869s  max: 2.399919s  std: 0.106733s
      CPU   mean: 0.020147s  medi: 0.014496s  min: 0.005852s  max: 2.357312s  std: 0.104823s




|


.. code-block:: python


    # Author: Tamas Gal <tgal@km3net.de>
    # License: BSD-3

    from collections import defaultdict, Counter

    import numpy as np
    import pandas as pd

    import matplotlib.pyplot as plt
    from matplotlib.colors import LogNorm

    import km3pipe as kp
    from km3pipe.dataclasses import Table
    from km3pipe.math import pld3
    from km3modules.common import StatusBar
    import km3pipe.style
    km3pipe.style.use("km3pipe")


    filename = "data/km3net_jul13_90m_muatm50T655.km3_v5r1.JTE_r2356.root.0-499.h5"
    cal = kp.calib.Calibration(filename="data/km3net_jul13_90m_r1494_corrected.detx")


    def filter_muons(blob):
        """Write all muons from McTracks to Muons."""
        tracks = blob['McTracks']
        muons = tracks[tracks.type == 5]
        blob["Muons"] = Table(muons)
        return blob


    class DOMHits(kp.Module):
        """Create histogram with n_hits and distance of hit to track."""

        def configure(self):
            self.hit_statistics = defaultdict(list)

        def process(self, blob):
            hits = blob['Hits']
            muons = blob['Muons']

            highest_energetic_muon = Table(muons[np.argmax(muons.energy)])
            muon = highest_energetic_muon

            triggered_hits = hits[hits.triggered.astype(bool)]

            dom_hits = Counter(triggered_hits.dom_id)
            for dom_id, n_hits in dom_hits.items():
                try:
                    distance = pld3(cal.detector.dom_positions[dom_id],
                                    muon.pos,
                                    muon.dir)
                except KeyError:
                    self.log.warning("DOM ID %s not found!" % dom_id)
                    continue
                self.hit_statistics['n_hits'].append(n_hits)
                self.hit_statistics['distance'].append(distance)
            return blob

        def finish(self):
            df = pd.DataFrame(self.hit_statistics)
            print(df)
            sdf = df[(df['distance'] < 200) & (df['n_hits'] < 50)]
            bins = (max(sdf['distance']) - 1, max(sdf['n_hits']) - 1)
            plt.hist2d(sdf['distance'], sdf['n_hits'], cmap='plasma', bins=bins,
                       norm=LogNorm())
            plt.xlabel('Distance between hit and muon track [m]')
            plt.ylabel('Number of hits on DOM')
            plt.show()


    pipe = kp.Pipeline()
    pipe.attach(kp.io.HDF5Pump, filename=filename)
    pipe.attach(StatusBar, every=100)
    pipe.attach(filter_muons)
    pipe.attach(DOMHits)
    pipe.drain()

**Total running time of the script:** ( 0 minutes  18.133 seconds)



.. container:: sphx-glr-footer


  .. container:: sphx-glr-download

     :download:`Download Python source code: plot_dom_hits.py <plot_dom_hits.py>`



  .. container:: sphx-glr-download

     :download:`Download Jupyter notebook: plot_dom_hits.ipynb <plot_dom_hits.ipynb>`

.. rst-class:: sphx-glr-signature

    `Generated by Sphinx-Gallery <https://sphinx-gallery.readthedocs.io>`_
