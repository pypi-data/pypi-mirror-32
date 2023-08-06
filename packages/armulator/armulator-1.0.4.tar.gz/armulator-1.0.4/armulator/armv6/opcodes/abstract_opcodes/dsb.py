from armulator.armv6.opcodes.abstract_opcode import AbstractOpcode
from armulator.armv6.enums import MBReqDomain, MBReqTypes
from armulator.armv6.configurations import have_virt_ext


class Dsb(AbstractOpcode):
    def __init__(self, option):
        super(Dsb, self).__init__()
        self.option = option

    def execute(self, processor):
        if processor.condition_passed():
            if self.option == "0b0010":
                domain = MBReqDomain.MBReqDomain_OuterShareable
                types = MBReqTypes.MBReqTypes_Writes
            elif self.option == "0b0011":
                domain = MBReqDomain.MBReqDomain_OuterShareable
                types = MBReqTypes.MBReqTypes_All
            elif self.option == "0b0110":
                domain = MBReqDomain.MBReqDomain_Nonshareable
                types = MBReqTypes.MBReqTypes_Writes
            elif self.option == "0b0111":
                domain = MBReqDomain.MBReqDomain_Nonshareable
                types = MBReqTypes.MBReqTypes_All
            elif self.option == "0b1010":
                domain = MBReqDomain.MBReqDomain_InnerShareable
                types = MBReqTypes.MBReqTypes_Writes
            elif self.option == "0b1011":
                domain = MBReqDomain.MBReqDomain_InnerShareable
                types = MBReqTypes.MBReqTypes_All
            elif self.option == "0b1110":
                domain = MBReqDomain.MBReqDomain_FullSystem
                types = MBReqTypes.MBReqTypes_Writes
            else:
                domain = MBReqDomain.MBReqDomain_FullSystem
                types = MBReqTypes.MBReqTypes_All
            if (have_virt_ext() and
                    not processor.registers.is_secure() and
                    not processor.registers.current_mode_is_hyp()):
                if processor.registers.hcr.get_bsu() == "0b11":
                    domain = MBReqDomain.MBReqDomain_FullSystem
                if processor.registers.hcr.get_bsu() == "0b10" and domain != MBReqDomain.MBReqDomain_FullSystem:
                    domain = MBReqDomain.MBReqDomain_OuterShareable
                if processor.registers.hcr.get_bsu() == "0b01" and domain == MBReqDomain.MBReqDomain_Nonshareable:
                    domain = MBReqDomain.MBReqDomain_InnerShareable
            processor.data_synchronization_barrier(domain, types)
