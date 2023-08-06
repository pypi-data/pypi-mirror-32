from armulator.armv6.opcodes.abstract_opcode import AbstractOpcode
from armulator.armv6.arm_exceptions import SMCException, UndefinedInstructionException
from bitstring import BitArray
from armulator.armv6.configurations import have_security_ext, have_virt_ext


class Smc(AbstractOpcode):
    def __init__(self):
        super(Smc, self).__init__()

    def execute(self, processor):
        if processor.condition_passed():
            if have_security_ext() and processor.registers.current_mode_is_not_user():
                if (have_virt_ext() and not processor.registers.is_secure() and
                        not processor.registers.current_mode_is_hyp() and
                        processor.registers.hcr.get_tsc()):
                    hsr_string = BitArray(25)
                    processor.write_hsr("010011", hsr_string)
                    processor.registers.take_hyp_trap_exception()
                else:
                    if processor.registers.scr.get_scd():
                        if processor.registers.is_secure():
                            print "unpredictable"
                        else:
                            raise UndefinedInstructionException()
                    else:
                        raise SMCException()
            else:
                raise UndefinedInstructionException()
