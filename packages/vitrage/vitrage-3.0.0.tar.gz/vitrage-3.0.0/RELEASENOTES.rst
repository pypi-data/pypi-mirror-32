=======
vitrage
=======

.. _vitrage_3.0.0:

3.0.0
=====

.. _vitrage_3.0.0_New Features:

New Features
------------

.. releasenotes/notes/datasource-scaffold-2f5ee6f0d9f83099.yaml @ b'9cc98e8e46ca0c6170861eacbb9dc507b96b322a'

- Add a command line tool used as scaffold for creating new datasource.

.. releasenotes/notes/mock-datasource-1c9b427bc16f4a59.yaml @ b'0a1017d7f003077d6d0ee7be8036beaefa53f5f5'

- Added a new ``Mock datasource``, which can mock an entire graph and allows testing large scale stability as well as performance.

.. releasenotes/notes/services-management-improvements-899c011e57002e84.yaml @ b'0a1017d7f003077d6d0ee7be8036beaefa53f5f5'

- The collector service was changed to run on demand instead of periodically, hence it can now be run in active-active mode. This is as part of a larger design to improve high availability.

.. releasenotes/notes/services-management-improvements-899c011e57002e84.yaml @ b'0a1017d7f003077d6d0ee7be8036beaefa53f5f5'

- Oslo service was replaced by cotyledon, so Vitrage uses real threads and multiprocessing. This change removes unnecessary complications of using eventlets and timers.

.. releasenotes/notes/services-management-improvements-899c011e57002e84.yaml @ b'0a1017d7f003077d6d0ee7be8036beaefa53f5f5'

- Created a dedicated process for the api handler, for better handling api calls under stress.

.. releasenotes/notes/static-datasource-changes-914f9a16ad7e46ed.yaml @ b'd5b4daf8aae7ee6ccf2ee9ce7a410e82f2d03921'

- Support get_changes in the static datasource

.. releasenotes/notes/support-get-changes-in-static-datasource-02715226f103455d.yaml @ b'd5b4daf8aae7ee6ccf2ee9ce7a410e82f2d03921'

- The static datasource now supports changes in existing yaml files, and updates the graph accordingly.


.. _vitrage_3.0.0_Bug Fixes:

Bug Fixes
---------

.. releasenotes/notes/performance-bug-fixes-a2b5e85ee37bba93.yaml @ b'0a1017d7f003077d6d0ee7be8036beaefa53f5f5'

- Many bug fixes related to performance and stability.

