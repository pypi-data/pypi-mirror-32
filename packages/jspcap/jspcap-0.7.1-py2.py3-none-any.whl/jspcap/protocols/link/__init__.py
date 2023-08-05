# -*- coding: utf-8 -*-
"""link layer protocols

`jspcap.protocols.link` is collection of all protocols in
link layer, with detailed implementation and methods.

"""
# Base Class for Link Layer
from jspcap.protocols.link.link import Link

# Utility Classes for Protocols
from jspcap.protocols.link.arp import ARP
from jspcap.protocols.link.ethernet import Ethernet
from jspcap.protocols.link.l2tp import L2TP
from jspcap.protocols.link.ospf import OSPF
from jspcap.protocols.link.rarp import RARP
from jspcap.protocols.link.vlan import VLAN

# Link-Layer Header Type Values
from jspcap.protocols.link.link import LINKTYPE


__all__ = [
    'LINKTYPE',                                         # Protocol Numbers
    'ARP', 'Ethernet', 'L2TP', 'OSPF', 'RARP', 'VLAN',  # Link Layer Protocols
]
