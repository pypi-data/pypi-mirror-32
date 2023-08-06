=================
python-novaclient
=================

.. _python-novaclient_10.3.0:

10.3.0
======

.. _python-novaclient_10.3.0_New Features:

New Features
------------

.. releasenotes/notes/microversion-v2_62-479a23f0d4307500.yaml @ 5483be7fe74a90e3a38428cfb436864ffeee4c54

- Adds support for microversion 2.62 which adds ``host`` (hostname)
  and ``hostId`` (an obfuscated hashed host id string) fields to the
  instance action ``GET /servers/{server_id}/os-instance-actions/{req_id}``
  API.
  
  The event columns are already included in the result of
  "nova instance-action <server> <request-id>" command, therefore does not
  have any CLI or python API binding impacts in the client.


.. _python-novaclient_10.3.0_Upgrade Notes:

Upgrade Notes
-------------

.. releasenotes/notes/bug-1767287-cc28d60d9e59f9bd.yaml @ 57e9a5d34cde8cef319487cb56eca383cff76059

- The ``nova server-group-create`` command now only supports specifying
  a single policy name when creating the server group. This is to match
  the server-side API validation.


.. _python-novaclient_10.2.0:

10.2.0
======

.. _python-novaclient_10.2.0_New Features:

New Features
------------

.. releasenotes/notes/show-instance-usage-audit-logs-7826b411fac1283b.yaml @ d418b5f245f4cef4d35b8795aa6af8b98cd60141

- Added new client API and CLI (``nova instance-usage-audit-log``) to get server usage audit logs. By default, it lists usage audits for all servers on all compute hosts where usage auditing is configured. If you specify the ``--before`` option, the result is filtered by the date and time before which to list server usage audits.

.. releasenotes/notes/strict_hostname_match-f37243f0520a09a2.yaml @ 9213ec2d32fa173ec9943c28fb6c3ba5c196015d

- Provides "--strict" option for  "nova host-servers-migrate", "nova host-evacuate",
  "nova host-evacuate-live" and "nova host-meta" commands. When "--strict" option is
  used, the action will be applied to a single compute with the exact hypervisor
  hostname string match rather than to the computes with hostname substring match.
  When the specified hostname does not exist in the system, "NotFound" error code
  will be returned.


.. _python-novaclient_10.2.0_Upgrade Notes:

Upgrade Notes
-------------

.. releasenotes/notes/bug-1764420-flavor-delete-output-7b80f73deee5a869.yaml @ bcc7d8f1138ea22207ac0d31c5be132d6f274b34

- The ``flavor-delete`` command no longer prints out the details of the
  deleted flavor. On successful deletion, there is no output.


.. _python-novaclient_10.2.0_Deprecation Notes:

Deprecation Notes
-----------------

.. releasenotes/notes/get-rid-off-redundant-methods-47e679c13e88f28a.yaml @ ca27736810a41080761e604e09c6763aaed07ed7

- ``novaclient.utils.add_resource_manager_extra_kwargs_hook`` and
  ``novaclient.utils.get_resource_manager_extra_kwargs`` were designed for
  supporting extensions in nova/novaclient. Nowadays, this "extensions"
  feature is abandoned and both ``add_resource_manager_extra_kwargs_hook``,
  ``add_resource_manager_extra_kwargs_hook`` are not used in novaclient's
  code. These methods are not documented, so we are removing them without
  standard deprecation cycle.


.. _python-novaclient_10.2.0_Other Notes:

Other Notes
-----------

.. releasenotes/notes/microversion-v2_61-9a8faa02fddf9ed6.yaml @ 229d0df752702700dd30ddbe6d94d5efd5477318

- Starting from microversion 2.61, the responses of the 'Flavor' APIs
  include the 'extra_specs' parameter. Therefore 'Flavors extra-specs'
  (os-extra_specs) API calls have been removed in the following commands
  since microversion 2.61.
  
  * ``nova flavor-list``
  * ``nova flavor-show``
  
  There are no behavior changes in the CLI. This is just a performance
  optimization.

