from nuaal.utils import get_logger
from nuaal.connections.cli import Cisco_IOS_Cli
import queue
import threading
import time

class CliMultiRunner(object):
    """
    This class allows running set of CLI commands on multiple devices in parallel, using Worker threads
    """
    def __init__(self, provider, ips, actions=None, workers=4, DEBUG=False):
        """

        :param dict provider: Dictionary with necessary info for creating connection
        :param list ips: List of IP addresses of the device
        :param list actions: List of actions to be run in each connection
        :param int workers: Number of worker threads to spawn
        :param bool DEBUG: Enables/disables debugging output
        """
        self.provider = provider
        self.ips = ips
        self.workers = workers
        self.DEBUG = DEBUG
        self.logger = get_logger(name="CliMultiRunner", DEBUG=self.DEBUG)
        self.queue = queue.Queue()
        self.threads = []
        self.actions = actions if isinstance(actions, list) else []
        self.data = []

    def fill_queue(self):
        """
        This function builds *device-specific* providers for CliConnection and adds them to Queue object.

        :return: ``None``
        """
        for i in range(len(self.ips)):
            provider = dict(self.provider)
            provider["ip"] = self.ips[i]
            self.queue.put(provider)

    def worker(self):
        """
        Worker function to handle individual connections. Based on provider object from Queue establishes connection to device and runs
        defined set of commands. Received data are stored in ``self.data`` as a list of dictionaries, each dictionary representing one device.

        :return: ``None``
        """
        self.logger.info(msg="Spawned new worker in {}".format(threading.current_thread()))
        while not self.queue.empty():
            provider = self.queue.get()
            if provider is None:
                break
            try:
                with Cisco_IOS_Cli(**provider) as device:
                    if "get_vlans" in self.actions:
                        device.get_vlans()
                    if "get_neighbors" in self.actions:
                        device.get_neighbors()
                    if "get_interfaces" in self.actions:
                        device.get_interfaces()
                    if "get_trunks" in self.actions:
                        device.get_trunks()
                    if "get_portchannels" in self.actions:
                        device.get_portchannels()
                    if "get_version" in self.actions:
                        device.get_version()
                    if "get_license" in self.actions:
                        device.get_license()
                    if "get_inventory" in self.actions:
                        device.get_inventory()
                    self.data.append(device.data)
            except Exception as e:
                self.logger.error(msg="Unhandled Exception occurred in thread '{}'. Exception: {}".format(threading.current_thread().getName(), repr(e)))
            self.queue.task_done()

    def thread_factory(self):
        """
        Function for spawning worker threads based on number of workers in ``self.workers``

        :return:
        """
        for i in range(self.workers):
            t = threading.Thread(name="WorkerThread-{}".format(i), target=self.worker, args=())
            self.threads.append(t)

    def run(self):
        """
        Main entry function, puts all pieces together. Firs populates Queue with IP addreses and connection information, spawns threads and starts them.
        Blocks until ``self.queue`` is empty.

        :return:
        """
        self.fill_queue()
        self.thread_factory()
        [t.start() for t in self.threads]
        self.queue.join()
