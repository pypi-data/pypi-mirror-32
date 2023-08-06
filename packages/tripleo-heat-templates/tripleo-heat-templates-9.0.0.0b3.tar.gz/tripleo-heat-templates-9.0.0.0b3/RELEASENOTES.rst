======================
tripleo-heat-templates
======================

.. _tripleo-heat-templates_9.0.0.0b3:

9.0.0.0b3
=========

.. _tripleo-heat-templates_9.0.0.0b3_Prelude:

Prelude
-------

.. releasenotes/notes/tls-inject-86ef6706e68f5740.yaml @ 59b762658d72977f838affa7fcf3d8d912e13678

TLS certificate injection is now done with Ansible instead of a bash script called by Heat.
It deprecates the NodeTLSData resource and script, while keeping the use of its variables (SSLCertificate, SSLIntermediateCertificate, SSLKey)


.. _tripleo-heat-templates_9.0.0.0b3_New Features:

New Features
------------

.. releasenotes/notes/Add-EnablePublicTLS-parameter-b3fcd01af6f3c101.yaml @ 1260da27461af826017afcbc765775832e5a9dde

- This adds a flag called EnablePublicTLS, which defaults to 'true'. It
  reflects that Public TLS is enabled by default, and it's read by
  the deployment workflow to let the public certificate generation happen.
  It can also be used to disable this feature, if it's set to 'false' as
  it's done in the no-tls-endpoints-public-ip.yaml environment
  file, which allows deployers to turn this feature off.

.. releasenotes/notes/Remove-version-from-KeystoneUrl-output-fe4ce6f1a45849d3.yaml @ 99263591319150ab286c712b6da87de901971fb7

- The KeystoneURL stack output is now versionless.

.. releasenotes/notes/add-cinder-backend-nvmeof-023d967980fcf7b8.yaml @ afcf2c71e31322d27b1b447409ee8a731847c64b

- Add NVMeOF as Cinder backend.

.. releasenotes/notes/collectd-overcloud-gnocchi-049a63bbd196a9bb.yaml @ 723e428f405a517ebec0e1ea77ba69a64d7d55ff

- Makes collectd deployment default output metrics data to Gnocchi instance
  running on overcloud nodes.

.. releasenotes/notes/collectd-polling-4aac123faaebd1bc.yaml @ 6c5b96c19202d3830b0a066a65198a7a909e7fe8

- Adds possibility to override default polling interval for collectd and set
  default value to 120 seconds, because current default (10s)
  was too aggressive.

.. releasenotes/notes/config-download-default-to-true-2331debd56c396eb.yaml @ f44e8d7bd2615dfdb0ed15b3bbd45a4475ae0152

- The mappings from environments/config-download-environment.yaml are now included by default in overcloud-resource-registry.j2.yaml. config-download is now the default way of deploying. An environment at environments/disable-config-download.yaml is added to enable the previous method, but that method is deprecated.

.. releasenotes/notes/containerize-neutron-lbaas-service-plugin-20562487d6631c88.yaml @ 9526cef547278a53c237f08c5b5e79948fc031dd

- Add support for Neutron LBaaSV2 service plugin in a containerized
  deployment.

.. releasenotes/notes/ctlplane_fixed_ip-81d14db5a01fa531.yaml @ 393476fda3652637bf6e5344f8815b3bb5398900

- Add ability to specify a fixed IP for the provisioning control plane
  (ctlplane) network. This works similarly to the existing fixed IPs
  for isolated networks, by including an environment file which includes
  an IP for each node in each role that should use a fixed IP. An example
  environment file is included in environments/ips-from-pool-ctlplane.yaml.

.. releasenotes/notes/merge_keys_from_services-cd17425d58b49840.yaml @ 8a8ad26435c43bc32e6cfe055ae69f208ab610ed

- It is now possible to specify values for any key in `config_settings`
  from multiple services; multiple values will be merged using YAQL
  mergeWith() function. For example, assuming two services defining
  a key as follows:
  
    config_settings:
      mykey:
      - val1
  
    config_settings:
      mykey:
      - val2
      - val3
  
  the content of the key, as seen by ansible or puppet on the nodes,
  will be:
  
    mykey: ['val1','val2','val3']

.. releasenotes/notes/octavia-amphora-image-defaults-0d9efe1a0222b76d.yaml @ 4d8a80f3860cffc9064c51cbe39b5a11ea150e04

- The Octavia amphora image name is now derived from the filename by default so the
  `OctaviaAmphoraImageName` now behaves as an override if set to a non-default value.

.. releasenotes/notes/octavia-amphora-image-defaults-0d9efe1a0222b76d.yaml @ 4d8a80f3860cffc9064c51cbe39b5a11ea150e04

- The Octavia amphora image file name default value is now an empty string resulting in a distribution specific default location being used. The `OctaviaAmphoraImageFilename` parameter now behaves as an override if set to a non-default value.

.. releasenotes/notes/octavia-amphora-ssh-245a21a35598440a.yaml @ 38eee383e52fa6e406c75f1e74b95493a10e54f8

- Allow users to specify SSH name and public key to add to Octavia amphorae.

.. releasenotes/notes/oslo-messaging-separate-backends-2d2221066f88f479.yaml @ 78bc45758563452ca5c3fd91afa901d93f9d8007

- Support separate oslo.messaging services for RPC and Notifications. Enable separate messaging backend servers.

.. releasenotes/notes/tripleo-nova-nfs-ead2827338aa9519.yaml @ 6b6ae966ba7880787d0584d7e8304e5d6c9c0093

- Allow NFS configuration of storage backend for Nova. This way
  the instance files will be stored on a shared NFS storage.

.. releasenotes/notes/update-lb-mgmt-subnet-to-class-b-1cd832ef08a30c85.yaml @ 5a28efc27da47782a10c5ba8450ee0e90527d908

- Enhance lb-mgmt-subnet to be a class B subnet, so the global amount of Octavia loadbalancers won't be constrained to a very low number.

.. releasenotes/notes/update_manila_unity_driver-43aeb041029c4e7f.yaml @ 79719a11ccde300aefcf3d98a8a55afdf3c09edd

- Adds network_plugin_ipv6_enabled, emc_ssl_cert_verify and
  emc_ssl_cert_path options for Manila Unity driver.

.. releasenotes/notes/update_manila_vnx_driver-678b22c4fcd81fcf.yaml @ 60796ebfc921fae598e125bfeef94a942e5a61b8

- Adds network_plugin_ipv6_enabled, emc_ssl_cert_verify and emc_ssl_cert_path options for Manila VNX driver.


.. _tripleo-heat-templates_9.0.0.0b3_Upgrade Notes:

Upgrade Notes
-------------

.. releasenotes/notes/direct-deploy-by-default-bc78a63f0a0c6e15.yaml @ 89de728acb7a734824ed61cc31cdf289da7e0f24

- Ironic in the containerized undercloud now uses the ``direct`` deploy
  interface by default for better performance and scalability. See
  `the direct deploy documentation
  <https://docs.openstack.org/ironic/latest/admin/interfaces/deploy.html#direct-deploy>`_
  for details.
  
  If undesired, this change can be reverted per node by setting the node's
  ``deploy_interface`` field to ``iscsi`` or globally by changing the new
  ``IronicDefaultDeployInterface`` to empty string.

.. releasenotes/notes/logrotate-containers-purge-a5587253fe6cbb28.yaml @ 62cdc3949f733b726fc1e25708b755e1a21dd9f7

- The 'LogrotatePurgeAfterDays'
  enforces cleaning up of information exceeded its life-time
  (defaults to a 14 days) in the /var/log/containers directory of
  bare metal overcloud hosts, including upgrade (from containers)
  cases, when leftovers may be remaining on the host systems.

.. releasenotes/notes/no-classic-drivers-9c59b696d8b50692.yaml @ a42373980f29c46f8655026329f21853ff61310f

- Support for deprecated classic drivers was removed from the Ironic
  templates. Please use ``IronicEnabledHardwareTypes`` and
  ``IronicEnabled***Interfaces`` parameters to enable/disable support
  for hardware types and interfaces.

.. releasenotes/notes/role-support-for-upgrade-to-dvr-containers-bc876f82f3e9f139.yaml @ f51f84e7818f7f70e4f6f298fff6d57509af4fbd

- Upgrading DVR deployments may require customization of the Compute role if
  they depend on the overcloud's external API network for floating IP
  connectivity. If necessary, please add "External" to the list of
  networks for the Compute role in roles_data.yaml before upgrading.

.. releasenotes/notes/tls-inject-86ef6706e68f5740.yaml @ 59b762658d72977f838affa7fcf3d8d912e13678

- All NodeTLSData related resources must be removed.

.. releasenotes/notes/tls-inject-86ef6706e68f5740.yaml @ 59b762658d72977f838affa7fcf3d8d912e13678

- SSLCertificate, SSLIntermediateCertificate, SSLKey are still used for the TLS configuration.

.. releasenotes/notes/update-lb-mgmt-subnet-to-class-b-1cd832ef08a30c85.yaml @ 5a28efc27da47782a10c5ba8450ee0e90527d908

- This fix is changing the default mask for lb-mgmt-subnet. Operators can either manually modify the already existing subnet or create a new one and update Octavia to use it. Newly create loadbalancer will use The newly created subnet.


.. _tripleo-heat-templates_9.0.0.0b3_Deprecation Notes:

Deprecation Notes
-----------------

.. releasenotes/notes/config-download-default-to-true-2331debd56c396eb.yaml @ f44e8d7bd2615dfdb0ed15b3bbd45a4475ae0152

- environments/disable-config-download.yaml can be used to disable config-download but is deprecated.

.. releasenotes/notes/deprecate_auth_uri_parameter-bdebdc6614ce8b7e.yaml @ 2b662be9a458d71101a1ba4c6a3b45c9cebdd272

- auth_uri is depreacted and will be removed in a future release. Please, use www_authenticate_uri instead.

.. releasenotes/notes/remove-support-for-puppet-ceph-bdafca24a59e7075.yaml @ 753a3504184d966c63121ca95bc0135afbc83a75

- Deployment of a managed Ceph cluster using puppet-ceph
  is not supported from the Pike release. From the Queens
  release it is not supported to use puppet-ceph when
  configuring OpenStack with an external Ceph cluster.
  In Rocky any support file necessary for the deployment
  with puppet-ceph is removed completely.

.. releasenotes/notes/remove-undercloud-specific-services-23046e607565d36d.yaml @ 64bc4a7683fab7e9d6feb67cd3252c4716722e6e

- The environment/services/undercloud-*.yaml files will be removed in the Stein
  release. These files relied on OS::TripleO::Services::Undercloud* services
  that have been removed.

.. releasenotes/notes/tls-inject-86ef6706e68f5740.yaml @ 59b762658d72977f838affa7fcf3d8d912e13678

- NodeTLSData is now deprecated.

.. releasenotes/notes/validate-no-config-outputs-used-8abcb673da6d373f.yaml @ a134b717dd787a5b6e28dc401260e22e77ef2162

- The use of outputs with Heat SoftwareConfig or StructuredConfig resources is now deprecated as they are no longer supported with config-download. Resources that depend on outputs and their values should be changed to use composable services with external_deploy_tasks or deploy_steps_tasks.


.. _tripleo-heat-templates_9.0.0.0b3_Security Issues:

Security Issues
---------------

.. releasenotes/notes/logrotate-containers-purge-a5587253fe6cbb28.yaml @ 62cdc3949f733b726fc1e25708b755e1a21dd9f7

- New heat parameters for containerized services 'LogrotateMaxsize',
  'LogrotateRotationInterval', 'LogrotateRotate' and
  'LogrotatePurgeAfterDays' allow customizing size/time-based rules
  for the containerized services logs rotation.
  The time based rules prevail over all.

.. releasenotes/notes/ssh_pass_auth-8cab3ca5a50d2a5a.yaml @ b749e027a031069625f0b71c2815499b686fbbf4

- PasswordAuthentication is enabled by default when deploying a containerized undercloud.
  We don't expect our operators to setup ssh keys during the initial deployment so we allow
  them to use the password to login into the undercloud node.


.. _tripleo-heat-templates_9.0.0.0b3_Bug Fixes:

Bug Fixes
---------

.. releasenotes/notes/add_site_id_cisco_ml2-60cfa450637d4fe0.yaml @ e52d7a552c9ca3b1eda344f4637405777cf9ad2d

- Add VTSSideId parameter to Cisco VTS ML2 template.

.. releasenotes/notes/convert-resource-name-to-number-80ada6c825554f56.yaml @ 49d072133563dde3c7693e8dacbfdaafac121329

- Previously, get-occ-config.sh could configure nodes out of order when deploying with more than 10 nodes. The script has been updated to properly sort the node resource names by first converting the names to a number.

.. releasenotes/notes/default-octavia-ssh-pub-key-to-keypair-70377d43bf76a407.yaml @ 0e87e640c88c316a8bc8d75974c8ac79aca868be

- Default Octavia SSH public key to 'default' keypair from undercloud.

.. releasenotes/notes/fix_nova_host-0b82c88597703353.yaml @ 31e4c0194dd1d6e049a728c876347df93ce89908

- The nova/neutron/ceilometer host parameter is now explicitly set to the
  same value that is written to /etc/hosts. On a correctly configured
  deployment they should be already be identical. However if the hostname
  or domainname is altered (e.g via DHCP) then the hostname is unlikely to
  resolve to the correct IP address for live-migraiton.
  Related bug: https://bugs.launchpad.net/tripleo/+bug/1758034

.. releasenotes/notes/live_migration_inbound_addr_all_transports-2fc9cd74d435a367.yaml @ 9faea7204ca56561a92b45426fc8047e5b48fe61

- Set live_migration_inbound_addr for ssh transport
  
  Previously this was only set when TLS is enabled, which means that with the ssh
  transport we could not control the network used, and were relying on DNS or
  hosts file to be correct, which is not guaranteed (especially with DNS).

.. releasenotes/notes/live_migration_port_range-54c28faf0a67a3fc.yaml @ 3da3f5d8de91181fa6c65ecfcf8d4733e000ace4

- By default, libvirtd uses ports from 49152 to 49215 for live-migration
  as specified in qemu.conf, that becomes a subset of ephemeral ports
  (from 32768 to 61000) used by many linux kernels.
  The issue here is that these ephemeral ports are used for outgoing TCP
  sockets. And live-migration might fail, if there are no port available
  from the specified range.
  Moving the port range out of ephemeral port range to be used only for
  live-migration.

.. releasenotes/notes/odl_delete_data_folder-b8c2f9a9382fd692.yaml @ 871e9619d5fdfa201736a2c6f0bafb5a4c56d89d

- Delete ODL data folder while updating/upgrading ODL.


.. _tripleo-heat-templates_9.0.0.0b3_Other Notes:

Other Notes
-----------

.. releasenotes/notes/add_neutron_segments_plugin_to_default-8acb69b112d4b31c.yaml @ bc3600b3628605436834d4712a745dfe9c43ddf4

- Add "segments" service plugin to the default list of
  neutron service plugins.


.. _tripleo-heat-templates_9.0.0.0b2:

9.0.0.0b2
=========

.. _tripleo-heat-templates_9.0.0.0b2_New Features:

New Features
------------

.. releasenotes/notes/add-cinder-backup-nfs-backend-0108fba91a3058ea.yaml @ e456e103fbfb5d7dd12f4dcbee4fd6686384117c

- Adds support for configuring the cinder-backup service with an NFS backend.

.. releasenotes/notes/add-purge-nova-tables-e0706cdcffa0f42e.yaml @ fb29f77987d9bb1009f99edf55758309e8c710dc

- Add the ability of fully purging the shadow
  tables whether in the archive or the purge
  cron.

.. releasenotes/notes/add_params_to_configure_ulimit-2359aa058da58054.yaml @ 70276931a4664d0bedcd6b0caa5a9cac2b73187b

- Add Parameters to Configure Ulimit for Containers.
  These parameters can be used to configure ulimit
  per container basis as per the requirement of the
  deployment.
  Following parameters are added for neutron, nova
  and cinder:-
  - DockerNeutronDHCPAgentUlimit defaults to nofile=1024
  - DockerNeutronL3AgentUlimit defaults to nofile=1024
  - DockerOpenvswitchUlimit defaults to nofile=1024
  - DockerNovaComputeUlimit defaults to nofile=1024
  - DockerCinderVolumeUlimit defaults to nofile=131072

.. releasenotes/notes/configure-ip-forward-268c165708cbd203.yaml @ 75ee85b1e45b09ac3093d3ace1112d5c3be18074

- Add KernelIpForward configuration to enable/disable the net.ipv4.ip_forward
  configuration.

.. releasenotes/notes/containerized-tempest-support-0ceaaf6427ce36e9.yaml @ 06638e76b74e58eccdfdee9c475dc529bb8bdf5a

- Added containerized tempest support in undercloud.

.. releasenotes/notes/containerized-tempest-support-0ceaaf6427ce36e9.yaml @ 06638e76b74e58eccdfdee9c475dc529bb8bdf5a

- Add DockerTempestImage parameter to add a fake tempest service which makes sure tempest container exists on the undercloud.

.. releasenotes/notes/containers-as-default-37bbe8afa0a60c2b.yaml @ 6c5f2b8f69d42b2201f29264360ca15a518f376b

- Containers are now the default way of deploying. There is still a way to
  deploy the baremetal services in environments/baremetal-services.yaml, but
  this is expected to eventually disappear.

.. releasenotes/notes/ffu-custom-script-to-switch-repo-a65db91760b46ec2.yaml @ 2587cb4acf44d9b64ff2cabab8b056f66812d45b

- The user can now use a custom script to switch repo during the
  fast forward upgrade.  He/She has to set ``FastForwardRepoType``
  to ``custom-script`` and set
  ``FastForwardCustomRepoScriptContent`` to a string representing a
  shell script.  That script will be executed on each node and given
  the upstream name of the release as the first argument (ocata,
  pike, queens in that order).  Here is an example that describes
  its interface.
  
  .. code-block:: bash
  
      #!/bin/bash
      case $1 in
        ocata)
          curl -o /etc/yum.repos.d/ocata.repo http://somewhere.com/my-Ocata.repo;
          yum clean metadata;
        pike)
          curl -o /etc/yum.repos.d/pike.repo http://somewhere.com/my-Pike.repo;
          yum clean metadata;
        queens)
          curl -o /etc/yum.repos.d/pike.repo http://somewhere.com/my-Queens.repo;
          yum clean metadata;
        *)
          echo "unknown release $1" >&2
          exit 1
      esac

.. releasenotes/notes/ipxe_timeout-5824c87e849b1b50.yaml @ 931067f9006fac97efd0493a79b33db168109054

- A new parameter IronicIPXETimeout can change the default iPXE timeout, set to
  60 seconds. Note that 0 would set an infinite timeout.

.. releasenotes/notes/ironic-inspector-use-dnsmasq_ip_subnets-abba77307e761b96.yaml @ c60489ecd989c07e27288b1973bd7df49304634d

- Adds support to configure ironic-inspector with multiple ip ranges. This enables ironic-inspector's DHCP server to serve request that came in via dhcp-relay agent.

.. releasenotes/notes/ironic-networking-baremetal-29d9ad465565bb87.yaml @ 5203e4397905d9d62ab2487b2fc5873937d8db42

- Adds support for Ironic Networking Baremetal. Networking Baremetal is used to integrate the Bare Metal service with the Networking service.

.. releasenotes/notes/ironic-rescue-cb1edecce357fc0b.yaml @ 3464547983c893292d1f3821aadf3e07e967f260

- Rescue mode is now enabled by default in ironic. To disable it, set
  ``IronicDefaultRescueInterface`` to ``no-rescue``.

.. releasenotes/notes/kernel_sysctl_role-d4f6a50d08b7a388.yaml @ da1ed3d19c54b5f60fdc883fd4f11f613914c854

- Allow to configure extra Kernel modules and extra sysctl settings per role
  and not only global to the whole deployment.
  The two parameters that can be role-specific are ExtraKernelModules and
  ExtraSysctlSettings.

.. releasenotes/notes/l2gw-driver-change-1f5b11d5676c5015.yaml @ add4ce1ccabd4b3e52e9fc89772af8ce7c29ed00

- L2GW driver changes to version 2 when using OpenDaylight.

.. releasenotes/notes/mistral_execs-5e1c363c9293504d.yaml @ a360759fd249dd50e62e9f5aad97c17950de54b3

- MistralEvaluationInterval is a new parameter that allow to configure
  how often will the Mistral Executions be evaluated.
  For example for value 120 the interval will be 2 hours (every 2 hours).

.. releasenotes/notes/mistral_execs-5e1c363c9293504d.yaml @ a360759fd249dd50e62e9f5aad97c17950de54b3

- MistralFinishedExecutionDuration is a new parameter that allow to configure
  how Mistral will evaluate from which time remove executions in minutes.
  For example when set to 60, remove all executions that finished a 60 minutes
  ago or more.
  Note that only final state execution will remove (SUCCESS/ERROR).

.. releasenotes/notes/ovs-dpdk-permissions-50c5b33334ff4711.yaml @ 825bd7d9e13ccbdf552e3bb01718d043d83487c6

- Till now, the ovs service file and ovs-ctl command files are patched to allow ovs to run with qemu group. In order to remove this workarounds, a new group hugetlbfs is created which will be shared between ovs and qemu. Vhostuser Socket Directory is changed from "/var/run/openvswitch" to "/var/lib/vhost_sockets" to avoid modifying the directory access by packaged scripts. Use env file ovs-dpdk-permissions.yaml while deploying.

.. releasenotes/notes/update_odl-cb997ce5c136ebb7.yaml @ 98faacad44e39a456d9fe1a1d21f5a65e8de4fc1

- Minor update ODL steps are added. ODL minor update (within same ODL release) can have 2 different workflow. These are called level 1 and level2. Level 1 is simple - stop, update and start ODL. Level 2 is complex and involved yang model changes. This requires wiping of DB and resync to repopulate the data. Steps involved in level 2 update are 1. Block OVS instances to connect to ODL 2. Set ODL upgrade flag to True 3. Start ODL 4. Start Neutron re-sync and wait for it to finish 5. Delete OVS groups and ports 6. Stop OVS 7. Unblock OVS ports 8. Start OVS 9. Unset ODL upgrade flag To achieve L2 update, use "-e environments/services-docker/ update-odl.yaml" along with other env files to the update command.

.. releasenotes/notes/vnc_tls-b3707d0134697cc7.yaml @ 37a339d2b0f0282bf1bac96587b10ca61868cec5

- If TLS on the internal network is enabled, the nova-novnc to libvirt vnc
  transport defaults to using TLS. This can be changed by setting the
  ``UseTLSTransportForVnc`` parameter, which is ``true`` by default.
  A dedicated IPA sub-CA can be specified by the ``LibvirtVncCACert``
  parameter. By default the main IPA CA will be used.

.. releasenotes/notes/xtremio_cinder_c5572898724a11e7.yaml @ a462d796a7a1efe17c399e81395a22610b016952

- Add support for Dell EMC XTREMIO ISCSI cinder driver


.. _tripleo-heat-templates_9.0.0.0b2_Upgrade Notes:

Upgrade Notes
-------------

.. releasenotes/notes/containers-as-default-37bbe8afa0a60c2b.yaml @ 6c5f2b8f69d42b2201f29264360ca15a518f376b

- Environment files originally referenced from `environments/services-docker`
  should be altered to the `environments/services` paths. If some of the
  deployed baremetal services need to be retained as non-containerized,
  update its references to `environments/services-baremetal` instead of
  `environments/services`.
  
  .. note:: Starting from Rocky, overcloud upgrades to baremetal services
    (non-containerized), or mixed services is no more tested nor verified.

.. releasenotes/notes/hiera_net_ip_map-ff866b443a28bdc4.yaml @ 3a7baa8fa6fa8dd6735f38d6236e8a2cb5d34659

- Per-service config_settings should now use hiera interpolation to set
  the bind IP for services, e.g "%{hiera('internal_api')}" whereas prior
  to this release we replaced e.g internal_api for the IP address internally.
  The network name can still be derived from the ServiceNetMap - all the
  in-tree templates have been converted to the new format, but any out
  of tree templates may require similar adjustment.

.. releasenotes/notes/mod_ssl-e7fd4db71189242e.yaml @ 628da8a37e544ae2783a7b2114c929e35c779003

- When a service is deployed in WSGI with Apache, make sure mode_ssl package is deployed during the upgrade process, it's now required by default so Apache can start properly.

.. releasenotes/notes/neutron_db_rename-bbfbce1c58cadc84.yaml @ d86025593be16c136a2f2104dc91f37fe7bca99f

- When the undercloud was not containerized, the neutron database name was called neutron.
  When we upgrade to a containerized undercloud, the database name is called ovs_neutron.

.. releasenotes/notes/pre_upgrade_rolling_tasks-6345e98e8283a907.yaml @ ae085825e22cb4ce7bf877087c2e324b8bec1f03

- pre_upgrade_rolling_tasks are added for use by the composable
  service templates. The resulting
  pre_upgrade_rolling_steps_playbook is intended to be run at the
  beginning of major update workflow (before running the
  upgrade_steps_playbook). As the name suggests, the tasks in this
  playbook will be executed in a node-by-node rolling fashion.

.. releasenotes/notes/remove_disable_upgrade_deployment_flag-872df40d7ff171b8.yaml @ 66df6bdb4666dbacd4a10d4433aede7bfa937018

- The disable_upgrade_deployment flag is now completely removed from
  the roles_data. It will have no effect if you continue to include
  this flag. It has not been used since the Pike upgrade. In Queens
  the upgrade workflow is delivered with ansible playbooks.

.. releasenotes/notes/zaqar-use-redis-by-default-930f542dda895a31.yaml @ e290824ce3cda176f9f71429d768369f715671f1

- Zaqar has been switched to use the redis backend by default from the
  mongodb backend. Mongodb has not been supported by TripleO since Pike.


.. _tripleo-heat-templates_9.0.0.0b2_Deprecation Notes:

Deprecation Notes
-----------------

.. releasenotes/notes/Deprecate-cinder-API-nova_catalog_admin_info-006ebda240f730a2.yaml @ 4268c8829cffd95cbfc23f9def4d30cc846a524a

- The nova_catalog_admin_info parameter is no longer being configured for
  cinder since it was deprecated.

.. releasenotes/notes/fix-odl-ovs-vhostusermode-7bc2b64fd2676ca2.yaml @ 2ecf3ac88bcdbab60b7f292ce635915c99acebc8

- Using 'client' for OvsVhostuserMode parameter. See 'vhost-user' section
  at http://docs.openvswitch.org/en/latest/topics/dpdk/vhost-user/

.. releasenotes/notes/ironic-inspector-use-dnsmasq_ip_subnets-abba77307e761b96.yaml @ c60489ecd989c07e27288b1973bd7df49304634d

- The parameter ``IronicInspectorIpRange`` is deprecated. Use the new ``IronicInspectorSubnets`` instead.

.. releasenotes/notes/remove-odl-dlux-gui-4728de06c973cd53.yaml @ f51f5336798751d9a23b69c4c301f748251a6064

- odl-dlux-all feature for OpenDaylight is no longer supported and removed
  from default installed OpenDaylightFeatures. See
  https://bugs.launchpad.net/tripleo/+bug/1751857


.. _tripleo-heat-templates_9.0.0.0b2_Security Issues:

Security Issues
---------------

.. releasenotes/notes/memcached_hardening-2529734099da27f4.yaml @ eaf77cb09c72fd1a9205c7a3266b99d6ce49d827

- Restrict memcached service to TCP and internal_api network (CVE-2018-1000115).


.. _tripleo-heat-templates_9.0.0.0b2_Bug Fixes:

Bug Fixes
---------

.. releasenotes/notes/add-cinder-backup-nfs-backend-0108fba91a3058ea.yaml @ e456e103fbfb5d7dd12f4dcbee4fd6686384117c

- Fixes `bug 1744174 <https://bugs.launchpad.net/tripleo/+bug/1744174>`__.

.. releasenotes/notes/fix-get-occ-config-with-role-count-greater-1-10ce2010556e5b76.yaml @ 154879b68dfe3ff3834babcfb1d619a4d05044af

- When using get-occ-config.sh with a role using a count greater than 1, the script will now configure all nodes that are of that role type instead of exiting after only configuring the first.

.. releasenotes/notes/fix-neutron-cert-key-perms.yaml-efcc17f188798cc4.yaml @ 16731819c5bb92d9bb33c6fb8086a6a776bdef8c

- Fixes Neutron certificate and key for TLS deployments to have the correct
  user/group IDs.

.. releasenotes/notes/fix-odl-gui-feature-6525b8c6807fb784.yaml @ e581c27e32202606e4580b6ba8b42c3bb9ae097e

- Fixes GUI feature loaded into OpenDaylight, which fixes the GUI as well
  as the URL used for Docker healthcheck.

.. releasenotes/notes/fix-odl-missing-etc-config-87c33bc05f692f44.yaml @ 97173caf8f68695140d207c2ab88226fd86659dc

- Fixes OpenDaylight container service not starting due to missing config
  files in /opt/opendaylight/etc directory.

.. releasenotes/notes/fix-odl-ovs-allowed-network-types-d196d6d40fadb1bc.yaml @ 186b03d72098858c9d8f1b0b014527ffe9864d72

- Fixes missing type "flat" from the default allowed network types for
  the ODL OVS parameter HostAllowedNetworkTypes.  See
  https://bugs.launchpad.net/tripleo/+bug/1762495

.. releasenotes/notes/fix-odl-ovs-vhostusermode-7bc2b64fd2676ca2.yaml @ 2ecf3ac88bcdbab60b7f292ce635915c99acebc8

- Fixes default of vhostuser_mode in ODL-OVS to be server, and clarifies
  the configuration parameter. See
  https://bugs.launchpad.net/tripleo/+bug/1762473

.. releasenotes/notes/fix-tls-neutron-agents-c40d5fc779d53bfa.yaml @ df31016a9af5003533f80989bcb8d3da42099953

- Fixes failure to create Neutron certificates for roles which do not
  contain Neutron DHCP agent, but include other Neutron agents
  (i.e. default Compute role).

.. releasenotes/notes/remove-pacemaker-passwords-default-values-dd0cfdf7922ecf90.yaml @ d57bd297a9470fedb15224e215733e01a198b286

- The default values for the PcsdPassword and PacemakerRemoteAuthkey parameters have been removed, as they did not result in a functioning pacemaker installation. These values are instead generated by tripleo-common, and in the cases where they are not (direct API), we want to fail explicitly if they are not provided.

.. releasenotes/notes/tripleo-ssh-known-hosts-5c64b1a90d61d7f2.yaml @ 088d5c12f0f37e24d836e6f8791f41fbeba3326d

- Add support for the SshKnownHostsDeployment resources to config-download. Since the deployment resources relied on Heat outputs, they were not supported with the default handling from tripleo-common that relies on the group_vars mechanism.  The templates have been refactored to add the known hosts entries as global_vars to deploy_steps_playbook.yaml, and then include the new tripleo-ssh-known-hosts role from tripleo-common to apply the same configuration that the Heat deployment did.

.. releasenotes/notes/use-role-name-ExtraConfig-with-deprecations-2688f34fbc6de74a.yaml @ fa4b3e2a3c634c8f1a18087e508d585a693aa84b

- ``{{role.name}}ExtraConfig`` will now be honored even when using deprecated
  params in roles_data.yaml. Previously, its value was ignored and never used
  even though it is defined as a valid parameter in the rendered template.


.. _tripleo-heat-templates_9.0.0.0b2_Other Notes:

Other Notes
-----------

.. releasenotes/notes/check-old-style-nic-config-4624a60e3303411b.yaml @ 0017b64560074032b09b18b349c2ba1d8b71196b

- Add check for nic config files using the old style format (os-apply-config) and list the script that can be used to convert the file.

