# -*- coding: utf-8 -*-
"""Casting utility: cast or punch ribbon, cast material for hand typesetting,
make a diecase proof, quickly compose and cast text.
"""

from collections import deque, OrderedDict
from functools import wraps

# common definitions
import librpi2caster
# QR code generating backend
try:
    import qrcode
except ImportError:
    qrcode = None

from .rpi2caster import UI, Abort, Finish, option, find_casters
from . import basic_models as bm, basic_controllers as bc, definitions as d
from . import main_controllers as mc
from .parsing import parse_record
from .typesetting import TypesettingContext


def cast_this(ribbon_source):
    """Get the ribbon from decorated routine and cast it"""
    @wraps(ribbon_source)
    def wrapper(self, *args, **kwargs):
        """Wrapper function"""
        ribbon = ribbon_source(self, *args, **kwargs)
        if not ribbon:
            return
        if self.machine.is_casting():
            return self.cast_ribbon(ribbon)
        else:
            return self.punch_ribbon(ribbon)
    return wrapper


class Casting(TypesettingContext):
    """Casting:

    Methods related to operating the composition caster.
    Requires configured caster.

    All methods related to operating a composition caster are here:
    -casting composition and sorts, punching composition,
    -calibrating the caster,
    -testing the interface,
    -sending an arbitrary combination of signals,
    -casting spaces to heat up the mould."""
    machine = None

    def __init__(self, interface_id, operation_mode=None):
        self.choose_machine(interface_id, operation_mode)

    def choose_machine(self, interface_id=None, operation_mode=None):
        """Choose a machine from the available interfaces."""
        def make_menu_entry(number, caster, url, name):
            """build a menu entry"""
            if caster:
                if number:
                    nums.append(number)
                modes = ', '.join(caster.supported_operation_modes)
                row16_modes = ', '.join(caster.supported_row16_modes) or 'none'
                text = ('{} - modes: {} - row 16 addressing modes: {}'
                        .format(caster, modes, row16_modes))
            else:
                text = '[Unavailable] {}\n\t\t\t{}'.format(url, name)
            return option(key=number, seq=number, value=caster, text=text)

        casters = find_casters(operation_mode)
        # look a caster up by the index
        try:
            url, caster, name = casters[interface_id]
            if not caster:
                raise ValueError
            UI.display('Using caster {} at {}'.format(name, url),
                       min_verbosity=1)
        except (KeyError, IndexError, TypeError, ValueError):
            # choose caster from menu
            # nums stores menu entry numbers for useful casters
            nums = []
            menu_options = [make_menu_entry(number, caster, url, name)
                            for number, (url, caster, name) in casters.items()]
            while True:
                caster = UI.simple_menu('Choose the caster:', menu_options,
                                        default_key=nums[0] if nums else 0)
                if not caster:
                    UI.pause('Tried to use the unavailable caster.')
                else:
                    break
        self.machine = caster

    def punch_ribbon(self, ribbon):
        """Punch the ribbon from start to end"""
        self.machine.punch(ribbon)

    def cast_ribbon(self, input_iterable):
        """Main casting routine.

        First check if ribbon needs rewinding (when it starts with pump stop).

        Ask user about number of repetitions (useful for e.g. business cards
        or souvenir lines), number of initial lines skipped for all runs
        and the upcoming run only.

        Ask user whether to pre-heat the mould
        to stabilize the temperature and cast good quality type.

        Cast the ribbon single or multiple times, displaying the statistics
        about current record and casting progress.

        If casting multiple runs, repeat until all are done,
        offer adding some more.
        """
        def rewind_if_needed():
            """Decide whether to rewind the ribbon or not.
            If casting and stop comes first, rewind. Otherwise not."""
            nonlocal ribbon
            for record in ribbon:
                if record.is_pump_stop:
                    # rewind
                    ribbon = [x for x in reversed(ribbon)]
                    return
                elif record.is_pump_start:
                    # no need to rewind
                    return

        def skip_lines(source):
            """Skip a definite number of lines"""
            # Apply constraints: 0 <= lines_skipped < lines in ribbon
            lines_skipped = stats.get_run_lines_skipped()
            if lines_skipped:
                UI.display('Skipping {} lines'.format(lines_skipped))
            # Take away combinations until we skip the desired number of lines
            # BEWARE: ribbon starts with galley trip!
            # We must give it back after lines are taken away
            record = parse_record('')
            sequence = deque(source)
            while lines_skipped > 0:
                record = sequence.popleft()
                lines_skipped -= 1 * record.is_newline
            # give the last code back
            sequence.appendleft(record)
            return sequence

        def set_lines_skipped(run=True, session=True):
            """Set the number of lines skipped for run and session
            (in case of multi-session runs)."""
            # allow skipping only if ribbon is more than 2 lines long
            limit = max(0, stats.get_ribbon_lines() - 2)
            if not limit:
                return
            # how many can we skip?
            if run or session:
                UI.display('We can skip up to {} lines.'.format(limit))

            if run:
                # run lines skipping
                # how many lines were successfully cast?
                lines_ok = stats.get_lines_done()
                if lines_ok:
                    UI.display('{} lines were cast in the last run.'
                               .format(lines_ok))
                # Ask user how many to skip (default: all successfully cast)
                r_skip = UI.enter('How many lines to skip for THIS run?',
                                  default=lines_ok, minimum=0, maximum=limit)
                stats.update(run_line_skip=r_skip)

            if session:
                # Skip lines effective for ALL runs
                # session line skipping affects multi-run sessions only
                # don't do it for single-run sessions
                if stats.runs > 1:
                    s_skip = UI.enter('How many lines to skip for ALL runs?',
                                      default=0, minimum=0, maximum=limit)
                    stats.update(session_line_skip=s_skip)

        def cast_queue(queue):
            """Casts the sequence of codes in given sequence.
            This function is executed until the sequence is exhausted
            or casting is stopped by machine or user."""
            # in punching mode, lack of row will trigger signal 15,
            # lack of column will trigger signal O
            # in punching and testing mode, signal O or 15 will be present
            # in the output combination as O15
            for record in queue:
                # check if signal will be cast at all
                if not record.has_signals:
                    UI.display_header(record.comment)
                    continue
                # display some info and cast the signals
                stats.update(record=record)
                # cast it and get current machine status
                status = self.machine.cast_one(record)
                # prepare data for display
                casting_status = OrderedDict()
                casting_status['Signals sent'] = ' '.join(status['signals'])
                casting_status['Pump'] = 'ON' if status['pump'] else 'OFF'
                casting_status['Wedge 0005 at'] = status['wedge_0005']
                casting_status['Wedge 0075 at'] = status['wedge_0075']
                casting_status['Speed'] = status['speed']
                # display status
                UI.clear()
                UI.display_parameters(stats.code_parameters, casting_status)

        # Ribbon pre-processing and casting parameters setup
        ribbon = [parse_record(code) for code in input_iterable]
        rewind_if_needed()
        # initialize statistics
        stats = bm.Stats()
        stats.update(ribbon=ribbon)
        # display some info for the user
        UI.display_parameters(stats.ribbon_parameters)
        # set the number of casting runs
        stats.update(runs=UI.enter('How many times do you want to cast this?',
                                   default=1, minimum=0))
        # initial line skipping
        set_lines_skipped(run=True, session=True)
        UI.display_parameters(stats.session_parameters)
        # check if the row 16 addressing will be used;
        # connect with the interface
        row16_in_use = any(record.uses_row_16 for record in ribbon)
        self.machine.choose_row16_mode(row16_in_use)
        with self.machine:
            while stats.get_runs_left():
                # Prepare the ribbon ad hoc
                queue = skip_lines(ribbon)
                stats.update(queue=queue)
                # Cast the run and check if it was successful
                try:
                    cast_queue(queue)
                    stats.update(casting_success=True)
                    if not stats.get_runs_left():
                        # make sure the machine will check
                        # whether it's running next time
                        self.machine.working = False
                        # user might want to re-run this
                        prm = 'Casting successfully finished. Any more runs?'
                        stats.update(runs=UI.enter(prm, default=0, minimum=0))

                except librpi2caster.MachineStopped:
                    stats.update(casting_success=False)
                    # aborted - ask if user wants to continue
                    runs_left = stats.get_runs_left()
                    if runs_left:
                        UI.confirm('{} runs left, continue?'.format(runs_left),
                                   default=True, abort_answer=False)
                    else:
                        UI.confirm('Retry casting?', default=True,
                                   abort_answer=False)
                    # offer to skip lines for re-casting the failed run
                    skip_successful = stats.get_lines_done() >= 2
                    set_lines_skipped(run=skip_successful, session=False)
                    # restart the machine after emergency stop
                    self.machine.start()

    @cast_this
    @mc.temp_wedge
    def cast_material(self):
        """Cast typesetting material: typecases, specified sorts, spaces"""
        def make_queue():
            """generate a sequence of items for casting"""
            while True:
                try:
                    UI.display('Specify characters to cast.\n'
                               'When done, press Enter to start casting.')
                    matrix = self.find_matrix(char='')
                    units = UI.enter('How many units?', datatype=int,
                                     default=round(self.get_units(matrix)))
                    qty = UI.enter('How many sorts?', default=10, minimum=0)
                    # ready to deliver
                    yield d.QueueItem(matrix, round(units, 2), qty)
                except (StopIteration, bm.MatrixNotFound):
                    break

        def make_ribbon(queue):
            """Take items from queue and adds them as long as there is
            space left in the galley. When space runs out, end a line.

            queue: [(code, quantity, units, pos_0075, pos_0005)]"""
            def new_mat():
                """matrix, width --> ribbon code, wedge positions"""
                # if the generator yields None first, try again
                queue_item = next(queue) or next(queue)
                sorts_left = queue_item.qty
                matrix, units = queue_item.matrix, queue_item.units
                # some info for the user
                UI.display('{} × {}, {} units {} set'
                           .format(sorts_left, matrix.code, units,
                                   self.wedge.set_width))
                # get code and wedge positions for the item
                positions = self.get_wedge_positions(matrix, units)
                record = matrix.get_ribbon_record(s_needle=positions != (3, 8))
                return record, units, positions, sorts_left

            def newline():
                """fill the line with quads, then spaces"""
                nonlocal units_left
                # single justification for the characters
                coarse, fine = wedges
                char_justification = ['NKS 0075 {}'.format(coarse),
                                      'NJS 0005 {}'.format(fine)]

                # add quads (one extra - last quad in the row)
                n_quads = 1 + int(units_left // self.quad.units)
                units_left %= self.quad.units
                try:
                    # variable space to adjust the width
                    mat = self.find_space(units=units_left)
                    coarse, fine = self.get_wedge_positions(mat, units_left)
                    use_s_needle = (coarse, fine) != (3, 8)
                    var_space = [mat.get_ribbon_record(s_needle=use_s_needle)]
                except bm.MatrixNotFound:
                    var_space = []
                    # no space to justify = just 0075+0005+NKJ
                    coarse, fine = '', ''
                # double justification sets the initial space width
                if coarse == fine:
                    # single code is enough
                    space_justification = ['NKJS 0005 0075 {}'.format(fine)]
                else:
                    space_justification = ['NKS 0075 {}'.format(coarse),
                                           'NKJS 0005 0075 {}'.format(fine)]

                units_left = self.measure.units - 2 * self.quad.units
                # single justification (for type), fillup spaces & quads,
                # double justification (for space), initial quad on new line
                return [*char_justification, *var_space, *[quad] * n_quads,
                        *space_justification, quad]

            def changeover():
                """use single-justification (0005+0075) to adjust wedges"""
                # add a quad between different series
                nonlocal units_left
                units_left -= self.quad.units * 2
                coarse, fine = wedges
                quads = [quad, quad]
                # single justification
                sjust = ([] if (coarse, fine) == (3, 8)
                         else ['NKS 0075 {}'.format(fine)] if fine == coarse
                         else ['NKS 0075 {}'.format(coarse),
                               'NJS 0005 {}'.format(fine)])
                return [*quads, *sjust]

            def add_code():
                """add codes to a ribbon, updating the number in the process"""
                nonlocal units_left, sorts_left
                # how many can we fit in the line? (until we've cast all)
                number = int(min(units_left // units, sorts_left))
                # update counters
                sorts_left -= number
                units_left -= number * units
                return [record] * number

            # em-quad for filling the line
            quad = self.quad.get_ribbon_record()
            # initialize the units
            self.measure = bc.set_measure(25, 'cc', 'galley width',
                                          self.wedge.set_width)
            units_left = self.measure.units - 2 * self.quad.units
            # get the first matrix
            # (if StopIteration is raised, no casting)
            record, units, wedges, sorts_left = new_mat()
            # first to set / last to cast last line out
            yield ['NJS 0005', 'NKJS 0005 0075', quad]
            # keep adding these characters
            while True:
                yield add_code()
                if sorts_left:
                    # still more to go...
                    yield newline()
                else:
                    # we're out of sorts... next character
                    try:
                        yield changeover()
                        record, units, wedges, sorts_left = new_mat()
                    except StopIteration:
                        # no more characters => fill the line and finish
                        # those are the first characters to cast
                        yield newline()
                        break
                    except (bm.TypesettingError, bm.MatrixNotFound) as exc:
                        UI.display('{}, omitting'.format(exc))
                        continue

        source = make_queue()
        ribbon = [code for chunk in make_ribbon(source) for code in chunk]
        return ribbon

    @cast_this
    @mc.temp_wedge
    def cast_qr_code(self):
        """Set up and cast a QR code which can be printed and then scanned
        with a mobile device."""
        def define_space(low):
            """find and set up a high or low space"""
            what = 'Low' if low else 'High'
            char = ' ' if low else '_'
            code = UI.enter('{} space matrix coordinates?'.format(what), '')
            space = bm.Matrix(char=char, code=code, diecase=self.diecase)
            wedges = self.get_wedge_positions(space, units)
            return space.get_ribbon_record(s_needle=wedges != (3, 8)), wedges

        def make_qr(data):
            """make a QR code matrix from data"""
            # QR rendering parameters
            border = UI.enter('QR code border width (squares)?', default=1,
                              minimum=1, maximum=10)
            ec_option = UI.enter('Error correction: 0 = lowest, 3 = highest?',
                                 default=1, minimum=0, maximum=3)
            # set up a QR code and generate a matrix
            modes = (qrcode.constants.ERROR_CORRECT_L,
                     qrcode.constants.ERROR_CORRECT_M,
                     qrcode.constants.ERROR_CORRECT_H,
                     qrcode.constants.ERROR_CORRECT_Q)
            engine = qrcode.QRCode(error_correction=modes[ec_option],
                                   border=border)
            engine.add_data(data)
            qr_matrix = engine.get_matrix()
            return qr_matrix

        def render(pattern):
            """translate a pattern into Monotype control codes,
            applying single justification if space widths differ,
            making spaces square in shape"""
            characters = {False: (low_space, ls_wedges),
                          True: (high_space, hs_wedges)}
            ribbon = ['NJS 0005', 'NKJS 0075 0005']
            for line in pattern:
                pairs = zip(line, [*line[1:], None])
                # newline (border is always low space)
                # double justification with low space wedges
                for current_item, next_item in pairs:
                    # add all spaces in a row
                    space, wedges = characters.get(current_item)
                    try:
                        _, next_wedges = characters.get(next_item)
                    except TypeError:
                        next_wedges = ls_wedges
                    ribbon.append(space)
                    if wedges != (3, 8) and wedges != next_wedges:
                        # set the wedges only if we need to
                        # use single justification in this case
                        ribbon.append('NKS 0075 {}'.format(wedges.pos_0075))
                        ribbon.append('NJS 0005 {}'.format(wedges.pos_0005))
                ribbon.append('NKS 0075 {}'.format(ls_wedges.pos_0075))
                ribbon.append('NKJS 0005 0075 {}'.format(ls_wedges.pos_0005))
            return ribbon

        # set the pixel size; smaler is preferred; depends on mould
        # (allow using different typesetting measures)
        px_size = bc.set_measure('6pt', what='pixel size (the same as mould)',
                                 set_width=self.wedge.set_width)
        units = px_size.units
        UI.display('The pixel is {} units {} set wide.'
                   .format(units, self.wedge.set_width))
        # determine the low and high space first
        try:
            low_space, ls_wedges = define_space(True)
            high_space, hs_wedges = define_space(False)
        except bm.TypesettingError as exc:
            UI.display(str(exc))
            UI.pause('Try again with a different pixel size or wedge')
            return self.cast_qr_code()
        # enter text and encode it
        text = UI.enter('Enter data to encode', '')
        qr_matrix = make_qr(text)
        # let the operator know how large the code is
        size = len(qr_matrix)
        prompt = ('The resulting QR code is {0} × {0} squares '
                  'or {1} × {1} inches.')
        UI.display(prompt.format(size, round(size * px_size.inches, 1)))
        UI.pause('Set your galley accordingly or abort.', allow_abort=True)
        # make a sequence of low and high spaces to cast
        return render(qr_matrix)

    @cast_this
    def cast_composition(self):
        """Casts or punches the ribbon contents if there are any"""
        if not self.ribbon.contents:
            raise Abort
        return self.ribbon.contents

    @cast_this
    @mc.temp_diecase
    @mc.temp_wedge
    def diecase_proof(self):
        """Tests the whole diecase, casting from each matrix.
        Casts spaces between characters to be sure that the resulting
        type will be of equal width."""
        def get_codes(matrix, wedges):
            """add codes with single justification if necessary"""
            combinations = []
            wedges_needed = wedges != (3, 8)
            code = matrix.get_ribbon_record(s_needle=wedges_needed)
            combinations.append(code)
            if wedges_needed:
                combinations.append('NKS 0075 {}'.format(wedges.pos_0075))
                combinations.append('NJS 0005 {}'.format(wedges.pos_0005))
            return combinations

        def add_matrix(matrix):
            """a matrix and a space"""
            # keep track of spaces too narrow to cast
            nonlocal leftover_units
            mat_units = self.get_units(matrix)
            space_units = 22 + leftover_units - mat_units
            leftover_units = 0
            try:
                mat_wedges = self.get_wedge_positions(matrix, mat_units)
            except bm.TypesettingError:
                # adjust char width as far as possible,
                # take away or add the rest to the space instead
                limits = self.wedge.get_adjustment_limits()
                row_units = matrix.get_units_from_row(wedge_used=self.wedge)
                if mat_units < row_units - limits.shrink:
                    # character too narrow for its row
                    space_units -= (row_units - mat_units - limits.shrink)
                    mat_units = row_units - limits.shrink
                elif mat_units > row_units + limits.stretch:
                    # character too wide for its row
                    space_units += (mat_units - row_units - limits.stretch)
                    mat_units = row_units + limits.stretch
                mat_wedges = self.get_wedge_positions(matrix, mat_units)
            # get these signals
            mat_codes = get_codes(matrix, mat_wedges)
            # omit spaces not wide enough
            if space_units >= 3:
                try:
                    # high spaces are preferred
                    space = self.find_space(space_units, low=False)
                except bm.MatrixNotFound:
                    # then we have to do with a low one...
                    space = self.find_space(space_units, low=True)
                sp_wedges = self.get_wedge_positions(space, space_units)
                space_codes = get_codes(space, sp_wedges)
            else:
                space_codes = []
                leftover_units += space_units
            # characters should be separated by at least one space
            return space_codes + mat_codes

        if self.diecase:
            self.display_diecase_layout()
        else:
            # select size for an empty layout
            self.resize_layout()

        if not UI.confirm('Proceed?', default=True, abort_answer=False):
            return

        queue = ['NJS 0005', 'NKJS 0075 0005']
        quad = self.quad.get_ribbon_record()
        pos_0075, pos_0005 = 3, 8
        leftover_units = 0
        # build the layout one by one
        for number, row in enumerate(self.diecase.layout.by_rows(), start=1):
            UI.display('Processing row {}'.format(number))
            queue.append(quad)
            for mat in row:
                queue.extend(add_matrix(mat))
            leftover_units = 0
            # single justification for line start
            queue.append(quad)
            queue.append('NKS 0075 {}'.format(pos_0075))
            queue.append('NKJS 0075 0005 {}'.format(pos_0005))

        return queue

    @mc.temp_wedge
    def calibrate_machine(self):
        """Casts the "en dash" characters for calibrating the character X-Y
        relative to type body."""
        def get_codes(char, default_position, comment):
            """Gets two mats for a given char and adjusts its parameters"""
            code = UI.enter('Where is the {}?'.format(char),
                            default=default_position)
            matrix = bm.Matrix(' ', code=code.upper())
            # try again recursively if wrong value
            if not matrix.position.row or not matrix.position.column:
                return get_codes(char, default_position, comment)
            # ask for unit width
            row_units = self.wedge[matrix.position.row]
            units = UI.enter('Unit width?', default=row_units)
            matrix.units = units
            # calculate the character width for measurement
            width = self.wedge.set_width / 12 * units / 18 * self.wedge.pica
            description = '{}, {:4f} inches wide'.format(comment, width)
            # calculate justifying wedge positions
            positions = mc.get_wedge_positions(matrix, self.wedge, units)
            pos_0075, pos_0005 = positions.pos_0075, positions.pos_0005
            use_s_needle = (pos_0075, pos_0005) != (3, 8)
            codes = matrix.get_ribbon_record(s_needle=use_s_needle,
                                             comment=description)
            sjust = ['NJS 0005 {}'.format(pos_0005),
                     'NKS 0075 {}'.format(pos_0075)]
            # use single justification to adjust character width, if needed
            return ([*sjust, codes, codes] if use_s_needle
                    else [codes, codes])

        UI.display('Mould blade opening and X-Y character calibration:\n'
                   'Cast G5, adjust the sort width to the value shown.\n'
                   '\nThen cast some lowercase "n" letters and n-dashes,\n'
                   'check the position of the character relative to the\n'
                   'type body and adjust the bridge X-Y.\n'
                   'Repeat if needed.\n')

        # build a casting queue and cast it repeatedly
        line_out = 'NKJS 0005 0075'
        pump_stop = 'NJS 0005'
        # use half-quad, quad, "n" and en-dash
        chars = [('half-quad', 'G5', 'full square'),
                 ('quad', 'O15', 'half square'),
                 ('n/h', None, 'serif overhanging test'),
                 ('dash or calibration mark', None, 'X-Y alignment test')]
        codes = (code for what in chars for code in get_codes(*what))
        sequence = [line_out, *codes, line_out, pump_stop]
        self.machine.cast(sequence)
        # G-8 calibration
        UI.display('\n\nCalibration done. Now adjust the matrix case draw rods'
                   '\nso that the diecase is not wobbling anymore.\n')
        with self.machine:
            self.machine.test_one('G8')
            UI.pause('Sending G8, press any key to stop...')

    def main_menu(self):
        """Main menu for the type casting utility."""
        def options():
            """Generate options based on current state of the program."""
            machine = self.machine
            is_casting = self.machine.is_casting()
            is_punching = not is_casting
            multiple_modes = len(self.machine.supported_operation_modes) >= 2

            got_ribbon = bool(self.ribbon)

            ret = [option(key='m', value=self.calibrate_machine, seq=5,
                          cond=is_casting,
                          text='Calibrate machine',
                          desc='Align the character width, then diecase'),

                   option(key='c', value=self.cast_composition, seq=10,
                          cond=is_casting and got_ribbon,
                          text='Cast composition',
                          desc='Cast type from a selected ribbon'),

                   option(key='r', value=self.choose_ribbon, seq=10,
                          text='Select ribbon',
                          desc='Select a ribbon from database or file'),

                   option(key='p', value=self.cast_composition, seq=30,
                          cond=is_punching and got_ribbon,
                          text='Punch ribbon',
                          desc='Punch a paper ribbon for casting'),

                   option(key='v', value=self.display_ribbon_contents, seq=80,
                          text='View codes', cond=got_ribbon,
                          desc='Display all codes in the selected ribbon'),

                   option(key='h', value=self.cast_material, seq=60,
                          cond=is_casting, text='Cast sorts or spaces',
                          desc='Cast characters from specified mats'),

                   option(key='q', value=self.cast_qr_code, seq=70,
                          cond=qrcode, text='Cast QR codes',
                          desc='Cast QR codes from high and low spaces'),

                   option(key='F3', value=self.choose_machine, seq=90,
                          text=('Change the machine in use (current: {})'
                                .format(machine))),

                   option(key='F4', value=self.machine.switch_operation_mode,
                          seq=91, cond=multiple_modes,
                          text=('Change operation mode (current: {})'
                                .format(machine.operation_mode))),

                   option(key='F5', value=self.display_details, seq=92,
                          text='Show details...',
                          desc='Display ribbon and interface information'),

                   option(key='F6', value=self.diecase_proof, seq=93,
                          text='Diecase proof',
                          desc='Cast every character from the diecase'),

                   option(key='F8', value=machine.diagnostics, seq=95,
                          text='Diagnostics menu...',
                          desc='Interface and machine diagnostic functions')]
            return ret

        header = ('rpi2caster - CAT (Computer-Aided Typecasting) '
                  'for Monotype Composition or Type and Rule casters.\n\n'
                  'This program reads a ribbon (from file or database) '
                  'and casts the type on a composition caster.'
                  '\n\nCasting / Punching Menu:')
        exceptions = (Finish, Abort, KeyboardInterrupt, EOFError)
        UI.dynamic_menu(options, header, catch_exceptions=exceptions)

    def display_details(self):
        """Collect ribbon, diecase and wedge data here"""
        data = [ribbon.parameters if self.ribbon else {},
                self.machine.parameters]
        UI.display_parameters(*data)
        UI.pause()
