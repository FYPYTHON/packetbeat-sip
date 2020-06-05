from packetbeat import BaseTest


class Test(BaseTest):
    def test_sip(self):
        self.render_config_template()
        self.run_packetbeat(pcap="sip.pcap",
                            debug_selectors=["tls"])
        objs = self.read_output()
        print(objs[0])


if __name__ == "__main__":
    a = Test()
    a.test_sip()
