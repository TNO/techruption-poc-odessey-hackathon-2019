
# Setup private variables (only callable from within the contract)
creator: address
gatekeeper_address: address

@public
def __init__():
    # Define owner as creator
    self.creator = msg.sender

@public
def set_gatekeeper_address(addr: address):
    # Future addition update: ask other contract whether sender is allowed to change stuff
    assert msg.sender == self.creator
    self.gatekeeper_address = addr

@public
@constant
def get_gatekeeper_address() -> address:
    return self.gatekeeper_address
    