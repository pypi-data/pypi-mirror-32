#   Copyright 2018 Red Hat, Inc.
#
#   Licensed under the Apache License, Version 2.0 (the "License"); you may
#   not use this file except in compliance with the License. You may obtain
#   a copy of the License at
#
#        http://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
#   WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
#   License for the specific language governing permissions and limitations
#   under the License.
#


from oslo_config import cfg
from tripleoclient.config.base import BaseConfig

NETCONFIG_TAGS_EXAMPLE = """
"network_config": [
 {
  "type": "ovs_bridge",
  "name": "br-ctlplane",
  "ovs_extra": [
   "br-set-external-id br-ctlplane bridge-id br-ctlplane"
  ],
  "members": [
   {
    "type": "interface",
    "name": "{{LOCAL_INTERFACE}}",
    "primary": "true",
    "mtu": {{LOCAL_MTU}},
    "dns_servers": {{UNDERCLOUD_NAMESERVERS}}
   }
  ],
  "addresses": [
    {
      "ip_netmask": "{{PUBLIC_INTERFACE_IP}}"
    }
  ],
  "routes": {{SUBNETS_STATIC_ROUTES}},
  "mtu": {{LOCAL_MTU}}
}
]
"""


class StandaloneConfig(BaseConfig):

    def get_enable_service_opts(self, cinder=False, ironic=False,
                                ironic_inspector=False, mistral=False,
                                novajoin=False, tempest=False,
                                telemetry=False, tripleo_ui=False,
                                validations=False, zaqar=False):
        _opts = [
            # service enablement
            cfg.BoolOpt('enable_cinder',
                        default=cinder,
                        help=(
                            'Whether to install the Volume service. It is not '
                            'currently used in the undercloud.')),
            cfg.BoolOpt('enable_ironic',
                        default=ironic,
                        help=('Whether to enable the ironic service.')),
            cfg.BoolOpt('enable_ironic_inspector',
                        default=ironic_inspector,
                        help=(
                            'Whether to enable the ironic inspector service.')
                        ),
            cfg.BoolOpt('enable_mistral',
                        default=mistral,
                        help=('Whether to enable the mistral service.')),
            cfg.BoolOpt('enable_novajoin',
                        default=novajoin,
                        help=('Whether to install novajoin metadata service '
                              'in the Undercloud.')
                        ),
            cfg.BoolOpt('enable_tempest',
                        default=tempest,
                        help=('Whether to install Tempest in the Undercloud.'
                              'This is a no-op for containerized undercloud.')
                        ),
            cfg.BoolOpt('enable_telemetry',
                        default=telemetry,
                        help=('Whether to install Telemetry services '
                              '(ceilometer, gnocchi, aodh, panko ) in the '
                              'Undercloud.')
                        ),
            cfg.BoolOpt('enable_ui',
                        default=tripleo_ui,
                        help=('Whether to install the TripleO UI.')
                        ),
            cfg.BoolOpt('enable_validations',
                        default=validations,
                        help=(
                            'Whether to install requirements to run the '
                            'TripleO validations.')
                        ),
            cfg.BoolOpt('enable_zaqar',
                        default=zaqar,
                        help=('Whether to enable the zaqar service.')),
        ]
        return self.sort_opts(_opts)

    def get_base_opts(self):
        _base_opts = super(StandaloneConfig, self).get_base_opts()
        _opts = [
            # deployment options
            cfg.StrOpt('deployment_user',
                       help=(
                           'User used to run openstack undercloud install '
                           'command which will be used to add the user to the '
                           'docker group, required to upload containers'),
                       ),
            cfg.StrOpt('hieradata_override',
                       default='',
                       help=(
                           'Path to hieradata override file. Relative paths '
                           'get computed inside of $HOME. When it points to a '
                           'heat env file, it is passed in t-h-t via "-e '
                           '<file>", as is. When the file contains legacy '
                           'instack data, it is wrapped with '
                           'UndercloudExtraConfig and also passed in for '
                           't-h-t as a temp file created in output_dir. Note, '
                           'instack hiera data may be not t-h-t compatible '
                           'and will highly likely require a manual revision.')
                       ),
            cfg.StrOpt('net_config_override',
                       default='',
                       help=(
                           'Path to network config override template.'
                           'Relative paths get computed inside of the '
                           'given heat templates directory. Must be in '
                           'json format.'
                           'Its content overrides anything in t-h-t '
                           'UndercloudNetConfigOverride. The processed '
                           'template is then passed in Heat via the top '
                           'scope undercloud_parameters.yaml file created in '
                           'output_dir and used to configure the networking '
                           'via run-os-net-config. If you wish to disable '
                           'you can set this location to an empty file.'
                           'Templated for instack j2 tags '
                           'may be used, '
                           'for example:\n%s ') % NETCONFIG_TAGS_EXAMPLE
                       ),
            cfg.StrOpt('templates',
                       default='',
                       help=('heat templates file to override.')
                       ),
            cfg.StrOpt('roles_file',
                       default=None,
                       help=('Roles file to override for heat. '
                             'The file path is related to the templates path')
                       ),
            cfg.BoolOpt('heat_native',
                        default=True,
                        help=('Use native heat templates.')),
            cfg.StrOpt('heat_container_image',
                       default='',
                       help=('URL for the heat container image to use.')
                       ),
            cfg.StrOpt('container_images_file',
                       default='',
                       help=(
                           'Heat environment file with parameters for all '
                           'required container images. Or alternatively, '
                           'parameter "ContainerImagePrepare" to drive the '
                           'required image preparation.')),
            cfg.ListOpt('custom_env_files',
                        default=[],
                        help=('List of any custom environment yaml files to '
                              'use')),
            # docker config bits
            cfg.StrOpt('docker_registry_mirror',
                       default='',
                       help=(
                           'An optional docker \'registry-mirror\' that will '
                           'beconfigured in /etc/docker/daemon.json.')
                       ),
            cfg.ListOpt('docker_insecure_registries',
                        default=[],
                        help=('Used to add custom insecure registries in '
                              '/etc/sysconfig/docker.')
                        ),
        ]
        return self.sort_opts(_base_opts + _opts)

    def get_opts(self):
        return self.sort_opts(self.get_base_opts() +
                              self.get_enable_service_opts())
