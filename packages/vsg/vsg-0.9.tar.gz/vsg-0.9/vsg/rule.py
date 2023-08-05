

class rule():

    def __init__(self, name=None, identifier=None):
        self.name = name
        self.identifier = identifier
        self.solution = None
        self.violations = []
        self.indentSize = 2
        self.phase = None
        self.disable = False
        self.fixable = True
        self.dFix = {}
        self.dFix['violations'] = {}

    def configure(self, dConfiguration):
        '''Configures attributes on rules using a dictionary of the following form:

            dConfiguration['rule'] = {}
            dConfiguration['rule']['xyz_001'] = {}
            dConfiguration['rule']['xyz_001']['disable'] = True
            dConfiguration['rule']['xyz_001']['solution'] = 'This is the new solution'
            dConfiguration['rule']['xyz_002'] = {}
            dConfiguration['rule']['xyz_002']['disable'] = False
            dConfiguration['rule']['global'] = {}
            dConfiguration['rule']['global']['indentSize'] = 4

          The rule:global dictionary will apply to all rules.
          Individual rule attributes can be modified with [self.name_self.identifier].
        '''

        self._configure_global_rule_attributes(dConfiguration)
        self._configure_rule_attributes(dConfiguration)

    def report_violations(self, iLineNumber, fQuiet=False):
        for sViolation in self.violations:
            if str(sViolation).startswith(str(iLineNumber) + '-') or str(iLineNumber) == str(sViolation):
                if not fQuiet:
                    print('  ' + (self.name + '_' + self.identifier).ljust(25) + ' | ' + str(sViolation).rjust(10) + ' | ' + self.solution)  # pragma: no coverage
                return 1
        return 0

    def fix(self, oFile):
        if self.fixable:
            self.analyze(oFile)
            self._fix_violations(oFile)
            self.violations = []
            self.dFix = {}
            self.dFix['violations'] = {}

    def add_violation(self, lineNumber):
        if lineNumber not in self.violations:
            self.violations.append(lineNumber)

    def _configure_global_rule_attributes(self, dConfiguration):
        try:
            for sAttributeName in dConfiguration['rule']['global']:
                if sAttributeName in self.__dict__:
                    self.__dict__[sAttributeName] = dConfiguration['rule']['global'][sAttributeName]
        except KeyError:
            pass

    def _configure_rule_attributes(self, dConfiguration):
        for sAttributeName in dConfiguration['rule'][self.name + '_' + self.identifier]:
            if sAttributeName in self.__dict__:
                self.__dict__[sAttributeName] = dConfiguration['rule'][self.name + '_' + self.identifier][sAttributeName]
