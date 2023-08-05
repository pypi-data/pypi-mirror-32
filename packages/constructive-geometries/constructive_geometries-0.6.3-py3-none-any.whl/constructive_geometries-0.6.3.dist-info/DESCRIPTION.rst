All rights reserved.

Redistribution and use in source and binary forms, with or without 
modification, are permitted provided that the following conditions are met:

Redistributions of source code must retain the above copyright notice, this 
list of conditions and the following disclaimer. Redistributions in binary 
form must reproduce the above copyright notice, this list of conditions and the 
following disclaimer in the documentation and/or other materials provided 
with the distribution.

Neither the name of Paul Scherrer Institut nor the names of its contributors 
may be used to endorse or promote products derived from this software without 
specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" 
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE 
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE 
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE 
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL 
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR 
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER 
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, 
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE 
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

Description-Content-Type: UNKNOWN
Description: Constructive Geometries - Python library
        ========================================
        
        |Documentation Status| |Build Status| |Coverage Status|
        
        Simple tools to define world locations from a set of topological faces
        and set algebra. For example, one could define a “rest of the world”
        which started from all countries, but excluded every country who name
        started with the letter “a”.
        
        `Documentation <http://constructive-geometries.readthedocs.io/?badge=latest>`__
        and `usage
        example <https://github.com/cmutel/constructive_geometries/blob/master/examples/Geomatching.ipynb>`__.
        
        Builds on top of `constructive
        geometries <https://github.com/cmutel/constructive_geometries>`__.
        
        Basic installation needs
        `wrapt <http://wrapt.readthedocs.io/en/latest/>`__ and
        `country_converter <https://github.com/konstantinstadler/country_converter>`__;
        GIS functions need `shapely <https://github.com/Toblerity/Shapely>`__,
        and `fiona <https://github.com/Toblerity/Fiona>`__.
        
        .. |Documentation Status| image:: https://readthedocs.org/projects/constructive-geometries/badge/?version=latest
           :target: http://constructive-geometries.readthedocs.io/?badge=latest
        .. |Build Status| image:: https://travis-ci.org/cmutel/constructive_geometries.svg?branch=master
           :target: https://travis-ci.org/cmutel/constructive_geometries
        .. |Coverage Status| image:: https://coveralls.io/repos/github/cmutel/constructive_geometries/badge.svg?branch=master
           :target: https://coveralls.io/github/cmutel/constructive_geometries?branch=master
        
Platform: UNKNOWN
Classifier: Intended Audience :: Developers
Classifier: Intended Audience :: Science/Research
Classifier: License :: OSI Approved :: BSD License
Classifier: Operating System :: MacOS :: MacOS X
Classifier: Operating System :: Microsoft :: Windows
Classifier: Operating System :: POSIX
Classifier: Programming Language :: Python
