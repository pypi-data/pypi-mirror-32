================
ironic-inspector
================

.. _ironic-inspector_7.3.0:

7.3.0
=====

.. _ironic-inspector_7.3.0_New Features:

New Features
------------

.. releasenotes/notes/pxe-filter-dnsmasq-not-known-hosts-filter-76ae5bd7a8db6f75.yaml @ 5e54e72136be734c5b4c0d4487c24ecc6bd984d5

- Adds wildcard ignore entry to ``dnsmasq`` PXE filter. When node
  introspection is active, or if ``node_not_found_hook`` is set in the
  configuration the ignore is removed from the wildcard entry. This ensures
  that unknown nodes do not accidentally boot into the introspection image
  when no node introspection is active.
  
  This brings ``dnsmasq`` PXE filter driver feature parity with the
  ``iptables`` PXE filter driver, which uses a firewall rule to block any
  DHCP request on the interface where Ironic Inspector's DHCP server is
  listening.

.. releasenotes/notes/sighup-support-e6eaec034d963108.yaml @ 76898b73824104486c243a8a6ac61c43a463184e

- Issuing a SIGHUP to the ironic-inspector service will cause the service to
  reload and use any changed values for *mutable* configuration options.
  
  Mutable configuration options are indicated as such in the `sample
  configuration file <https://docs.openstack.org/ironic-inspector/latest/configuration/sample-config.html>`_
  by ``Note: This option can be changed without restarting``.
  
  A warning is logged for any changes to immutable configuration options.


.. _ironic-inspector_7.3.0_Upgrade Notes:

Upgrade Notes
-------------

.. releasenotes/notes/discovery-default-driver-94f990bb0676369b.yaml @ 66f318b339ec4250ffa42c6239c76ca96942ecd6

- The ``[discovery]enroll_node_driver`` option, specifying the hardware type
  or driver to use for newly discovered nodes, was changed from ``fake``
  classic driver to ``fake-hardware`` hardware type.

.. releasenotes/notes/port-list-retry-745d1cf41780e961.yaml @ 3237511cc630c0f1da4cb89a71ab2d945080bc7c

- Adds dependency on the `retrying <https://github.com/rholder/retrying>`_
  python library.


.. _ironic-inspector_7.3.0_Bug Fixes:

Bug Fixes
---------

.. releasenotes/notes/fix-llc-switch-id-not-mac-e2de3adc0945ee70.yaml @ 97282c64e97485fb40d7c64413745c74e6ead942

- Fixes bug in which the ``switch_id`` field in a port's ``local_link_connection`` can be set to
  a non-MAC address if the processed LLDP has a value other than a
  MAC address for ``ChassisID``. The bare metal API requires the ``switch_id``
  field to be a MAC address, and will return an error otherwise.
  See `bug 1748022 <https://bugs.launchpad.net/ironic-inspector/+bug/1748022>`_
  for details.

.. releasenotes/notes/keystone-noauth-9ba5ad9884c6273c.yaml @ fbeb0783e4b2f1fc5f44da7f7b704ce48bfe0087

- Ironic introspection no longer tries to access the Identity service if the
  ``auth_strategy`` option is set to ``noauth`` and the ``auth_type`` option
  is not set to ``none``.

.. releasenotes/notes/port-list-retry-745d1cf41780e961.yaml @ 3237511cc630c0f1da4cb89a71ab2d945080bc7c

- The periodic PXE filter update task now retries fetching port list from
  the Bare Metal service 5 times (with 1 second delay) before giving up.
  This ensures that a temporary networking glitch will not result in
  the ironic-inspector service stopping.


.. _ironic-inspector_7.1.0:

7.1.0
=====

.. _ironic-inspector_7.1.0_Deprecation Notes:

Deprecation Notes
-----------------

.. releasenotes/notes/ksadapters-abc9edc63cafa405.yaml @ 918775cb0109209561fd62f1b52a3d9f7cc226c6

- Several configuration options related to ironic API access
  are deprecated and will be removed in the Rocky release.
  These include:
  
  - ``[ironic]/os_region`` - use ``[ironic]/region_name`` option instead
  - ``[ironic]/auth_strategy`` - set ``[ironic]/auth_type`` option to
    ``none`` to access ironic API in noauth mode
  - ``[ironic]/ironic_url`` - use ``[ironic]/endpoint_override`` option
    to set specific ironic API endpoint address if discovery of ironic API
    endpoint is not desired or impossible (for example in standalone mode)
  - ``[ironic]/os_service_type`` - use ``[ironic]/service_type`` option
  - ``[ironic]/os_endpoint_type`` - use ``[ironic]/valid_interfaces``
    option to set ironic endpoint types that will be attempted to be used

.. releasenotes/notes/ksadapters-abc9edc63cafa405.yaml @ 918775cb0109209561fd62f1b52a3d9f7cc226c6

- Several configuration options related to swift API access are deprecated
  and will be removed in Rocky release.
  These include:
  
  - ``[swift]/os_service_type`` - use ``[swift]/service_type`` option
  - ``[swift]/os_endpoint_type`` - use ``[swift]/valid_interfaces`` option
  - ``[swift]/os_region`` - use ``[swift]region_name`` option


.. _ironic-inspector_7.1.0_Other Notes:

Other Notes
-----------

.. releasenotes/notes/remove-policy-json-b4746d64c1511023.yaml @ 5c54c0938e8c61fe96b23a99549f3a35865c5cf8

- The sample configuration file located at ``example.conf``
  and the sample policy file located at ``policy.yaml.sample``
  were removed in this release, as they are now published with documentation.
  See `the sample configuration file
  <https://docs.openstack.org/ironic-inspector/latest/configuration/sample-config.html>`_
  and `the sample policy file
  <https://docs.openstack.org/ironic-inspector/latest/configuration/sample-policy.html>`_.


.. _ironic-inspector_7.0.0:

7.0.0
=====

.. _ironic-inspector_7.0.0_New Features:

New Features
------------

.. releasenotes/notes/dnsmasq-pxe-filter-37928d3fdb1e8ec3.yaml @ 8ddfacdf341670c923f7e41e7c3bb1986dc3bcaf

- Introduces the **dnsmasq** PXE filter driver. This driver takes advantage of
  the ``inotify`` facility to reconfigure the **dnsmasq** service in real time
  to implement a caching black-/white-list of port MAC addresses.


.. _ironic-inspector_7.0.0_Upgrade Notes:

Upgrade Notes
-------------

.. releasenotes/notes/db-status-consistency-enhancements-f97fbaccfc81a60b.yaml @ 7e72ceffd193afc268417df817c90a4ae89518d8

- A new state ``aborting`` was introduced to distinguish between the node
  introspection abort precondition (being able to perform the state
  transition from the ``waiting`` state) from the activities necessary to
  abort an ongoing node introspection (power-off, set finished timestamp
  etc.)

.. releasenotes/notes/local_gb-250bd415684a7855.yaml @ 6e82571cf367c1515c9e7b208077d912f9e9e873

- Handling of ``local_gb`` property was moved from the ``scheduler`` hook
  to ``root_disk_selection``.


.. _ironic-inspector_7.0.0_Bug Fixes:

Bug Fixes
---------

.. releasenotes/notes/db-status-consistency-enhancements-f97fbaccfc81a60b.yaml @ 7e72ceffd193afc268417df817c90a4ae89518d8

- The ``node_info.finished(<transition>, error=<error>)`` now updates node
  state together with other status attributes in a single DB transaction.


.. _ironic-inspector_7.0.0_Other Notes:

Other Notes
-----------

.. releasenotes/notes/tempest_plugin_removal-91a01f5950f543e1.yaml @ d7e18416805f699f9e259f30a9f98aa06ccb9499

- The tempest plugin code that was in
  ``ironic_inspector/test/inspector_tempest_plugin/`` has been removed.
  Tempest plugin code has been migrated to the project
  `openstack/ironic-tempest-plugin
  <https://git.openstack.org/cgit/openstack/ironic-tempest-plugin>`_. This was
  an OpenStack wide `goal for the Queens cycle
  <https://governance.openstack.org/tc/goals/queens/split-tempest-plugins.html>`_.


.. _ironic-inspector_6.1.0:

6.1.0
=====

.. _ironic-inspector_6.1.0_New Features:

New Features
------------

.. releasenotes/notes/firewall-refactoring-17e8ad764f2cde8d.yaml @ 4aeeaf7397e65978a6ea7d00268f006ab11442a5

- The PXE filter drivers mechanism is now enabled. The firewall-based
  filtering was re-implemented as the ``iptables`` PXE filter driver.

.. releasenotes/notes/policy-engine-c44828e3131e6c62.yaml @ 65d0213e84443586e8b1af68594c58c5cc0ca762

- Adds an API access policy enforcment based on **oslo.policy** rules.
  Similar to other OpenStack services, operators now can configure
  fine-grained access policies using ``policy.yaml`` file. See
  `policy.yaml.sample`_ in the code tree for the list of available policies
  and their default rules. This file can also be generated from the code tree
  with the following command::
  
      tox -egenpolicy
  
  See the `oslo.policy package documentation`_ for more information
  on using and configuring API access policies.
  
  .. _policy.yaml.sample: https://git.openstack.org/cgit/openstack/ironic-inspector/plain/policy.yaml.sample
  .. _oslo.policy package documentation: https://docs.openstack.org/oslo.policy/latest/


.. _ironic-inspector_6.1.0_Upgrade Notes:

Upgrade Notes
-------------

.. releasenotes/notes/policy-engine-c44828e3131e6c62.yaml @ 65d0213e84443586e8b1af68594c58c5cc0ca762

- Due to the choice of default values for API access policies rules,
  some API parts of the **ironic-inspector** service will become available
  to wider range of users after upgrade:
  
  - general access to the whole API is by default granted to a user
    with either ``admin``, ``administrator`` or ``baremetal_admin`` role
    (previously it allowed access only to a user with ``admin`` role)
  - listing of current introspection statuses and showing a given
    introspection is by default also allowed to a user with the
    ``baremetal_observer`` role
  
  If these access policies are not appropriate for your deployment, override
  them in a ``policy.json`` file in the **ironic-inspector** configuration
  directory (usually ``/etc/ironic-inspector``).
  
  See the `oslo.policy package documentation`_ for more information
  on using and configuring API access policies.
  
  .. _oslo.policy package documentation: https://docs.openstack.org/oslo.policy/latest/


.. _ironic-inspector_6.1.0_Deprecation Notes:

Deprecation Notes
-----------------

.. releasenotes/notes/firewall-refactoring-17e8ad764f2cde8d.yaml @ 4aeeaf7397e65978a6ea7d00268f006ab11442a5

- The firewall-specific configuration options were moved from the
  ``firewall`` to the ``iptables`` group. All options in the ``iptables``
  group are now deprecated.

.. releasenotes/notes/firewall-refactoring-17e8ad764f2cde8d.yaml @ 4aeeaf7397e65978a6ea7d00268f006ab11442a5

- The generic firewall options ``firewall_update_period`` and
  ``manage_firewall`` were moved under the ``pxe_filter`` group as
  ``sync_period`` and ``driver=iptables/noop`` respectively.


.. _ironic-inspector_6.1.0_Bug Fixes:

Bug Fixes
---------

.. releasenotes/notes/add_node-with-version_id-24f51e5888480aa0.yaml @ 82000e48ecdaa5738f6a7e69d94386977e714493

- A ``version_id`` is now explicitly generated during the
  ``node_cache.start_introspection/.add_node`` call to avoid race conditions
  such as in case of the `two concurrent introspection calls bug`_.
  
  .. _two concurrent introspection calls bug: https://bugs.launchpad.net/ironic-inspector/+bug/1719627

.. releasenotes/notes/empty-ipmi-address-2-4d57c34aec7d14e2.yaml @ a1d19d97b5973898b07ac05bb8c2c3cf5c199605

- The older ``ipmi_address`` field in the introspection data no longer has
  priority over the newer ``bmc_address`` inventory field during lookup.
  This fixes lookup based on MAC addresses, when the BMC address is reported
  as ``0.0.0.0`` for any reason (see `bug 1714944
  <https://bugs.launchpad.net/ironic-python-agent/+bug/1714944>`_).

.. releasenotes/notes/firewall-refactoring-17e8ad764f2cde8d.yaml @ 4aeeaf7397e65978a6ea7d00268f006ab11442a5

- Should the ``iptables`` PXE filter encounter an unexpected exception in the
  periodic ``sync`` call, the exception will be logged and the filter driver
  will be reset in order to make subsequent ``sync`` calls fail (and
  propagate the failure, exiting the **ironic-inspector** process eventually).


.. _ironic-inspector_6.1.0_Other Notes:

Other Notes
-----------

.. releasenotes/notes/allow-periodics-shutdown-inspector-ac28ea5ba3224279.yaml @ 65d0213e84443586e8b1af68594c58c5cc0ca762

- Allows a periodic task to shut down an **ironic-inspector** process
  upon a failure.

