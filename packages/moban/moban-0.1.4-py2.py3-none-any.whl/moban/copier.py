import os
import shutil

import moban.reporter as reporter


class Copier(object):
    def __init__(self, template_dirs):
        self.template_dirs = template_dirs
        self._count = 0

    def copy_files(self, file_list):
        for dest_src_pair in file_list:
            for dest, src in dest_src_pair.items():
                src_path = self._get_src_file(src)
                if src_path:
                    reporter.report_copying(src_path, dest)
                    shutil.copy(src_path, dest)
                    self._count = self._count + 1
                else:
                    reporter.report_error_message(
                        "{0} cannot be found".format(src))

    def number_of_copied_files(self):
        return self._count

    def report(self):
        if self._count:
            reporter.report_copying_summary(self._count)
        else:
            reporter.report_no_action()

    def _get_src_file(self, src):
        for folder in self.template_dirs:
            path = os.path.join(folder, src)
            if os.path.exists(path):
                return path
        else:
            return None
