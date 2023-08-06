from armulator.armv6.opcodes.abstract_opcode import AbstractOpcode
from bitstring import BitArray
from armulator.armv6.bits_ops import zeros
from armulator.armv6.configurations import have_virt_ext, jazelle_accepts_execution
from armulator.armv6.enums import InstrSet


class Bxj(AbstractOpcode):
    def __init__(self, m):
        super(Bxj, self).__init__()
        self.m = m

    def execute(self, processor):
        if processor.condition_passed():
            if (have_virt_ext() and not processor.registers.is_secure() and
                    not processor.registers.current_mode_is_hyp() and
                    processor.registers.hstr.get_tjdbx()):
                hsr_string = zeros(25)
                hsr_string[-4:] = self.m
                processor.write_hsr(BitArray(bin="001010"), hsr_string)
                processor.registers.take_hyp_trap_exception()
            elif (not processor.registers.jmcr.get_je() or
                  processor.registers.current_instr_set() == InstrSet.InstrSet_ThumbEE):
                processor.bx_write_pc(processor.registers.get(self.m))
            else:
                if jazelle_accepts_execution():
                    processor.switch_to_jazelle_execution()
                else:
                    # SUBARCHITECTURE_DEFINED handler call
                    raise NotImplementedError()
