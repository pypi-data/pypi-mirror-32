import pip.utils as utils
import pytest

from plenum.test.helper import randomText, mockGetInstalledDistributions, \
    mockImportModule
from functools import partial
import importlib

nodeCount = 1


def testPluginManagerFindsPlugins(monkeypatch, pluginManager):
    validPackagesCnt = 5
    invalidPackagesCnt = 10
    validPackages = [pluginManager.prefix + randomText(10)
                     for _ in range(validPackagesCnt)]
    invalidPackages = [randomText(10) for _ in range(invalidPackagesCnt)]

    monkeypatch.setattr(
        utils,
        'get_installed_distributions',
        partial(
            mockGetInstalledDistributions,
            packages=validPackages +
                     invalidPackages))
    assert len(pluginManager._findPlugins()) == validPackagesCnt


def testPluginManagerImportsPlugins(monkeypatch, pluginManager):
    packagesCnt = 3
    packages = [pluginManager.prefix + randomText(10)
                for _ in range(packagesCnt)]

    monkeypatch.setattr(utils, 'get_installed_distributions',
                        partial(mockGetInstalledDistributions,
                                packages=packages))
    monkeypatch.setattr(importlib, 'import_module', mockImportModule)

    importedPlugins, foundPlugins = pluginManager.importPlugins()
    assert importedPlugins == foundPlugins and importedPlugins == packagesCnt


def testPluginManagerSendsMessage(pluginManagerWithImportedModules):
    topic = randomText(10)
    message = randomText(20)
    sent, pluginCnt = pluginManagerWithImportedModules._sendMessage(
        topic, message)
    assert sent == 3


def testPluginManagerSendMessageUponSuspiciousSpikeFailsOnMinCnt(
        pluginManagerWithImportedModules):
    topic = randomText(10)
    name = randomText(10)
    historicalData = {
        'value': 0,
        'cnt': 0
    }
    newVal = 10
    config = {
        'borders_coeff': 2,
        'min_cnt': 10,
        'min_activity_threshold': 0,
        'use_weighted_borders_coeff': True,
        'enabled': True
    }
    assert pluginManagerWithImportedModules \
               .sendMessageUponSuspiciousSpike(topic, historicalData,
                                               newVal, config, name, enabled=True) is None


def testPluginManagerSendMessageUponSuspiciousSpikeFailsOnBordersCoefficient(
        pluginManagerWithImportedModules):
    topic = randomText(10)
    name = randomText(10)
    historicalData = {
        'value': 10,
        'cnt': 10
    }
    newVal = 15
    config = {
        'borders_coeff': 2,
        'min_cnt': 10,
        'min_activity_threshold': 0,
        'use_weighted_borders_coeff': True,
        'enabled': True
    }
    assert pluginManagerWithImportedModules \
               .sendMessageUponSuspiciousSpike(topic, historicalData,
                                               newVal, config, name, enabled=True) is None


def testPluginManagerSendMessageUponSuspiciousSpike(
        pluginManagerWithImportedModules):
    topic = randomText(10)
    name = randomText(10)
    historicalData = {
        'value': 10,
        'cnt': 10
    }
    newVal = 25
    config = {
        'borders_coeff': 2,
        'min_cnt': 10,
        'min_activity_threshold': 0,
        'use_weighted_borders_coeff': True,
        'enabled': True
    }
    sent, found = pluginManagerWithImportedModules \
        .sendMessageUponSuspiciousSpike(topic, historicalData,
                                        newVal, config, name, enabled=True)
    assert sent == 3


def testNodeSendNodeRequestSpike(pluginManagerWithImportedModules, testNode):
    def mockProcessRequest(obj, inc=1):
        obj.nodeRequestSpikeMonitorData['accum'] += inc

    N = 15
    testNode.config.SpikeEventsEnabled = True
    testNode.config.notifierEventTriggeringConfig['nodeRequestSpike'] = {
        'borders_coeff': 10,
        'min_cnt': N,
        'freq': 60,
        'min_activity_threshold': 0,
        'use_weighted_borders_coeff': True,
        'enabled': True
    }

    # The learning period is necessary
    for _ in range(0, N):
        mockProcessRequest(testNode)
        assert testNode.sendNodeRequestSpike() is None
    mockProcessRequest(testNode, 2)
    assert testNode.sendNodeRequestSpike() is None
    mockProcessRequest(testNode, 8)
    assert testNode.sendNodeRequestSpike() is None
    mockProcessRequest(testNode, 100)
    sent, found = testNode.sendNodeRequestSpike()
    assert sent == 3


def testMonitorSendClusterThroughputSpike(pluginManagerWithImportedModules,
                                          testNode):
    N = 15
    testNode.monitor.notifierEventTriggeringConfig['clusterThroughputSpike'] = {
        'borders_coeff': 10,
        'min_cnt': N,
        'freq': 60,
        'min_activity_threshold': 0,
        'use_weighted_borders_coeff': True,
        'enabled': True}

    # The learning period is necessary
    for _ in range(0, N):
        testNode.monitor.clusterThroughputSpikeMonitorData['accum'] = [1]
        assert testNode.monitor.sendClusterThroughputSpike() is None
    testNode.monitor.clusterThroughputSpikeMonitorData['accum'] = [2]
    assert testNode.monitor.sendClusterThroughputSpike() is None
    testNode.monitor.clusterThroughputSpikeMonitorData['accum'] = [7, 9]
    assert testNode.monitor.sendClusterThroughputSpike() is None
    testNode.monitor.clusterThroughputSpikeMonitorData['accum'] = [100]
    sent, found = testNode.monitor.sendClusterThroughputSpike()
    assert sent == 3


def test_suspicious_spike_check_disabled_config(pluginManagerWithImportedModules):
    topic = randomText(10)
    name = randomText(10)
    hdata = {'value': 10, 'cnt': 10}
    nval = 25
    config = {'borders_coeff': 2,
              'min_cnt': 10,
              'min_activity_threshold': 2,
              'use_weighted_borders_coeff': True,
              'enabled': True}

    sent, _ = pluginManagerWithImportedModules.sendMessageUponSuspiciousSpike(topic, hdata, nval,
                                                                              config, name, enabled=True)
    assert sent == 3

    config['enabled'] = False
    hdata['value'] = 10
    hdata['cnt'] = 10
    assert pluginManagerWithImportedModules.sendMessageUponSuspiciousSpike(topic, hdata, nval,
                                                                           config, name, enabled=True) is None


def test_suspicious_spike_check_disabled_func(pluginManagerWithImportedModules):
    topic = randomText(10)
    name = randomText(10)
    hdata = {'value': 10, 'cnt': 10}
    nval = 25
    config = {'borders_coeff': 2,
              'min_cnt': 10,
              'min_activity_threshold': 2,
              'use_weighted_borders_coeff': True,
              'enabled': True}

    sent, _ = pluginManagerWithImportedModules.sendMessageUponSuspiciousSpike(topic, hdata, nval,
                                                                              config, name, enabled=True)
    assert sent == 3

    hdata['value'] = 10
    hdata['cnt'] = 10
    assert pluginManagerWithImportedModules.sendMessageUponSuspiciousSpike(topic, hdata, nval,
                                                                           config, name, enabled=False) is None


def test_suspicious_spike_check_weighted_borders(pluginManagerWithImportedModules):
    topic = randomText(10)
    name = randomText(10)
    hdata = {'value': 100, 'cnt': 100}
    nval = 700
    config = {'borders_coeff': 10,
              'min_cnt': 10,
              'min_activity_threshold': 10,
              'use_weighted_borders_coeff': True,
              'enabled': True}

    sent, _ = pluginManagerWithImportedModules.sendMessageUponSuspiciousSpike(topic, hdata, nval,
                                                                              config, name, enabled=True)
    assert sent == 3

    hdata['value'] = 100
    hdata['cnt'] = 100
    config['use_weighted_borders_coeff'] = False
    assert pluginManagerWithImportedModules.sendMessageUponSuspiciousSpike(topic, hdata, nval,
                                                                           config, name, enabled=True) is None


def test_no_message_from_0_to_1(pluginManagerWithImportedModules):
    topic = randomText(10)
    name = randomText(10)
    hdata = {'value': 0, 'cnt': 10}
    nval = 1
    config = {'borders_coeff': 2,
              'min_cnt': 10,
              'min_activity_threshold': 0,
              'use_weighted_borders_coeff': True,
              'enabled': True}

    sent, _ = pluginManagerWithImportedModules.sendMessageUponSuspiciousSpike(topic, hdata, nval,
                                                                              config, name, enabled=True)
    assert sent == 3

    hdata = {'value': 0, 'cnt': 10}
    config = {'borders_coeff': 2,
              'min_cnt': 10,
              'min_activity_threshold': 2,
              'use_weighted_borders_coeff': True,
              'enabled': True}
    assert pluginManagerWithImportedModules.sendMessageUponSuspiciousSpike(topic, hdata, nval,
                                                                           config, name, enabled=True) is None


def test_no_message_from_1_to_0(pluginManagerWithImportedModules):
    topic = randomText(10)
    name = randomText(10)
    hdata = {'value': 1, 'cnt': 10}
    nval = 0
    config = {'borders_coeff': 2,
              'min_cnt': 10,
              'min_activity_threshold': 0,
              'use_weighted_borders_coeff': True,
              'enabled': True}

    sent, _ = pluginManagerWithImportedModules.sendMessageUponSuspiciousSpike(topic, hdata, nval,
                                                                              config, name, enabled=True)
    assert sent == 3

    hdata = {'value': 1, 'cnt': 10}
    config = {'borders_coeff': 2,
              'min_cnt': 10,
              'min_activity_threshold': 2,
              'use_weighted_borders_coeff': True,
              'enabled': True}
    assert pluginManagerWithImportedModules.sendMessageUponSuspiciousSpike(topic, hdata, nval,
                                                                           config, name, enabled=True) is None
