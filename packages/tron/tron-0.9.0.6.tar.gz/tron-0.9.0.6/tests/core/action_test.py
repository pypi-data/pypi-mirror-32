from __future__ import absolute_import
from __future__ import unicode_literals

import mock
from testify import assert_equal
from testify import run
from testify import setup
from testify import TestCase

from tron import node
from tron.core import action


class TestAction(TestCase):
    @setup
    def setup_action(self):
        self.node_pool = mock.create_autospec(node.NodePool)
        self.action = action.Action("my_action", "doit", self.node_pool)

    def test_from_config(self):
        config = mock.Mock(
            name="ted",
            command="do something",
            node="first",
            executor="ssh",
            cluster="prod",
            pool="default",
            cpus=1,
            mem=100,
            service="bar",
            deploy_group="test",
        )
        new_action = action.Action.from_config(config)
        assert_equal(new_action.name, config.name)
        assert_equal(new_action.command, config.command)
        assert_equal(new_action.executor, config.executor)
        assert_equal(new_action.cluster, config.cluster)
        assert_equal(new_action.pool, config.pool)
        assert_equal(new_action.cpus, config.cpus)
        assert_equal(new_action.mem, config.mem)
        assert_equal(new_action.service, config.service)
        assert_equal(new_action.deploy_group, config.deploy_group)
        assert_equal(new_action.node_pool, None)
        assert_equal(new_action.required_actions, [])

    def test__eq__(self):
        new_action = action.Action(
            self.action.name,
            self.action.command,
            self.node_pool,
        )
        assert_equal(new_action, self.action)


if __name__ == '__main__':
    run()
