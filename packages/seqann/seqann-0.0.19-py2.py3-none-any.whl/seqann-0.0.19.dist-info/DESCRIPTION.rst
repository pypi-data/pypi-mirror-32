===============================
SeqAnn
===============================


.. image:: https://img.shields.io/pypi/v/seqann.svg
        :target: https://pypi.python.org/pypi/seqann

.. image:: https://img.shields.io/travis/nmdp-bioinformatics/seqann.svg
        :target: https://travis-ci.org/nmdp-bioinformatics/seqann

.. image:: https://readthedocs.org/projects/seqann/badge/?version=latest
        :target: https://seqann.readthedocs.io/en/latest/?badge=latest
        :alt: Documentation Status

.. image:: https://pyup.io/repos/github/nmdp-bioinformatics/gfe/shield.svg
     :target: https://pyup.io/repos/github/nmdp-bioinformatics/seqann/
     :alt: Updates


Sequence Annotation


* Free software: LGPL 3.0
* Documentation: https://seqann.readthedocs.io.


Docker
--------
* docker pull nmdpbioinformatics/pygfe

.. code-block:: 

	docker run -it --rm -v $PWD:/opt nmdpbioinformatics/pygfe seq2gfe \
		-f /opt/your_fastafile.fasta -l HLA-A



Features
--------

With mysql connection:

.. code-block:: python3

	from seqann
	from Bio import SeqIO
	from BioSQL import BioSeqDatabase

	server = BioSeqDatabase.open_database(driver="pymysql", user="root",
	                                      passwd="", host="localhost",
	                                      db="bioseqdb")
	seqann = seqann.BioSeqAnn(server=server)
	for seq in SeqIO.parse(input_seq, "fasta"):
		annotation = seqann.annotate(seq, "HLA-A")
		for feat in annotation.annotation:
			print(feat, annotation.annotation[feat], sep="\t")


Without mysql connection:

.. code-block:: python3

	import seqann
	from Bio import SeqIO

	# ** If you don't have a copy of the hla.dat
	# ** file it will download it
	seqann = seqann.BioSeqAnn()
	for seq in SeqIO.parse(input_seq, "fasta"):
		annotation = seqann.annotate(seq, "HLA-A")
		for feat in annotation.annotation:
			print(feat, annotation.annotation[feat], sep="\t")


Dependencies
------------
* `Clustal Omega`_ 1.2.0 or higher
* `Python 3.6`_
* blastn_

Credits
---------

This package was created with Cookiecutter_ and the `audreyr/cookiecutter-pypackage`_ project template.

.. _`Python 3.6`: https://www.python.org/downloads
.. _`Clustal Omega`: http://www.clustal.org/omega/
.. _blastn: https://blast.ncbi.nlm.nih.gov/Blast.cgi?PAGE_TYPE=BlastDocs&DOC_TYPE=Download
.. _Cookiecutter: https://github.com/audreyr/cookiecutter
.. _`audreyr/cookiecutter-pypackage`: https://github.com/audreyr/cookiecutter-pypackage



=======
History
=======

0.0.1 (2017-10-19)
------------------

* First release on PyPI.


